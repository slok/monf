package configuration

// Configuration names
const (
	Debug         = "Debug"
	ListenPort    = "ListenPort"
	ListenHost    = "ListenHost"
	TemplatesPath = "TemplatesPath"
	StaticURL     = "StaticUrl"
	StaticPath    = "StaticPath"
	DisableLogs   = "DisableLogs"
	LogLevel      = "LogLevel"
	SettingsPath  = "Settings" // on env var MONF_SETTINGS
)

// Default app settings for the application
var defaultAppSettings = map[string]interface{}{

	// sets the app on debug level
	Debug:         false,
	ListenPort:    3000,
	ListenHost:    "",
	TemplatesPath: "./templates",
	StaticURL:     "/static",
	StaticPath:    "./static",
	DisableLogs:   false,
	LogLevel:      "info",
	SettingsPath:  "",
}
