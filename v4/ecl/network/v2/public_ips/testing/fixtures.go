package testing

import (
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/public_ips"
)

const ListResponse = `
{
	"public_ips": [
	  {
		"cidr": "100.127.255.80",
		"description": "",
		"id": "0718a31b-67be-4349-946b-61a0fc38e4cd",
		"internet_gw_id": "2a75cfa6-89af-425b-bce5-2a85197ef04f",
		"name": "seinou-test-public",
		"status": "PENDING_CREATE",
		"submask_length": 29,
		"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
	  },
	  {
		"cidr": "100.127.254.56",
		"description": "",
		"id": "110846c3-3a20-42ff-ad3d-25ba7b0272bb",
		"internet_gw_id": "05db9b0e-65ed-4478-a6b3-d3fc259c8d07",
		"name": "6_Public",
		"status": "ACTIVE",
		"submask_length": 29,
		"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
	  }
	]
}
`

const GetResponse = `
{
	"public_ip": {
	  "cidr": "100.127.255.80",
	  "description": "",
	  "id": "0718a31b-67be-4349-946b-61a0fc38e4cd",
	  "internet_gw_id": "2a75cfa6-89af-425b-bce5-2a85197ef04f",
	  "name": "seinou-test-public",
	  "status": "PENDING_CREATE",
	  "submask_length": 29,
	  "tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
	}
}
`

const CreateRequest = `
{
	"public_ip": {
	  "internet_gw_id": "2a75cfa6-89af-425b-bce5-2a85197ef04f",
	  "name": "seinou-test-public",
	  "submask_length": 29,
	  "tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
	}
}
`

const CreateResponse = `
{
	"public_ip": {
	  "cidr": "100.127.255.80",
	  "description": "",
	  "id": "0718a31b-67be-4349-946b-61a0fc38e4cd",
	  "internet_gw_id": "2a75cfa6-89af-425b-bce5-2a85197ef04f",
	  "name": "seinou-test-public",
	  "status": "PENDING_CREATE",
	  "submask_length": 29,
	  "tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
	}
}
`

const UpdateRequest = `
{
	"public_ip": {
		"name": "seinou-test-public",
		"description": ""
	}
}
 `

const UpdateResponse = `
{
	"public_ip": {
	  "cidr": "100.127.255.80",
	  "description": "",
	  "id": "0718a31b-67be-4349-946b-61a0fc38e4cd",
	  "internet_gw_id": "2a75cfa6-89af-425b-bce5-2a85197ef04f",
	  "name": "seinou-test-public",
	  "status": "PENDING_UPDATE",
	  "submask_length": 29,
	  "tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
	}
}
`

var PublicIP1 = public_ips.PublicIP{
	Cidr:          "100.127.255.80",
	Description:   "",
	ID:            "0718a31b-67be-4349-946b-61a0fc38e4cd",
	InternetGwID:  "2a75cfa6-89af-425b-bce5-2a85197ef04f",
	Name:          "seinou-test-public",
	Status:        "PENDING_CREATE",
	SubmaskLength: 29,
	TenantID:      "19ab165c7a664abe9c217334cd0e9cc9",
}

var PublicIP2 = public_ips.PublicIP{
	Cidr:          "100.127.254.56",
	Description:   "",
	ID:            "110846c3-3a20-42ff-ad3d-25ba7b0272bb",
	InternetGwID:  "05db9b0e-65ed-4478-a6b3-d3fc259c8d07",
	Name:          "6_Public",
	Status:        "ACTIVE",
	SubmaskLength: 29,
	TenantID:      "19ab165c7a664abe9c217334cd0e9cc9",
}

var ExpectedPublicIPSlice = []public_ips.PublicIP{PublicIP1, PublicIP2}
