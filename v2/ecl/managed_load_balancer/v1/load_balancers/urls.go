package load_balancers

import (
	"github.com/nttcom/eclcloud/v2"
)

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("load_balancers")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancers", id)
}

func actionURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancers", id, "action")
}

func stagedURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancers", id, "staged")
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func createURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

func showURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func updateURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createStagedURL(c *eclcloud.ServiceClient, id string) string {
	return stagedURL(c, id)
}

func showStagedURL(c *eclcloud.ServiceClient, id string) string {
	return stagedURL(c, id)
}

func updateStagedURL(c *eclcloud.ServiceClient, id string) string {
	return stagedURL(c, id)
}

func cancelStagedURL(c *eclcloud.ServiceClient, id string) string {
	return stagedURL(c, id)
}
