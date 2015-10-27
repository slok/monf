package configuration

// Configuration names
const (
	Debug         = "Debug"
	ListenPort    = "ListenPort"
	ListenHost    = "ListenHost"
	TemplatesPath = "TemplatesPath"
	StaticURL     = "StaticUrl"
	StaticPath    = "StaticPath"
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
}
