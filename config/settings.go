package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Database string `mapstructure:"name"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

func New() *Config {

	viper.SetConfigName("config")
	viper.AddConfigPath("config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")

	var config Config

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error while reading config file: %v", err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Configuration file changed")
	})

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode config file to struct, err: %v", err)
	}

	// Load .ENV data
	// serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	// dbHost := os.Getenv("DB_HOST")
	// dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	// dbDatabase := os.Getenv("DB_NAME")
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")

	// fmt.Printf("ISTO %s", dbHost)

	// database := DatabaseConfig{dbHost, dbPort, dbDatabase, dbUser, dbPassword}
	// server := ServerConfig{serverPort}

	// config := Config{server, database}

	fmt.Print("Loaded configuration files \n")
	return &config
}
