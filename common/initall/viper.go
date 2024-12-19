package initall

import (
	"demo-go/common/global"
	"fmt"
	"github.com/spf13/viper"
)

// LoadConfig 加载配置
func loadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error resources file: %w \n", err))
	}
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("unable to decode into struct %w \n", err))
	}
	return nil
}
