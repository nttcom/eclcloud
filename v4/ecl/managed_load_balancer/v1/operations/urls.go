package operations

import (
	"github.com/nttcom/eclcloud/v4"
)

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("operations")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("operations", id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func showURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}
