package testing

import(
	"github.com/nttcom/eclcloud/ecl/network/v2/gateway_interfaces"
)

const ListResponse = `
{
    "gw_interfaces": [
    	{
    	  "aws_gw_id": null,
    	  "azure_gw_id": null,
    	  "description": "",
    	  "gcp_gw_id": null,
    	  "gw_vipv4": "100.127.254.49",
    	  "gw_vipv6": null,
    	  "id": "09771fbb-6496-4ae1-9b53-226b6edcc1be",
    	  "interdc_gw_id": null,
    	  "internet_gw_id": "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
    	  "name": "5_Gateway",
    	  "netmask": 29,
    	  "network_id": "0200a550-82cf-4d6d-b564-a87eb63e2b75",
    	  "primary_ipv4": "100.127.254.53",
    	  "primary_ipv6": null,
    	  "secondary_ipv4": "100.127.254.54",
    	  "secondary_ipv6": null,
    	  "service_type": "internet",
    	  "status": "PENDING_CREATE",
    	  "tenant_id": "19ab165c7a664abe9c217334cd0e9cc9",
    	  "vpn_gw_id": null,
    	  "vrid": 1
    	},
    	{
    	  "aws_gw_id": null,
    	  "azure_gw_id": null,
    	  "description": "lab3-test-user-internet-gateway-interface, role : member",
    	  "gcp_gw_id": null,
    	  "gw_vipv4": "100.127.254.1",
    	  "gw_vipv6": null,
    	  "id": "165ed64c-b9d4-46b1-afc1-cbbdc356ddcb",
    	  "interdc_gw_id": null,
    	  "internet_gw_id": "0ed09319-fe94-4618-8482-23ee254f3d2c",
    	  "name": "lab3-hara-cfg-20151204",
    	  "netmask": 29,
    	  "network_id": "cce5c9a1-1ec3-40b1-bfc7-634bb914646b",
    	  "primary_ipv4": "100.127.254.3",
    	  "primary_ipv6": null,
    	  "secondary_ipv4": "100.127.254.4",
    	  "secondary_ipv6": null,
    	  "service_type": "internet",
    	  "status": "ACTIVE",
    	  "tenant_id": "fe1f6fb95b0e48ba8c59be2121a58adc",
    	  "vpn_gw_id": null,
    	  "vrid": 10
    	}
	]
}`

const GetResponse = `
{
	"gw_interface": {
		"aws_gw_id": null,
		"azure_gw_id": null,
		"description": "",
		"gcp_gw_id": null,
		"gw_vipv4": "100.127.254.49",
		"gw_vipv6": null,
		"id": "09771fbb-6496-4ae1-9b53-226b6edcc1be",
		"interdc_gw_id": null,
		"internet_gw_id": "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
		"name": "5_Gateway",
		"netmask": 29,
		"network_id": "0200a550-82cf-4d6d-b564-a87eb63e2b75",
		"primary_ipv4": "100.127.254.53",
		"primary_ipv6": null,
		"secondary_ipv4": "100.127.254.54",
		"secondary_ipv6": null,
		"service_type": "internet",
		"status": "PENDING_CREATE",
		"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9",
		"vpn_gw_id": null,
		"vrid": 1
	}
}`

const CreateRequest = `
{
	"gw_interface": {
		"description": "",
		"gw_vipv4": "100.127.254.49",
		"internet_gw_id": "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
		"name": "5_Gateway",
		"netmask": 29,
		"network_id": "0200a550-82cf-4d6d-b564-a87eb63e2b75",
		"primary_ipv4": "100.127.254.53",
		"secondary_ipv4": "100.127.254.54",
		"service_type": "internet",
		"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9",
		"vrid": 1
	}
}
`

const CreateResponse = `
{
	"gw_interface": {
		"description": "",
		"gw_vipv4": "100.127.254.49",
		"id": "09771fbb-6496-4ae1-9b53-226b6edcc1be",
		"internet_gw_id": "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
		"name": "5_Gateway",
		"netmask": 29,
		"network_id": "0200a550-82cf-4d6d-b564-a87eb63e2b75",
		"primary_ipv4": "100.127.254.53",
		"secondary_ipv4": "100.127.254.54",
		"service_type": "internet",
		"status": "PENDING_CREATE",
		"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9",
		"vrid": 1
	}
}`


const UpdateRequest = `
{
	"gw_interface": {
		"description": "Updated",
		"name": "5_Gateway"
	}
}`

const UpdateResponse = `
{
	"gw_interface": {
		"description": "Updated",
		"gw_vipv4": "100.127.254.49",
		"id": "09771fbb-6496-4ae1-9b53-226b6edcc1be",
		"internet_gw_id": "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
		"name": "5_Gateway",
		"netmask": 29,
		"network_id": "0200a550-82cf-4d6d-b564-a87eb63e2b75",
		"primary_ipv4": "100.127.254.53",
		"secondary_ipv4": "100.127.254.54",
		"service_type": "internet",
		"status": "PENDING_UPDATE",
		"tenant_id": "19ab165c7a664abe9c217334cd0e9cc9",
		"vrid": 1
	}
}`


var GatewayInterface1 = gateway_interfaces.GatewayInterface{
	Description: "",
	GwVipv4: "100.127.254.49",
	ID: "09771fbb-6496-4ae1-9b53-226b6edcc1be",
	InternetGwID: "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
	Name: "5_Gateway",
	Netmask: 29,
	NetworkID: "0200a550-82cf-4d6d-b564-a87eb63e2b75",
	PrimaryIpv4: "100.127.254.53",
	SecondaryIpv4: "100.127.254.54",
	ServiceType: "internet",
	Status: "PENDING_CREATE",
	TenantID: "19ab165c7a664abe9c217334cd0e9cc9",
	VRID: 1,
}

var GatewayInterface2 = gateway_interfaces.GatewayInterface{
	AwsGwID: "",
	AzureGwID: "",
	Description: "lab3-test-user-internet-gateway-interface, role : member",
	GcpGwID: "",
	GwVipv4: "100.127.254.1",
	GwVipv6: "",
	ID: "165ed64c-b9d4-46b1-afc1-cbbdc356ddcb",
	InterdcGwID: "",
	InternetGwID: "0ed09319-fe94-4618-8482-23ee254f3d2c",
	Name: "lab3-hara-cfg-20151204",
	Netmask: 29,
	NetworkID: "cce5c9a1-1ec3-40b1-bfc7-634bb914646b",
	PrimaryIpv4: "100.127.254.3",
	PrimaryIpv6: "",
	SecondaryIpv4: "100.127.254.4",
	SecondaryIpv6: "",
	ServiceType: "internet",
	Status: "ACTIVE",
	TenantID: "fe1f6fb95b0e48ba8c59be2121a58adc",
	VpnGwID: "",
	VRID: 10,
}

var ExpectedGatewayInterfaceSlice = []gateway_interfaces.GatewayInterface{GatewayInterface1, GatewayInterface2}
