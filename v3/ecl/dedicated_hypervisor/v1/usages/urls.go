package usages

import "github.com/nttcom/eclcloud/v2"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("usages")
}

func getHistoriesURL(client *eclcloud.ServiceClient, usageID string) string {
	return client.ServiceURL("usages", usageID, "histories")
}
