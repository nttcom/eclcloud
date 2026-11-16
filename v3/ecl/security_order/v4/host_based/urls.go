package host_based

import (
	"github.com/nttcom/eclcloud/v3"
)

func getURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/ScreenEventHBSOrderInfoGet")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/SoEntryHBS")
}

func deleteURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/SoEntryHBS")
}

func updateURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/v4/SoEntryHBS")
}
