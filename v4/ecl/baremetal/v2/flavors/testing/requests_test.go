package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/baremetal/v2/flavors"
	"github.com/nttcom/eclcloud/v4/pagination"

	th "github.com/nttcom/eclcloud/v4/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListFlavors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := flavors.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := flavors.ExtractFlavors(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedFlavors, actual)
		return true, nil
	})

	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListFlavorsAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/flavors/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := flavors.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)

	allFlavors, err := flavors.ExtractFlavors(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allFlavors))
}

func TestGetFlavor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/flavors/%s", "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79")
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := flavors.Get(fakeclient.ServiceClient(), "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &expectedFlavor1, actual)
}
