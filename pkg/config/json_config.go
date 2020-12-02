package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hlxwell/packet-monitor/pkg/utils"
)

// Policy struct of each policy
type Policy struct {
	Index    int    `json:"index"`
	SrcIP    string `json:"srcIp"`
	DstIP    string `json:"dstIp"`
	DstPort  string `json:"dstPort"`
	Protocol string `json:"protocol"`
}

// ConfigData is a global accessable config
var ConfigData = map[string]Policy{}

// UpdateConfig download and parse json config and store it to global var.
func UpdateConfig() {
	url := fmt.Sprintf("%s%s?ip=%s", Host, ConfigGetPath, utils.GetLocalIP().String())
	log.Println("=== Get config from: ", url)
	resp, err := http.Get(url)

	if err != nil {
		log.Println("Upload data error: ", err)
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	policies := []Policy{}
	json.Unmarshal(result, &policies)
	for _, policy := range policies {
		id := fmt.Sprintf("tcp-%s-%s", policy.DstIP, policy.DstPort)
		ConfigData[id] = policy
	}
}
