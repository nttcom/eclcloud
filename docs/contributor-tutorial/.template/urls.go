package RESOURCE

import "github.com/nttcom/eclcloud"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("resource")
}

func getURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("resource")
}

func deleteURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}

func updateURL(client *eclcloud.ServiceClient, id string) string {
	return client.ServiceURL("resource", id)
}
