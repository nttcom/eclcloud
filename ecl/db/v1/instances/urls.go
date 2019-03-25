package instances

import "github.com/nttcom/eclcloud"

func baseURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("instances")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("instances", id)
}

func userRootURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "root")
}

func actionURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("instances", id, "action")
}
