package testing

import(
	"github.com/nttcom/eclcloud/ecl/network/v2/static_routes"
)

const ListResponse = `
{
	"static_routes": [
    {
      "aws_gw_id": null,
      "azure_gw_id": null,
      "description": "",
      "destination": "100.127.254.152/29",
      "gcp_gw_id": null,
      "id": "93aaec0f-1546-4062-88c5-93c397b93c03",
      "interdc_gw_id": null,
      "internet_gw_id": "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
      "name": "TEST-01",
      "nexthop": "100.127.254.153",
      "service_type": "internet",
      "status": "PENDING_CREATE",
      "tenant_id": "60ed68071ca14fff8a6c28458379864b",
      "vpn_gw_id": null
    },
    {
      "aws_gw_id": null,
      "azure_gw_id": null,
      "description": "jbd test",
      "destination": "100.127.254.70/31",
      "gcp_gw_id": null,
      "id": "a0caa2c8-5d30-49cd-a5f2-99517a4d848b",
      "interdc_gw_id": null,
      "internet_gw_id": "b684b775-63cc-40e6-b615-cc88c804df19",
      "name": "jbd_test_StaticRoute001",
      "nexthop": "192.168.1.0",
      "service_type": "internet",
      "status": "ACTIVE",
      "tenant_id": "28ebaf1b212f4052a1a88d27f61deaa3",
      "vpn_gw_id": null
    }
  ]
}
`

const GetResponse = `
{
	"static_route": {
    "aws_gw_id": null,
    "azure_gw_id": null,
    "description": "",
    "destination": "100.127.254.152/29",
    "gcp_gw_id": null,
    "id": "93aaec0f-1546-4062-88c5-93c397b93c03",
    "interdc_gw_id": null,
    "internet_gw_id": "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
    "name": "TEST-01",
    "nexthop": "100.127.254.153",
    "service_type": "internet",
    "status": "PENDING_CREATE",
    "tenant_id": "60ed68071ca14fff8a6c28458379864b",
    "vpn_gw_id": null
  }
}
`

const CreateRequest = `
{
	"static_route": {
    "description": "",
    "destination": "100.127.254.152/29",
    "internet_gw_id": "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
    "name": "TEST-01",
    "nexthop": "100.127.254.153",
    "service_type": "internet",
    "tenant_id": "60ed68071ca14fff8a6c28458379864b"
  }
}
`

const CreateResponse = `
{
	"static_route": {
    "description": "",
    "destination": "100.127.254.152/29",
    "id": "93aaec0f-1546-4062-88c5-93c397b93c03",
    "internet_gw_id": "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
    "name": "TEST-01",
    "nexthop": "100.127.254.153",
    "service_type": "internet",
    "status": "PENDING_CREATE",
    "tenant_id": "60ed68071ca14fff8a6c28458379864b"
  }
}
`


const UpdateRequest = `
{
	"static_route": {
    "name": "TEST-02"
  }
}
 `

const UpdateResponse = `
{
	"static_route": {
    "description": "",
    "destination": "100.127.254.152/29",
    "id": "93aaec0f-1546-4062-88c5-93c397b93c03",
    "internet_gw_id": "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
    "name": "TEST-02",
    "nexthop": "100.127.254.153",
    "service_type": "internet",
    "status": "PENDING_UPDATE",
    "tenant_id": "60ed68071ca14fff8a6c28458379864b"
  }
}
`


var StaticRoute1 = static_routes.StaticRoute {
  Description: "",
  Destination: "100.127.254.152/29",
  ID: "93aaec0f-1546-4062-88c5-93c397b93c03",
  InternetGwID: "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
  Name: "TEST-01",
  Nexthop: "100.127.254.153",
  ServiceType: "internet",
  Status: "PENDING_CREATE",
  TenantID: "60ed68071ca14fff8a6c28458379864b",
}

var StaticRoute2 = static_routes.StaticRoute {
	Description: "jbd test",
  Destination: "100.127.254.70/31",
  ID: "a0caa2c8-5d30-49cd-a5f2-99517a4d848b",
  InternetGwID: "b684b775-63cc-40e6-b615-cc88c804df19",
  Name: "jbd_test_StaticRoute001",
  Nexthop: "192.168.1.0",
  ServiceType: "internet",
  Status: "ACTIVE",
	TenantID: "28ebaf1b212f4052a1a88d27f61deaa3",
}

var ExpectedStaticRouteSlice = []static_routes.StaticRoute{StaticRoute1, StaticRoute2}
