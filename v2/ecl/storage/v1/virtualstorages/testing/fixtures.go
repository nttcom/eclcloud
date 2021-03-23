package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/storage/v1/virtualstorages"
)

// Define parameters which are used in assertion.
// Additionally, kind of IDs are defined here.
const idVirtualStorage1 = "fb3efc23-ca8c-4eb5-b7f6-6fc66ff24f9c"
const idVirtualStorage2 = "3535de20-192d-4f5a-a74a-cd1a9c1bf747"

const idVolumeType = "4f4971a5-899d-42b4-8442-24f17eac9683"

const nameVirtualStorage1 = "virtual_storage_name_1"
const descriptionVirtualStorage1 = "virtual_storage_description_1"

const nameVirtualStorage1Update = "virtual_storage_name_1-update"
const descriptionVirtualStorage1Update = "virtual_storage_description_1-update"

const tenantID = "2d5b878c-147a-4d7c-87fd-90a8be9d255f"

const networkID = "511f266e-a8bf-4547-ab2a-fc4d2bda9f81"
const subnetID = "9f3fd369-e4d4-4c3a-84f1-9c5ba7686297"

const storageTime = "2015-05-17T18:14:34+0000"

const hostRoute1Destination = "0.0.0.0/0"
const hostRoute1Nexthop = "123.123.123.1"
const hostRoute2Destination = "192.168.0.0/24"
const hostRoute2Nexthop = "123.123.123.1"
const hostRoute3Destination = "192.168.1.0/24"
const hostRoute3Nexthop = "123.123.123.1"

const ipAddrPoolStart = "192.168.1.10"
const ipAddrPoolEnd = "192.168.1.20"

const ipAddrPoolStartUpdate = "192.168.1.9"
const ipAddrPoolEndUpdate = "192.168.1.21"

// ListResponse is mocked response of virtualstorages.List
var ListResponse = fmt.Sprintf(`
{
    "virtual_storages": [
        {
            "id" : "%s",
            "volume_type_id" : "%s",
            "name" : "%s",
            "description" : "%s",
            "tenant_id" : "%s",
            "network_id" : "%s",
            "subnet_id" : "%s",
            "ip_addr_pool" : {
                "start" : "%s",
                "end" : "%s"
            },
            "host_routes":[{
                    "destination": "%s",
                    "nexthop": "%s"
                },
                {
                    "destination":"%s",
                    "nexthop": "%s"
            }],
            "status"      : "available",
            "links": [{
                    "href": "http://storage.sdp.url:port/v1.0/virtual_storages/440cf918-3ee0-4143-b289-f63e1d2000e6",
                    "rel": "self"
            }],
            "created_at" : "%s",
            "updated_at" : "%s"
        },
        {
            "id" : "%s",
            "volume_type_id" : "%s",
            "name" : "virtual_storage_name_2",
            "description" : "virtual_storage_description_2",
            "tenant_id" : "%s",
            "network_id" : "%s",
            "subnet_id" : "%s",
            "ip_addr_pool" : {
                "start" : "%s",
                "end" : "%s"
            },
            "host_routes":[{
                    "destination": "%s",
                    "nexthop": "%s"
                },
                {
                    "destination":"%s",
                    "nexthop": "%s"
            }],
            "status": "available",
            "links": [{
                    "href": "http://storage.sdp.url:port/v1.0/virtual_storages/440cf918-3ee0-4143-b289-f63e1d2000e6",
                    "rel": "self"
            }],
            "created_at" : "%s",
            "updated_at" : "%s"
        }
    ]
}`,
	// for virtual storage 1
	idVirtualStorage1,
	idVolumeType,
	nameVirtualStorage1,
	descriptionVirtualStorage1,
	tenantID,
	networkID,
	subnetID,
	ipAddrPoolStart,
	ipAddrPoolEnd,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
	storageTime,
	storageTime,
	// for virtual storage 2
	idVirtualStorage1,
	idVolumeType,
	tenantID,
	networkID,
	subnetID,
	ipAddrPoolStart,
	ipAddrPoolEnd,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
	storageTime,
	storageTime)

