package configuration

// Basic imports
import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	log "github.com/Sirupsen/logrus"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestSettingsLoad(t *testing.T) {
	// test suite globals
	testFilePath := "/tmp/monf_test.yml"

	Convey("Given a unconfigured app", t, func() {

		// Reset settings on each test
		viper.Reset()
		viper.SetConfigType(configType)
		fmt.Println("Hi")

		//---------------------------------------------------- Test LoadDefaults
		Convey("When loading the default settings", func() {
			LoadAppDefaultSettings()

			Convey("Then the default settings should be loaded", func() {
				So(viper.GetBool(Debug), ShouldEqual, defaultAppSettings[Debug])
				So(viper.GetInt(ListenPort), ShouldEqual, defaultAppSettings[ListenPort])
				So(viper.GetString(ListenHost), ShouldEqual, defaultAppSettings[ListenHost])
				So(viper.GetString(TemplatesPath), ShouldEqual, defaultAppSettings[TemplatesPath])
				So(viper.GetString(StaticURL), ShouldEqual, defaultAppSettings[StaticURL])
				So(viper.GetString(StaticPath), ShouldEqual, defaultAppSettings[StaticPath])
			})
		})

		//----------------------------------------- Test LoadDefaults debug true
		Convey("And creating a settings file with the debug setting active", func() {
			testData := []byte(`debug: true`)
			err := ioutil.WriteFile("./settings.yaml", testData, 0644)
			So(err, ShouldBeNil)
			LoadSettings("")

			Convey("The log debug level should be debug", func() {
				So(log.GetLevel(), ShouldEqual, log.DebugLevel)
			})
		})

		//--------------------------------------------- Default File path tests
		Convey("And creating a settings file on default path", func() {
			testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
			err := ioutil.WriteFile("./settings.yaml", testData, 0644)
			So(err, ShouldBeNil)

			//---------------------------------------- Test LoadFilePathDefaults
			Convey("When loading the settings file from a preset default path", func() {
				LoadDefaultFilePathSettings()
				Convey("The file settings should be loaded", func() {
					So(viper.GetInt("test_setting1"), ShouldEqual, 1)
					So(viper.GetString("test_setting2"), ShouldEqual, "batman")
					So(viper.GetBool("test_setting3"), ShouldBeTrue)
				})
			})
			//----------------------------------------------- Test LoadSettings
			Convey("When loading defaults and the settings file from a preset default path", func() {
				LoadSettings("")
				Convey("The file settings and the app default settings should be loaded", func() {
					So(viper.GetInt("test_setting1"), ShouldEqual, 1)
					So(viper.GetString("test_setting2"), ShouldEqual, "batman")
					So(viper.GetBool("test_setting3"), ShouldBeTrue)
					So(viper.GetBool(Debug), ShouldEqual, defaultAppSettings[Debug])
				})
			})
		})
		//---------------------------------------------- Custom file path tests
		Convey("And creating a settings file on a custom path", func() {
			testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
			err := ioutil.WriteFile(testFilePath, testData, 0644)
			So(err, ShouldBeNil)

			//--------------------------------------- Test LoadFromFileSettings
			Convey("When loading the settings file from a custom path", func() {
				LoadFromFileSettings(testFilePath)

				Convey("The file settings should be loaded", func() {
					So(viper.GetInt("test_setting1"), ShouldEqual, 1)
					So(viper.GetString("test_setting2"), ShouldEqual, "batman")
					So(viper.GetBool("test_setting3"), ShouldBeTrue)
				})
			})
			//----------------------------------------------- Test LoadSettings
			Convey("When loading defaults and the settings file from a custom path", func() {
				LoadSettings(testFilePath)
				Convey("The file settings and the app default settings should be loaded", func() {
					So(viper.GetInt("test_setting1"), ShouldEqual, 1)
					So(viper.GetString("test_setting2"), ShouldEqual, "batman")
					So(viper.GetBool("test_setting3"), ShouldBeTrue)
					So(testFilePath, ShouldEqual, SpecificConfigPath)
				})
			})
		})

		//------------------------------------------------------------- Teardown
		Reset(func() {
			os.Remove("./settings.yaml")
			os.Remove(testFilePath)
		})
	})
}
