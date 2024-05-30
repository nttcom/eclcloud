package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/operations"
	"github.com/nttcom/eclcloud/v3/pagination"
	"github.com/nttcom/eclcloud/v3/testhelper/client"

	th "github.com/nttcom/eclcloud/v3/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"

	return sc
}

func TestListOperations(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/operations",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := operations.ListOpts{}

	err := operations.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := operations.ExtractOperations(page)
		if err != nil {
			t.Errorf("Failed to extract operations: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestShowOperation(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/operations/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()

	actual, err := operations.Show(cli, id).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}
