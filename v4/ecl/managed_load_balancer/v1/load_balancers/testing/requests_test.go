package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/load_balancers"
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

func TestListLoadBalancers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/load_balancers",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := load_balancers.ListOpts{}

	err := load_balancers.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := load_balancers.ExtractLoadBalancers(page)
		if err != nil {
			t.Errorf("Failed to extract load balancers: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestCreateLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/load_balancers",
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
	reservedFixedIP1 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.CreateOptsReservedFixedIP{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.CreateOptsInterface{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: &[]load_balancers.CreateOptsReservedFixedIP{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.CreateOptsSyslogServer{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	createOpts := load_balancers.CreateOpts{
		Name:          "load_balancer",
		Description:   "description",
		Tags:          tags,
		PlanID:        "00713021-9aea-41da-9a88-87760c08fa72",
		SyslogServers: &[]load_balancers.CreateOptsSyslogServer{syslogServer1},
		Interfaces:    &[]load_balancers.CreateOptsInterface{interface1},
	}

	actual, err := load_balancers.Create(cli, createOpts).Extract()

	th.CheckDeepEquals(t, createResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()
	showOpts := load_balancers.ShowOpts{}

	actual, err := load_balancers.Show(cli, id, showOpts).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s", id),
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

	name := "load_balancer"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	updateOpts := load_balancers.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := load_balancers.Update(cli, id, updateOpts).Extract()

	th.CheckDeepEquals(t, updateResult(), actual)
	th.AssertNoErr(t, err)
}

func TestDeleteLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := load_balancers.Delete(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestApplyConfigurationsLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/action", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, applyConfigurationsRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()
	actionOpts := load_balancers.ActionOpts{
		ApplyConfigurations: true,
	}
	err := load_balancers.Action(cli, id, actionOpts).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestSystemUpdateLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/action", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, systemUpdateRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()
	systemUpdate := load_balancers.ActionOptsSystemUpdate{
		SystemUpdateID: "31746df7-92f9-4b5e-ad05-59f6684a54eb",
	}
	actionOpts := load_balancers.ActionOpts{
		SystemUpdate: &systemUpdate,
	}

	err := load_balancers.Action(cli, id, actionOpts).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestApplyConfigurationsAndSystemUpdateLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/action", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, applyConfigurationsAndSystemUpdateRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()
	systemUpdate := load_balancers.ActionOptsSystemUpdate{
		SystemUpdateID: "31746df7-92f9-4b5e-ad05-59f6684a54eb",
	}
	actionOpts := load_balancers.ActionOpts{
		ApplyConfigurations: true,
		SystemUpdate:        &systemUpdate,
	}

	err := load_balancers.Action(cli, id, actionOpts).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCancelConfigurationsLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/action", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, cancelConfigurationsRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()
	err := load_balancers.CancelConfigurations(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCreateStagedLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/staged", id),
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
	reservedFixedIP1 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.CreateStagedOptsReservedFixedIP{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.CreateStagedOptsInterface{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: &[]load_balancers.CreateStagedOptsReservedFixedIP{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.CreateStagedOptsSyslogServer{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}
	createStagedOpts := load_balancers.CreateStagedOpts{
		SyslogServers: &[]load_balancers.CreateStagedOptsSyslogServer{syslogServer1},
		Interfaces:    &[]load_balancers.CreateStagedOptsInterface{interface1},
	}

	actual, err := load_balancers.CreateStaged(cli, id, createStagedOpts).Extract()

	th.CheckDeepEquals(t, createStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowStagedLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showStagedResponse)
		})

	cli := ServiceClient()
	actual, err := load_balancers.ShowStaged(cli, id).Extract()

	th.CheckDeepEquals(t, showStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateStagedLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/staged", id),
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

	reservedFixedIP1IPAddress := "192.168.0.2"
	reservedFixedIP1 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP1IPAddress,
	}

	reservedFixedIP2IPAddress := "192.168.0.3"
	reservedFixedIP2 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP2IPAddress,
	}

	reservedFixedIP3IPAddress := "192.168.0.4"
	reservedFixedIP3 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP3IPAddress,
	}

	reservedFixedIP4IPAddress := "192.168.0.5"
	reservedFixedIP4 := load_balancers.UpdateStagedOptsReservedFixedIP{
		IPAddress: &reservedFixedIP4IPAddress,
	}

	interface1NetworkID := "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3"
	interface1VirtualIPAddress := "192.168.0.1"
	interface1 := load_balancers.UpdateStagedOptsInterface{
		NetworkID:        &interface1NetworkID,
		VirtualIPAddress: &interface1VirtualIPAddress,
		ReservedFixedIPs: &[]load_balancers.UpdateStagedOptsReservedFixedIP{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}

	syslogServer1IPAddress := "192.168.0.6"
	syslogServer1Port := 514
	syslogServer1Protocol := "udp"
	syslogServer1 := load_balancers.UpdateStagedOptsSyslogServer{
		IPAddress: &syslogServer1IPAddress,
		Port:      &syslogServer1Port,
		Protocol:  &syslogServer1Protocol,
	}

	updateStagedOpts := load_balancers.UpdateStagedOpts{
		SyslogServers: &[]load_balancers.UpdateStagedOptsSyslogServer{syslogServer1},
		Interfaces:    &[]load_balancers.UpdateStagedOptsInterface{interface1},
	}

	actual, err := load_balancers.UpdateStaged(cli, id, updateStagedOpts).Extract()

	th.CheckDeepEquals(t, updateStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestCancelStagedLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/load_balancers/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := load_balancers.CancelStaged(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}
