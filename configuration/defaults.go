package configuration

// Configuration names
const (
	Debug         = "debug"
	ListenPort    = "listen-port"
	ListenHost    = "listen-host"
	TemplatesPath = "templates-path"
	StaticURL     = "static-url"
	StaticPath    = "static-path"
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
