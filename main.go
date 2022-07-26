package main

import (
	"AdGuardDHCPFailover/dhcp"
	"AdGuardDHCPFailover/status"
	"AdGuardDHCPFailover/utils"
	"log"
	"time"
)

func init() {
	utils.InitEnvVars()
}

func main() {
	log.Println("AdGuardHome DHCP FailOver starting...")
	go dhcp.StartFailoverControl()
	if status.EnableStatusWebServer {
		go status.StartWebServer()
	}
	for {
		time.Sleep(time.Second)
	}

}
