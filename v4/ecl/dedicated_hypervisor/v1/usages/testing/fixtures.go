package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/v4/ecl/dedicated_hypervisor/v1/usages"

	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

// ListResult provides a single page of Usage results.
const ListResult = `
{
	"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
	"usages": [
		{
			"description": "vCenter Server 6.x Standard",
			"has_license_key": true,
			"id": "9ada4c06-a2a4-46d5-b969-72ac12433a79",
			"name": "vCenter Server 6.x Standard",
			"resource_id": "1bc271e7a8af4d988ff91612f5b122f8",
			"type": "dedicated-hypervisor.guest-image.vcenter-server-6-0-standard",
			"unit": "License",
			"value": "2"
		},
		{
			"description": "SQL Server 2014 Standard Edition",
			"has_license_key": false,
			"id": "9da9116d-cc44-4ad8-aca5-7db398fcb478",
			"name": "SQL Server 2014 Standard Edition",
			"resource_id": "d-cc44-4ad8-aca5-7db398fcb477bbbbbb",
			"type": "dedicated-hypervisor.guest-image.sql-server-2014-standard",
			"unit": "vCPU",
			"value": "6"
		}
	]
}
`

// FirstUsage is the first Usage in the List request.
var FirstUsage = usages.Usage{
	ID:            "9ada4c06-a2a4-46d5-b969-72ac12433a79",
	Type:          "dedicated-hypervisor.guest-image.vcenter-server-6-0-standard",
	Value:         "2",
	Unit:          "License",
	Name:          "vCenter Server 6.x Standard",
	Description:   "vCenter Server 6.x Standard",
	HasLicenseKey: true,
	ResourceID:    "1bc271e7a8af4d988ff91612f5b122f8",
}

// SecondUsage is the second Usage in the List request.
var SecondUsage = usages.Usage{
	ID:            "9da9116d-cc44-4ad8-aca5-7db398fcb478",
	Type:          "dedicated-hypervisor.guest-image.sql-server-2014-standard",
	Value:         "6",
	Unit:          "vCPU",
	Name:          "SQL Server 2014 Standard Edition",
	Description:   "SQL Server 2014 Standard Edition",
	HasLicenseKey: false,
	ResourceID:    "d-cc44-4ad8-aca5-7db398fcb477bbbbbb",
}

// ExpectedUsagesSlice is the slice of LicenseTypes expected to be returned from ListResult.
var ExpectedUsagesSlice = []usages.Usage{FirstUsage, SecondUsage}

// HandleListUsagesSuccessfully creates an HTTP handler at `/usages` on the
// test handler mux that responds with a list of two Usages.
func HandleListUsagesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/usages", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListResult)
	})
}

const usageID = "9ada4c06-a2a4-46d5-b969-72ac12433a79"

// GetHistoriesResult provides a single page of Usage results.
const GetHistoriesResult = `
{
	"description": "vCenter Server 6.x Standard",
	"histories": [
		{
			"time": "2019-10-10 00:00:00",
			"value": "1"
		},
		{
			"time": "2019-10-10 01:00:00",
			"value": "1"
		}
	],
	"license_type": "vCenter Server 6.x Standard",
	"resource_id": "1bc271e7a8af4d988ff91612f5b122f8",
	"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
	"unit": "License"
}
`

// FirstHistory is the first History in the Get histories request.
var FirstHistory = usages.History{
	Time:  time.Date(2019, 10, 10, 0, 0, 0, 0, time.UTC),
	Value: "1",
}

// SecondHistory is the second History in the Get histories request.
var SecondHistory = usages.History{
	Time:  time.Date(2019, 10, 10, 1, 0, 0, 0, time.UTC),
	Value: "1",
}

// ExpectedHistories is the UsageHistories expected to be returned from GetHistoriesResult.
var ExpectedHistories = &usages.UsageHistories{
	Unit:        "License",
	ResourceID:  "1bc271e7a8af4d988ff91612f5b122f8",
	LicenseType: "vCenter Server 6.x Standard",
	Histories:   []usages.History{FirstHistory, SecondHistory},
}

// HandleGetHistoriesSuccessfully creates an HTTP handler at `/usages/<usageID>/histories` on the
// test handler mux that responds with usage histories.
func HandleGetHistoriesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/usages/%s/histories", usageID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetHistoriesResult)
	})
}
