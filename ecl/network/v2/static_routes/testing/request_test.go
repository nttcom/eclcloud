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

func TestList(t *testing.T) {
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

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes/93aaec0f-1546-4062-88c5-93c397b93c03", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	i, err := static_routes.Get(fake.ServiceClient(), "93aaec0f-1546-4062-88c5-93c397b93c03").Extract()
	t.Logf("%s", err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &StaticRoute1, i)
}

func TestCreate(t *testing.T) {
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
		Name:         "TEST-01",
		Description:  "",
		Destination:  "100.127.254.152/29",
		InternetGwID: "3c5703b7-e783-42fe-ba23-5b0fe872cccb",
		Nexthop:      "100.127.254.153",
		ServiceType:  "internet",
		TenantID:     "60ed68071ca14fff8a6c28458379864b",
	}
	i, err := static_routes.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Status, "PENDING_CREATE")
	th.AssertDeepEquals(t, &StaticRoute1, i)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes/93aaec0f-1546-4062-88c5-93c397b93c03", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	name := "TEST-02"
	options := static_routes.UpdateOpts{Name: &name}
	i, err := static_routes.Update(fake.ServiceClient(), "93aaec0f-1546-4062-88c5-93c397b93c03", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Name, "TEST-02")
	th.AssertEquals(t, i.Description, "")
	th.AssertEquals(t, i.ID, "93aaec0f-1546-4062-88c5-93c397b93c03")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/static_routes/93aaec0f-1546-4062-88c5-93c397b93c03", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := static_routes.Delete(fake.ServiceClient(), "93aaec0f-1546-4062-88c5-93c397b93c03")
	th.AssertNoErr(t, res.Err)
}
