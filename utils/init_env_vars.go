package utils

import (
	"AdGuardDHCPFailover/dhcp"
	"AdGuardDHCPFailover/status"
	"encoding/base64"
	"log"
	"strconv"
	"time"
)

func InitEnvVars() {
	log.Println("AdGuardHome DHCP FailOver initializing...")
	log.Println("Processing environment variables...")
	// Read the environment variables into the config struct
	sleepVal := ProcessEnvVars(DHCPCheckInterval, dhcp.DefaultCheckInterval, false)
	sleepInt, err := strconv.Atoi(sleepVal)
	if err == nil {
		dhcp.SleepDuration = time.Duration(sleepInt) * time.Second
	} else {
		log.Fatal("Unable to convert sleep duration to a usable value\n")
	}
	dhcp.Primary.Address = ProcessEnvVars(PrimaryServerAddress, "", true)
	// default active server to the primary server
	dhcp.Primary.Port = ProcessEnvVars(PrimaryServerPort, "80", false)
	dhcp.Primary.Username = ProcessEnvVars(PrimaryUserName, "", false)
	dhcp.Primary.Password = ProcessEnvVars(PrimaryPassword, "", false)
	if len(dhcp.Primary.Username) > 1 && len(dhcp.Primary.Password) < 1 {
		log.Fatal("If a username is specified, a password must also be specified\n")
	} else {
		dhcp.Primary.Base64Auth = base64.StdEncoding.EncodeToString([]byte(dhcp.Primary.Username + ":" + dhcp.Primary.Password))
	}
	dhcp.ActiveServer = dhcp.Primary

	dhcp.Secondary.Address = ProcessEnvVars(SecondaryServerAddress, "", false)
	dhcp.Secondary.Port = ProcessEnvVars(SecondaryServerPort, "80", false)
	dhcp.Secondary.Username = ProcessEnvVars(SecondaryUserName, "", false)
	dhcp.Secondary.Password = ProcessEnvVars(SecondaryPassword, "", false)
	if len(dhcp.Secondary.Username) > 1 && len(dhcp.Secondary.Password) < 1 {
		log.Fatal("If a username is specified, a password must also be specified\n")
	} else {
		dhcp.Secondary.Base64Auth = base64.StdEncoding.EncodeToString([]byte(dhcp.Secondary.Username + ":" + dhcp.Secondary.Password))
	}
	status.EnableStatusWebServer, _ = strconv.ParseBool(ProcessEnvVars(EnableStatusWebServer, "true", false))
	status.StatusWebServerPort = ProcessEnvVars(StatusWebServerPort, "5050", false)
	log.Println("Environment variables processed.")
}
