package config

// Config 全局配置
type Config struct {
	Server    ServerConfig   `mapstructure:"server"`
	Redis     RedisConfig    `mapstructure:"redis"`
	Databases DatabaseConfig `mapstructure:"databases"`
	Log       LogConfig      `mapstructure:"log"`
}

type LogConfig struct {
	Path                 string `mapstructure:"path"`                    // 日志文件路径
	Level                string `mapstructure:"level"`                   // 日志级别：debug, info, warn, error, fatal, panic
	Format               string `mapstructure:"format"`                  // 日志格式：json, text
	AccessFile           string `mapstructure:"access_file"`             // 请求记录日志名
	ErrorFile            string `mapstructure:"error_file"`              // 错误记录日志名
	SlowQueryFile        string `mapstructure:"slow_query_file"`         // 慢查询记录日志名
	SlowQueryThresholdMs int    `mapstructure:"slow_query_threshold_ms"` // 慢查询阈值（毫秒）
}

//// 默认日志配置
//var DefaultLogConfig = LogConfig{
//	Path: "logs",
//	Level: "debug",
//	access_file: "access.log"
//	error_file: "error.log"
//	slow_query_file: "slow_query.log"
//	slow_query_threshold_ms: 500 # 慢查询阈值（毫秒）
//}

// ServerConfig 服务启动端口号配置
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// DatabaseConfig gorm数据库配置
type DatabaseConfig struct {
	Driver       string `mapstructure:"driver"`
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	DBName       string `mapstructure:"db_name"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	ConnMaxLife  int    `mapstructure:"conn_max_lifetime"`
}

// RedisConfig redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}
