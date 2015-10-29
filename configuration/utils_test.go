package configuration

import (
	"fmt"
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestConfigurationUtils(t *testing.T) {
	Convey("Given a not configured app", t, func() {
		viper.Reset()
		viper.SetConfigType(configType)

		Convey("And loading default settings", func() {
			// Load as normal
			LoadSettings("")

			//--------------------------------------------------- Test ResetSettings
			Convey("And adding new settings", func() {
				// Add some settings
				settingsLengthBefore := len(viper.AllSettings())
				for i := 0; i < 20; i++ {
					viper.Set(fmt.Sprintf("%d", i), i)
				}
				Convey("There should be new settings", func() {
					for i := 0; i < 20; i++ {
						So(viper.GetInt(fmt.Sprintf("%d", i)), ShouldEqual, i)
					}
				})
				Convey("When using the settings reset util", func() {
					// Reset and check
					ResetSettings()
					settingsLengthAfter := len(viper.AllSettings())
					Convey("Then the context should be the same as the begging settings", func() {
						So(settingsLengthAfter, ShouldEqual, settingsLengthBefore)
					})
				})
			})

			//------------------------------------------------ Test OverrideSettings
			Convey("And setting some specific setting", func() {
				viper.Set("special", true)
				settingsLengthBefore := len(viper.AllSettings())

				Convey("the setted settings should be available", func() {
					So(viper.GetBool("special"), ShouldBeTrue)
				})

				Convey("When using the settings override util", func() {
					// Override and check
					OverrideSettings(map[string]interface{}{
						"special":  false,
						"special1": 1,
						"special2": "test",
					})
					settingsLengthAfter := len(viper.AllSettings())

					Convey("Then the override values should override or add the old settings", func() {
						So(settingsLengthAfter, ShouldEqual, settingsLengthBefore+2)
						So(viper.GetBool("special"), ShouldBeFalse)
						So(viper.GetInt("special1"), ShouldEqual, 1)
						So(viper.GetString("special2"), ShouldEqual, "test")
					})

				})
			})
		})

		Convey("And loading file settings", func() {
			testFilePath := "/tmp/monf_test.yml"
			testData := []byte(`test_setting1: 1
test_setting2: batman
test_setting3: true`)
			err := ioutil.WriteFile(testFilePath, testData, 0644)
			So(err, ShouldBeNil)
			LoadSettings(testFilePath)
			So(viper.GetInt("test_setting1"), ShouldEqual, 1)
			So(viper.GetString("test_setting2"), ShouldEqual, "batman")
			So(viper.GetBool("test_setting3"), ShouldBeTrue)

			Convey("And adding new settings", func() {
				// Add some settings
				settingsLengthBefore := len(viper.AllSettings())
				for i := 0; i < 20; i++ {
					viper.Set(fmt.Sprintf("%d", i), i)
				}
				Convey("There should be new settings", func() {
					for i := 0; i < 20; i++ {
						So(viper.GetInt(fmt.Sprintf("%d", i)), ShouldEqual, i)
					}
				})
				Convey("When using the settings reset util", func() {
					// Reset and check
					ResetSettings()
					settingsLengthAfter := len(viper.AllSettings())
					Convey("Then the context should be the same as the beggining settings including the ones loaded from the config file", func() {
						So(settingsLengthAfter, ShouldEqual, settingsLengthBefore)
						So(viper.GetInt("test_setting1"), ShouldEqual, 1)
						So(viper.GetString("test_setting2"), ShouldEqual, "batman")
						So(viper.GetBool("test_setting3"), ShouldBeTrue)
					})
				})
			})
		})
	})
}
