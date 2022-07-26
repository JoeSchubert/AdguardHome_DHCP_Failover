package dhcp

type ServerSettings struct {
	Address    string
	Port       string
	Username   string
	Password   string
	Base64Auth string
}

type ServerConfig struct {
	Enabled bool `json:"enabled"`
	//InterfaceName string `json:"interface_name"`  // Ignore this field as AdguardHome currently returns "0" for all interface selections
	V4 struct {
		GatewayIp     string `json:"gateway_ip"`
		LeaseDuration int64  `json:"lease_duration"`
		RangeEnd      string `json:"range_end"`
		RangeStart    string `json:"range_start"`
		SubnetMask    string `json:"subnet_mask"`
	} `json:"v4"`
	V6 struct {
		LeaseDuration int64  `json:"lease_duration"`
		RangeStart    string `json:"range_start"`
	} `json:"v6"`
}

type StaticLeaseList struct {
	StaticLeases []StaticLease `json:"static_leases"`
}

type StaticLease struct {
	Hostname string `json:"hostname"`
	IP       string `json:"ip"`
	MAC      string `json:"mac"`
}
