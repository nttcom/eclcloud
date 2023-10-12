package license_types

import "github.com/nttcom/eclcloud/v4"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("license_types")
}
