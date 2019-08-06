package network_based_device_single

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

func getURLPartFromDeviceType(deviceType string) string {
	if deviceType == "WAF" {
		return "WAF"
	}
	return "S"
}

func listURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/ScreenEventFG%sDeviceGet", part)
	return client.ServiceURL(url)
}

// func getURL(client *eclcloud.ServiceClient) string {
// 	return client.ServiceURL("API/SoEntryFGS")
// 	// return client.ServiceURL("API/SoEntryFGWAF")
// }

func createURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/SoEntryFG%s", part)
	return client.ServiceURL(url)
	// return client.ServiceURL("API/SoEntryFGWAF")
}

func deleteURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/SoEntryFG%s", part)
	return client.ServiceURL(url)
	// return client.ServiceURL("API/SoEntryFGWAF")
}

func updateURL(client *eclcloud.ServiceClient, deviceType string) string {
	part := getURLPartFromDeviceType(deviceType)
	url := fmt.Sprintf("API/SoEntryFG%s", part)
	return client.ServiceURL(url)
	// return client.ServiceURL("API/SoEntryFGWAF")
}
