package testing

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/testhelper"
)

func createClient() *eclcloud.ServiceClient {
	return &eclcloud.ServiceClient{
		ProviderClient: &eclcloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
