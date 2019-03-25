package configurations

import "github.com/nttcom/eclcloud"

func baseURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("configurations")
}

func resourceURL(c *eclcloud.ServiceClient, configID string) string {
	return c.ServiceURL("configurations", configID)
}

func instancesURL(c *eclcloud.ServiceClient, configID string) string {
	return c.ServiceURL("configurations", configID, "instances")
}

func listDSParamsURL(c *eclcloud.ServiceClient, datastoreID, versionID string) string {
	return c.ServiceURL("datastores", datastoreID, "versions", versionID, "parameters")
}

func getDSParamURL(c *eclcloud.ServiceClient, datastoreID, versionID, paramID string) string {
	return c.ServiceURL("datastores", datastoreID, "versions", versionID, "parameters", paramID)
}

func listGlobalParamsURL(c *eclcloud.ServiceClient, versionID string) string {
	return c.ServiceURL("datastores", "versions", versionID, "parameters")
}

func getGlobalParamURL(c *eclcloud.ServiceClient, versionID, paramID string) string {
	return c.ServiceURL("datastores", "versions", versionID, "parameters", paramID)
}
