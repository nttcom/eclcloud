package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/tenants"
	"github.com/nttcom/eclcloud/v4/pagination"

	th "github.com/nttcom/eclcloud/v4/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListTenant(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/tenants", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListResponse)
	})

	count := 0
	err := tenants.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := tenants.ExtractTenants(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedTenantsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListTenantAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/tenants", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, ListResponse)
	})

	allPages, err := tenants.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allZones, err := tenants.ExtractTenants(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allZones))
}

func TestGetTenant(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/tenants/%s", idTenant1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetResponse)
	})

	actual, err := tenants.Get(fakeclient.ServiceClient(), idTenant1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponseStruct, actual)
}

func TestCreateTenant(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/tenants", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateResponse)
	})

	createOpts := tenants.CreateOpts{
		WorkspaceID:  workspaceID1,
		TenantRegion: "jp1",
	}

	// clone FirstTenant into createdTenant(Used as assertion target)
	// and initialize StartTime
	createdTenant := FirstTenant
	createdTenant.StartTime = time.Time{}

	actual, err := tenants.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createdTenant, actual)
}
