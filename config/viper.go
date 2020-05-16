package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

func PopulateConfig(filename string, dst interface{}, watchCallback func()) {
	viper.SetConfigName(filename) // name of config file (without extension)
	viper.AddConfigPath(".")      // path to look for the config file in

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("fatal error config file: %s \n", err)
	}

	loadConfiguration(dst)

	if watchCallback != nil {
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			loadConfiguration(dst)
			watchCallback()
		})
	}
}

func loadConfiguration(dst interface{}) {
	err := viper.Unmarshal(&dst)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
