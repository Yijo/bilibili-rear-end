package viper

import (
	"testing"
	"fmt"
)

// Test Init viper config.
func TestInitViperConfig(t *testing.T) {
	if err := InitViperConfig("app"); err != nil {
		fmt.Println("Init viper config failure, error is: ", err)
		return
	}


	fmt.Println("Init viper config success.")
}

// Test return viper all settings.
func TestAllViperSettings(t *testing.T) {

	TestInitViperConfig(t)

	allSettings := AllViperSettings()

	mysqlC := allSettings["mysql"].(map[string] interface{})

	data := mysqlC["bi_member"]
	fmt.Println("data is: ", data)

	fmt.Println("Viper all settings is: ", allSettings)
}

// Test init config.
func TestInitConfig(t *testing.T) {
	c := Config{
		Name: "app",
	}
	err := c.initConfig()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	fmt.Println("Init config success.")
}

// Test hot update.
func TestWatchConfig(t *testing.T) {
	c := Config{
		Name: "app",
	}
	err := c.initConfig()
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	// hot update operation

}
