package containers

import "github.com/nttcom/eclcloud"

func listURL(c *eclcloud.ServiceClient) string {
	return c.Endpoint
}

func createURL(c *eclcloud.ServiceClient, container string) string {
	return c.ServiceURL(container)
}

func getURL(c *eclcloud.ServiceClient, container string) string {
	return createURL(c, container)
}

func deleteURL(c *eclcloud.ServiceClient, container string) string {
	return createURL(c, container)
}

func updateURL(c *eclcloud.ServiceClient, container string) string {
	return createURL(c, container)
}
