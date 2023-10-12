package volumeactions

import "github.com/nttcom/eclcloud/v4"

func actionURL(c *eclcloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
