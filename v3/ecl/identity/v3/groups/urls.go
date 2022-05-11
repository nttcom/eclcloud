package groups

import "github.com/nttcom/eclcloud/v2"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

func getURL(client *eclcloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("groups")
}

func updateURL(client *eclcloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}

func deleteURL(client *eclcloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID)
}
