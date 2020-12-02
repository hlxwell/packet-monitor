package cmd

import (
	"log"
	"net"
	"time"

	"github.com/hlxwell/packet-monitor/pkg/config"
	"github.com/hlxwell/packet-monitor/pkg/packet"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var (
	packetSendingFrequency time.Duration
	configUpdateFrequency  time.Duration
	host                   string
	startCmd               = &cobra.Command{
		Use:   "start",
		Short: "Start Packet Mon Service",
		Long: `Start a packet mon service, you can adjust packet sending frequency and config update frequency.
e.g.
packet-mon --configUpdateFrequency 60 --packetSendingFrequency 10`,
		Run: func(cmd *cobra.Command, args []string) {
			// Set global host
			config.Host = host

			// Update Policy Periodically
			go func() {
				for {
					config.UpdateConfig()
					time.Sleep(configUpdateFrequency)
				}
			}()

			// Send packet according to policies
			for {
				time.Sleep(packetSendingFrequency)
				log.Printf("=== Start Send Packet (%d policies)", len(config.ConfigData))

				for _, policy := range config.ConfigData {
					dstIP := net.ParseIP(policy.DstIP)
					if policy.Protocol == config.UDP {
						log.Printf("Sending UDP packet: %v", policy)
						packet.SendUDPPacket(dstIP, policy.DstPort)
					} else if policy.Protocol == config.TCP {
						log.Printf("Sending TCP packet: %v", policy)
						packet.SendTCPPacket(dstIP, policy.DstPort)
					}
				}
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(startCmd)

	// cmd.PersistentFlags().Lookup("foo").Value
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	startCmd.Flags().StringVarP(&host, "host", "o", "http://127.0.0.1", "Server host")
	startCmd.Flags().DurationVarP(&configUpdateFrequency, "configUpdateFrequency", "u", 30*time.Second, "Frequency for loading config from server. unit is second.")
	startCmd.Flags().DurationVarP(&packetSendingFrequency, "packetSendingFrequency", "s", 10*time.Second, "Frequency for sending packet to server. unit is second.")
}
