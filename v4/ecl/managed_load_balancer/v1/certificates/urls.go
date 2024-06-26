package certificates

import (
	"github.com/nttcom/eclcloud/v4"
)

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("certificates")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("certificates", id)
}

func filesURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("certificates", id, "files")
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

func uploadFileURL(c *eclcloud.ServiceClient, id string) string {
	return filesURL(c, id)
}
