package middleware

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/spf13/viper"
)

var settings map[string]interface{}

func init() {
	settings = viper.AllSettings()
}

// SettingsContext sets all the app settings on the request context
type SettingsContext struct{}

// NewSettingsContext returns a new instance of AppSettingsContext
func NewSettingsContext() *SettingsContext {
	return &SettingsContext{}
}

func (sc *SettingsContext) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// Set all settings on the context
	for k, v := range settings {
		context.Set(r, k, v)
	}

	next(rw, r)
}
