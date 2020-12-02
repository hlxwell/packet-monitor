package packet

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/hlxwell/packet-monitor/pkg/config"
)

// UploadResult can upload connectivity test result
func UploadResult(dstIP string, port string, success bool) {
	id := fmt.Sprintf("tcp-%s-%s", dstIP, port)
	policy := config.ConfigData[id]
	postURL := fmt.Sprintf("%s%s", config.Host, config.DataPostPath)
	resp, err := http.PostForm(
		postURL,
		url.Values{
			"index":  {strconv.Itoa(policy.Index)},
			"result": {strconv.FormatBool(success)},
		},
	)

	if err != nil {
		log.Println("Upload data error: ", err)
	}

	if resp.StatusCode >= 300 {
		log.Println("Upload data status code is not 20x but:", resp.StatusCode)
	}

	log.Println("Upload data success: ", policy.Index, policy.DstIP, policy.DstPort, success, postURL)
}
