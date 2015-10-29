package configuration

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// ResetSettings resets the viper settings
func ResetSettings() {
	log.Debug("Resetting all the settings")
	viper.Reset()

	// We will use yaml settings
	viper.SetConfigType(configType)

	// Get the settigns file from env var
	viper.SetEnvPrefix("monf")
	viper.BindEnv(SettingsPath)

	// start App configuration
	// TODO: Specify settings file
	filePath := ""
	if SpecificConfigPath != "" {
		filePath = SpecificConfigPath
	}
	LoadSettings(filePath)
}

// OverrideSettings overrides the passed ones, useful for tests
func OverrideSettings(settings map[string]interface{}) {
	for k, v := range settings {
		viper.Set(k, v)
	}
}
