package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/ecl/computevolume/v2/volumes"
)

const idVolume1 = "251df9eb-c088-4e71-808b-75a690e8814b"
const idVolume2 = "7e0b432b-c922-49d7-b85a-28ac88164328"

const sizeVolume1 = 40

const nameVolume1 = "volume1"
const nameVolume1Update = "volume1-update"

const descriptionVolume1 = "test volume 1"
const descriptionVolume1Update = "test volume 1-update"

const instanceID = "83ec2e3b-4321-422b-8706-a84185f52a0a"
const tenantID = "9ee80f2a926c49f88f166af47df4e9f5"
const az = "zone1-groupa"

const createdAt = "2019-02-06T08:06:57.000000"

var timeCreatedAt = time.Date(2019, 2, 6, 8, 6, 57, 0, time.UTC)

var listResponse = fmt.Sprintf(`{
	"volumes": [{
		"id": "%s",
		"name": "%s",
		"status": "in-use",
		"size": %d,
		"availability_zone": "%s",
		"created_at": "%s",
		"os-vol-tenant-attr:tenant_id": "%s",
		"description": "%s",
		"attachments": [{
			"host_name": null,
			"device": "/dev/vdb",
			"server_id": "%s",
			"id": "%s",
			"volume_id": "%s"
		}],
		"links": [{
			"href": "dummy_self_link",
			"rel": "self"
		}, {
			"href": "dummy_bookmark_link",
			"rel": "bookmark"
		}],
		"encrypted": false,
		"os-volume-replication:extended_status": null,
		"volume_type": "nfsdriver",
		"snapshot_id": null,
		"user_id": "2a5719084bc9457c93e659f4f13c6bfc",
		"metadata": {
			"readonly": "False",
			"attached_mode": "rw"
		},
		"volume_image_metadata": {
			".edition": "none",
			".major.version": "7",
			".official_image_template": "CentOS-7.1-1503_64_virtual-server_12",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "qcow2",
			".is_license": "False",
			"image_name": "CentOS-7.1-1503_64_virtual-server_12",
			"image_id": "df1944a7-ca45-4709-9ec6-e31664133650",
			".os.type": "centos",
			".enable.download": "True",
			"checksum": "a828b6ba68b9d13d2da0a0cb3cfaa950",
			"min_disk": "15",
			".service.type": "virtual-server",
			".virtual_server.os.pod": "other",
			".minor.version": "1-1503",
			"size": "461504512"
		},
		"source_volid": null,
		"consistencygroup_id": null,
		"bootable": "true",
		"os-volume-replication:driver_data": null,
        "replication_status": "disabled"
	}, {
		"id": "%s",
		"name": "volume 2",
		"status": "available",
		"size": 40,
		"availability_zone": "%s",
		"created_at": "%s",
		"os-vol-tenant-attr:tenant_id": "%s",
		"description": "test volume 2",
		"attachments": [],
		"links": [{
			"href": "dummy_self_link",
			"rel": "self"
		}, {
			"href": "dummy_bookmark_link",
			"rel": "bookmark"
		}],
		"encrypted": false,
		"os-volume-replication:extended_status": null,
		"volume_type": "nfsdriver",
		"snapshot_id": null,
		"user_id": "2a5719084bc9457c93e659f4f13c6bfc",
		"metadata": {},
		"volume_image_metadata": {
			".edition": "none",
			".major.version": "7",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "qcow2",
			".is_license": "True",
			"image_name": "RedHatEnterpriseLinux-7.1_64_include-license_virtual-server_42",
			"image_id": "f304bc07-056a-406f-85fc-9f97c7b8ef95",
			".os.type": "rhel",
			".enable.download": "False",
			"checksum": "85851188a680c5bddecb664914917a81",
			"min_disk": "40",
			".service.type": "virtual-server",
			".virtual_server.os.pod": "rhel",
			".minor.version": "1",
			"size": "515899392"
		},
		"source_volid": null,
		"consistencygroup_id": null,
		"bootable": "true",
		"os-volume-replication:driver_data": null,
		"replication_status": "disabled"
	}]
}`,
	// For volume 1
	idVolume1,
	nameVolume1,
	sizeVolume1,
	az,
	createdAt,
	tenantID,
	descriptionVolume1,
	instanceID,
	idVolume1,
	idVolume1,
	// For volume 2
	idVolume2,
	az,
	createdAt,
	tenantID,
)

var structVolume1 = volumes.Volume{
	ID:               idVolume1,
	Status:           "in-use",
	Size:             sizeVolume1,
	AvailabilityZone: az,
	CreatedAt:        timeCreatedAt,
	Attachments: []volumes.Attachment{{
		// AttachedAt:   time.Date(2016, 8, 6, 14, 48, 20, 0, time.UTC),
		// AttachmentID: idVolume1,
		Device:   "/dev/vdb",
		HostName: "",
		ID:       idVolume1,
		ServerID: instanceID,
		VolumeID: idVolume1,
	}},
	Name:        nameVolume1,
	Description: descriptionVolume1,
	VolumeType:  "nfsdriver",
	SnapshotID:  "",
	SourceVolID: "",
	Metadata: map[string]string{
		"readonly":      "False",
		"attached_mode": "rw",
	},
	UserID:            "2a5719084bc9457c93e659f4f13c6bfc",
	Bootable:          "true",
	Encrypted:         false,
	ReplicationStatus: "disabled",
	TenantID:          tenantID,
}

