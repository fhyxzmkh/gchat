package config

import (
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strings"
)

func NewViperConfig() error {
	// 1. 先加载 .env 文件（如果有）
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found, will use system environment variables")
	}

	// 2. 设置 Viper 配置
	viper.SetConfigName("global")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./internal/config")
	viper.AutomaticEnv()

	// 3. 使 Viper 能读取嵌套的环境变量（如 ${MYSQL_PASSWORD}）
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AllowEmptyEnv(true)

	// 4. 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
