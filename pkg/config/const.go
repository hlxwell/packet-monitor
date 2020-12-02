package config

var (
	// Host is the server host
	Host = "http://localhost:8080"
)

const (
	// ConfigGetPath is the packet sender config get url
	ConfigGetPath = "/getNetworkPolicyConfigs"

	// DataPostPath is the result upload url
	DataPostPath = "/postNetworkPolicyResult"

	// UDP enum value
	UDP string = "udp"
	// TCP enum value
	TCP = "tcp"
)
