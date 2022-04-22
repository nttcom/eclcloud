package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/health_monitors"
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

	err := health_monitors.List(cli, health_monitors.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := health_monitors.ExtractHealthMonitors(page)
		if err != nil {
			t.Errorf("Failed to extract health monitors: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedHealthMonitorsSlice, actual)

		return true, nil
	})

	if err != nil {
		t.Errorf("Failed to get health monitor list: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/health_monitors/%s", idHealthMonitor1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, getResponse)
	})

	hm, err := health_monitors.Get(ServiceClient(), idHealthMonitor1, health_monitors.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &healthMonitor1, hm)
}

func TestGetChangesHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/health_monitors/%s", idHealthMonitor1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, getChangesResponse)
	})

	getOpts := health_monitors.GetOpts{
		Changes: true,
	}

	hm, err := health_monitors.Get(ServiceClient(), idHealthMonitor1, getOpts).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &current, hm.Current)
	th.AssertEquals(t, (*health_monitors.Configuration)(nil), hm.Staged)
	th.CheckDeepEquals(t, &changesHealthMonitor1, hm)
}

func TestCreateHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/health_monitors",
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

	createOpts := health_monitors.CreateOpts{
		Name:           "health_monitor_1",
		Description:    "health_monitor_1_description",
		Tags:           map[string]string{"key1": "value1"},
		Port:           0,
		Protocol:       "icmp",
		Interval:       5,
		Retry:          3,
		Timeout:        5,
		LoadBalancerID: idLoadBalancer,
	}
  th.AssertEquals(t, createOpts.Port, 0)
	hm, err := health_monitors.Create(ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, hm.ConfigurationStatus, "CREATE_STAGED")
	th.AssertEquals(t, hm.OperationStatus, "NONE")
	th.AssertDeepEquals(t, &healthMonitor1, hm)
}

func TestDeleteHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/health_monitors/%s", idHealthMonitor1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := health_monitors.Delete(ServiceClient(), idHealthMonitor1)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateHealthMonitor(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/health_monitors/%s", idHealthMonitor1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, updateResponse)
	})

	updateOpts := health_monitors.UpdateOpts{
		Name:        "health_monitor_1-update",
		Description: "health_monitor_1_description-update",
		Tags:        map[string]string{"key1": "value1", "key2": "value2"},
	}
	hm, err := health_monitors.Update(
		ServiceClient(), idHealthMonitor1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, hm.Name, "health_monitor_1-update")
	th.AssertEquals(t, hm.Description, "health_monitor_1_description-update")
	th.AssertEquals(t, hm.ID, idHealthMonitor1)
}
