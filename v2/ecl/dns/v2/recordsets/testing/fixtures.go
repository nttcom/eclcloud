package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/dns/v2/recordsets"
)

const idZone = "4eb5c333-7031-48ff-a247-0eeccc57472e"

const idRecordSet1 = "f1dcc528-6b94-47aa-9b13-b62a356ed44f"
const idRecordSet2 = "a0b9df1e-fa38-47a5-855d-f15927fea579"

const nameRecordSet1 = "rs1.zone1.com."
const nameRecordSet1Update = "rs1update.zone1.com."

const descriptionRecordSet1 = "a record set 1"
const descriptionRecordSet1Update = "a record set 1-updated"

const ipRecordSet1 = "10.1.0.0"
const ipRecordSet1Update = "10.1.0.1"

const TTLRecordSet1 = 3000
const TTLRecordSet1Update = 3600

const recordSetCreatedAt = "2019-02-05T23:41:57"
const recordSetUpdatedAt = "2019-02-05T23:42:13"

// ListResponse is a sample response to a TestListDNSRecordSet call.
var ListResponse = fmt.Sprintf(`{
	"recordsets": [{
		"id": "%s",
		"action": "",
		"name": "%s",
		"ttl": %d,
		"description": "%s",
		"records": ["%s"],
		"type": "A",
		"version": 1,
		"status": "ACTIVE",
		"created_at": "%s",
		"updated_at": "%s",
		"zone_id": "%s",
		"links": {
			"self": "dummylink"
		}
	}, {
		"id": "%s",
		"action": "",
		"name": "rs2.zone1.com.",
		"ttl": 3000,
		"description": "a record set 2",
		"records": ["20.1.0.0"],
		"type": "A",
		"version": 1,
		"status": "ACTIVE",
		"created_at": "%s",
		"updated_at": "%s",
		"zone_id": "%s",
		"links": {
			"self": "dummylink"
		}
	}],
	"links": {
        "self": "dummylink"
	},
	"metadata": {
		"total_count": 2
	}
}`,
	// For recordSet1
	idRecordSet1,
	nameRecordSet1,
	TTLRecordSet1,
	descriptionRecordSet1,
	ipRecordSet1,
	recordSetCreatedAt,
	recordSetUpdatedAt,
	idZone,
	// For recordSet2
	idRecordSet2,
	recordSetCreatedAt,
	recordSetUpdatedAt,
	idZone,
)

// ListResponseLimited is a sample response with limit query option.
var ListResponseLimited = fmt.Sprintf(`{
	"recordsets": [{
		"id": "%s",
		"action": "",
		"name": "rs2.zone1.com.",
		"ttl": 3000,
		"description": "a record set 2",
		"records": ["20.1.0.0"],
		"type": "A",
		"version": 1,
		"status": "ACTIVE",
		"created_at": "%s",
		"updated_at": "%s",
		"zone_id": "%s",
		"links": {
			"self": "dummylink"
		}
	}],
	"links": {
        "self": "dummylink"
	},
	"metadata": {
		"total_count": 1
	}
}`,
	idRecordSet2,
	recordSetCreatedAt,
	recordSetUpdatedAt,
	idZone,
)

// RecordSetCreatedAt is mocked created time of each records.
var RecordSetCreatedAt, _ = time.Parse(eclcloud.RFC3339MilliNoZ, recordSetCreatedAt)

// RecordSetUpdatedAt is mocked updated time of each records.
var RecordSetUpdatedAt, _ = time.Parse(eclcloud.RFC3339MilliNoZ, recordSetUpdatedAt)

// FirstRecordSet is initialized struct as actual response
var FirstRecordSet = recordsets.RecordSet{
	ID:          idRecordSet1,
	Description: descriptionRecordSet1,
	Records:     []string{ipRecordSet1},
	TTL:         TTLRecordSet1,
	Name:        nameRecordSet1,
	ZoneID:      idZone,
	CreatedAt:   RecordSetCreatedAt,
	UpdatedAt:   RecordSetUpdatedAt,
	Version:     1,
	Type:        "A",
	Status:      "ACTIVE",
	Action:      "",
	Links: []eclcloud.Link{
		{
			Rel:  "self",
			Href: "dummylink",
		},
	},
}

