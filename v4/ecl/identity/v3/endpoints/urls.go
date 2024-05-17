package endpoints

import "github.com/nttcom/eclcloud/v4"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("endpoints")
}

func endpointURL(client *eclcloud.ServiceClient, endpointID string) string {
	return client.ServiceURL("endpoints", endpointID)
}
