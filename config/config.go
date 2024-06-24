package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"sync"
)

type Server struct {
	Port int
}

type Database struct {
	Driver   string
	Host     string
	Port     int
	User     string
	Password string
	Dbname   string
	ssl      string
	Timezone string
}

type Config struct {
	Server   *Server
	Database *Database
}

// LoadConfig
// Load configuration from file
func LoadConfig(path string) (*Config, error) {
	var (
		once         sync.Once
		err          error
		loadedConfig = &Config{
			Server:   &Server{},
			Database: &Database{},
		}
	)
	once.Do(func() {
		viper.AddConfigPath(path)
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AutomaticEnv()
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		err = viper.ReadInConfig()
		if err == nil {
			err = viper.Unmarshal(&loadedConfig)
		}
	})

	return loadedConfig, err
}

// GetDatabaseSource
// Method of Config struct
// Use config to prepare data source connection string
func (config *Config) GetDatabaseSource() string {
	dbSource := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s&timezone=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Dbname,
		config.Database.ssl,
		config.Database.Timezone,
	)
	return dbSource
}
