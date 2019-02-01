package viper

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
)

// Init viper config.
func InitViperConfig(configName string) error {
	config := Config{
		Name: configName,
	}

	// Init
	if err := config.initConfig(); err != nil {
		return err
	}

	// Viper hot update
	config.watchConfig()

	return nil
}

// Return viper all settings.
func AllViperSettings() map[string]interface{} {
	return viper.AllSettings()
}

type Config struct {
	Name string
}

// Init config.
func (c *Config) initConfig() error {

	var configPath string

	// Get system GOPATH
	GOPATH := os.Getenv("GOPATH")

	if GOPATH != "" {
		configPath = configPath + GOPATH + "/src/bilibili-rear-end/configer/"
	} else {
		configPath = "./configer/"
	}

	// If config's name is empty, set default
	if c.Name != "" {
		// name of config file (without extension)
		viper.SetConfigName(c.Name)
	} else {
		viper.SetConfigName("app")
	}

	viper.AddConfigPath(configPath)
	// Current use yaml config file.
	viper.SetConfigType("yaml")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// Hot update.
func (c *Config) watchConfig() {
	viper.WatchConfig()  // Mark sure you add all of the configPaths prior to calling WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
	})
}