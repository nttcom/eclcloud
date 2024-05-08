package testing

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/testhelper"
)

func createClient() *eclcloud.ServiceClient {
	return &eclcloud.ServiceClient{
		ProviderClient: &eclcloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
