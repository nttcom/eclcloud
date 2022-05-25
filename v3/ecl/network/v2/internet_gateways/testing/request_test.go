package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v3/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/internet_gateways"
	"github.com/nttcom/eclcloud/v3/pagination"
	th "github.com/nttcom/eclcloud/v3/testhelper"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/internet_gateways", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	tmp := internet_gateways.List(client, internet_gateways.ListOpts{})
	err := tmp.EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := internet_gateways.ExtractInternetGateways(page)
		if err != nil {
			t.Errorf("Failed to extract internet gateways: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedInternetGatewaySlice, actual)

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

	th.Mux.HandleFunc("/v2.0/internet_gateways/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, GetResponse)
	})

	i, err := internet_gateways.Get(fake.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22").Extract()
	t.Logf("%s", err)
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &InternetGateway1, i)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/internet_gateways", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprint(w, CreateResponse)
	})

	options := internet_gateways.CreateOpts{
		Name:              "Lab3-Internet-Service-Provider-01",
		TenantID:          "6c0bdafab1914ab2b2b6c415477defc7",
		Description:       "test",
		InternetServiceID: "5536154d-9a00-4b11-81fb-b185c9111d90",
		QoSOptionID:       "e497bbc3-1127-4490-a51d-93582c40ab40",
	}
	i, err := internet_gateways.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Status, "PENDING_CREATE")
	th.AssertDeepEquals(t, &InternetGateway1, i)
}

func TestUpdate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/internet_gateways/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, UpdateResponse)
	})

	name := "Lab3-Internet-Service-Provider-01"
	description := "test2"
	qosOptionId := "e497bbc3-1127-4490-a51d-93582c40ab40"
	options := internet_gateways.UpdateOpts{
		Name:        &name,
		Description: &description,
		QoSOptionID: &qosOptionId,
	}
	i, err := internet_gateways.Update(fake.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, i.Name, "Lab3-Internet-Service-Provider-01")
	th.AssertEquals(t, i.Description, "test2")
	th.AssertEquals(t, i.QoSOptionID, "e497bbc3-1127-4490-a51d-93582c40ab40")
	th.AssertEquals(t, i.ID, "d32019d3-bc6e-4319-9c1d-6722fc136a22")
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/internet_gateways/d32019d3-bc6e-4319-9c1d-6722fc136a22", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := internet_gateways.Delete(fake.ServiceClient(), "d32019d3-bc6e-4319-9c1d-6722fc136a22")
	th.AssertNoErr(t, res.Err)
}
