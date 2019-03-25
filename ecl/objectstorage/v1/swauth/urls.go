package swauth

import "github.com/nttcom/eclcloud"

func getURL(c *eclcloud.ProviderClient) string {
	return c.IdentityBase + "auth/v1.0"
}
