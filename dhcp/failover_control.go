package dhcp

import (
	"fmt"
	"log"
	"time"
)

var ServerDownTime time.Time

func StartFailoverControl() {
	emptyDhcpConfig := ServerConfig{}
	// Currently this is only retrieving settings from the primary server once. I am debating on having it replicate the settings on a schedule.
	for {
		serverUp := IsAdguardRunning(ActiveServer)
		if serverUp {
			if ActiveServer == Primary {
				if DHCPSettings == emptyDhcpConfig {
					tempSettings, tempLeases, err := GetServerSettings(ActiveServer, DHCPSettings, StaticLeases)
					if err != nil {
						log.Println("Error getting settings: ", err)
					} else {
						log.Println("Primary server up, settings retrieved and stored.")
						DHCPSettings = tempSettings
						StaticLeases = tempLeases
						DHCPSettings.Enabled = true
						// the server has to have dhcp enabled to store the dhcp static leases
						ChangeDHCPSettings(Secondary, DHCPSettings)
						AddStaticLeases(Secondary, StaticLeases)
						DHCPSettings.Enabled = false
						ChangeDHCPSettings(Secondary, DHCPSettings)
						fmt.Println("Replicated DHCP settings and static leases to secondary server.")
						DHCPSettings.Enabled = true
					}
				}
			} else {
				if DHCPSettings == emptyDhcpConfig {
					log.Println("Settings were never retrieved from the primary server, unable to switch.")
				} else {
					log.Println("Checking if primary server is up...")
					serverUp := IsAdguardRunning(Primary)
					if serverUp {
						log.Printf("Primary server up while on secondary server, switching back to primary server. Server down time: %s ", time.Since(ServerDownTime).String())
						ServerDownTime = switchServers()
					} else {
						log.Printf("Primary server is down and has been for: %s", time.Since(ServerDownTime).String())
					}
				}
			}
		} else {
			if DHCPSettings.V4.GatewayIp == emptyDhcpConfig.V4.GatewayIp {
				fmt.Println("Unable to switch to secondary server as settings were unable to be retrieved from primary server.")
			} else {
				log.Println("Primary server is down, switching to secondary server.")
				ServerDownTime = switchServers()
			}
		}
		time.Sleep(SleepDuration)
	}
}

func switchServers() time.Time {
	if ActiveServer == Primary {
		// Switch to secondary server
		ActiveServer = Secondary
		DHCPSettings.Enabled = true
		ChangeDHCPSettings(Secondary, DHCPSettings)
		DHCPSettings.Enabled = false
		return time.Now()
	} else {
		// Switch to primary server
		ActiveServer = Primary
		DHCPSettings.Enabled = true
		ChangeDHCPSettings(Primary, DHCPSettings)
		DHCPSettings.Enabled = false
		ChangeDHCPSettings(Secondary, DHCPSettings)
		return ServerDownTime
	}
}
