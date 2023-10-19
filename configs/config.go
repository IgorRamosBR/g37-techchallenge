package configs

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	viper *viper.Viper
}

type AppConfig struct {
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) GetViperConfig() *viper.Viper {
	return c.viper
}

func (c *Config) ReadConfig() (AppConfig, error) {
	log.Info("Reading Environment Variables")
	c.setupEnvironment()

	log.Info("Setting Config Path")
	c.setupConfigPath()

	log.Info("Reading Config File")
	err := c.viper.ReadInConfig()
	if err != nil {
		return AppConfig{}, fmt.Errorf("error reading config file or env variable, error: %v", err)
	}

	appConfig, err := c.extractConfigVars()
	if err != nil {
		return AppConfig{}, err
	}

	return appConfig, nil
}

func (c *Config) setupConfigPath() {
	// Get proper config directory
	configDirPath := c.viper.GetString("CONFIG_DIR_PATH")
	if configDirPath == "" {
		configDirPath = "./configs"
	}
	log.Infof("ConfigPath %v", configDirPath)
	c.viper.AddConfigPath(configDirPath)
}

func (c *Config) setupEnvironment() {
	c.viper = viper.New()
	c.viper.AutomaticEnv()

	// Next, Read yaml config file based on the AF_ENVIRONMENT
	environment := c.viper.GetString("ENVIRONMENT")
	log.Infof("ENVIRONMENT %s", environment)
	c.viper.SetConfigType("yaml")
	c.viper.SetConfigName(environment)
}

func (c *Config) extractConfigVars() (AppConfig, error) {
	appConfig := AppConfig{}

	appConfig.DatabaseHost = c.viper.GetString("postgres.host")
	appConfig.DatabasePort = c.viper.GetString("postgres.port")
	appConfig.DatabaseName = c.viper.GetString("postgres.dbname")
	appConfig.DatabaseSSLMode = c.viper.GetString("postgres.sslmode")
	appConfig.DatabaseUser = c.viper.GetString("postgres.user")
	appConfig.DatabasePassword = c.viper.GetString("postgres.portpassword")

	return appConfig, nil
}
