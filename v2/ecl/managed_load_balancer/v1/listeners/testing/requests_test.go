package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2"
	listeners "github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/listeners"
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

func TestListListeners(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/listeners",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0

	err := listeners.List(cli, listeners.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := listeners.ExtractListeners(page)
		if err != nil {
			t.Errorf("Failed to extract listeners: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedListenersSlice, actual)

		return true, nil
	})

	if err != nil {
		t.Errorf("Failed to get listener list: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getResponse)
	})

	l, err := listeners.Get(ServiceClient(), idListener1, listeners.GetOpts{}).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &listener1, l)
}

func TestGetChangesListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, getChangesResponse)
	})

	getOpts := listeners.GetOpts{
		Changes: true,
	}

	l, err := listeners.Get(ServiceClient(), idListener1, getOpts).Extract()
	th.AssertNoErr(t, err)

	// Current has same fields as inline parameters only when changes=true
	listener1.Current = &listeners.Listener{
		IPAddress: listener1.IPAddress,
		Port:      listener1.Port,
		Protocol:  listener1.Protocol,
	}

	// Staged is nil until updating some fielads for the listener
	listener1.Staged = (*listeners.Listener)(nil)

	th.CheckDeepEquals(t, &listener1, l)
}

func TestCreateListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/listeners",
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

	createOpts := listeners.CreateOpts{
		Name:           "listener_3",
		Description:    "listener_3_description",
		Tags:           map[string]string{"key3": "value3"},
		IPAddress:      "10.0.0.3",
		Port:           443,
		Protocol:       "tcp",
		LoadBalancerID: idLoadBalancer,
	}

	l, err := listeners.Create(ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, l.ConfigurationStatus, "CREATE_STAGED")
	th.AssertEquals(t, l.OperationStatus, "NONE")
	th.AssertDeepEquals(t, &listener3, l)
}

func TestUpdateListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s", idListener1)
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

	updateOpts := listeners.UpdateOpts{
		Name:        "listener_1-update",
		Description: "listener_1_description-update",
	}
	l, err := listeners.Update(ServiceClient(), idListener1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, l.Name, "listener_1-update")
	th.AssertEquals(t, l.Description, "listener_1_description-update")
	th.AssertEquals(t, l.ID, idListener1)
}

func TestDeleteListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := listeners.Delete(ServiceClient(), idListener1)
	th.AssertNoErr(t, res.Err)
}

func TestCreateStagedListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s/staged", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createStagedRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, createStagedResponse)
	})

	createStagedOpts := listeners.CreateStagedOpts{
		IPAddress: "10.0.1.1",
		Port:      443,
		Protocol:  "tcp",
	}

	l, err := listeners.CreateStaged(ServiceClient(), idListener1, createStagedOpts).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &stagedListener, l)
}

func TestGetStagedListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s/staged", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, getStagedResponse)
	})

	l, err := listeners.GetStaged(ServiceClient(), idListener1).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &stagedListener, l)
}

func TestUpdateStagedListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s/staged", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateStagedRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, updateStagedResponse)
	})

	updateStagedOpts := listeners.UpdateStagedOpts{
		Port: 8080,
	}

	hm, err := listeners.UpdateStaged(ServiceClient(), idListener1, updateStagedOpts).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &updateStagedListener, hm)
}

func TestDeleteStagedListener(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/listeners/%s/staged", idListener1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := listeners.DeleteStaged(ServiceClient(), idListener1)
	th.AssertNoErr(t, res.Err)
}
