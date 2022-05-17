package workspaces

import "github.com/nttcom/eclcloud/v3"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("workspaces")
}

func getURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("workspaces", tenantID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("workspaces")
}

func deleteURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("workspaces", tenantID)
}

func updateURL(client *eclcloud.ServiceClient, tenantID string) string {
	return client.ServiceURL("workspaces", tenantID)
}
