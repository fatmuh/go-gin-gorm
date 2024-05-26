package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	DBHost     string `mapstructure:"MYSQL_HOST"`
	DBUsername string `mapstructure:"MYSQL_USER"`
	DBPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName     string `mapstructure:"MYSQL_DB"`
	DBPort     string `mapstructure:"MYSQL_PORT"`

	ServerPort string `mapstructure:"PORT"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}
