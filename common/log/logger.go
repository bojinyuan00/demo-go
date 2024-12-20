package log

import (
	"demo-go/common/global"
	"demo-go/common/utils"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	accessLogger    = logrus.New()
	errorLogger     = logrus.New()
	slowQueryLogger = logrus.New()
)

// InitLogger 初始化日志
func InitLogger() {
	cfg := global.Config
	// 读取配置
	logPath := cfg.Log.Path
	logLevel := cfg.Log.Level
	logFormat := cfg.Log.Format

	// 确保日志目录存在
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, os.ModePerm)
	}

	// 设置日志级别
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		level = logrus.InfoLevel
	}

	// 配置访问日志
	accessFile := fmt.Sprintf("%s/access_%s.log", logPath, time.Now().Format("2006-01-02"))
	setupLogger(accessLogger, accessFile, level, logFormat)

	// 配置错误日志
	errorFile := fmt.Sprintf("%s/error_%s.log", logPath, time.Now().Format("2006-01-02"))
	setupLogger(errorLogger, errorFile, level, logFormat)

	// 配置慢查询日志
	slowQueryFile := fmt.Sprintf("%s/slow_query_%s.log", logPath, time.Now().Format("2006-01-02"))
	setupLogger(slowQueryLogger, slowQueryFile, level, logFormat)
}

// checkLogFile 检查日志文件是否存在，不存在则重新创建
func checkLogFile(filePath, fileName string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		switch fileName {
		case "access":
			setupLogger(accessLogger, filePath, accessLogger.Level, global.Config.Log.Format) // 使用默认文本格式
		case "error":
			setupLogger(errorLogger, filePath, errorLogger.Level, global.Config.Log.Format) // 使用默认文本格式
		case "slow_query":
			setupLogger(slowQueryLogger, filePath, slowQueryLogger.Level, global.Config.Log.Format) // 使用默认文本格式
		}
	}
}

// setupLogger 配置日志实例
func setupLogger(logger *logrus.Logger, filePath string, level logrus.Level, format string) {
	logger.SetOutput(&lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    10,   // 单个日志文件最大大小（MB）
		MaxBackups: 7,    // 最大保留备份数
		MaxAge:     30,   // 最长保存时间（天）
		Compress:   true, // 是否压缩备份文件
	})

	// 设置日志格式
	if format == "json" {
		logger.SetFormatter(&logrus.JSONFormatter{
			//TimestampFormat: time.RFC3339,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		logger.SetFormatter(&logrus.TextFormatter{
			//TimestampFormat: time.RFC3339,
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		})
	}

	// 设置日志级别
	logger.SetLevel(level)
}

// AccessLogger 记录访问日志
func AccessLogger(method, path string, statusCode int, latency time.Duration, clientIP string) {
	// 检查日志文件是否存在，不存在则创建（服务启动后，删除日志文件需要重建）
	accessFile := fmt.Sprintf("%s/access_%s.log", global.Config.Log.Path, time.Now().Format("2006-01-02"))
	checkLogFile(accessFile, "access") // 检查日志文件是否存在

	//file, line, function := utils.GetCallerInfo(1)
	accessLogger.WithFields(logrus.Fields{
		"logType":     "Access-log",
		"method":      method,
		"path":        path,
		"status_code": statusCode,
		"latency_ms":  latency.Milliseconds(),
		"client_ip":   clientIP,
		//"file":        file,
		//"line":        line,
		//"function":    function,
	}).Info("Access-log")
}

// ErrorLogger 记录错误日志
func ErrorLogger(err error, message string) {
	// 检查日志文件是否存在，不存在则创建（服务启动后，删除日志文件需要重建）
	errorFile := fmt.Sprintf("%s/error_%s.log", global.Config.Log.Path, time.Now().Format("2006-01-02"))
	checkLogFile(errorFile, "error") // 检查日志文件是否存在
	file, line, function := utils.GetCallerInfo(2)
	errorLogger.WithFields(logrus.Fields{
		"logType":  "Error-log",
		"error":    err,
		"message":  message,
		"file":     file,
		"line":     line,
		"function": function,
	}).Error("Error-log")
}

// SlowQueryLogger 记录慢查询日志
func SlowQueryLogger(method, path string, latency time.Duration) {
	threshold := viper.GetInt64("log.slow_query_threshold_ms")
	if latency.Milliseconds() > threshold {
		// 检查日志文件是否存在，不存在则创建（服务启动后，删除日志文件需要重建）
		slowQueryFile := fmt.Sprintf("%s/slow_query_%s.log", global.Config.Log.Path, time.Now().Format("2006-01-02"))
		checkLogFile(slowQueryFile, "slow_query") // 检查日志文件是否存在

		file, line, function := utils.GetCallerInfo(1)
		slowQueryLogger.WithFields(logrus.Fields{
			"logType":    "Slow-query-log",
			"method":     method,
			"path":       path,
			"latency_ms": latency.Milliseconds(),
			"file":       file,
			"line":       line,
			"function":   function,
		}).Warn("Slow-query-detected")
	}
}

func TimeTracker(start time.Time, params interface{}, result interface{}, err error) {
	// 检查日志文件是否存在，不存在则创建（服务启动后，删除日志文件需要重建）
	accessFile := fmt.Sprintf("%s/access_%s.log", global.Config.Log.Path, time.Now().Format("2006-01-02"))
	checkLogFile(accessFile, "access") // 检查日志文件是否存在
	elapsed := time.Since(start)       //时间差

	//有error时，记录相关的调用栈信息（不需要记录file、line、function信息，因为err内已经包含）且单独记录error到error日志文件内
	if err != nil {
		accessLogger.WithFields(logrus.Fields{
			"logType":    "Tracker-log",
			"latency_ms": elapsed.Milliseconds(),
			"params":     params,
			"result":     result,
			"error":      err,
		}).Info("Tracker-log")
		ErrorLogger(err, "TimeTracker-ErrorLog")
		return
	}

	//无error时，只记录相关的调用栈信息--增加日志记录的灵活性（file、line、function）具体信息
	file, line, function := utils.GetCallerInfo(3)
	accessLogger.WithFields(logrus.Fields{
		"logType":    "Tracker-log",
		"latency_ms": elapsed.Milliseconds(),
		"params":     params,
		"result":     result,
		"error":      err,
		"file":       file,
		"line":       line,
		"function":   function,
	}).Info("Tracker-log")

}

// CustomLogger 自定义日志记录
func CustomLogger(start time.Time, fields logrus.Fields) {
	// 检查日志文件是否存在，不存在则创建（服务启动后，删除日志文件需要重建）
	accessFile := fmt.Sprintf("%s/access_%s.log", global.Config.Log.Path, time.Now().Format("2006-01-02"))
	checkLogFile(accessFile, "access")
	elapsed := time.Since(start)
	file, line, function := utils.GetCallerInfo(2)
	fields["logType"] = "Custom-logger"
	fields["latency_ms"] = elapsed.Milliseconds()
	fields["file"] = file
	fields["line"] = line
	fields["function"] = function
	accessLogger.WithFields(fields).Info("Custom-logger")
}

// getCallerInfo 获取调用日志函数的文件名、行号和方法名
func getCallerInfo() (file string, line int, function string) {
	pc, file, line, ok := runtime.Caller(2)
	if !ok {
		return "unknown", 0, "unknown"
	}
	fn := runtime.FuncForPC(pc)
	function = "unknown"
	if fn != nil {
		function = fn.Name()
	}
	return file, line, function
}
