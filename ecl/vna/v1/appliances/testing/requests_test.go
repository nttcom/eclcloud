package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
	"github.com/nttcom/eclcloud/pagination"
	"github.com/nttcom/eclcloud/testhelper/client"

	th "github.com/nttcom/eclcloud/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"
	return sc
}

func TestListAppliances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/virtual_network_appliances",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, listResponse)
		})

	client := ServiceClient()
	count := 0

	appliances.List(client, appliances.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := appliances.ExtractAppliances(page)
		if err != nil {
			t.Errorf("Failed to extract virtual network appliances: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedAppliancesSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}
