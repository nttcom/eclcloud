package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/internet_services"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/internet_services", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	tmp := internet_services.List(client, internet_services.ListOpts{})
	err := tmp.EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := internet_services.ExtractInternetServices(page)
		if err != nil {
			t.Errorf("Failed to extract internet services: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedInternetServiceSlice, actual)

		return true, nil
	})

	if err != nil {
		fmt.Printf("%s", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/internet_services/a7791c79-19b0-4eb6-9a8f-ea739b44e8d5", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	i, err := internet_services.Get(fake.ServiceClient(), "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5").Extract()
	t.Logf("%s", err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &InternetService1, i)
}
