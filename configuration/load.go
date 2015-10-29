package configuration

import (
	"bytes"
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

// SpecificConfigPath will be the global var to track the custom config file
var SpecificConfigPath = ""

func init() {
	ResetSettings()
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
		log.Warningf("Error loading '%s' settings file", path)
	} else {
		viper.ReadConfig(bytes.NewBuffer(data))
		SpecificConfigPath = path
	}

}

// LoadAppDefaultSettings loads the default settings
func LoadAppDefaultSettings() {
	for k, v := range defaultAppSettings {
		viper.SetDefault(k, v)
	}
}

// LoadSettings Load all the app settings
// If specific path then don't look at the env var, nor default places
func LoadSettings(path string) {

	// 1 - Default settings
	LoadAppDefaultSettings()

	// 2 - Specific file settings
	if path != "" {
		LoadFromFileSettings(path)
		SpecificConfigPath = path
	} else {
		// 3 - Env var settings file
		envSettingsFile := viper.GetString(SettingsPath)
		if envSettingsFile != "" {
			LoadFromFileSettings(envSettingsFile)
		} else { // 4 finally load from one of the default places
			LoadDefaultFilePathSettings()
		}

	}

	// the first thing after loading the settings is to set the log level
	ConfigureLogs()

	log.Info("Loaded Settings: ")
	for k, v := range viper.AllSettings() {
		log.Infof(" - %s: %v", k, v)
	}
}
