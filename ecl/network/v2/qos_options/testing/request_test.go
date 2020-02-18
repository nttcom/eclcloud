package testing

import (
	"fmt"
	"github.com/nttcom/eclcloud/ecl/network/v2/qos_options"
	"github.com/nttcom/eclcloud/pagination"

	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestListQoS(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/qos_options", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	qos_options.List(client, qos_options.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := qos_options.ExtractQoSOptions(page)
		if err != nil {
			t.Errorf("Failed to extrace ports: %v", err)
			return false, nil
		}
		th.CheckDeepEquals(t, ExpectedQosSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetQoS(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/qos_options/2c649b8e-f007-4d90-b208-9b8710937a94", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := qos_options.Get(fake.ServiceClient(), "2c649b8e-f007-4d90-b208-9b8710937a94").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &Qos1, n)
}
