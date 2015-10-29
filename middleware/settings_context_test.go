package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/slok/monf/configuration"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestSettingsContextMiddleware(t *testing.T) {
	Convey("Given a server with settings context middleware", t, func() {
		// Reset settings on each test
		configuration.ResetSettings()

		// Set the negroni server with our middleware
		response := httptest.NewRecorder()
		n := negroni.New()
		n.Use(NewSettingsContext())
		req, _ := http.NewRequest("GET", "http://localhost:3000/", nil)

		// Check no context before running the middleware
		Convey("The starting context should be empty", func() {
			So(len(context.GetAll(req)), ShouldEqual, 0)
		})

		Convey("When making a request", func() {

			n.ServeHTTP(response, req)
			requestContextAfter := context.GetAll(req)
			settings := viper.AllSettings()

			Convey("The context should be the same as the app settings", func() {
				So(len(requestContextAfter), ShouldEqual, len(requestContextAfter))
				for k := range settings {
					So(requestContextAfter[k], ShouldEqual, settings[k])
				}
			})
		})

		Reset(func() {
		})
	})
}