// GetResponse is mocked format of virtualstorages.Get
var GetResponse = fmt.Sprintf(`
{
    "virtual_storage": {
		"id": "%s",
		"volume_type_id": "%s",
		"name": "%s",
		"description": "%s",
		"network_id": "%s",
		"subnet_id": "%s",
		"ip_addr_pool": {
			"start": "%s",
			"end": "%s"
		},
        "host_routes":[{
            "destination": "%s",
            "nexthop": "%s"
        },
        {
            "destination": "%s",
            "nexthop": "%s"
        }],
        "status": "available",
		"created_at": "%s",
        "updated_at" : "%s",
		"error_message": ""
    }
}`, idVirtualStorage1,
	idVolumeType,
	nameVirtualStorage1,
	descriptionVirtualStorage1,
	networkID,
	subnetID,
	ipAddrPoolStart,
	ipAddrPoolEnd,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
	storageTime,
	storageTime)

// CreateRequest is mocked request for virtualstorages.Create
var CreateRequest = fmt.Sprintf(`
{
	"virtual_storage": {
		"volume_type_id": "%s",
		"name": "%s",
		"description": "%s",
		"network_id": "%s",
		"subnet_id": "%s",
		"ip_addr_pool": {
			"start": "%s",
			"end": "%s"
        },
        "host_routes":[{
            "destination": "%s",
            "nexthop": "%s"
        },
        {
            "destination": "%s",
            "nexthop": "%s"
        }]
	}
}`, idVolumeType,
	nameVirtualStorage1,
	descriptionVirtualStorage1,
	networkID,
	subnetID,
	ipAddrPoolStart,
	ipAddrPoolEnd,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
)

// CreateResponse is mocked response of virtualstorages.Create
var CreateResponse = fmt.Sprintf(`
{
	"virtual_storage": {
		"id": "%s",
		"volume_type_id": "%s",
		"name": "%s",
		"description": "%s",
		"network_id": "%s",
		"subnet_id": "%s",
		"ip_addr_pool": {
			"start": "%s",
			"end": "%s"
		},
        "host_routes":[{
                "destination": "%s",
                "nexthop": "%s"
            },
            {
                "destination": "%s",
                "nexthop": "%s"
        }],
        "status": "creating",
		"created_at": "null",
		"error_message": ""
	}
}`, idVirtualStorage1,
	idVolumeType,
	nameVirtualStorage1,
	descriptionVirtualStorage1,
	networkID,
	subnetID,
	ipAddrPoolStart,
	ipAddrPoolEnd,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
)

// UpdateRequest is mocked request of virtualstorages.Update
var UpdateRequest = fmt.Sprintf(`
{
	"virtual_storage": {
		"name": "%s",
		"description": "%s",
		"ip_addr_pool": {
			"start": "%s",
			"end": "%s"
        },
        "host_routes":[{
            "destination": "%s",
            "nexthop": "%s"
        },
        {
            "destination": "%s",
            "nexthop": "%s"
        },
        {
            "destination": "%s",
            "nexthop": "%s"
        }]
	}
}`, nameVirtualStorage1Update,
	descriptionVirtualStorage1Update,
	ipAddrPoolStartUpdate,
	ipAddrPoolEndUpdate,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
	hostRoute3Destination,
	hostRoute3Nexthop,
)

