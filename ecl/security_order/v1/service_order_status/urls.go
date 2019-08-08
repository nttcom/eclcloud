package service_order_status

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

func getURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := "S"
	if deviceType == "WAF" {
		part = "WAF"
	}
	url := fmt.Sprintf("API/ScreenEventFG%sOrderProgressRate", part)
	return client.ServiceURL(url)
}
