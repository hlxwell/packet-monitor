package sniffer

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/hlxwell/packet-monitor/pkg/config"
	result "github.com/hlxwell/packet-monitor/pkg/packet"
)

// ListenToPacket use libpcap to monitoring all the packet
//
// go sniffer.ListenToPacket(receiverPolicies)
//
func ListenToPacket(policies []config.Policy) {
	if len(policies) == 0 {
		return
	}

	var handle *pcap.Handle
	handle, err := pcap.OpenLive("any", 65535, false, -1*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	var filters []string
	for _, policy := range policies {
		filters = append(filters, strings.ToLower(policy.Protocol)+" dst port "+policy.DstPort)
	}

	err = handle.SetBPFFilter(strings.Join(filters, " "))
	if err != nil {
		log.Fatal(err)
	}

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		ipLayer := packet.Layer(layers.LayerTypeIPv4) // parse IP header
		ipPacket := ipLayer.(*layers.IPv4)

		if ipPacket.NextLayerType() == layers.LayerTypeTCP {
			dataLayer := packet.Layer(layers.LayerTypeTCP) // parse TCP header
			dataPacket := dataLayer.(*layers.TCP)

			fmt.Println(
				"--- Received TCP Packet: ",
				ipPacket.SrcIP, ipPacket.DstIP,
				dataPacket.SrcPort, dataPacket.DstPort,
				dataPacket.Payload,
			)
			result.UploadResult(ipPacket.DstIP.String(), dataPacket.DstPort.String(), true)
		} else if ipPacket.NextLayerType() == layers.LayerTypeUDP {
			dataLayer := packet.Layer(layers.LayerTypeUDP) // parse TCP header
			dataPacket := dataLayer.(*layers.UDP)

			fmt.Println(
				"--- Received UDP Packet: ",
				ipPacket.SrcIP, ipPacket.DstIP,
				dataPacket.SrcPort, dataPacket.DstPort,
				dataPacket.Payload,
			)
			result.UploadResult(ipPacket.DstIP.String(), dataPacket.DstPort.String(), true)
		}
	}
}
