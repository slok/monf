package configuration

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// ResetSettings resets the viper settings
func ResetSettings() {
	log.Debug("Resetting all the settings")
	viper.Reset()
	viper.SetConfigType(configType)
	LoadSettings("")
}

// OverrideSettings overrides the passed ones, useful for tests
func OverrideSettings(settings map[string]interface{}) {
	for k, v := range settings {
		viper.Set(k, v)
	}
}
