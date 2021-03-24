package testing

import (
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/networks"
)

const ListResponse = `
{
	"networks": [
	  {
		"admin_state_up": true,
		"description": "",
		"id": "8f36b88a-443f-4d97-9751-34d34af9e782",
		"name": "",
		"plane": "data",
		"shared": false,
		"status": "ACTIVE",
		"subnets": [
		  "ab49eb24-667f-4a4e-9421-b4d915bff416",
		  "f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678"
		],
		"tags": {},
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	  },
	  {
		"admin_state_up": true,
		"description": "Example Network 2",
		"id": "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2",
		"name": "Example Network 2",
		"plane": "data",
		"shared": false,
		"status": "ACTIVE",
		"subnets": [],
		"tags": {
		  "keyword1": "value1",
		  "keyword2": "value2"
		},
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	  }
	]
  }`
const GetResponse = `{
	"network": {
	  "admin_state_up": true,
	  "description": "Example Network 2",
	  "id": "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2",
	  "name": "Example Network 2",
	  "plane": "data",
	  "shared": false,
	  "status": "ACTIVE",
	  "subnets": [],
	  "tags": {
		"keyword1": "value1",
		"keyword2": "value2"
	  },
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }`
const CreateResponse = `
{
	"network": {
	  "admin_state_up": true,
	  "description": "Example Network 2",
	  "id": "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2",
	  "name": "Example Network 2",
	  "plane": "data",
	  "shared": false,
	  "status": "ACTIVE",
	  "subnets": [],
	  "tags": {
		"keyword1": "value1",
		"keyword2": "value2"
	  },
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }`
const CreateRequest = `
{
	"network": {
	  "admin_state_up": true,
	  "description": "Example Network 2",
	  "name": "Example Network 2",
	  "plane": "data",
	  "tags": {
		"keyword1": "value1",
		"keyword2": "value2"
	  },
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }
`
const UpdateResponse = `
{
	"network": {
	  "admin_state_up": false,
	  "description": "UPDATED",
	  "id": "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2",
	  "name": "UPDATED",
	  "plane": "data",
	  "shared": false,
	  "status": "PENDING_UPDATE",
	  "subnets": [],
	  "tags": {
		"keyword1": "UPDATED",
		"keyword3": "CREATED"
	  },
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }`
const UpdateRequest = `
{
	"network": {
	  "admin_state_up": false,
	  "description": "UPDATED",
	  "name": "UPDATED",
	  "tags": {
		"keyword1": "UPDATED",
		"keyword3": "CREATED"
	  }
	}
  }`

var Network1 = networks.Network{
	AdminStateUp: true,
	Description:  "",
	ID:           "8f36b88a-443f-4d97-9751-34d34af9e782",
	Name:         "",
	Plane:        "data",
	Shared:       false,
	Status:       "ACTIVE",
	Subnets: []string{
		"ab49eb24-667f-4a4e-9421-b4d915bff416",
		"f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678",
	},
	Tags:     map[string]string{},
	TenantID: "dcb2d589c0c646d0bad45c0cf9f90cf1",
}

var Network2 = networks.Network{
	AdminStateUp: true,
	Description:  "Example Network 2",
	ID:           "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2",
	Name:         "Example Network 2",
	Plane:        "data",
	Shared:       false,
	Status:       "ACTIVE",
	Subnets:      []string{},
	Tags: map[string]string{
		"keyword1": "value1",
		"keyword2": "value2",
	},
	TenantID: "dcb2d589c0c646d0bad45c0cf9f90cf1",
}

var ExpectedNetworkSlice = []networks.Network{Network1, Network2}
