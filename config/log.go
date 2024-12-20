package config

type LogConfig struct {
	Path                 string `mapstructure:"path"`                    // 日志文件路径
	Level                string `mapstructure:"level"`                   // 日志级别：debug, info, warn, error, fatal, panic
	Format               string `mapstructure:"format"`                  // 日志格式：json, text
	AccessFile           string `mapstructure:"access_file"`             // 请求记录日志名
	ErrorFile            string `mapstructure:"error_file"`              // 错误记录日志名
	SlowQueryFile        string `mapstructure:"slow_query_file"`         // 慢查询记录日志名
	SlowQueryThresholdMs int    `mapstructure:"slow_query_threshold_ms"` // 慢查询阈值（毫秒）
}
