package testing

import (
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/ports"
)

const ListResponse = `
{
	"ports": [
	  {
		"admin_state_up": true,
		"allowed_address_pairs": [],
		"description": "DHCP Server Port",
		"device_id": "ab49eb24-667f-4a4e-9421-b4d915bff416",
		"device_owner": "network:dhcp",
		"fixed_ips": [
		  {
			"ip_address": "192.168.2.2",
			"subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		  }
		],
		"id": "8db1ba30-be40-4943-a7be-ed5b98f053b3",
		"mac_address": "00:00:5e:00:01:00",
		"managed_by_service": false,
		"name": "dhcp-server-port",
		"network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
		"segmentation_id": null,
		"segmentation_type": null,
		"status": "ACTIVE",
		"tags": {},
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	  },
	  {
		"admin_state_up": true,
		"allowed_address_pairs": [{
			"ip_address": "192.168.2.100",
			"mac_address": "00:00:5e:00:01:01"
		}],
		"description": "",
		"device_id": "",
		"device_owner": "",
		"fixed_ips": [
		  {
			"ip_address": "192.168.2.30",
			"subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		  }
		],
		"id": "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730",
		"mac_address": "fa:16:3e:b0:ca:f1",
		"managed_by_service": false,
		"name": "port_12",
		"network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
		"segmentation_id": 0,
		"segmentation_type": "flat",
		"status": "PENDING_CREATE",
		"tags": {},
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	  }
	]
  }`
const GetResponse = `
{
	"port": {
	  "admin_state_up": true,
	  "allowed_address_pairs": [{
			"ip_address": "192.168.2.100",
			"mac_address": "00:00:5e:00:01:01"
		}],
	  "description": "",
	  "device_id": "",
	  "device_owner": "",
	  "fixed_ips": [
		{
		  "ip_address": "192.168.2.30",
		  "subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		}
	  ],
	  "id": "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730",
	  "mac_address": "fa:16:3e:b0:ca:f1",
	  "managed_by_service": false,
	  "name": "port_12",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "segmentation_id": 0,
	  "segmentation_type": "flat",
	  "status": "PENDING_CREATE",
	  "tags": {},
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }
`

const CreateResponse = `
{
	"port": {
	  "admin_state_up": true,
	  "allowed_address_pairs": [{
			"ip_address": "192.168.2.100",
			"mac_address": "00:00:5e:00:01:01"
		}],
	  "description": "",
	  "device_id": "",
	  "device_owner": "",
	  "fixed_ips": [
		{
		  "ip_address": "192.168.2.30",
		  "subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		}
	  ],
	  "id": "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730",
	  "mac_address": "fa:16:3e:b0:ca:f1",
	  "name": "port_12",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "segmentation_id": 0,
	  "segmentation_type": "flat",
	  "status": "PENDING_CREATE",
	  "tags": {},
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }`
const CreateRequest = `
{
	"port":
	{
		"admin_state_up": true,
		"allowed_address_pairs": [{
			"ip_address": "192.168.2.100",
			"mac_address": "00:00:5e:00:01:01"
		}],
		"fixed_ips": [
		{
			"ip_address": "192.168.2.30",
			"subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		}
		],
		"name": "port_12",
		"network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1",
		"segmentation_type": "flat"
	}
}`
const UpdateResponse = `
{
	"port": {
	  "admin_state_up": true,
	  "allowed_address_pairs": [{
			"ip_address": "192.168.2.100",
			"mac_address": "00:00:5e:00:01:01"
		},{
		"ip_address": "192.168.2.255",
		"mac_address": "26:8d:42:f6:c2:c4"
	}],
	  "description": "UPDATED",
	  "device_id": "b269b8c0-1a42-4464-9314-4396e51e5107",
	  "device_owner": "UPDATED",
	  "fixed_ips": [
		{
		  "ip_address": "192.168.2.30",
		  "subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		}, {
			"ip_address": "192.168.2.31",
			"subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff417"
		}
	  ],
	  "id": "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730",
	  "mac_address": "fa:16:3e:b0:ca:f1",
	  "name": "UPDATED",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "segmentation_id": 2,
	  "segmentation_type": "vlan",
	  "status": "PENDING_CREATE",
	  "tags": {
		  "some-key":"UPDATED"
		},
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }`

