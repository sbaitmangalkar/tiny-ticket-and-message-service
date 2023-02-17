package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
	"time"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	AppVersion   string
	Host         string
	Port         string
	Mode         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func GetConfigPath(configPath string) string {
	return "./config/config"
}

var once sync.Once
var config Config

func loadConfig() error {
	configPath := GetConfigPath(os.Getenv("environment"))

	v := viper.New()
	v.SetConfigName(configPath)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config file not found")
		}
		return err
	}

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return err
	}

	config = c
	return nil
}

func Get() (Config, error) {
	var err error
	once.Do(func() {
		err = loadConfig()
	})
	if err != nil {
		log.Printf("Failed to initialise configuration due to - " + err.Error())
		return config, err
	} else {
		return config, nil
	}

}
