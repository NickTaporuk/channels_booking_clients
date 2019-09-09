package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func NewConfig() (*Configuration, error) {
	var (
		err            error
		configFileName = "config"
	)
	v := viper.New()
	// Set the path to look for the configurations file
	v.AddConfigPath("./data")
	v.AddConfigPath("$HOME/.cbc")

	v.SetConfigName(configFileName)

	// Enable VIPER to read Environment Variables
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var configuration Configuration

	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = v.Unmarshal(&configuration)
	if err != nil {
		return nil, errors.Wrap(err, "unable to decode into struct", )
	}

	return &configuration, nil
}
