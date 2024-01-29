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
	Environment string

	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string

	SQSRegion   string
	SQSEndpoint string

	PaymentBrokerURL string
	NotificationURL  string
	SponsorId        string
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

	environment := c.viper.GetString("ENVIRONMENT")
	log.Infof("ENVIRONMENT %s", environment)
	c.viper.SetConfigType("yaml")
	c.viper.SetConfigName(environment)
}

func (c *Config) extractConfigVars() (AppConfig, error) {
	appConfig := AppConfig{}

	appConfig.Environment = c.viper.GetString("ENVIRONMENT")

	appConfig.DatabaseHost = c.viper.GetString("POSTGRES_HOST")
	appConfig.DatabasePort = c.viper.GetString("POSTGRES_PORT")
	appConfig.DatabaseName = c.viper.GetString("POSTGRES_DB")
	appConfig.DatabaseSSLMode = c.viper.GetString("POSTGRES_SSLMODE")
	appConfig.DatabaseUser = c.viper.GetString("POSTGRES_USER")
	appConfig.DatabasePassword = c.viper.GetString("POSTGRES_PASSWORD")

	appConfig.PaymentBrokerURL = c.viper.GetString("paymentBroker.url")
	appConfig.NotificationURL = c.viper.GetString("paymentBroker.notificationUrl")
	appConfig.SponsorId = c.viper.GetString("paymentBroker.sponsorId")

	return appConfig, nil
}
