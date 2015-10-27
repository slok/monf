package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/slok/monf/configuration"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type SettingsContextTestSuite struct {
	suite.Suite
}

func (suite *SettingsContextTestSuite) SetupTest() {
	// Reset settings on each test
	configuration.ResetSettings()
}

func (suite *SettingsContextTestSuite) TestSettingsContext() {
	// Create a request recorder
	response := httptest.NewRecorder()

	// Set the negroni server with our middleware
	n := negroni.New()
	n.Use(NewSettingsContext())

	// Create the request
	req, err := http.NewRequest("GET", "http://localhost:3000/", nil)
	suite.Nil(err)

	// Check no context before running the middleware
	suite.Equal(0, len(context.GetAll(req)))

	// Make call and check context
	n.ServeHTTP(response, req)
	requestContextAfter := context.GetAll(req)
	settings := viper.AllSettings()

	suite.Equal(len(settings), len(requestContextAfter))
	for k := range settings {
		suite.Equal(settings[k], requestContextAfter[k])
	}
}

func TestSettingsContextTestSuite(t *testing.T) {
	suite.Run(t, new(SettingsContextTestSuite))
}
