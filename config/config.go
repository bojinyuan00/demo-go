package config

// Config 全局配置结构体
type Config struct {
	Server    ServerConfig   `mapstructure:"server"`
	Redis     RedisConfig    `mapstructure:"redis"`
	Databases DatabaseConfig `mapstructure:"databases"`
	Log       LogConfig      `mapstructure:"log"`
}
