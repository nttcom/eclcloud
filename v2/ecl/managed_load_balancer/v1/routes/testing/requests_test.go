package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/routes"
	"github.com/nttcom/eclcloud/v2/pagination"
	"github.com/nttcom/eclcloud/v2/testhelper/client"

	th "github.com/nttcom/eclcloud/v2/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"

	return sc
}

func TestListRoutes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/routes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := routes.ListOpts{}

	err := routes.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := routes.ExtractRoutes(page)
		if err != nil {
			t.Errorf("Failed to extract routes: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestCreateRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/routes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, createResponse)
		})

	cli := ServiceClient()

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	createOpts := routes.CreateOpts{
		Name:             "route",
		Description:      "description",
		Tags:             tags,
		DestinationCidr:  "172.16.0.0/24",
		NextHopIPAddress: "192.168.0.254",
		LoadBalancerID:   "67fea379-cff0-4191-9175-de7d6941a040",
	}

	actual, err := routes.Create(cli, createOpts).Extract()

	th.CheckDeepEquals(t, createResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()
	showOpts := routes.ShowOpts{}

	actual, err := routes.Show(cli, id, showOpts).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, updateResponse)
		})

	cli := ServiceClient()

	name := "route"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	updateOpts := routes.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := routes.Update(cli, id, updateOpts).Extract()

	th.CheckDeepEquals(t, updateResult(), actual)
	th.AssertNoErr(t, err)
}

func TestDeleteRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := routes.Delete(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCreateStagedRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createStagedRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, createStagedResponse)
		})

	cli := ServiceClient()
	createStagedOpts := routes.CreateStagedOpts{
		NextHopIPAddress: "192.168.0.254",
	}

	actual, err := routes.CreateStaged(cli, id, createStagedOpts).Extract()

	th.CheckDeepEquals(t, createStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowStagedRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showStagedResponse)
		})

	cli := ServiceClient()
	actual, err := routes.ShowStaged(cli, id).Extract()

	th.CheckDeepEquals(t, showStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateStagedRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateStagedRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, updateStagedResponse)
		})

	cli := ServiceClient()

	nextHopIPAddress := "192.168.0.254"
	updateStagedOpts := routes.UpdateStagedOpts{
		NextHopIPAddress: &nextHopIPAddress,
	}

	actual, err := routes.UpdateStaged(cli, id, updateStagedOpts).Extract()

	th.CheckDeepEquals(t, updateStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestCancelStagedRoute(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/routes/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := routes.CancelStaged(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}
