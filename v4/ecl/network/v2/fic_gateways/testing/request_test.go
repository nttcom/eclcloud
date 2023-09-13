package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/network/v2/fic_gateways"
	"github.com/nttcom/eclcloud/v4/pagination"

	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v3/ecl/network/v2/common"
	th "github.com/nttcom/eclcloud/v3/testhelper"
)

func TestListFICGateway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/fic_gateways", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	fic_gateways.List(client, fic_gateways.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := fic_gateways.ExtractFICGateways(page)
		if err != nil {
			t.Errorf("Failed to extract FIC Gateways: %v", err)
			return false, nil
		}
		th.CheckDeepEquals(t, ExpectedFICGatewaySlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetFICGateway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	id := "07f97269-e616-4dff-a73f-ca80bc5682dc"
	th.Mux.HandleFunc(fmt.Sprintf("/v2.0/fic_gateways/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, GetResponse)
		})

	n, err := fic_gateways.Get(fake.ServiceClient(), id).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ficgw1, n)
}
