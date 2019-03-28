package internet_services

import "github.com/nttcom/eclcloud"

func resourceURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("internet_services", id)
}

func rootURL(c *eclcloud.ServiceClient) string {
	return c.ServiceURL("internet_services")
}

func getURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func listURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}

/*func createURL(c *eclcloud.ServiceClient) string {
	return rootURL(c)
}*/

/*func updateURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}*/

/*func deleteURL(c *eclcloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}*/