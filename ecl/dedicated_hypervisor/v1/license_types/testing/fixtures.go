package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/ecl/dedicated_hypervisor/v1/license_types"

	th "github.com/nttcom/eclcloud/testhelper"
	"github.com/nttcom/eclcloud/testhelper/client"
)

// ListResult provides a single page of LicenseType results.
const ListResult = `
{
	"license_types": [
		{
			"description": "Windows Server 2016 Standard Edition",
			"has_license_key": false,
			"id": "9c54c437-5f0f-46f5-8270-ddf450a44135",
			"license_switch": true,
			"name": "Windows Server 2016 Standard Edition",
			"unit": "VM"
		},
		{
			"description": "vCenter Server 6.x Standard",
			"has_license_key": true,
			"id": "e37c05ba-8fd0-493e-93d2-688833363a74",
			"license_switch": false,
			"name": "vCenter Server 6.x Standard",
			"unit": "License"
		}
	]
}
`

// FirstLicenseType is the first LicenseType in the List request.
var FirstLicenseType = license_types.LicenseType{
	ID:            "9c54c437-5f0f-46f5-8270-ddf450a44135",
	Name:          "Windows Server 2016 Standard Edition",
	HasLicenseKey: false,
	Unit:          "VM",
	LicenseSwitch: true,
	Description:   "Windows Server 2016 Standard Edition",
}

// SecondLicenseType is the second LicenseType in the List request.
var SecondLicenseType = license_types.LicenseType{
	ID:            "e37c05ba-8fd0-493e-93d2-688833363a74",
	Name:          "vCenter Server 6.x Standard",
	HasLicenseKey: true,
	Unit:          "License",
	LicenseSwitch: false,
	Description:   "vCenter Server 6.x Standard",
}

// ExpectedLicenseTypesSlice is the slice of LicenseTypes expected to be returned from ListResult.
var ExpectedLicenseTypesSlice = []license_types.LicenseType{FirstLicenseType, SecondLicenseType}

// HandleListLicenseTypesSuccessfully creates an HTTP handler at `/license_types` on the
// test handler mux that responds with a list of two LicenseType.
func HandleListLicenseTypesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/license_types", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListResult)
	})
}
