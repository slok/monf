package main

import (
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/meatballhat/negroni-logrus"
	"github.com/slok/monf/configuration"
	"github.com/slok/monf/panel"
	"github.com/spf13/viper"
)

func createServer() *negroni.Negroni {
	n := negroni.New()

	// Recovery stuff for the server
	n.Use(negroni.NewRecovery())

	// logger for the server
	n.Use(negronilogrus.NewMiddleware())

	// static file handler for the server
	static := &negroni.Static{
		Dir:    http.Dir(viper.GetString(configuration.StaticPath)),
		Prefix: viper.GetString(configuration.StaticURL),
	}
	n.Use(static)

	// Our server is finished
	return n
}

func main() {
	log.Info("Starting Monf...")

	// Configuration load is executed when configuration package is imported
	// (see configuration/load.go:init)
	// We need this because some code executed before main needs the
	// configuration (like template compilation)

	// Bind all the routes to the handlers
	router := panel.BindRoutes(nil)

	// Server stuff creation
	n := createServer()
	n.UseHandler(router)

	dPort := viper.GetInt(configuration.ListenPort)
	dHost := viper.GetString(configuration.ListenHost)
	listenAddr := fmt.Sprintf("%s:%d", dHost, dPort)

	// Execute monf!
	n.Run(listenAddr)
}
