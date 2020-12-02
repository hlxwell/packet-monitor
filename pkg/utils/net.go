package utils

import (
	"net"
	"os"
)

func GetLocalIP() net.IP {
	hostname, _ := os.Hostname()
	addrs, _ := net.LookupIP(hostname)

	return addrs[0]
}