// SecondRecordSet is initialized struct as actual response
var SecondRecordSet = recordsets.RecordSet{
	ID:          idRecordSet2,
	Description: "a record set 2",
	Records:     []string{"20.1.0.0"},
	TTL:         3000,
	Name:        "rs2.zone1.com.",
	ZoneID:      idZone,
	CreatedAt:   RecordSetCreatedAt,
	UpdatedAt:   RecordSetUpdatedAt,
	Version:     1,
	Type:        "A",
	Status:      "ACTIVE",
	Action:      "",
	Links: []eclcloud.Link{
		{
			Rel:  "self",
			Href: "dummylink",
		},
	},
}

// ExpectedRecordSetSlice is the slice of results that should be parsed
// from ListByZoneOutput, in the expected order.
var ExpectedRecordSetSlice = []recordsets.RecordSet{FirstRecordSet, SecondRecordSet}

// ExpectedRecordSetSliceLimited is the slice of limited results that should be parsed
// from ListByZoneOutput.
var ExpectedRecordSetSliceLimited = []recordsets.RecordSet{SecondRecordSet}

// GetResponse is a sample response to a Get call.
var GetResponse = fmt.Sprintf(`{
	"id": "%s",
	"name": "%s",
	"ttl": %d,
	"description": "%s",
	"records": ["%s"],
	"type": "A",
	"version": 1,
	"created_at": "%s",
	"updated_at": "%s",
	"zone_id": "%s",
	"status": "ACTIVE",
	"links": {
		"self": "dummylink"
    }
}`,
	idRecordSet1,
	nameRecordSet1,
	TTLRecordSet1,
	descriptionRecordSet1,
	ipRecordSet1,
	recordSetCreatedAt,
	recordSetUpdatedAt,
	idZone,
)

const selfURL = "http://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets?limit=1"
const nextURL = "http://127.0.0.1:9001/v2/zones/2150b1bf-dee2-4221-9d85-11f7886fb15f/recordsets?limit=1&marker=f7b10e9b-0cae-4a91-b162-562bc6096648"

// NextPageRequest is a sample request to test pagination.
var NextPageRequest = fmt.Sprintf(`
{
  "links": {
    "self": "%s",
    "next": "%s"
  }
}`, selfURL, nextURL)

// CreateRequest is a sample request to create a resource record.
var CreateRequest = fmt.Sprintf(`{
    "name" : "%s",
    "description" : "%s",
    "type" : "A",
    "ttl" : %d,
    "records" : ["%s"]
}`,
	nameRecordSet1,
	descriptionRecordSet1,
	TTLRecordSet1,
	ipRecordSet1,
)

// CreateResponse is a sample response to a create request.
var CreateResponse = fmt.Sprintf(`{
	"recordsets": [{
		"id": "%s",
		"zone_id": "%s",
		"records": ["%s"],
		"ttl": %d,
		"name": "%s",
		"description": "%s",
		"type": "A",
		"version": 1,
		"created_at": "",
		"updated_at": null,
		"links": {
			"self": "dummylink"
		}
	}],
	"links": {
		"self": "dummylink"
	},
	"metadata": {
		"total_count": 1
	}
}`, idRecordSet1,
	idZone,
	ipRecordSet1,
	TTLRecordSet1,
	nameRecordSet1,
	descriptionRecordSet1,
)

// UpdateRequest is a sample request to update a record set.
var UpdateRequest = fmt.Sprintf(`{
    "name": "%s",
    "description" : "%s",
    "ttl" : %d,
    "records" : ["%s"]
}`,
	nameRecordSet1Update,
	descriptionRecordSet1Update,
	TTLRecordSet1Update,
	ipRecordSet1Update,
)

// UpdateResponse is a sample response to an update request.
var UpdateResponse = fmt.Sprintf(`{
	"id": "%s",
	"name": "%s",
	"ttl": %d,
	"description": "%s",
	"records": "%s",
	"type": "A",
	"version": 1,
	"created_at": null,
	"updated_at": null,
	"zone_id": "%s",
	"links": {
		"self": "dummylink"
	} 
}`,
	idRecordSet1,
	nameRecordSet1Update,
	TTLRecordSet1Update,
	descriptionRecordSet1Update,
	ipRecordSet1Update,
	idZone,
)

// UpdatedRecordSet is initialized struct as actual response of update
var UpdatedRecordSet = recordsets.RecordSet{
	ID:          idRecordSet1,
	Name:        nameRecordSet1Update,
	TTL:         TTLRecordSet1Update,
	Description: descriptionRecordSet1Update,
	Records:     ipRecordSet1Update,
	Type:        "A",
	Version:     1,
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
	ZoneID:      idZone,
	// Status:      "",
	// Action:      "",
	Links: []eclcloud.Link{
		{
			Rel:  "self",
			Href: "dummylink",
		},
	},
}
