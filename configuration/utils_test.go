package configuration

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

type ConfigurationUtilsTestSuite struct {
	suite.Suite
}

func (suite *ConfigurationUtilsTestSuite) SetupTest() {
	viper.Reset()
	viper.SetConfigType(configType)
}

func (suite *ConfigurationUtilsTestSuite) TestResetSettings() {
	// Load as normal
	LoadSettings("")

	// Add some settings and check
	for i := 0; i < 20; i++ {
		viper.Set(fmt.Sprintf("%d", i), i)
	}
	settingsLengthBefore := len(viper.AllSettings())
	for i := 0; i < 20; i++ {
		suite.Equal(i, viper.GetInt(fmt.Sprintf("%d", i)))
	}
	fmt.Printf("%v", viper.AllSettings())
	// Reset and check
	ResetSettings()
	settingsLengthAfter := len(viper.AllSettings())

	suite.Equal(settingsLengthBefore-20, settingsLengthAfter)
}

func (suite *ConfigurationUtilsTestSuite) TestOverrideSettings() {
	// Load as normal and add special vars
	LoadSettings("")
	viper.Set("special", true)

	// check status
	suite.True(viper.GetBool("special"))
	settingsLengthBefore := len(viper.AllSettings())

	// Override and check
	OverrideSettings(map[string]interface{}{
		"special":  false,
		"special1": 1,
		"special2": "test",
	})

	settingsLengthAfter := len(viper.AllSettings())
	suite.Equal(settingsLengthBefore+2, settingsLengthAfter)
	suite.False(viper.GetBool("special"))
	suite.Equal(1, viper.GetInt("special1"))
	suite.Equal("test", viper.GetString("special2"))
}

func TestConfigurationUtilsTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigurationUtilsTestSuite))
}
