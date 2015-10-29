package configuration

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

// logrus formatter to disable the logs
type noFormatter struct{}

func (f *noFormatter) Format(entry *log.Entry) ([]byte, error) { return nil, nil }

func disableLogs() {
	//log.Warning("Disabled logging")
	log.SetFormatter((log.Formatter)(&noFormatter{}))
}

func setDebugLevel() {
	log.Info("Logging set on debug mode")
	log.SetLevel(log.DebugLevel)
}

func setLevel() {
	level, err := log.ParseLevel(viper.GetString(LogLevel))
	// If error then default info mode
	if err != nil {
		level = log.InfoLevel
	}
	log.SetLevel(level)
}

// ConfigureLogs configure the logging system based on teh log settings
func ConfigureLogs() {

	// If logging is disabled then
	if viper.GetBool(DisableLogs) {
		disableLogs()
	}

	// Set Log level
	setLevel()

	// Set debug level logging on debug mode
	if viper.GetBool(Debug) {
		setDebugLevel()
	}

	log.Infof("Logging level: %v", log.GetLevel())
}