// UpdateResponse is mocked response of virtualstorages.Update
var UpdateResponse = fmt.Sprintf(`
{
    "virtual_storage": {
		"id": "%s",
		"volume_type_id": "%s",
		"name": "%s",
		"description": "%s",
		"network_id": "%s",
		"subnet_id": "%s",
		"ip_addr_pool": {
			"start": "%s",
			"end": "%s"
		},
        "host_routes":[{
            "destination": "%s",
            "nexthop": "%s"
        },
        {
            "destination": "%s",
            "nexthop": "%s"
        },
        {
            "destination": "%s",
            "nexthop": "%s"
        }],
        "status": "available",
		"created_at": "%s",
        "updated_at" : "%s",
		"error_message": ""
    }
}`, idVirtualStorage1,
	idVolumeType,
	nameVirtualStorage1Update,
	descriptionVirtualStorage1Update,
	networkID,
	subnetID,
	ipAddrPoolStartUpdate,
	ipAddrPoolEndUpdate,
	hostRoute1Destination,
	hostRoute1Nexthop,
	hostRoute2Destination,
	hostRoute2Nexthop,
	hostRoute3Destination,
	hostRoute3Nexthop,
	storageTime,
	storageTime)

func getExpectedVirtualStoragesSlice() []virtualstorages.VirtualStorage {
	storageParsedTime, _ := time.Parse(eclcloud.ISO8601, storageTime)

	var virtualStorage1 = virtualstorages.VirtualStorage{
		ID:           idVirtualStorage1,
		VolumeTypeID: idVolumeType,
		Name:         nameVirtualStorage1,
		Description:  descriptionVirtualStorage1,
		NetworkID:    networkID,
		SubnetID:     subnetID,
		CreatedAt:    storageParsedTime,
		UpdatedAt:    storageParsedTime,
		IPAddrPool:   getIPAddrPool(false),
		HostRoutes:   getHostRoutes(false),
		Status:       "available",
	}

	var virtualStorage2 = virtualstorages.VirtualStorage{
		ID:           idVirtualStorage1,
		VolumeTypeID: idVolumeType,
		Name:         "virtual_storage_name_2",
		Description:  "virtual_storage_description_2",
		NetworkID:    networkID,
		SubnetID:     subnetID,
		CreatedAt:    storageParsedTime,
		UpdatedAt:    storageParsedTime,
		IPAddrPool:   getIPAddrPool(false),
		HostRoutes:   getHostRoutes(false),
		Status:       "available",
	}

	// ExpectedVirtualStoragesSlice is expected assertion target
	ExpectedVirtualStoragesSlice := []virtualstorages.VirtualStorage{
		virtualStorage1,
		virtualStorage2,
	}

	return ExpectedVirtualStoragesSlice
}

func getHostRoutes(isUpdate bool) []virtualstorages.HostRoute {
	hostRoutes := []virtualstorages.HostRoute{
		{
			Destination: hostRoute1Destination,
			Nexthop:     hostRoute1Nexthop,
		},
		{
			Destination: hostRoute2Destination,
			Nexthop:     hostRoute2Nexthop,
		},
	}

	if isUpdate {
		hostRoutes = append(
			hostRoutes,
			virtualstorages.HostRoute{
				Destination: hostRoute3Destination,
				Nexthop:     hostRoute3Nexthop,
			},
		)
	}

	return hostRoutes
}

func getIPAddrPool(isUpdate bool) virtualstorages.IPAddressPool {
	var ipAddrPool virtualstorages.IPAddressPool

	if isUpdate {
		ipAddrPool = virtualstorages.IPAddressPool{
			Start: ipAddrPoolStartUpdate,
			End:   ipAddrPoolEndUpdate,
		}
		return ipAddrPool
	}

	ipAddrPool = virtualstorages.IPAddressPool{
		Start: ipAddrPoolStart,
		End:   ipAddrPoolEnd,
	}
	return ipAddrPool
}

func getExpectedCreateVirtualStorage() virtualstorages.VirtualStorage {

	result := virtualstorages.VirtualStorage{
		ID:           idVirtualStorage1,
		VolumeTypeID: idVolumeType,
		Name:         nameVirtualStorage1,
		Description:  descriptionVirtualStorage1,
		NetworkID:    networkID,
		SubnetID:     subnetID,
		IPAddrPool:   getIPAddrPool(false),
		HostRoutes:   getHostRoutes(false),
		Status:       "creating",
		ErrorMessage: "",
	}
	return result
}
