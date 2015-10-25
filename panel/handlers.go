package panel

import (
	"net/http"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
	"github.com/slok/monf/configuration"
	"github.com/spf13/viper"
)

var (
	templatesPrefix string

	// Precompiled templates are faster
	tplRoot *pongo2.Template
)

func init() {
	// Load varialbes with the settings loaded
	templatesPrefix = filepath.Join(viper.GetString(configuration.TemplatesPath), "panel")
	tplRoot = pongo2.Must(pongo2.FromFile(filepath.Join(templatesPrefix, "root.html")))
}

func logHandler(r *http.Request, handler string, params map[string]string, queryParams map[string][]string) {
	log.WithFields(log.Fields{
		"url":     r.RequestURI,
		"handler": handler,
		"method":  r.Method,
		"params":  params,
		"Query":   queryParams,
	}).Debug("Calling handler")
}

// RootHandler processes the logic of the root panel
func RootHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// Headers
	w.WriteHeader(http.StatusOK) // 200

	// body
	err := tplRoot.ExecuteWriter(nil, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	logHandler(r, "RootHandler", params, nil)
}
