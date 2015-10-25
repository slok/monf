package configuration

// Basic imports
import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SettingsTestSuite struct {
	suite.Suite
	testFilePath string
}

func (suite *SettingsTestSuite) SetupTest() {
	suite.testFilePath = "/tmp/monf_test.yml"

	// Reset settings on each test
	viper.Reset()
	viper.SetConfigType(configType)
}
func (suite *SettingsTestSuite) TearDownSuite() {
	os.Remove("./settings.yaml")
	os.Remove(suite.testFilePath)
}

func (suite *SettingsTestSuite) TestLoadDefaults() {
	// Load default settings
	LoadAppDefaultSettings()

	// Check default settings are loaded correctly
	suite.Equal(defaultAppSettings[Debug], viper.GetBool(Debug))
	suite.Equal(defaultAppSettings[ListenPort], viper.GetInt(ListenPort))
	suite.Equal(defaultAppSettings[ListenHost], viper.GetString(ListenHost))
	suite.Equal(defaultAppSettings[TemplatesPath], viper.GetString(TemplatesPath))
	suite.Equal(defaultAppSettings[StaticURL], viper.GetString(StaticURL))
	suite.Equal(defaultAppSettings[StaticPath], viper.GetString(StaticPath))

}

func (suite *SettingsTestSuite) TestLoadFileDefaults() {
	// Create a file with custom settings
	testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
	err := ioutil.WriteFile("./settings.yaml", testData, 0644)
	suite.Nil(err)

	LoadDefaultFilePathSettings()

	// Check default settings are loaded correctly
	suite.Equal(1, viper.GetInt("test_setting1"))
	suite.Equal("batman", viper.GetString("test_setting2"))
	suite.True(viper.GetBool("test_setting3"))

}

func (suite *SettingsTestSuite) TestLoadFromFile() {
	// Create a file with custom settings
	testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
	err := ioutil.WriteFile(suite.testFilePath, testData, 0644)
	suite.Nil(err)

	// Load default settings
	LoadFromFileSettings(suite.testFilePath)

	// Check default settings are loaded correctly
	suite.Equal(1, viper.GetInt("test_setting1"))
	suite.Equal("batman", viper.GetString("test_setting2"))
	suite.True(viper.GetBool("test_setting3"))

}

func (suite *SettingsTestSuite) TestLoadSettingsWithPath() {
	// Create a file with custom settings
	testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
	err := ioutil.WriteFile(suite.testFilePath, testData, 0644)
	suite.Nil(err)

	LoadSettings(suite.testFilePath)

	// Check default settings are loaded correctly
	suite.Equal(1, viper.GetInt("test_setting1"))
	suite.Equal("batman", viper.GetString("test_setting2"))
	suite.True(viper.GetBool("test_setting3"))
	suite.Equal(defaultAppSettings[Debug], viper.GetBool(Debug))
}

func (suite *SettingsTestSuite) TestLoadSettingsWithoutPath() {
	// Create a file with custom settings
	testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
	err := ioutil.WriteFile("./settings.yaml", testData, 0644)
	suite.Nil(err)

	LoadSettings("")

	// Check default settings are loaded correctly
	suite.Equal(1, viper.GetInt("test_setting1"))
	suite.Equal("batman", viper.GetString("test_setting2"))
	suite.True(viper.GetBool("test_setting3"))
	suite.Equal(defaultAppSettings[Debug], viper.GetBool(Debug))
}

func (suite *SettingsTestSuite) TestLoadSettingsDebugTrue() {
	// Create a file with custom settings
	testData := []byte(`debug: true`)
	err := ioutil.WriteFile("./settings.yaml", testData, 0644)
	suite.Nil(err)

	LoadSettings("")
	suite.True(viper.GetBool(Debug))
}

func TestSettingsTestSuite(t *testing.T) {
	suite.Run(t, new(SettingsTestSuite))
}
