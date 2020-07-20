package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_interfaces"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestListLoadBalancerInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_interfaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	load_balancer_interfaces.List(client, load_balancer_interfaces.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := load_balancer_interfaces.ExtractLoadBalancerInterfaces(page)
		if err != nil {
			t.Errorf("Failed to extract Load Balancer Interfaces: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedLoadBalancerInterfaceSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetLoadBalancerInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_interfaces/5f3cae7c-58a5-4124-b622-9ca3cfbf2525", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := load_balancer_interfaces.Get(fake.ServiceClient(), "5f3cae7c-58a5-4124-b622-9ca3cfbf2525").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LoadBalancerInterfaceDetail, s)
}

func TestUpdateLoadBalancerInterface(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_interfaces/ab49eb24-667f-4a4e-9421-b4d915bff416", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	description := "test"
	ipAddress := "100.64.64.34"
	name := "Interface 1/2"
	var networkID interface{}
	networkID = "e6106a35-d79b-44a3-bda0-6009b2f8775a"
	var virtualIPAddress interface{}
	virtualIPAddress = "100.64.64.101"
	virtualIPProperties := load_balancer_interfaces.VirtualIPProperties{
		Protocol: "vrrp",
		Vrid:     10,
	}

	id := "2897f333-3554-4099-a638-64d7022bf9ae"
	slotNumber := 2

	status := "PENDING_UPDATE"

	tenantID := "6a156ddf2ecd497ca786ff2da6df5aa8"

	loadBalancerID := "9f872504-36ab-46af-83ce-a4991c669edd"

	options := load_balancer_interfaces.UpdateOpts{
		Description:         &description,
		IPAddress:           ipAddress,
		Name:                &name,
		NetworkID:           &networkID,
		VirtualIPAddress:    &virtualIPAddress,
		VirtualIPProperties: &virtualIPProperties,
	}

	s, err := load_balancer_interfaces.Update(fake.ServiceClient(), "ab49eb24-667f-4a4e-9421-b4d915bff416", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckEquals(t, description, s.Description)
	th.CheckEquals(t, id, s.ID)
	th.CheckEquals(t, ipAddress, *s.IPAddress)
	th.CheckEquals(t, loadBalancerID, s.LoadBalancerID)
	th.CheckEquals(t, name, s.Name)
	th.CheckEquals(t, networkID, *s.NetworkID)
	th.CheckEquals(t, slotNumber, s.SlotNumber)
	th.CheckEquals(t, status, s.Status)
	th.CheckEquals(t, tenantID, s.TenantID)
	th.CheckEquals(t, virtualIPAddress, *s.VirtualIPAddress)
	th.CheckDeepEquals(t, virtualIPProperties, *s.VirtualIPProperties)
}

func TestIDFromName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_interfaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	expectedID := "b409f68e-9307-4649-9073-bb3cb776bda5"
	actualID, err := load_balancer_interfaces.IDFromName(client, "Interface 1/2")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedID, actualID)
}

func TestIDFromNameNoResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_interfaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	_, err := load_balancer_interfaces.IDFromName(client, "Interface X")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}

func TestIDFromNameDuplicated(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_interfaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponseDuplicatedNames)
	})

	client := fake.ServiceClient()

	_, err := load_balancer_interfaces.IDFromName(client, "Interface 1/2")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
