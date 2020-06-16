package testing

import "github.com/nttcom/eclcloud/ecl/network/v2/fic_gateways"

const ListResponse = `
{
	"fic_gateways": [
    {
      "description": "fic_gateway_inet_test, 10M-BE, member role",
      "fic_service_id": "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
      "id": "07f97269-e616-4dff-a73f-ca80bc5682dc",
      "name": "lab3-test-member-user-fic-gateway",
      "qos_option_id": "e41f6a2f-e197-41c8-9f71-ef19cfd2a85a",
      "status": "ACTIVE",
      "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8"
    },
    {
      "description": "",
      "fic_service_id": "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
      "id": "4c842674-60e4-48eb-b5a3-b902f832d0af",
      "name": "N000001996_V15000001",
      "qos_option_id": "aa776ce4-08a8-4cc1-9a2c-bb95e547916b",
      "status": "ACTIVE",
      "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8"
    }
  ]
}
`

const GetResponse = `
{
	"fic_gateway": {
		"description": "fic_gateway_inet_test, 10M-BE, member role",
		"fic_service_id": "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
		"id": "07f97269-e616-4dff-a73f-ca80bc5682dc",
		"name": "lab3-test-member-user-fic-gateway",
		"qos_option_id": "e41f6a2f-e197-41c8-9f71-ef19cfd2a85a",
		"status": "ACTIVE",
		"tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8"
	  }
}
`

var ficgw1 = fic_gateways.FICGateway{
	Description:  "fic_gateway_inet_test, 10M-BE, member role",
	FICServiceID: "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
	ID:           "07f97269-e616-4dff-a73f-ca80bc5682dc",
	Name:         "lab3-test-member-user-fic-gateway",
	QoSOptionID:  "e41f6a2f-e197-41c8-9f71-ef19cfd2a85a",
	Status:       "ACTIVE",
	TenantID:     "6a156ddf2ecd497ca786ff2da6df5aa8",
}

var ficgw2 = fic_gateways.FICGateway{
	Description:  "",
	FICServiceID: "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
	ID:           "4c842674-60e4-48eb-b5a3-b902f832d0af",
	Name:         "N000001996_V15000001",
	QoSOptionID:  "aa776ce4-08a8-4cc1-9a2c-bb95e547916b",
	Status:       "ACTIVE",
	TenantID:     "6a156ddf2ecd497ca786ff2da6df5aa8",
}

var ExpectedFICGatewaySlice = []fic_gateways.FICGateway{ficgw1, ficgw2}
