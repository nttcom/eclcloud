package accounts

import "github.com/nttcom/eclcloud"

func getURL(c *eclcloud.ServiceClient) string {
	return c.Endpoint
}

func updateURL(c *eclcloud.ServiceClient) string {
	return getURL(c)
}
