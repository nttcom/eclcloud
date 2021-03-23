package testing

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/testhelper"
)

func createClient() *eclcloud.ServiceClient {
	return &eclcloud.ServiceClient{
		ProviderClient: &eclcloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
