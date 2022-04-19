package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/vna/v1/appliance_plans"
	"github.com/nttcom/eclcloud/v2/pagination"
	th "github.com/nttcom/eclcloud/v2/testhelper"
	cli "github.com/nttcom/eclcloud/v2/testhelper/client"
)

const TokenID = cli.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := cli.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"
	return sc
}

func TestListAppliancePlans(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/virtual_network_appliance_plans",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, ListResponse)
		})

	client := ServiceClient()
	count := 0

	appliance_plans.List(client, appliance_plans.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := appliance_plans.ExtractVirtualNetworkAppliancePlans(page)
		if err != nil {
			t.Errorf("Failed to extract Virtual Network Appliance Plans: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedVirtualNetworkAppliancePlanSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetAppliancePlan(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/virtual_network_appliance_plans/6589b37a-cf82-4918-96fe-255683f78e76", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := appliance_plans.Get(ServiceClient(), "6589b37a-cf82-4918-96fe-255683f78e76", appliance_plans.GetOpts{}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &VirtualNetworkApplianceDetail, s)
}

func TestIDFromName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/virtual_network_appliance_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := ServiceClient()

	expectedID := "37556569-87f2-4699-b5ff-bf38e7cbf8a7"
	actualID, err := appliance_plans.IDFromName(client, "appliance_plans_name")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedID, actualID)
}

func TestIDFromNameNoResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/virtual_network_appliance_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := ServiceClient()

	_, err := appliance_plans.IDFromName(client, "appliance_plans_nameX")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}

func TestIDFromNameDuplicated(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/virtual_network_appliance_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponseDuplicatedNames)
	})

	client := ServiceClient()

	_, err := appliance_plans.IDFromName(client, "appliance_plans_name")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
