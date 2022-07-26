package status

import (
	"AdGuardDHCPFailover/dhcp"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

var EnableStatusWebServer bool = true
var StatusWebServerPort string = "5050"

func StartWebServer() {
	log.Println("AdGuardHome DHCP FailOver Status WebServer initializing... on port " + StatusWebServerPort)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if dhcp.ActiveServer == dhcp.Primary {
			fmt.Fprintf(w, "Active Server: Primary (%s)\n", html.EscapeString(dhcp.Primary.Address))
		} else if dhcp.ActiveServer == dhcp.Secondary {
			fmt.Fprintf(w, "Active Server: Secondary (%s). Primary has been down for: %s\n", html.EscapeString(dhcp.Secondary.Address), html.EscapeString(time.Since(dhcp.ServerDownTime).String()))
		}
	})

	log.Printf(http.ListenAndServe(":"+StatusWebServerPort, nil).Error())
}
