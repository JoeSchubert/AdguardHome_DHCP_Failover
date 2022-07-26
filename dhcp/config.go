package dhcp

import (
	"time"
)

var SleepDuration time.Duration
var ActiveServer ServerSettings
var Primary ServerSettings
var Secondary ServerSettings
var DHCPSettings ServerConfig
var StaticLeases StaticLeaseList

var DefaultCheckInterval = "60"
