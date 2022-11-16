/*
Generated by https://github.com/tamac-io/openapi-to-eclcloud-rb
*/
package target_groups

import (
	"github.com/nttcom/eclcloud/v2"
)

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("target_groups")
}

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("target_groups", id)
}

func stagedURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("target_groups", id, "staged")
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
