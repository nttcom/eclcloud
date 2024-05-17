package load_balancer_actions

import "github.com/nttcom/eclcloud/v3"

func rebootURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancers", id, "reboot")
}

func resetPasswordURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("load_balancers", id, "reset_password")
}
