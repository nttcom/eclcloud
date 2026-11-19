package network_based_device_ha

import (
	"github.com/nttcom/eclcloud/v3"
)

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/ScreenEventFGHADeviceGet")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/SoEntryFGHA")
}

func deleteURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/SoEntryFGHA")
}

func updateURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/SoEntryFGHA")
}