var structVolume2 = volumes.Volume{
	ID:                idVolume2,
	Status:            "available",
	Size:              40,
	AvailabilityZone:  az,
	CreatedAt:         timeCreatedAt,
	Attachments:       []volumes.Attachment{},
	Name:              "volume 2",
	Description:       "test volume 2",
	VolumeType:        "nfsdriver",
	SnapshotID:        "",
	SourceVolID:       "",
	Metadata:          map[string]string{},
	UserID:            "2a5719084bc9457c93e659f4f13c6bfc",
	Bootable:          "true",
	Encrypted:         false,
	ReplicationStatus: "disabled",
	TenantID:          tenantID,
}

var expectedVolumesSlice = []volumes.Volume{
	structVolume1,
	structVolume2,
}

var getResponse = fmt.Sprintf(`{
    "volume": {
		"id": "%s",
		"name": "%s",
		"status": "in-use",
		"size": %d,
		"availability_zone": "%s",
		"created_at": "%s",
		"os-vol-tenant-attr:tenant_id": "%s",
		"description": "%s",
		"attachments": [{
			"host_name": null,
			"device": "/dev/vdb",
			"server_id": "%s",
			"id": "%s",
			"volume_id": "%s"
		}],
		"links": [{
			"href": "dummy_self_link",
			"rel": "self"
		}, {
			"href": "dummy_bookmark_link",
			"rel": "bookmark"
		}],
		"encrypted": false,
		"os-volume-replication:extended_status": null,
		"volume_type": "nfsdriver",
		"snapshot_id": null,
		"user_id": "2a5719084bc9457c93e659f4f13c6bfc",
		"metadata": {
			"readonly": "False",
			"attached_mode": "rw"
		},
		"volume_image_metadata": {
			".edition": "none",
			".major.version": "7",
			".official_image_template": "CentOS-7.1-1503_64_virtual-server_12",
			"container_format": "bare",
			"min_ram": "0",
			"disk_format": "qcow2",
			".is_license": "False",
			"image_name": "CentOS-7.1-1503_64_virtual-server_12",
			"image_id": "df1944a7-ca45-4709-9ec6-e31664133650",
			".os.type": "centos",
			".enable.download": "True",
			"checksum": "a828b6ba68b9d13d2da0a0cb3cfaa950",
			"min_disk": "15",
			".service.type": "virtual-server",
			".virtual_server.os.pod": "other",
			".minor.version": "1-1503",
			"size": "461504512"
		},
		"source_volid": null,
		"consistencygroup_id": null,
		"bootable": "true",
		"os-volume-replication:driver_data": null,
        "replication_status": "disabled"
	}
}`,
	idVolume1,
	nameVolume1,
	sizeVolume1,
	az,
	createdAt,
	tenantID,
	descriptionVolume1,
	instanceID,
	idVolume1,
	idVolume1,
)

var createRequest = fmt.Sprintf(`{
	"volume": {
		"size": 15,
		"availability_zone": "%s",
		"description": "%s",
		"name": "%s",
		"imageRef": "dummyimage"
	}
}`,
	az,
	descriptionVolume1,
	nameVolume1,
)

var createResponse = fmt.Sprintf(`{
	"volume": {
		"status": "creating",
		"user_id": "2a5719084bc9457c93e659f4f13c6bfc",
		"attachments": [],
		"links": [{
			"href": "https://cinder-jp4-ecl.api.ntt.com/v2/9ee80f2a926c49f88f166af47df4e9f5/volumes/251df9eb-c088-4e71-808b-75a690e8814b",
			"rel": "self"
		}, {
			"href": "https://cinder-jp4-ecl.api.ntt.com/9ee80f2a926c49f88f166af47df4e9f5/volumes/251df9eb-c088-4e71-808b-75a690e8814b",
			"rel": "bookmark"
		}],
		"availability_zone": "%s",
		"bootable": "false",
		"encrypted": false,
		"created_at": "2019-02-06T08:06:57.581271",
		"description": "%s",
		"volume_type": "nfsdriver",
		"name": "%s",
		"replication_status": "disabled",
		"consistencygroup_id": null,
		"source_volid": null,
		"snapshot_id": null,
		"metadata": {},
		"id": "%s",
		"size": 15
	}
}`, az,
	descriptionVolume1,
	nameVolume1,
	idVolume1,
)

var updateResponse = fmt.Sprintf(`{
	"volume": {
		"status": "available",
		"user_id": "2a5719084bc9457c93e659f4f13c6bfc",
		"attachments": [],
		"links": [{
			"href": "https://cinder-jp4-ecl.api.ntt.com/v2/9ee80f2a926c49f88f166af47df4e9f5/volumes/7e0b432b-c922-49d7-b85a-28ac88164328",
			"rel": "self"
		}, {
			"href": "https://cinder-jp4-ecl.api.ntt.com/9ee80f2a926c49f88f166af47df4e9f5/volumes/7e0b432b-c922-49d7-b85a-28ac88164328",
			"rel": "bookmark"
		}],
		"availability_zone": "%s",
		"bootable": "true",
		"encrypted": false,
		"created_at": "2019-02-06T08:08:32.000000",
		"description": "%s",
		"volume_type": "nfsdriver",
		"name": "%s",
		"replication_status": "disabled",
		"consistencygroup_id": null,
		"source_volid": null,
		"snapshot_id": null,
		"metadata": {},
		"id": "%s",
		"size": 40
	}
}`,
	az,
	descriptionVolume1Update,
	nameVolume1Update,
	idVolume1,
)
