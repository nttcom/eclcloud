package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/gateway_interfaces"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/gw_interfaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	tmp := gateway_interfaces.List(client, gateway_interfaces.ListOpts{})
	err := tmp.EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := gateway_interfaces.ExtractGatewayInterfaces(page)
		if err != nil {
			t.Errorf("Failed to extract gateway interfaces: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedGatewayInterfaceSlice, actual)

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

	th.Mux.HandleFunc("/v2.0/gw_interfaces/09771fbb-6496-4ae1-9b53-226b6edcc1be", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	i, err := gateway_interfaces.Get(fake.ServiceClient(), "09771fbb-6496-4ae1-9b53-226b6edcc1be").Extract()
	t.Logf("%s", err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GatewayInterface1, i)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/gw_interfaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := gateway_interfaces.CreateOpts{
		Description:   "",
		GwVipv4:       "100.127.254.49",
		InternetGwID:  "e72ef35a-c96f-45f8-aeee-e7547c5b94b3",
		Name:          "5_Gateway",
		Netmask:       29,
		NetworkID:     "0200a550-82cf-4d6d-b564-a87eb63e2b75",
		PrimaryIpv4:   "100.127.254.53",
		SecondaryIpv4: "100.127.254.54",
		ServiceType:   "internet",
		TenantID:      "19ab165c7a664abe9c217334cd0e9cc9",
		VRID:          1,
	}
	i, err := gateway_interfaces.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Status, "PENDING_CREATE")
	th.AssertDeepEquals(t, &GatewayInterface1, i)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/gw_interfaces/09771fbb-6496-4ae1-9b53-226b6edcc1be", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	description := "Updated"
	name := "5_Gateway"
	options := gateway_interfaces.UpdateOpts{
		Description: &description,
		Name:        &name,
	}
	i, err := gateway_interfaces.Update(fake.ServiceClient(), "09771fbb-6496-4ae1-9b53-226b6edcc1be", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Name, "5_Gateway")
	th.AssertEquals(t, i.Description, "Updated")
	th.AssertEquals(t, i.ID, "09771fbb-6496-4ae1-9b53-226b6edcc1be")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/gw_interfaces/09771fbb-6496-4ae1-9b53-226b6edcc1be", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := gateway_interfaces.Delete(fake.ServiceClient(), "09771fbb-6496-4ae1-9b53-226b6edcc1be")
	th.AssertNoErr(t, res.Err)
}
