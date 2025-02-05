// config.go

package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort          string
	DatabaseHost     string
	DatabasePort     int
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseSSLMode  string
}

func LoadConfig() (Config, error) {
	// viper.SetEnvPrefix("POSTGRES")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Set default values
	viper.SetDefault("POSTGRES.HOST", "localhost")
	viper.SetDefault("POSTGRES.PORT", 5432)
	viper.SetDefault("POSTGRES.USER", "postgres")
	viper.SetDefault("POSTGRES.PASSWORD", "")
	viper.SetDefault("POSTGRES.DBNAME", "bookstore")
	viper.SetDefault("POSTGRES.SSLMODE", "disable")

	// Set config values
	config := Config{
		AppPort:          viper.GetString("APP.PORT"),
		DatabaseHost:     viper.GetString("POSTGRES.HOST"),
		DatabasePort:     viper.GetInt("POSTGRES.PORT"),
		DatabaseUser:     viper.GetString("POSTGRES.USER"),
		DatabasePassword: viper.GetString("POSTGRES.PASSWORD"),
		DatabaseName:     viper.GetString("POSTGRES.DBNAME"),
		DatabaseSSLMode:  viper.GetString("POSTGRES.SSLMODE"),
	}

	return config, nil
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseName,
		c.DatabaseSSLMode)
}