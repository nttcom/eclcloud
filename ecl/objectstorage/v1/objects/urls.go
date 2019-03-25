package objects

import (
	"github.com/nttcom/eclcloud"
)

func listURL(c *eclcloud.ServiceClient, container string) string {
	return c.ServiceURL(container)
}

func copyURL(c *eclcloud.ServiceClient, container, object string) string {
	return c.ServiceURL(container, object)
}

func createURL(c *eclcloud.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func getURL(c *eclcloud.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func deleteURL(c *eclcloud.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func downloadURL(c *eclcloud.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}

func updateURL(c *eclcloud.ServiceClient, container, object string) string {
	return copyURL(c, container, object)
}
