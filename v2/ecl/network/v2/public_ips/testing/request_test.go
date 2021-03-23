package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/public_ips"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/public_ips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	tmp := public_ips.List(client, public_ips.ListOpts{})
	err := tmp.EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := public_ips.ExtractPublicIPs(page)
		if err != nil {
			t.Errorf("Failed to extract public ips: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedPublicIPSlice, actual)

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

	th.Mux.HandleFunc("/v2.0/public_ips/0718a31b-67be-4349-946b-61a0fc38e4cd", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	i, err := public_ips.Get(fake.ServiceClient(), "0718a31b-67be-4349-946b-61a0fc38e4cd").Extract()
	t.Logf("%s", err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &PublicIP1, i)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/public_ips", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := public_ips.CreateOpts{
		Name:          "seinou-test-public",
		Description:   "",
		InternetGwID:  "2a75cfa6-89af-425b-bce5-2a85197ef04f",
		SubmaskLength: 29,
		TenantID:      "19ab165c7a664abe9c217334cd0e9cc9",
	}
	i, err := public_ips.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Status, "PENDING_CREATE")
	th.AssertDeepEquals(t, &PublicIP1, i)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/public_ips/0718a31b-67be-4349-946b-61a0fc38e4cd", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	name := "seinou-test-public"
	description := ""
	options := public_ips.UpdateOpts{Name: &name, Description: &description}
	i, err := public_ips.Update(fake.ServiceClient(), "0718a31b-67be-4349-946b-61a0fc38e4cd", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Name, "seinou-test-public")
	th.AssertEquals(t, i.Description, "")
	th.AssertEquals(t, i.ID, "0718a31b-67be-4349-946b-61a0fc38e4cd")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/public_ips/0718a31b-67be-4349-946b-61a0fc38e4cd", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := public_ips.Delete(fake.ServiceClient(), "0718a31b-67be-4349-946b-61a0fc38e4cd")
	th.AssertNoErr(t, res.Err)
}
