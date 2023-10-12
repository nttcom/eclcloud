package workspaces

import "github.com/nttcom/eclcloud/v4"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("workspaces")
}

func getURL(client *eclcloud.ServiceClient, workspaceID string) string {
	return client.ServiceURL("workspaces", workspaceID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("workspaces")
}

func deleteURL(client *eclcloud.ServiceClient, workspaceID string) string {
	return client.ServiceURL("workspaces", workspaceID)
}

func updateURL(client *eclcloud.ServiceClient, workspaceID string) string {
	return client.ServiceURL("workspaces", workspaceID)
}
