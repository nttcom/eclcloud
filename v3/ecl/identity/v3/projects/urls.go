package projects

import "github.com/nttcom/eclcloud/v3"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

func getURL(client *eclcloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("projects")
}

func deleteURL(client *eclcloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}

func updateURL(client *eclcloud.ServiceClient, projectID string) string {
	return client.ServiceURL("projects", projectID)
}
