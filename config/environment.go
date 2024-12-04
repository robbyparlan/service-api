package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DBConfig struct {
	DBHost            string `mapstructure:"DB_HOST"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	DBPort            int    `mapstructure:"DB_PORT"`
	DBSSL             string `mapstructure:"DB_SSL"`
	DBMaxOpenConn     int    `mapstructure:"DB_SET_MAX_OPEN_CONN"`
	DBIdleConn        int    `mapstructure:"DB_SET_IDLE_CONN"`
	DBConnMaxLifeTime int    `mapstructure:"DB_CONN_MAX_LIFE_TIME"`
}

type Config struct {
	DB              DBConfig `mapstructure:"DATABASE"`
	AppPort         string   `mapstructure:"APP_PORT"`
	AppReadTimeout  int      `mapstructure:"APP_READ_TIME_OUT"`
	AppWriteTimeout int      `mapstructure:"APP_WRITE_TIME_OUT"`
	LogsDir         string   `mapstructure:"LOGS_DIR"`
	JwtSecretKey    string   `mapstructure:"JWT_SECRET_KEY"`
}

// LoadConfig initializes and loads the configuration from a config.yml file
func LoadConfig() (*Config, error) {
	log.Println("--------------------- : Loading config", os.Getenv("APP_ENV"))
	env := os.Getenv("APP_ENV") // Default: empty
	var fileConfig string
	if env == "local" {
		fileConfig = "config-local"
		log.Println("--------------------- : Using local config")
	} else {
		fileConfig = "config"
		log.Println("--------------------- : Using production config")
	}

	viper.SetConfigName(fileConfig)
	viper.SetConfigType("yml")
	viper.AddConfigPath(".") // Current directory
	viper.AutomaticEnv()     // Enable automatic environment variable binding

	bindEnvironmentVariables()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	// Watch for changes in the config file
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		if err := viper.Unmarshal(&config); err != nil {
			log.Printf("Error reloading config: %v", err)
		} else {
			log.Println("Config successfully reloaded")
		}
	})

	return &config, nil
}

// bindEnvironmentVariables binds specific environment variables to viper
func bindEnvironmentVariables() {
	envBindings := map[string]string{
		"DATABASE.DB_HOST":               "DB_HOST",
		"DATABASE.DB_USER":               "DB_USER",
		"DATABASE.DB_PASSWORD":           "DB_PASSWORD",
		"DATABASE.DB_NAME":               "DB_NAME",
		"DATABASE.DB_PORT":               "DB_PORT",
		"DATABASE.DB_SSL":                "DB_SSL",
		"DATABASE.DB_SET_MAX_OPEN_CONN":  "DB_SET_MAX_OPEN_CONN",
		"DATABASE.DB_SET_IDLE_CONN":      "DB_SET_IDLE_CONN",
		"DATABASE.DB_CONN_MAX_LIFE_TIME": "DB_CONN_MAX_LIFE_TIME",
		"APP_PORT":                       "APP_PORT",
		"APP_READ_TIME_OUT":              "APP_READ_TIME_OUT",
		"APP_WRITE_TIME_OUT":             "APP_WRITE_TIME_OUT",
		"JWT_SECRET_KEY":                 "JWT_SECRET_KEY",
		"LOGS_DIR":                       "LOGS_DIR",
	}

	for key, env := range envBindings {
		if err := viper.BindEnv(key, env); err != nil {
			log.Printf("Error binding environment variable: %s -> %s", key, env)
		}
	}
}
