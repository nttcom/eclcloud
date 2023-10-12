package testing

import (
	az "github.com/nttcom/eclcloud/v4/ecl/compute/v2/extensions/availabilityzones"
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
}
