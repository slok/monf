package configuration

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
)

func TestConfigurationUtils(t *testing.T) {
	Convey("Given a fresh configured app", t, func() {
		viper.Reset()
		viper.SetConfigType(configType)
		// Load as normal
		LoadSettings("")

		//--------------------------------------------------- Test ResetSettings
		Convey("And adding new settings", func() {
			// Add some settings
			for i := 0; i < 20; i++ {
				viper.Set(fmt.Sprintf("%d", i), i)
			}
			settingsLengthBefore := len(viper.AllSettings())
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
					So(settingsLengthAfter, ShouldEqual, settingsLengthBefore-20)
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
}
