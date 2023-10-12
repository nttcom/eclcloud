package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/dns/v2/zones"
)

const idZone1 = "dcbd3d17-26ce-461d-b77c-a8774cafee75"
const idZone2 = "db511e50-3b1d-4805-98c6-a00adeb6ded0"

const nameZone1 = "myzone.com."

const descriptionZone1 = "this is my zone"
const descriptionZone1Update = "this is my zone-update"

const tenantID = "9b8f16df551e42f3b905859f28a33d55"

const zoneCreatedAt = "2019-02-05T06:09:41"
const zoneUpdatedAt = "2019-02-05T06:09:45"

// ListResponse is a sample response to a List call.
var ListResponse = fmt.Sprintf(`
{
	"api_result": "success",
	"zones": [{
		"id": "%s",
		"description": "%s",
		"project_id": "%s",
		"created_at": "%s",
		"updated_at": "%s",
		"name": "%s",
		"pool_id": "",
		"email": "",
		"ttl": 0,
		"serial": 0,
		"status": "ACTIVE",
		"masters": [],
		"type": "",
		"transferred_at": null,
		"version": 1,
		"links": {
			"self": "dummylink"
		},
		"action": "",
		"attributes": []
	}, {
		"id": "%s",
		"description": "This is my zone 2",
		"project_id": "%s",
		"created_at": "%s",
		"updated_at": "%s",
		"name": "myzone2.com.",
		"pool_id": "",
		"email": "",
		"ttl": 0,
		"serial": 0,
		"status": "ACTIVE",
		"masters": [],
		"type": "",
		"transferred_at": null,
		"version": 1,
		"links": {
			"self": "dummylink"
		},
		"action": "",
		"attributes": []
	}],
	"links": {
		"self": "dummylink"
	},
	"metadata": {
		"total_count": 2
	}
}`,
	// for Zone1
	idZone1,
	descriptionZone1,
	tenantID,
	zoneCreatedAt,
	zoneUpdatedAt,
	nameZone1,
	// for Zone2
	idZone2,
	tenantID,
	zoneCreatedAt,
	zoneUpdatedAt,
)

// ExpectedZonesSlice is the slice of results that should be parsed
// from ListOutput, in the expected order.
var ExpectedZonesSlice = []zones.Zone{FirstZone, SecondZone}

// ZoneCreatedAt is parsed zone creation time
var ZoneCreatedAt, _ = time.Parse(eclcloud.RFC3339MilliNoZ, zoneCreatedAt)

// ZoneUpdatedAt is parsed zone update time
var ZoneUpdatedAt, _ = time.Parse(eclcloud.RFC3339MilliNoZ, zoneUpdatedAt)

// FirstZone is the mock object of expected zone-1
var FirstZone = zones.Zone{
	ID:            idZone1,
	PoolID:        "",
	ProjectID:     tenantID,
	Name:          nameZone1,
	Email:         "",
	TTL:           0,
	Serial:        0,
	Status:        "ACTIVE",
	Description:   descriptionZone1,
	Masters:       []string{},
	Type:          "",
	TransferredAt: time.Time{},
	Version:       1,
	CreatedAt:     ZoneCreatedAt,
	UpdatedAt:     ZoneUpdatedAt,
	Action:        "",
	Attributes:    []string{},
	Links: map[string]interface{}{
		"self": "dummylink",
	},
}

// SecondZone is the mock object of expected zone-2
var SecondZone = zones.Zone{
	ID:            idZone2,
	PoolID:        "",
	ProjectID:     tenantID,
	Name:          "myzone2.com.",
	Email:         "",
	TTL:           0,
	Serial:        0,
	Status:        "ACTIVE",
	Description:   "This is my zone 2",
	Masters:       []string{},
	Type:          "",
	TransferredAt: time.Time{},
	Version:       1,
	CreatedAt:     ZoneCreatedAt,
	UpdatedAt:     ZoneUpdatedAt,
	Action:        "",
	Attributes:    []string{},
	Links: map[string]interface{}{
		"self": "dummylink",
	}}

