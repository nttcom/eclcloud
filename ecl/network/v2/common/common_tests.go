package common

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/testhelper/client"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v2.0/"
	return sc
}
