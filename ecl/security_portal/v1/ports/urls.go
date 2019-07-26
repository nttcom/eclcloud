package ports

import (
	"fmt"

	"github.com/nttcom/eclcloud"
)

// func listURL(client *eclcloud.ServiceClient) string {
// 	return client.ServiceURL("API/ScreenEventFGSDeviceGet")
// }

// func getURL(client *eclcloud.ServiceClient) string {
// 	return client.ServiceURL("API/SoEntryFGS")
// }

// func createURL(client *eclcloud.ServiceClient) string {
// 	return client.ServiceURL("API/SoEntryFGS")
// }

// func deleteURL(client *eclcloud.ServiceClient) string {
// 	return client.ServiceURL("API/SoEntryFGS")
// }

func updateURL(client *eclcloud.ServiceClient, deviceType string) string {
	url := fmt.Sprintf("ecl-api/ports/%s/", deviceType)
	return client.ServiceURL(url)
}
