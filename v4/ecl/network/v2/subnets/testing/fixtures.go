package testing

import (
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/subnets"
)

const ListResponse = `
{
	"subnets": [
	  {
		"allocation_pools": [
		  {
			"end": "192.168.2.254",
			"start": "192.168.2.2"
		  }
		],
		"cidr": "192.168.2.0/24",
		"description": "",
		"dns_nameservers": [
		  "0.0.0.0"
		],
		"enable_dhcp": true,
		"gateway_ip": "192.168.2.1",
		"host_routes": [],
		"id": "ab49eb24-667f-4a4e-9421-b4d915bff416",
		"ip_version": 4,
		"ipv6_address_mode": null,
		"ipv6_ra_mode": null,
		"name": "",
		"network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
		"ntp_servers": [],
		"status": "ACTIVE",
		"tags": {},
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	  },
	  {
		"allocation_pools": [
		  {
			"end": "192.168.10.254",
			"start": "192.168.10.2"
		  }
		],
		"cidr": "192.168.10.0/24",
		"description": "",
		"dns_nameservers": [
		  "0.0.0.0"
		],
		"enable_dhcp": true,
		"gateway_ip": "192.168.10.1",
		"host_routes": [],
		"id": "f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678",
		"ip_version": 4,
		"ipv6_address_mode": null,
		"ipv6_ra_mode": null,
		"name": "",
		"network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
		"ntp_servers": [],
		"status": "ACTIVE",
		"tags": {},
		"tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	  }
	]
  }
`
const GetResponse = `
{
	"subnet": {
	  "allocation_pools": [
		{
		  "end": "192.168.2.254",
		  "start": "192.168.2.2"
		}
	  ],
	  "cidr": "192.168.2.0/24",
	  "description": "",
	  "dns_nameservers": [
		"0.0.0.0"
	  ],
	  "enable_dhcp": true,
	  "gateway_ip": "192.168.2.1",
	  "host_routes": [],
	  "id": "ab49eb24-667f-4a4e-9421-b4d915bff416",
	  "ip_version": 4,
	  "ipv6_address_mode": null,
	  "ipv6_ra_mode": null,
	  "name": "",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "ntp_servers": [],
	  "status": "ACTIVE",
	  "tags": {},
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }
  `
const CreateResponse = `
{
	"subnet": {
	  "allocation_pools": [
		{
		  "end": "192.168.10.254",
		  "start": "192.168.10.2"
		}
	  ],
	  "cidr": "192.168.10.0/24",
	  "description": "",
	  "dns_nameservers": [
		"0.0.0.0"
	  ],
	  "enable_dhcp": true,
	  "gateway_ip": "192.168.10.1",
	  "host_routes": [],
	  "id": "f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678",
	  "ip_version": 4,
	  "name": "",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "ntp_servers": [],
	  "status": "ACTIVE",
	  "tags": {},
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }
  `
const CreateRequest = `
{
	"subnet": {
		"cidr": "192.168.10.0/24",
		"network_id": "8f36b88a-443f-4d97-9751-34d34af9e782"
	}
}
`
const UpdateResponse = `
{
	"subnet": {
	  "allocation_pools": [
		{
		  "end": "192.168.2.254",
		  "start": "192.168.2.2"
		}
	  ],
	  "cidr": "192.168.2.0/24",
	  "description": "UPDATED",
	  "dns_nameservers": [
		"0.0.0.0",
		"1.1.1.1"
	  ],
	  "enable_dhcp": false,
	  "gateway_ip": "192.168.10.1",
	  "host_routes": [
		{
		  "destination": "10.2.0.0/24",
		  "nexthop": "10.1.0.10"
		}
	  ],
	  "id": "ab49eb24-667f-4a4e-9421-b4d915bff416",
	  "ip_version": 4,
	  "ipv6_address_mode": null,
	  "ipv6_ra_mode": null,
	  "name": "UPDATED",
	  "network_id": "8f36b88a-443f-4d97-9751-34d34af9e782",
	  "ntp_servers": [
		"2.2.2.2"
	  ],
	  "status": "PENDING_UPDATE",
	  "tags": {
		"updated": "true"
	  },
	  "tenant_id": "dcb2d589c0c646d0bad45c0cf9f90cf1"
	}
  }
`
const UpdateRequest = `
{
	"subnet": {
      "description": "UPDATED",
      "dns_nameservers": [
        "0.0.0.0",
		"1.1.1.1"
      ],
      "enable_dhcp": false,
      "gateway_ip": "192.168.10.1",
      "host_routes": [{
        "destination": "10.2.0.0/24",
        "nexthop": "10.1.0.10"
	  }],
      "name": "UPDATED",
      "ntp_servers": [
				"2.2.2.2"
			],
      "tags": {
				"updated": "true"
			}
    }
}
`

var Subnet1 = subnets.Subnet{
	AllocationPools: []subnets.AllocationPool{
		{
			End:   "192.168.2.254",
			Start: "192.168.2.2",
		},
	},
	CIDR:        "192.168.2.0/24",
	Description: "",
	DNSNameservers: []string{
		"0.0.0.0",
	},
	EnableDHCP: true,
	GatewayIP:  "192.168.2.1",
	HostRoutes: []subnets.HostRoute{},
	ID:         "ab49eb24-667f-4a4e-9421-b4d915bff416",
	IPVersion:  4,
	Name:       "",
	NetworkID:  "8f36b88a-443f-4d97-9751-34d34af9e782",
	NTPServers: []string{},
	Status:     "ACTIVE",
	Tags:       map[string]string{},
	TenantID:   "dcb2d589c0c646d0bad45c0cf9f90cf1",
}

var Subnet2 = subnets.Subnet{
	AllocationPools: []subnets.AllocationPool{
		{
			End:   "192.168.10.254",
			Start: "192.168.10.2",
		},
	},
	CIDR:        "192.168.10.0/24",
	Description: "",
	DNSNameservers: []string{
		"0.0.0.0",
	},
	EnableDHCP: true,
	GatewayIP:  "192.168.10.1",
	HostRoutes: []subnets.HostRoute{},
	ID:         "f6aa2d33-f3ae-4c4e-82f7-0d4ab4c67678",
	IPVersion:  4,
	Name:       "",
	NetworkID:  "8f36b88a-443f-4d97-9751-34d34af9e782",
	NTPServers: []string{},
	Status:     "ACTIVE",
	Tags:       map[string]string{},
	TenantID:   "dcb2d589c0c646d0bad45c0cf9f90cf1",
}

var ExpectedSubnetSlice = []subnets.Subnet{Subnet1, Subnet2}
