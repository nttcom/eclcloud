package testing

import (
	"github.com/nttcom/eclcloud/ecl/network/v2/static_routes"
)

const ListResponse = `
{
  "static_routes": [
    {
      "aws_gw_id": null,
      "azure_gw_id": null,
      "description": "SRT2",
      "destination": "100.127.254.116/30",
      "fic_gw_id": "5af4f343-91a7-4956-aabb-9ac678d215e5",
      "gcp_gw_id": null,
      "id": "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a",
      "interdc_gw_id": null,
      "internet_gw_id": null,
      "name": "SRT2",
      "nexthop": "100.127.254.117",
      "service_type": "fic",
      "status": "PENDING_CREATE",
      "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
      "vpn_gw_id": null
    },
    {
      "aws_gw_id": null,
      "azure_gw_id": null,
      "description": "StaticRoute for Scenario-test.",
      "destination": "100.127.255.184/29",
      "fic_gw_id": "1331e6a7-2876-4d34-b12f-5aac9517b034",
      "gcp_gw_id": null,
      "id": "e58162ca-9fef-4f27-898f-af0d495b780c",
      "interdc_gw_id": null,
      "internet_gw_id": null,
      "name": "StaticRoute_INGW_02_01",
      "nexthop": "100.127.255.189",
      "service_type": "fic",
      "status": "PENDING_CREATE",
      "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
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
    "description": "SRT2",
    "destination": "100.127.254.116/30",
    "fic_gw_id": "5af4f343-91a7-4956-aabb-9ac678d215e5",
    "gcp_gw_id": null,
    "id": "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a",
    "interdc_gw_id": null,
    "internet_gw_id": null,
    "name": "SRT2",
    "nexthop": "100.127.254.117",
    "service_type": "fic",
    "status": "PENDING_CREATE",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "vpn_gw_id": null
  }
}
`

const CreateRequest = `
{
  "static_route": {
    "description": "SRT2",
    "destination": "100.127.254.116/30",
    "fic_gw_id": "5af4f343-91a7-4956-aabb-9ac678d215e5",
    "name": "SRT2",
    "nexthop": "100.127.254.117",
    "service_type": "fic",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8"
  }
}
`

const CreateResponse = `
{
  "static_route": {
    "description": "SRT2",
    "destination": "100.127.254.116/30",
    "fic_gw_id": "5af4f343-91a7-4956-aabb-9ac678d215e5",
    "id": "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a",
    "name": "SRT2",
    "nexthop": "100.127.254.117",
    "service_type": "fic",
    "status": "PENDING_CREATE",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8"
  }
}
`

const UpdateRequest = `
{
  "static_route": {
    "description": "SRT2",
    "name": "SRT2"
  }
}
 `

const UpdateResponse = `
{
  "static_route": {
    "aws_gw_id": null,
    "azure_gw_id": null,
    "description": "SRT2",
    "destination": "100.127.254.116/30",
    "fic_gw_id": "5af4f343-91a7-4956-aabb-9ac678d215e5",
    "gcp_gw_id": null,
    "id": "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a",
    "interdc_gw_id": null,
    "internet_gw_id": null,
    "name": "SRT2",
    "nexthop": "100.127.254.117",
    "service_type": "fic",
    "status": "PENDING_UPDATE",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "vpn_gw_id": null
  }
}
`

var StaticRoute1 = static_routes.StaticRoute{
	Description:  "SRT2",
	Destination:  "100.127.254.116/30",
	FICGatewayID: "5af4f343-91a7-4956-aabb-9ac678d215e5",
	ID:           "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a",
	Name:         "SRT2",
	Nexthop:      "100.127.254.117",
	ServiceType:  "fic",
	Status:       "PENDING_CREATE",
	TenantID:     "6a156ddf2ecd497ca786ff2da6df5aa8",
}

var StaticRoute2 = static_routes.StaticRoute{
	Description:  "StaticRoute for Scenario-test.",
	Destination:  "100.127.255.184/29",
	FICGatewayID: "1331e6a7-2876-4d34-b12f-5aac9517b034",
	ID:           "e58162ca-9fef-4f27-898f-af0d495b780c",
	Name:         "StaticRoute_INGW_02_01",
	Nexthop:      "100.127.255.189",
	ServiceType:  "fic",
	Status:       "PENDING_CREATE",
	TenantID:     "6a156ddf2ecd497ca786ff2da6df5aa8",
}

var ExpectedStaticRouteSlice = []static_routes.StaticRoute{StaticRoute1, StaticRoute2}
