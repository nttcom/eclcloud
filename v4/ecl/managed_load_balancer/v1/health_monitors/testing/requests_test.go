package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/health_monitors"
	"github.com/nttcom/eclcloud/v4/pagination"
	"github.com/nttcom/eclcloud/v4/testhelper/client"

	th "github.com/nttcom/eclcloud/v4/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"

	return sc
}

func TestListHealthMonitors(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/health_monitors",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := health_monitors.ListOpts{}

	err := health_monitors.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := health_monitors.ExtractHealthMonitors(page)
		if err != nil {
			t.Errorf("Failed to extract health monitors: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestCreateHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/health_monitors",
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

	createOpts := health_monitors.CreateOpts{
		Name:           "health_monitor",
		Description:    "description",
		Tags:           tags,
		Port:           80,
		Protocol:       "http",
		Interval:       5,
		Retry:          3,
		Timeout:        5,
		Path:           "/health",
		HttpStatusCode: "200-299",
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
	}

	actual, err := health_monitors.Create(cli, createOpts).Extract()

	th.CheckDeepEquals(t, createResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()
	showOpts := health_monitors.ShowOpts{}

	actual, err := health_monitors.Show(cli, id, showOpts).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s", id),
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

	name := "health_monitor"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	updateOpts := health_monitors.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := health_monitors.Update(cli, id, updateOpts).Extract()

	th.CheckDeepEquals(t, updateResult(), actual)
	th.AssertNoErr(t, err)
}

func TestDeleteHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := health_monitors.Delete(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCreateStagedHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s/staged", id),
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
	createStagedOpts := health_monitors.CreateStagedOpts{
		Port:           80,
		Protocol:       "http",
		Interval:       5,
		Retry:          3,
		Timeout:        5,
		Path:           "/health",
		HttpStatusCode: "200-299",
	}

	actual, err := health_monitors.CreateStaged(cli, id, createStagedOpts).Extract()

	th.CheckDeepEquals(t, createStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowStagedHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showStagedResponse)
		})

	cli := ServiceClient()
	actual, err := health_monitors.ShowStaged(cli, id).Extract()

	th.CheckDeepEquals(t, showStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateStagedHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s/staged", id),
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

	port := 80
	protocol := "http"
	interval := 5
	retry := 3
	timeout := 5
	path := "/health"
	httpStatusCode := "200-299"
	updateStagedOpts := health_monitors.UpdateStagedOpts{
		Port:           &port,
		Protocol:       &protocol,
		Interval:       &interval,
		Retry:          &retry,
		Timeout:        &timeout,
		Path:           &path,
		HttpStatusCode: &httpStatusCode,
	}

	actual, err := health_monitors.UpdateStaged(cli, id, updateStagedOpts).Extract()

	th.CheckDeepEquals(t, updateStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestCancelStagedHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/health_monitors/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := health_monitors.CancelStaged(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}
