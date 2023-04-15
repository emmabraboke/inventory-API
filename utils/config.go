package utils

import (
	"github.com/spf13/viper"
)

// Config file stores configuration of application

// Values read by viper viper from a config file or enviroment variable

type Config struct {
	DatabaseURI    string `mapstructure:"DATABASE_URI"`
	PaystackSecret string `mapstructure:"PAYSTACK_SECRET"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
	CloudinaryUrl	string `mapstructure:"CLOUDINARY_URL"`
}

func GetConfig(path string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var config Config

	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
