package testing

import (
	"github.com/nttcom/eclcloud/ecl/network/v2/internet_gateways"
)

const ListResponse = `
{
	"internet_gateways": [
		{
			"description": "test",
			"id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
			"internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
			"name": "Lab3-Internet-Service-Provider-01",
			"qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
			"status": "PENDING_CREATE",
			"tenant_id": "6c0bdafab1914ab2b2b6c415477defc7"
	  	},
		{
			"description": "",
			"id": "05db9b0e-65ed-4478-a6b3-d3fc259c8d07",
			"internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
			"name": "6_performance",
			"qos_option_id": "be985a60-e918-4cca-98f1-8886333f6f5e",
			"status": "ACTIVE",
			"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9"
		}
	]
}`

const GetResponse = `{
	"internet_gateway": {
	  	"description": "test",
	  	"id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
	  	"internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
		"name": "Lab3-Internet-Service-Provider-01",
		"qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
		"status": "PENDING_CREATE",
		"tenant_id": "6c0bdafab1914ab2b2b6c415477defc7"
	}
}`

const CreateRequest = `
{
	"internet_gateway": {
		"description": "test",
		"internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
		"name": "Lab3-Internet-Service-Provider-01",
		"qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
		"tenant_id": "6c0bdafab1914ab2b2b6c415477defc7"
	}
}
`

const CreateResponse = `
{
	"internet_gateway": {
	  "description": "test",
	  "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
	  "internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
	  "name": "Lab3-Internet-Service-Provider-01",
	  "qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
	  "status": "PENDING_CREATE",
	  "tenant_id": "6c0bdafab1914ab2b2b6c415477defc7"
	}
 }`

const UpdateRequest = `
 {
	"internet_gateway": {
	  "description": "test2",
	  "name": "Lab3-Internet-Service-Provider-01",
	  "qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40"
	}
}`

const UpdateResponse = `
{
	"internet_gateway": {
	  "description": "test2",
	  "id": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
	  "internet_service_id": "5536154d-9a00-4b11-81fb-b185c9111d90",
	  "name": "Lab3-Internet-Service-Provider-01",
	  "qos_option_id": "e497bbc3-1127-4490-a51d-93582c40ab40",
	  "status": "PENDING_UPDATE",
	  "tenant_id": "6c0bdafab1914ab2b2b6c415477defc7"
	}
}`

var InternetGateway1 = internet_gateways.InternetGateway{
	Description:       "test",
	ID:                "d32019d3-bc6e-4319-9c1d-6722fc136a22",
	InternetServiceID: "5536154d-9a00-4b11-81fb-b185c9111d90",
	Name:              "Lab3-Internet-Service-Provider-01",
	QoSOptionID:       "e497bbc3-1127-4490-a51d-93582c40ab40",
	Status:            "PENDING_CREATE",
	TenantID:          "6c0bdafab1914ab2b2b6c415477defc7",
}

var InternetGateway2 = internet_gateways.InternetGateway{
	Description:       "",
	ID:                "05db9b0e-65ed-4478-a6b3-d3fc259c8d07",
	InternetServiceID: "5536154d-9a00-4b11-81fb-b185c9111d90",
	Name:              "6_performance",
	QoSOptionID:       "be985a60-e918-4cca-98f1-8886333f6f5e",
	Status:            "ACTIVE",
	TenantID:          "19ab165c7a664abe9c217334cd0e9cc9",
}

var ExpectedInternetGatewaySlice = []internet_gateways.InternetGateway{InternetGateway1, InternetGateway2}