const UpdateRequest = `
{
	"port": {
	  "allowed_address_pairs": [{
			"ip_address": "192.168.2.100",
			"mac_address": "00:00:5e:00:01:01"
		},{
		"ip_address": "192.168.2.255",
		"mac_address": "26:8d:42:f6:c2:c4"
	}],
	"description": "UPDATED",
	"device_id": "b269b8c0-1a42-4464-9314-4396e51e5107",
	"device_owner": "UPDATED",
	  "fixed_ips": [
		{
		  "ip_address": "192.168.2.30",
		  "subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		}, {
			"ip_address": "192.168.2.31",
			"subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff417"
		}
	  ],
	  "name": "UPDATED",
	  "segmentation_id": 2,
	  "segmentation_type": "vlan",
	  "tags": {
		"some-key":"UPDATED"
	  }
	}
  }`

const RemoveAllowedAddressPairsRequest = `
{
	"port": {
		"allowed_address_pairs": [],
		"name": "new_port_name"
	}
  }
`

const RemoveAllowedAddressPairsResponse = `
{
	"port": {
	  "admin_state_up": true,
	  "allowed_address_pairs": [],
	  "description": "",
	  "device_id": "",
	  "device_owner": "",
	  "fixed_ips": [
		{
		  "ip_address": "192.168.2.30",
		  "subnet_id": "ab49eb24-667f-4a4e-9421-b4d915bff416"
		}
	  ],
	  "id": "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730",
	  "mac_address": "fa:16:3e:b0:ca:f1",
	  "name": "new_port_name",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "segmentation_id": 0,
	  "segmentation_type": "flat",
	  "status": "PENDING_CREATE",
	  "tags": {},
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }
`

var Port1 = ports.Port{
	AdminStateUp:        true,
	AllowedAddressPairs: []ports.AddressPair{},
	Description:         "DHCP Server Port",
	DeviceID:            "ab49eb24-667f-4a4e-9421-b4d915bff416",
	DeviceOwner:         "network:dhcp",
	FixedIPs: []ports.IP{{
		IPAddress: "192.168.2.2",
		SubnetID:  "ab49eb24-667f-4a4e-9421-b4d915bff416",
	}},
	ID:               "8db1ba30-be40-4943-a7be-ed5b98f053b3",
	MACAddress:       "00:00:5e:00:01:00",
	ManagedByService: false,
	Name:             "dhcp-server-port",
	NetworkID:        "8f36b88a-443f-4d97-9751-34d34af9e782",
	Status:           "ACTIVE",
	Tags:             map[string]string{},
	TenantID:         "dcb2d589c0c646d0bad45c0cf9f90cf1",
}

var Port2 = ports.Port{
	AdminStateUp: true,
	AllowedAddressPairs: []ports.AddressPair{{
		IPAddress:  "192.168.2.100",
		MACAddress: "00:00:5e:00:01:01",
	}},
	Description: "",
	DeviceID:    "",
	DeviceOwner: "",
	FixedIPs: []ports.IP{{
		IPAddress: "192.168.2.30",
		SubnetID:  "ab49eb24-667f-4a4e-9421-b4d915bff416",
	}},
	ID:               "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730",
	MACAddress:       "fa:16:3e:b0:ca:f1",
	ManagedByService: false,
	Name:             "port_12",
	NetworkID:        "8f36b88a-443f-4d97-9751-34d34af9e782",
	SegmentationID:   0,
	SegmentationType: "flat",
	Status:           "PENDING_CREATE",
	Tags:             map[string]string{},
	TenantID:         "dcb2d589c0c646d0bad45c0cf9f90cf1",
}

var ExpectedPortSlice = []ports.Port{Port1, Port2}
