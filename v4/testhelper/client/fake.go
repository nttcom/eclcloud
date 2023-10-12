package client

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/testhelper"
)

// Fake token to use.
const TokenID = "cbc36478b0bd8e67e89469c7749d4127"

// ServiceClient returns a generic service client for use in tests.
func ServiceClient() *eclcloud.ServiceClient {
	return &eclcloud.ServiceClient{
		ProviderClient: &eclcloud.ProviderClient{TokenID: TokenID},
		Endpoint:       testhelper.Endpoint(),
	}
}
