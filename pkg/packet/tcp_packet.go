package packet

import (
	"log"
	"net"
	"time"
)

// SendTCPPacket will send a tcp packet
// packet.SendTCPPacket(net.IP{10, 128, 117, 157}, '8888')
func SendTCPPacket(dstIP net.IP, port string) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(dstIP.String(), port), 1*time.Second)
	if err != nil {
		// if strings.Contains(err.Error(), "refused") {
		// 	// success
		// 	log.Printf("TCP check success %s:%s", dstIP, port)
		// 	UploadResult(dstIP.String(), port, true)
		// } else if strings.Contains(err.Error(), "timeout") {
		// 	// failed
		// 	log.Printf("TCP check failed %s:%s", dstIP, port)
		// 	UploadResult(dstIP.String(), port, false)
		// }

		log.Printf("TCP check failed %s:%s with error: %s", dstIP, port, err.Error())
		UploadResult(dstIP.String(), port, false)
		return
	}
	defer conn.Close()

	log.Printf("TCP check Success %s:%s", dstIP, port)
	UploadResult(dstIP.String(), port, true)
}
