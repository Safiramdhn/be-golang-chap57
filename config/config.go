package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppDebug   bool
	Email      EmailConfig
	ServerPort string
}

type EmailConfig struct {
	ApiKey    string
	FromName  string
	FromEmail string
}

func LoadConfig() (Config, error) {
	// Set default values
	setDefaultValues()

	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.SetConfigType("dotenv")
	viper.SetConfigName(".env")

	// Allow Viper to read environment variables
	viper.AutomaticEnv()

	// Read the configuration file
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file: %s, using default values or environment variables", err)
	}

	// add value to the config
	config := Config{
		Email:      loadEmailConfig(),
		AppDebug:   viper.GetBool("APP_DEBUG"),
		ServerPort: viper.GetString("SERVER_PORT"),
	}
	return config, nil
}
func loadEmailConfig() EmailConfig {
	return EmailConfig{
		ApiKey:    viper.GetString("MAILERSEND_API_KEY"),
		FromName:  viper.GetString("MAILERSEND_FROM_NAME"),
		FromEmail: viper.GetString("MAILERSEND_FROM_EMAIL"),
	}
}

func setDefaultValues() {
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "admin")
	viper.SetDefault("DB_NAME", "database")
	viper.SetDefault("APP_DEBUG", true)
	viper.SetDefault("APP_SECRET", "team-1")
	viper.SetDefault("SERVER_PORT", ":8080")
	viper.SetDefault("SHUTDOWN_TIMEOUT", 5)

	viper.SetDefault("PROFIT_MARGIN", 10.00)

	viper.SetDefault("DB_MIGRATE", false)
	viper.SetDefault("DB_SEEDING", false)
}
