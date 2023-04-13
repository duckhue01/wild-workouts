package service

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
	"github.com/ory/viper"
)

// ReadConfig read config from configuration file
func ReadConfig[C any](path, name, typ string) (*C, error) {
	c := new(C)
	viper.SetConfigName(name)
	viper.SetConfigType(typ)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return c, fmt.Errorf("read config file: %w", err)
	}

	if err := viper.UnmarshalExact(c, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		return c, fmt.Errorf("unmarshal config: %w", err)
	}
	return c, nil
}

// ReadSecret read config from configuration file
func ReadSecret[C any](path, name, typ string) (*C, error) {
	c := new(C)
	viper.SetConfigName(name)
	viper.SetConfigType(typ)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		return c, fmt.Errorf("read secret file: %w", err)
	}

	if err := viper.UnmarshalExact(c, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
	}); err != nil {
		return c, fmt.Errorf("unmarshal secret: %w", err)
	}
	return c, nil
}
