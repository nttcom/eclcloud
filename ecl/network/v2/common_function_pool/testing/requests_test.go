package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v2/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/common_function_pool"
	"github.com/nttcom/eclcloud/v2/pagination"
	th "github.com/nttcom/eclcloud/v2/testhelper"
)

func TestListCommonFunctionPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	th.Mux.HandleFunc("/v2.0/common_function_pools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	common_function_pool.List(client, common_function_pool.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := common_function_pool.ExtractCommonFunctionPools(page)
		if err != nil {
			t.Errorf("Failed to extract Common Function Pools: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedCommonFunctionPoolSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetCommonFunctionPool(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/common_function_pools/c57066cc-9553-43a6-90de-c847231bc70b", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := common_function_pool.Get(fake.ServiceClient(), "c57066cc-9553-43a6-90de-c847231bc70b").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CommonFunctionDetail, s)
}

func TestIDFromName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/common_function_pools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	expectedID := "c57066cc-9553-43a6-90de-asfdfesfffff"
	actualID, err := common_function_pool.IDFromName(client, "CF_Pool1")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedID, actualID)
}

func TestIDFromNameNoResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/common_function_pools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	_, err := common_function_pool.IDFromName(client, "CF_PoolX")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}

func TestIDFromNameDuplicated(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/common_function_pools", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponseDuplicatedNames)
	})

	client := fake.ServiceClient()

	_, err := common_function_pool.IDFromName(client, "CF_Pool1")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
