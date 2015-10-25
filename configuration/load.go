package configuration

import (
	"bytes"
	"fmt"
	"io/ioutil"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// Settings for the setting manager
var (
	configName  = "settings"
	configPaths = []string{
		"~/.monf",
		"./",
		"/etc/monf",
	}
	configType = "yaml"
)

func init() {
	// We will use yaml settings
	viper.SetConfigType(configType)
}

// LoadDefaultFilePathSettings Sets the configuration default file paths
func LoadDefaultFilePathSettings() {
	// Configure settings files
	viper.SetConfigName(configName)
	for _, i := range configPaths {
		viper.AddConfigPath(i)
	}

	// Load the settings
	err := viper.ReadInConfig()
	if err != nil {
		log.Warning("Default settings paths not found")
	}
}

// LoadFromFileSettings loads the settings from an specified path
func LoadFromFileSettings(path string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		log.Warning("Error loading '%s' settings file", path)
	} else {
		viper.ReadConfig(bytes.NewBuffer(data))
	}

}

// LoadAppDefaultSettings loads the default settings
func LoadAppDefaultSettings() {
	for k, v := range defaultAppSettings {
		viper.SetDefault(k, v)
	}
	log.Info("Loaded default settings")
}

// LoadSettings Load all the app settings
func LoadSettings(path string) {
	// Load settings
	if path != "" {
		LoadFromFileSettings(path)
	} else {
		LoadDefaultFilePathSettings()
	}
	LoadAppDefaultSettings()

	log.Info("Loaded Settings: ")
	for k, v := range viper.AllSettings() {
		log.Infof(" - %s: %v", k, v)
	}
	fmt.Println("")

	// Set debug level
	debug := viper.GetBool(DEBUG)
	if debug {
		log.SetLevel(log.DebugLevel)
		log.Info("Logging level: debug")
	} else {
		log.SetLevel(log.InfoLevel)
		log.Info("Logging level: info")
	}
}
