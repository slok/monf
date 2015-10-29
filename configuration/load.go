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
	// We will use yaml settings
	viper.SetConfigType(configType)

	// Get the settigns file from env var
	viper.SetEnvPrefix("monf")
	viper.BindEnv(SettingsPath)
	settingsFile := viper.GetString(SettingsPath)

	// start App configuration
	LoadSettings(settingsFile)
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
func LoadSettings(path string) {
	// Load settings
	if path != "" {
		LoadFromFileSettings(path)
	} else {
		if path != "" {
			LoadFromFileSettings(SpecificConfigPath)
		}
		LoadDefaultFilePathSettings()
	}
	LoadAppDefaultSettings()

	// the first thing after loading the settings is to set the log level
	ConfigureLogs()

	log.Info("Loaded Settings: ")
	for k, v := range viper.AllSettings() {
		log.Infof(" - %s: %v", k, v)
	}
}
