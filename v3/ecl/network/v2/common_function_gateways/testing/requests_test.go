package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v3/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/common_function_gateways"
	"github.com/nttcom/eclcloud/v3/pagination"
	th "github.com/nttcom/eclcloud/v3/testhelper"
)

func TestListCommonFunctionGatway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v2.0/common_function_gateways",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, ListResponse)
		})

	client := fake.ServiceClient()
	count := 0

	common_function_gateways.List(client, common_function_gateways.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := common_function_gateways.ExtractCommonFunctionGateways(page)
		if err != nil {
			t.Errorf("Failed to extract common function gateways: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, ExpectedCommonFunctionGatewaysSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetCommonFunctionGatway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v2.0/common_function_gateways/%s", idCommonFunctionGatway1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	cfGw, err := common_function_gateways.Get(fake.ServiceClient(), idCommonFunctionGatway1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &commonFunctionGateway1, cfGw)
}

func TestCreateCommonFunctionGatway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/common_function_gateways",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, CreateRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)

			fmt.Fprintf(w, CreateResponse)
		})

	createOpts := common_function_gateways.CreateOpts{
		Name:                 nameCommonFunctionGateway1,
		Description:          descriptionCommonFunctionGateway1,
		CommonFunctionPoolID: idCommonFunctionPool,
		TenantID:             tenantID,
	}
	cfGw, err := common_function_gateways.Create(fake.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, cfGw.Status, "ACTIVE")
	th.AssertDeepEquals(t, &commonFunctionGateway1, cfGw)
}

func TestUpdateCommonFunctionGatway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v2.0/common_function_gateways/%s", idCommonFunctionGatway1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	name := nameCommonFunctionGateway1Update
	description := descriptionCommonFunctionGateway1Update
	updateOpts := common_function_gateways.UpdateOpts{
		Name:        &name,
		Description: &description,
	}
	cfGw, err := common_function_gateways.Update(
		fake.ServiceClient(), idCommonFunctionGatway1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, cfGw.Name, nameCommonFunctionGateway1Update)
	th.AssertEquals(t, cfGw.Description, descriptionCommonFunctionGateway1Update)
	th.AssertEquals(t, cfGw.ID, idCommonFunctionGatway1)
}

func TestDeleteCommonFunctionGatway(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v2.0/common_function_gateways/%s", idCommonFunctionGatway1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := common_function_gateways.Delete(fake.ServiceClient(), idCommonFunctionGatway1)
	th.AssertNoErr(t, res.Err)
}
