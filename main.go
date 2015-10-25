package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/slok/monf/configuration"
	"github.com/slok/monf/panel"
	"github.com/spf13/viper"
)

func main() {
	log.Info("Starting Monf...")

	// Configuration load is executed when configuration package is imported
	// (see configuration/load.go:init)
	// We need this because some code executed before main needs the
	// configuration (like template compilation)

	// Bind all the routes to the handlers
	router := panel.BindRoutes(nil)

	// Server stuff creation
	n := negroni.Classic()
	n.UseHandler(router)

	dPort := viper.GetInt(configuration.ListenPort)
	dHost := viper.GetString(configuration.ListenHost)
	listenAddr := fmt.Sprintf("%s:%d", dHost, dPort)

	// Execute monf!
	n.Run(listenAddr)
}
