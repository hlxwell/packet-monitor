package packet

import (
	"net"
	"strconv"
)

// SendUDPPacket Send normal UDP Packet
// packet.SendUDPPacket(net.IP{10, 128, 117, 157}, '8888')
func SendUDPPacket(dstIP net.IP, port string) {
	portInt, _ := strconv.Atoi(port)
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: dstIP, Port: portInt}
	conn, _ := net.DialUDP("udp", srcAddr, dstAddr)
	defer conn.Close()
	conn.Write([]byte("hlxwell"))
}
