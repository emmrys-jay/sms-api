package config

import (
	"log"

	"github.com/Emmrys-Jay/altschool-sms/utility"
	"github.com/spf13/viper"
)

type Configuration struct {
	Port          string `mapstructure:"PORT"`
	SecretKey     string `mapstructure:"SECRET_KEY"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBUsername    string `mapstructure:"DB_USERNAME"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	AdminEmail    string `mapstructure:"ADMIN_EMAIL"`
	AdminPassword string `mapstructure:"ADMIN_PASSWORD"`
}

// Setup initialize configuration
var (
	Config *Configuration
)

// Params = getConfig.Params
func Setup() {
	var configuration *Configuration
	logger := utility.NewLogger()

	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Params = configuration.Params
	Config = configuration
	logger.Info("configurations loading successfully")
}

// GetConfig helps you to get configuration data
func GetConfig() *Configuration {
	return Config
}
