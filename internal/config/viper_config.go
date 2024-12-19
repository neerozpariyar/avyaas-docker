package config

import (
	"log"

	"github.com/spf13/viper"
)

/*
ConfigureViper initializes Viper configuration by setting the path to the configuration file and
attempting to read its content. If the read operation encounters an error, the function prints an
error message and exits with a fatal log entry.
*/
func ConfigureViper() {
	println("[+] Processing: Initializing Viper Configuration [+]")

	// Specify the path to the configuration file
	viper.SetConfigFile("./config.json")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		println("[-] Error: Failed to read config file [-]")
		log.Fatal(err.Error())
	}
}
