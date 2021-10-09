package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Files struct {
		PokemonCsv   string
		HtmlTemplate string
	}

	ExternalUrl struct {
		GetPokemonById string
		TimeOut        time.Duration
	}
}

func ReadConfig() *Config {
	config := &Config{}
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		fmt.Println(err)
	}

	return config
}
