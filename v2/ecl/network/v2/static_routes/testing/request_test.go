package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/static_routes"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestListStaticRoutes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	tmp := static_routes.List(client, static_routes.ListOpts{})
	err := tmp.EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := static_routes.ExtractStaticRoutes(page)
		if err != nil {
			t.Errorf("Failed to extract public ips: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedStaticRouteSlice, actual)

		return true, nil
	})

	if err != nil {
		fmt.Printf("%s", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetStaticRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes/cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	i, err := static_routes.Get(fake.ServiceClient(), "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a").Extract()
	t.Logf("%s", err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &StaticRoute1, i)
}

func TestCreateStaticRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := static_routes.CreateOpts{
		Name:         "SRT2",
		Description:  "SRT2",
		Destination:  "100.127.254.116/30",
		FICGatewayID: "5af4f343-91a7-4956-aabb-9ac678d215e5",
		Nexthop:      "100.127.254.117",
		ServiceType:  "fic",
		TenantID:     "6a156ddf2ecd497ca786ff2da6df5aa8",
	}
	i, err := static_routes.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Status, "PENDING_CREATE")
	th.AssertDeepEquals(t, &StaticRoute1, i)
}

func TestUpdateStaticRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes/cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	name := "SRT2"
	description := "SRT2"
	options := static_routes.UpdateOpts{Name: &name, Description: &description}
	i, err := static_routes.Update(fake.ServiceClient(), "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Name, "SRT2")
	th.AssertEquals(t, i.Description, "SRT2")
	th.AssertEquals(t, i.ID, "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a")
}

func TestDeleteStaticRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes/cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := static_routes.Delete(fake.ServiceClient(), "cd1dacf1-0838-4ffc-bbb8-54d3152b9a5a")
	th.AssertNoErr(t, res.Err)
}
