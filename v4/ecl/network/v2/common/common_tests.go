package common

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