// GetResponse is a sample response to a Get call.
// This get result does not have action, attributes in ECL2.0
var GetResponse = fmt.Sprintf(`
{
	"id": "%s",
	"name": "%s",
	"description": "%s",
	"project_id": "%s",
	"pool_id": "",
	"email": "",
	"ttl": 0,
	"serial": 0,
	"status": "ACTIVE",
	"masters": [],
	"type": "",
	"transferred_at": null,
	"version": 1,
	"created_at": "%s",
	"updated_at": "%s",
	"links": {
		"self": "dummylink"
	}
}`, idZone1,
	nameZone1,
	descriptionZone1,
	tenantID,
	zoneCreatedAt,
	zoneUpdatedAt,
)

// GetResponseStruct mocked actual
var GetResponseStruct = zones.Zone{
	ID:            idZone1,
	PoolID:        "",
	ProjectID:     tenantID,
	Name:          nameZone1,
	Email:         "",
	TTL:           0,
	Serial:        0,
	Status:        "ACTIVE",
	Description:   descriptionZone1,
	Masters:       []string{},
	Type:          "",
	TransferredAt: time.Time{},
	Version:       1,
	CreatedAt:     ZoneCreatedAt,
	UpdatedAt:     ZoneUpdatedAt,
	Action:        "",
	Links: map[string]interface{}{
		"self": "dummylink",
	},
}

// CreateZoneRequest is a sample request to create a zone.
var CreateZoneRequest = fmt.Sprintf(`{
    "description": "%s",
    "email": "joe@example.org",
    "name": "%s",
    "ttl": 7200,
    "type": "PRIMARY"
}`,
	descriptionZone1,
	nameZone1,
)

// CreateZoneResponse is a sample response to a create request.
var CreateZoneResponse = fmt.Sprintf(`{
	"id": "%s",
	"name": "%s",
	"description": "%s",
	"project_id": "%s",
	"pool_id": "",
	"email": "",
	"ttl": 0,
	"serial": 0,
	"status": "CREATING",
	"masters": [],
	"type": "",
	"transferred_at": null,
	"version": 1,
	"created_at": "%s",
	"updated_at": null,
	"links": {
		"self": "dummylink"
	}
}`, idZone1,
	nameZone1,
	descriptionZone1,
	tenantID,
	zoneCreatedAt,
)

// CreatedZone is the expected created zone
var CreatedZone = zones.Zone{
	ID:            idZone1,
	Name:          nameZone1,
	Description:   descriptionZone1,
	ProjectID:     tenantID,
	PoolID:        "",
	Email:         "",
	TTL:           0,
	Serial:        0,
	Status:        "CREATING",
	Masters:       []string{},
	Type:          "",
	TransferredAt: time.Time{},
	Version:       1,
	CreatedAt:     ZoneCreatedAt,
	UpdatedAt:     time.Time{},
	// Action:        "",
	Links: map[string]interface{}{
		"self": "dummylink",
	},
}

// UpdateZoneRequest is a sample request to update a zone.
var UpdateZoneRequest = fmt.Sprintf(`
{
    "ttl": 600,
	"description": "%s",
	"masters": [],
	"email": ""
}`,
	descriptionZone1Update,
)

// UpdateZoneResponse is a sample response to update a zone.
var UpdateZoneResponse = fmt.Sprintf(`{
	"id": "%s",
	"name": "%s",
	"description": "%s",
	"project_id": "%s",
	"pool_id": "",
	"email": "",
	"ttl": 0,
	"serial": 0,
	"status": "ACTIVE",
	"masters": [],
	"type": "",
	"transferred_at": null,
	"version": 1,
	"created_at": "%s",
	"updated_at": "%s",
	"links": {
		"self": "dummylink"
	}
}`, idZone1,
	nameZone1,
	descriptionZone1Update,
	tenantID,
	zoneCreatedAt,
	zoneUpdatedAt,
)

// UpdatedZone is the expected updated zone
var UpdatedZone = zones.Zone{
	ID:            idZone1,
	Name:          nameZone1,
	Description:   descriptionZone1Update,
	ProjectID:     tenantID,
	PoolID:        "",
	Email:         "",
	TTL:           0,
	Serial:        0,
	Status:        "ACTIVE",
	Masters:       []string{},
	Type:          "",
	TransferredAt: time.Time{},
	Version:       1,
	CreatedAt:     ZoneCreatedAt,
	UpdatedAt:     ZoneUpdatedAt,
	// Action:        "",
	Links: map[string]interface{}{
		"self": "dummylink",
	},
}
