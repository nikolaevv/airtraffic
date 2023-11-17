package config

import (
	"errors"

	"github.com/spf13/viper"
)

func Init(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigFile(filename)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if errors.As(err, &viper.ConfigFileNotFoundError{}) {
			return nil, errors.New("config file not found")
		}

		return nil, err
	}

	return v, nil
}
