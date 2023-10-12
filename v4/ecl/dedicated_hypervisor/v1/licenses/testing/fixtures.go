package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/v4/ecl/dedicated_hypervisor/v1/licenses"

	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

// ListResult provides a single page of License results.
const ListResult = `
{
	"licenses": [
		{
			"id": "02471b45-3de0-4fc8-8469-a7cc52c378df",
			"key": "5H69L-8C3D7-K8292-03926-CREMN",
			"assigned_from": "2017-04-27T09:20:47Z",
			"expires_at": null,
			"license_type": "vCenter Server 6.x Standard"
		},
		{
			"id": "0801a388-68e8-4e41-9158-73571117c915",
			"key": "0021L-8CJ47-2829A-0A8K2-CXN4J",
			"assigned_from": "2017-06-01T04:13:31Z",
			"expires_at": null,
			"license_type": "vCenter Server 6.x Standard"
		}
	]
}
`

// GetResult provides a Get result.
const GetResult = `
{
	"license": {
		"id": "0801a388-68e8-4e41-9158-73571117c915",
		"key": "0021L-8CJ47-2829A-0A8K2-CXN4J",
		"assigned_from": "2017-06-01T04:13:31Z",
		"expires_at": null,
		"license_type": "vCenter Server 6.x Standard"
	}
}
`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
	"license_type": "vCenter Server 6.x Standard"
}
`

// FirstLicense is the first License in the List request.
var FirstLicense = licenses.License{
	ID:           "02471b45-3de0-4fc8-8469-a7cc52c378df",
	Key:          "5H69L-8C3D7-K8292-03926-CREMN",
	AssignedFrom: time.Date(2017, 4, 27, 9, 20, 47, 0, time.UTC),
	ExpiresAt:    nil,
	LicenseType:  "vCenter Server 6.x Standard",
}

// SecondLicense is the second License in the List request.
var SecondLicense = licenses.License{
	ID:           "0801a388-68e8-4e41-9158-73571117c915",
	Key:          "0021L-8CJ47-2829A-0A8K2-CXN4J",
	AssignedFrom: time.Date(2017, 6, 1, 4, 13, 31, 0, time.UTC),
	ExpiresAt:    nil,
	LicenseType:  "vCenter Server 6.x Standard",
}

// ExpectedLicensesSlice is the slice of Licenses expected to be returned from ListResult.
var ExpectedLicensesSlice = []licenses.License{FirstLicense, SecondLicense}

// HandleListLicenseSuccessfully creates an HTTP handler at `/licenses` on the
// test handler mux that responds with a list of two Licenses.
func HandleListLicensesSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/licenses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListResult)
	})
}

// HandleCreateLicenseSuccessfully creates an HTTP handler at `/licenses` on the
// test handler mux that tests License creation.
func HandleCreateLicenseSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/licenses", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetResult)
	})
}

// HandleDeleteLicenseSuccessfully creates an HTTP handler at `/licenses` on the
// test handler mux that tests License deletion.
func HandleDeleteLicenseSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/licenses/%s", FirstLicense.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}
