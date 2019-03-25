package datastores

import "github.com/nttcom/eclcloud"

func baseURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("datastores")
}

func resourceURL(c *eclcloud.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID)
}

func versionsURL(c *eclcloud.ServiceClient, dsID string) string {
	return c.ServiceURL("datastores", dsID, "versions")
}

func versionURL(c *eclcloud.ServiceClient, dsID, versionID string) string {
	return c.ServiceURL("datastores", dsID, "versions", versionID)
}
