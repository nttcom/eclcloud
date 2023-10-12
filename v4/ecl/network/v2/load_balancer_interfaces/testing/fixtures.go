package testing

import (
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/load_balancer_interfaces"
)

const ListResponse = `
{
  "load_balancer_interfaces": [
    {
      "description": "test1",
      "id": "b409f68e-9307-4649-9073-bb3cb776bda5",
      "ip_address": "100.64.64.34",
      "load_balancer_id": "5a109f4a-ebd8-4998-8410-98629e2bd5cd",
      "name": "Interface 1/2",
      "network_id": "30b665e3-db2b-473b-a09a-8940148b6491",
      "slot_number": 2,
      "status": "ACTIVE",
      "tenant_id": "8fe1cc29-ff7d4773bced6cb02fc8002f",
      "virtual_ip_address": "100.64.64.101",
      "virtual_ip_properties": {
        "protocol": "vrrp",
        "vrid": 10
      }
    },
    {
      "description": "test2",
      "id": "0aaef2e9-b4a0-4c31-bd98-496e0a8fed4f",
      "ip_address": null,
      "load_balancer_id": "12efe0b1-02b6-4e97-ad93-9dc1f7b5c0fc",
      "name": "Interface 1/1",
      "network_id": null,
      "slot_number": 1,
      "status": "DOWN",
      "tenant_id": "44777b33f0ee474ab1466ebee9fa369f",
      "virtual_ip_address": null,
      "virtual_ip_properties": null
    }
  ]
}
`
const GetResponse = `
{
  "load_balancer_interface": {
    "description": "test3",
    "id": "da3f99e8-a949-40e7-a0e4-4609b705a7c7",
    "ip_address": "100.64.64.34",
    "load_balancer_id": "79378a5d-bc2f-4a74-ab4b-ceae8693dca5",
    "name": "Interface 1/2",
    "network_id": "30b665e3-db2b-473b-a09a-8940148b6491",
    "slot_number": 2,
    "status": "ACTIVE",
    "tenant_id": "401c9473a52b4ee486d17ea76f466f66",
    "virtual_ip_address": "100.64.64.101",
    "virtual_ip_properties": {
      "protocol": "vrrp",
      "vrid": 10
    }
  }
}
  `

const UpdateRequest = `
{
  "load_balancer_interface": {
    "description": "test",
    "ip_address": "100.64.64.34",
    "name": "Interface 1/2",
    "network_id": "e6106a35-d79b-44a3-bda0-6009b2f8775a",
    "virtual_ip_address": "100.64.64.101",
    "virtual_ip_properties": {
      "protocol": "vrrp",
      "vrid": 10
    }
  }
}
`
const UpdateResponse = `
{
  "load_balancer_interface": {
    "description": "test",
    "id": "2897f333-3554-4099-a638-64d7022bf9ae",
    "ip_address": "100.64.64.34",
    "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
    "name": "Interface 1/2",
    "network_id": "e6106a35-d79b-44a3-bda0-6009b2f8775a",
    "slot_number": 2,
    "status": "PENDING_UPDATE",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "virtual_ip_address": "100.64.64.101",
    "virtual_ip_properties": {
      "protocol": "vrrp",
      "vrid": 10
    }
  }
}
`

var LoadBalancerInterface1 = load_balancer_interfaces.LoadBalancerInterface{
	Description:      "test1",
	ID:               "b409f68e-9307-4649-9073-bb3cb776bda5",
	IPAddress:        &DetailIPAddress,
	LoadBalancerID:   "5a109f4a-ebd8-4998-8410-98629e2bd5cd",
	Name:             "Interface 1/2",
	NetworkID:        &DetailNetworkID,
	SlotNumber:       2,
	Status:           "ACTIVE",
	TenantID:         "8fe1cc29-ff7d4773bced6cb02fc8002f",
	VirtualIPAddress: &DetailVirtualIPAddress,
	VirtualIPProperties: &load_balancer_interfaces.VirtualIPProperties{
		Protocol: "vrrp",
		Vrid:     10,
	},
}

var DetailIPAddress = "100.64.64.34"
var DetailNetworkID = "30b665e3-db2b-473b-a09a-8940148b6491"
var DetailVirtualIPAddress = "100.64.64.101"

var LoadBalancerInterface2 = load_balancer_interfaces.LoadBalancerInterface{
	Description:    "test2",
	ID:             "0aaef2e9-b4a0-4c31-bd98-496e0a8fed4f",
	LoadBalancerID: "12efe0b1-02b6-4e97-ad93-9dc1f7b5c0fc",
	Name:           "Interface 1/1",
	SlotNumber:     1,
	Status:         "DOWN",
	TenantID:       "44777b33f0ee474ab1466ebee9fa369f",
}

var LoadBalancerInterfaceDetail = load_balancer_interfaces.LoadBalancerInterface{
	Description:      "test3",
	ID:               "da3f99e8-a949-40e7-a0e4-4609b705a7c7",
	IPAddress:        &DetailIPAddress,
	LoadBalancerID:   "79378a5d-bc2f-4a74-ab4b-ceae8693dca5",
	Name:             "Interface 1/2",
	NetworkID:        &DetailNetworkID,
	SlotNumber:       2,
	Status:           "ACTIVE",
	TenantID:         "401c9473a52b4ee486d17ea76f466f66",
	VirtualIPAddress: &DetailVirtualIPAddress,
	VirtualIPProperties: &load_balancer_interfaces.VirtualIPProperties{
		Protocol: "vrrp",
		Vrid:     10,
	},
}

var ExpectedLoadBalancerInterfaceSlice = []load_balancer_interfaces.LoadBalancerInterface{LoadBalancerInterface1, LoadBalancerInterface2}

const ListResponseDuplicatedNames = `
{
  "load_balancer_interfaces": [
    {
      "description": "test1",
      "id": "b409f68e-9307-4649-9073-bb3cb776bda5",
      "ip_address": "100.64.64.34",
      "load_balancer_id": "5a109f4a-ebd8-4998-8410-98629e2bd5cd",
      "name": "Interface 1/2",
      "network_id": "30b665e3-db2b-473b-a09a-8940148b6491",
      "slot_number": 2,
      "status": "ACTIVE",
      "tenant_id": "8fe1cc29-ff7d4773bced6cb02fc8002f",
      "virtual_ip_address": "100.64.64.101",
      "virtual_ip_properties": {
        "protocol": "vrrp",
        "vrid": 10
      }
    },
    {
      "description": "test2",
      "id": "0aaef2e9-b4a0-4c31-bd98-496e0a8fed4f",
      "ip_address": null,
      "load_balancer_id": "12efe0b1-02b6-4e97-ad93-9dc1f7b5c0fc",
      "name": "Interface 1/2",
      "network_id": null,
      "slot_number": 1,
      "status": "DOWN",
      "tenant_id": "44777b33f0ee474ab1466ebee9fa369f",
      "virtual_ip_address": null,
      "virtual_ip_properties": null
    }
  ]
}
`
