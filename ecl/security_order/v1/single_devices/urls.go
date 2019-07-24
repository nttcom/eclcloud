package single_devices

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/ScreenEventFGSDeviceGet")
}

func getURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGS")
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGS")
}

func deleteURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGS")
}

func updateURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("API/SoEntryFGS")
}
