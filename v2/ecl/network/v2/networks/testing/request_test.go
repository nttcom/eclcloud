package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v2/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/networks"
	"github.com/nttcom/eclcloud/v2/pagination"
	th "github.com/nttcom/eclcloud/v2/testhelper"
)

func TestListNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	networks.List(client, networks.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := networks.ExtractNetworks(page)
		if err != nil {
			t.Errorf("Failed to extrace ports: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedNetworkSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks/a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	n, err := networks.Get(fake.ServiceClient(), "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &Network2, n)
}

func TestCreateNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	asu := true

	options := &networks.CreateOpts{
		AdminStateUp: &asu,
		Description:  "Example Network 2",
		Name:         "Example Network 2",
		Plane:        "data",
		Tags: map[string]string{
			"keyword1": "value1",
			"keyword2": "value2",
		},
		TenantID: "dcb2d589c0c646d0bad45c0cf9f90cf1",
	}
	n, err := networks.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &Network2, n)
}

func TestRequiredCreateOptsNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := networks.Create(fake.ServiceClient(), networks.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestUpdateNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks/a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	asu := false
	description := "UPDATED"
	name := "UPDATED"
	tags := map[string]string{
		"keyword1": "UPDATED",
		"keyword3": "CREATED",
	}

	options := &networks.UpdateOpts{
		AdminStateUp: &asu,
		Description:  &description,
		Name:         &name,
		Tags:         &tags,
	}
	n, err := networks.Update(fake.ServiceClient(), "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckEquals(t, asu, n.AdminStateUp)
	th.CheckEquals(t, description, n.Description)
	th.CheckEquals(t, name, n.Name)
	th.CheckDeepEquals(t, tags, n.Tags)
}

func TestDeleteNetwork(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/networks/a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := networks.Delete(fake.ServiceClient(), "a033d04b-b1fe-4ff4-a7c7-5f4b6da981d2")
	th.AssertNoErr(t, res.Err)
}
