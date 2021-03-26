package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/storage/v1/volumes"
)

// Define parameters which are used in assertion.
// Additionally, kind of IDs are defined here.
const idVolume1 = "fb3efc23-ca8c-4eb5-b7f6-6fc66ff24f9c"
const idVolume2 = "3535de20-192d-4f5a-a74a-cd1a9c1bf747"

const idVirtualStorage = "4f4971a5-899d-42b4-8442-24f17eac9683"

const nameVolume1 = "virtual_storage_name_1"
const descriptionVolume1 = "virtual_storage_description_1"

const nameVolume1Update = "virtual_storage_name_1-update"
const descriptionVolume1Update = "virtual_storage_description_1-update"

const storageTime = "2015-05-17T18:14:34+0000"

const idVolumeType = "3f4971a5-899d-42b4-8442-24f17eac9684"

const IQN1 = "iqn.1986-03.com.nttcom:iscsihost.0"
const IQN2 = "iqn.1986-03.com.nttcom:iscsihost.1"

// ListResponse is mocked response of volumes.List
var ListResponse = fmt.Sprintf(`
{
    "volumes": [
        {
			"id" : "%s",
			"virtual_storage_id": "%s",
			"name" : "%s",
			"description": "%s",
			"size": 100,
			"iops_per_gb": "2",
			"initiator_iqns": [
				"%s"
			],
			"snapshot_ids": [],
			"availability_zone": "zone1_groupa",
			"created_at": "%s",
			"updated_at": "%s",
			"links": [
				{
					"href": "http://storage.sdp.url:port/v1.0/0c2eba2c5af04d3f9e9d0d410b371fde/volumes/13fea5a0-a36f-43e8-92ef-1cf472725dbe",
					"rel": "self"
				}
			],
			"metadata": {"lun_id": "1"},
			"error_message": "",
			"status": "available"
		},
        {
			"id" : "%s",
			"virtual_storage_id": "%s",
			"name" : "virtual_storage_name_2",
			"description": "virtual_storage_description_2",
			"size": 100,
			"iops_per_gb": "2",
			"initiator_iqns": [
				"%s"
			],
			"snapshot_ids": [],
			"availability_zone": "zone1_groupa",
			"created_at": "%s",
			"updated_at": "%s",
			"links": [
				{
					"href": "http://storage.sdp.url:port/v1.0/0c2eba2c5af04d3f9e9d0d410b371fde/volumes/13fea5a0-a36f-43e8-92ef-1cf472725dbe",
					"rel": "self"
				}
			],
			"metadata": {"lun_id": "1"},
			"error_message": "",
			"status": "available"
		}
    ]
}`,
	// for volume 1
	idVolume1,
	idVirtualStorage,
	nameVolume1,
	descriptionVolume1,
	IQN1,
	storageTime,
	storageTime,
	// for volume 2
	idVolume2,
	idVirtualStorage,
	IQN1,
	storageTime,
	storageTime,
)

// GetResponse is mocked format of volumes.Get
var GetResponse = fmt.Sprintf(`
{
    "volume": {
		"id" : "%s",
		"virtual_storage_id": "%s",
		"name" : "%s",
		"description": "%s",
		"size": 100,
		"iops_per_gb": "2",
		"initiator_iqns": [
			"%s"
		],
		"snapshot_ids": [],
		"availability_zone": "zone1_groupa",
		"created_at": "%s",
		"updated_at": "%s",
		"links": [
			{
				"href": "http://storage.sdp.url:port/v1.0/0c2eba2c5af04d3f9e9d0d410b371fde/volumes/13fea5a0-a36f-43e8-92ef-1cf472725dbe",
				"rel": "self"
			}
		],
		"metadata": {"lun_id": "1"},
		"error_message": "",
		"status": "available"
	}
}`, idVolume1,
	idVirtualStorage,
	nameVolume1,
	descriptionVolume1,
	IQN1,
	storageTime,
	storageTime,
)

// CreateRequestBlock is mocked request for volumes.Create
var CreateRequestBlock = fmt.Sprintf(`
{
	"volume": {
		"virtual_storage_id": "%s",
		"name" : "%s",
		"description": "%s",
		"size": 100,
		"iops_per_gb": "2",
		"initiator_iqns": [
			"%s"
		],
		"availability_zone": "zone1_groupa"
	}
}`, idVirtualStorage,
	nameVolume1,
	descriptionVolume1,
	IQN1,
)

// CreateResponseBlock is mocked response of volumes.Create
var CreateResponseBlock = fmt.Sprintf(`
{
	"volume": {
		"id" : "%s",
		"virtual_storage_id": "%s",
		"name" : "%s",
		"description": "%s",
		"size": 100,
		"iops_per_gb": "2",
		"initiator_iqns": [
			"%s"
		],
		"snapshot_ids": [],
		"availability_zone": "zone1_groupa",
		"created_at": "null",
		"links": [
			{
				"href": "http://storage.sdp.url:port/v1.0/0c2eba2c5af04d3f9e9d0d410b371fde/volumes/13fea5a0-a36f-43e8-92ef-1cf472725dbe",
				"rel": "self"
			}
		],
		"metadata": {"lun_id": "1"},
		"error_message": "",
		"status": "creating"
	}
}`, idVolume1,
	idVirtualStorage,
	nameVolume1,
	descriptionVolume1,
	IQN1,
)

