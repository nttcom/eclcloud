package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v4/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/load_balancer_plans"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
)

func TestListLoadBalancerPlan(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	load_balancer_plans.List(client, load_balancer_plans.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := load_balancer_plans.ExtractLoadBalancerPlans(page)
		if err != nil {
			t.Errorf("Failed to extract Load Balancer Plans: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedLoadBalancerPlanSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetLoadBalancerPlan(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_plans/5f3cae7c-58a5-4124-b622-9ca3cfbf2525", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := load_balancer_plans.Get(fake.ServiceClient(), "5f3cae7c-58a5-4124-b622-9ca3cfbf2525").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LoadBalancerDetail, s)
}

func TestIDFromName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	expectedID := "58ab4df4-10f2-4fa0-b374-74b06dd648ee"
	actualID, err := load_balancer_plans.IDFromName(client, "LB_Plan1")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedID, actualID)
}

func TestIDFromNameNoResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	_, err := load_balancer_plans.IDFromName(client, "LB_PlanX")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}

func TestIDFromNameDuplicated(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_plans", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponseDuplicatedNames)
	})

	client := fake.ServiceClient()

	_, err := load_balancer_plans.IDFromName(client, "LB_Plan1")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
