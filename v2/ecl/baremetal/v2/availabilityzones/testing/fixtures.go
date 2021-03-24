package testing

import (
	az "github.com/nttcom/eclcloud/v2/ecl/baremetal/v2/availabilityzones"
)

const getResponse = `
{
	"availabilityZoneInfo": [{
		"zoneState": {
			"available": true
		},
		"hosts": null,
		"zoneName": "zone1-groupa"
	}, {
		"zoneState": {
			"available": true
		},
		"hosts": null,
		"zoneName": "zone1-groupb"
	}]
}
`

var azResult = []az.AvailabilityZone{
	{
		Hosts:     nil,
		ZoneName:  "zone1-groupa",
		ZoneState: az.ZoneState{Available: true},
	},
	{
		Hosts:     nil,
		ZoneName:  "zone1-groupb",
		ZoneState: az.ZoneState{Available: true},
	},
}