// CreateRequestFile is mocked request for volumes.Create
var CreateRequestFile = fmt.Sprintf(`
{
	"volume": {
		"virtual_storage_id": "%s",
		"name" : "%s",
		"description": "%s",
		"size": 256,
		"throughput": "50",
		"availability_zone": "zone1_groupa"
	}
}`, idVirtualStorage,
	nameVolume1,
	descriptionVolume1,
)

// CreateResponseFile is mocked response of volumes.Create
var CreateResponseFile = fmt.Sprintf(`
{
	"volume": {
		"id" : "%s",
		"virtual_storage_id": "%s",
		"name" : "%s",
		"description": "%s",
		"size": 256,
		"throughput": "50",
		"snapshot_ids": [],
		"availability_zone": "zone1_groupa",
		"created_at": "null",
		"links": [
			{
				"href": "http://storage.sdp.url:port/v1.0/0c2eba2c5af04d3f9e9d0d410b371fde/volumes/13fea5a0-a36f-43e8-92ef-1cf472725dbe",
				"rel": "self"
			}
		],
		"metadata": {"lun_id": "1"},
		"error_message": "",
		"status": "creating"
	}
}`, idVolume1,
	idVirtualStorage,
	nameVolume1,
	descriptionVolume1,
)

// UpdateRequest is mocked request of volumes.Update
var UpdateRequest = fmt.Sprintf(`
{
	"volume": {
		"name": "%s",
		"description": "%s",
		"initiator_iqns": [
			"%s",
			"%s"
		]
	}
}`, nameVolume1Update,
	descriptionVolume1Update,
	IQN1,
	IQN2,
)

// UpdateResponse is mocked response of volumes.Update
var UpdateResponse = fmt.Sprintf(`
{
    "volume": {
		"id" : "%s",
		"virtual_storage_id": "%s",
		"name" : "%s",
		"description": "%s",
		"size": 100,
		"iops_per_gb": "2",
		"initiator_iqns": [
			"%s",
			"%s"
		],
		"snapshot_ids": [],
		"availability_zone": "zone1_groupa",
		"created_at": "%s",
		"updated_at": "%s",
		"links": [
			{
				"href": "http://storage.sdp.url:port/v1.0/0c2eba2c5af04d3f9e9d0d410b371fde/volumes/13fea5a0-a36f-43e8-92ef-1cf472725dbe",
				"rel": "self"
			}
		],
		"metadata": {"lun_id": "1"},
		"error_message": "",
		"status": "updating"
    }
}`, idVolume1,
	idVirtualStorage,
	nameVolume1Update,
	descriptionVolume1Update,
	IQN1,
	IQN2,
	storageTime,
	storageTime,
)

func getExpectedVolumesSlice() []volumes.Volume {
	storageParsedTime, _ := time.Parse(eclcloud.ISO8601, storageTime)

	var volume1 = volumes.Volume{
		ID:               idVolume1,
		VirtualStorageID: idVirtualStorage,
		Name:             nameVolume1,
		Description:      descriptionVolume1,
		Size:             100,
		IOPSPerGB:        "2",
		InitiatorIQNs:    []string{IQN1},
		SnapshotIDs:      []string{},
		Metadata:         map[string]string{"lun_id": "1"},
		CreatedAt:        storageParsedTime,
		UpdatedAt:        storageParsedTime,
		AvailabilityZone: "zone1_groupa",
		Status:           "available",
		ErrorMessage:     "",
	}

	var volume2 = volumes.Volume{
		ID:               idVolume2,
		VirtualStorageID: idVirtualStorage,
		Name:             "virtual_storage_name_2",
		Description:      "virtual_storage_description_2",
		Size:             100,
		IOPSPerGB:        "2",
		InitiatorIQNs:    []string{IQN1},
		SnapshotIDs:      []string{},
		Metadata:         map[string]string{"lun_id": "1"},
		CreatedAt:        storageParsedTime,
		UpdatedAt:        storageParsedTime,
		AvailabilityZone: "zone1_groupa",
		Status:           "available",
		ErrorMessage:     "",
	}

	// ExpectedVolumesSlice is expected assertion target
	ExpectedVolumesSlice := []volumes.Volume{
		volume1,
		volume2,
	}

	return ExpectedVolumesSlice
}

func getExpectedCreateBlockStorageTypeVolume() volumes.Volume {

	result := volumes.Volume{
		ID:               idVolume1,
		VirtualStorageID: idVirtualStorage,
		Name:             nameVolume1,
		Description:      descriptionVolume1,
		Size:             100,
		IOPSPerGB:        "2",
		InitiatorIQNs:    []string{IQN1},
		AvailabilityZone: "zone1_groupa",
		SnapshotIDs:      []string{},
		Metadata:         map[string]string{"lun_id": "1"},
		Status:           "creating",
		ErrorMessage:     "",
	}
	return result
}

func getExpectedCreateFileStorageTypeVolume() volumes.Volume {

	result := volumes.Volume{
		ID:               idVolume1,
		VirtualStorageID: idVirtualStorage,
		Name:             nameVolume1,
		Description:      descriptionVolume1,
		Size:             256,
		Throughput:       "50",
		AvailabilityZone: "zone1_groupa",
		SnapshotIDs:      []string{},
		Metadata:         map[string]string{"lun_id": "1"},
		Status:           "creating",
		ErrorMessage:     "",
	}
	return result
}
