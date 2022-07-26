package utils

import (
	"log"
	"os"
)

const (
	// DHCPCheckInterval - The number of seconds between checks
	DHCPCheckInterval = "DHCPCheckInterval" // Optional
	// Primary_Server_Address - The IP address or url of the primary server
	PrimaryServerAddress = "PRIMARY_SERVER"
	PrimaryServerPort    = "PRIMARY_SERVER_PORT" // Optional
	PrimaryUserName      = "PRIMARY_USERNAME"    // Optional
	PrimaryPassword      = "PRIMARY_PASSWORD"    // Optional
	// Secondary_Server_Address - The IP address or url of the secondary server
	SecondaryServerAddress = "SECONDARY_SERVER"
	SecondaryServerPort    = "SECONDARY_SERVER_PORT"   // Optional
	SecondaryUserName      = "SECONDARY_USERNAME"      // Optional
	SecondaryPassword      = "SECONDARY_PASSWORD"      // Optional
	EnableStatusWebServer  = "ENABLE_STATUS_WEBSERVER" // Optional
	StatusWebServerPort    = "STATUS_WEBSERVER_PORT"   // Optional
)

// helper function to process environment variables without too much code duplication
func ProcessEnvVars(envVarName string, defaultValue string, required bool) string {
	envVarValue, present := os.LookupEnv(envVarName)
	if present {
		return envVarValue
	} else {
		if required {
			log.Fatalf("%s not set. Unable to continue.", envVarName)
		} else {
			if defaultValue != "" {
				log.Printf("%s not set. Defaulting to %s\n", envVarName, defaultValue)
			} else {
				log.Printf("%s not set. Defaulting to blank\n", envVarName)
			}
			return defaultValue
		}
	}
	return ""
}
