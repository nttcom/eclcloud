package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_interfaces"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_syslog_servers"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancers"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestListLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	load_balancers.List(client, load_balancers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := load_balancers.ExtractLoadBalancers(page)
		if err != nil {
			t.Errorf("Failed to extract Load Balancers: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedLoadBalancerSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers/5f3cae7c-58a5-4124-b622-9ca3cfbf2525", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := load_balancers.Get(fake.ServiceClient(), "5f3cae7c-58a5-4124-b622-9ca3cfbf2525").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LoadBalancerDetail, s)
}

func TestCreateLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := load_balancers.CreateOpts{
		AvailabilityZone:   "zone1-groupa",
		Description:        "abcdefghijklmnopqrstuvwxyz",
		LoadBalancerPlanID: "bd12784a-c66e-4f13-9f72-5143d64762b6",
		Name:               "abcdefghijklmnopqrstuvwxyz",
		TenantID:           "6a156ddf2ecd497ca786ff2da6df5aa8",
	}
	s, err := load_balancers.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &LoadBalancerDetail, s)
}

func TestRequiredCreateOptsLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := load_balancers.Create(fake.ServiceClient(), load_balancers.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestUpdateLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers/ab49eb24-667f-4a4e-9421-b4d915bff416", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	adminUsername := "user-admin"
	availabilityZone := "zone1-groupa"
	defaultGateway := interface{}("100.127.253.1")
	description := "UPDATED"
	id := "5f3cae7c-58a5-4124-b622-9ca3cfbf2525"

	ipAddress1 := "100.127.253.173"
	networkID1 := "c7f88fab-573e-47aa-b0b4-257db28dae23"
	ipAddress2 := "192.168.110.1"
	networkID2 := "1839d290-721c-49ba-99f1-3d7aa37811b0"

	interfaces := []load_balancer_interfaces.LoadBalancerInterface{
		{
			ID:         "ee335c69-b50f-4a32-9d0f-f44cef84a456",
			IPAddress:  &ipAddress1,
			Name:       "Interface 1/1",
			NetworkID:  &networkID1,
			SlotNumber: 1,
			Status:     "ACTIVE",
		},
		{
			ID:         "b39b61e4-00b1-4698-aed0-1928beb90abe",
			IPAddress:  &ipAddress2,
			Name:       "Interface 1/2",
			NetworkID:  &networkID2,
			SlotNumber: 2,
			Status:     "ACTIVE",
		},
	}

	loadBalancerPlanID := "bd12784a-c66e-4f13-9f72-5143d64762b6"
	name := "abcdefghijklmnopqrstuvwxyz"
	status := "PENDING_UPDATE"

	syslogServers := []load_balancer_syslog_servers.LoadBalancerSyslogServer{
		{
			ID:          "11001101-2edf-1844-1ff7-12ba5b7e566a",
			IPAddress:   "177.77.07.215",
			LogFacility: "LOCAL0",
			LogLevel:    "ALERT|INFO|ERROR",
			Name:        "syslog_server_main",
			PortNumber:  514,
			Status:      "ACTIVE",
		},
		{
			ID:          "22002202-2edf-1844-1ff7-12ba5b7e566a",
			IPAddress:   "177.77.07.211",
			LogFacility: "LOCAL1",
			LogLevel:    "ERROR",
			Name:        "syslog_server_backup_fst",
			PortNumber:  514,
			Status:      "ACTIVE",
		},
	}

	tenantID := "6a156ddf2ecd497ca786ff2da6df5aa8"
	userUsername := "user-read"

	options := load_balancers.UpdateOpts{
		DefaultGateway:     &defaultGateway,
		Description:        &description,
		LoadBalancerPlanID: loadBalancerPlanID,
		Name:               &name,
	}

	s, err := load_balancers.Update(fake.ServiceClient(), "ab49eb24-667f-4a4e-9421-b4d915bff416", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckEquals(t, adminUsername, s.AdminUsername)
	th.CheckEquals(t, availabilityZone, s.AvailabilityZone)
	th.CheckEquals(t, defaultGateway, *s.DefaultGateway)
	th.CheckEquals(t, description, s.Description)
	th.CheckEquals(t, id, s.ID)
	th.CheckDeepEquals(t, interfaces, s.Interfaces)
	th.CheckEquals(t, loadBalancerPlanID, s.LoadBalancerPlanID)
	th.CheckEquals(t, name, s.Name)
	th.CheckEquals(t, status, s.Status)
	th.CheckDeepEquals(t, syslogServers, s.SyslogServers)
	th.CheckEquals(t, tenantID, s.TenantID)
	th.CheckEquals(t, userUsername, s.UserUsername)
}

func TestDeleteLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers/ab49eb24-667f-4a4e-9421-b4d915bff416", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := load_balancers.Delete(fake.ServiceClient(), "ab49eb24-667f-4a4e-9421-b4d915bff416")
	th.AssertNoErr(t, res.Err)
}

func TestIDFromName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	expectedID := "5f3cae7c-58a5-4124-b622-9ca3cfbf2525"
	actualID, err := load_balancers.IDFromName(client, "Load Balancer 1")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedID, actualID)
}

func TestIDFromNameNoResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	_, err := load_balancers.IDFromName(client, "Load Balancer X")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}

func TestIDFromNameDuplicated(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponseDuplicatedNames)
	})

	client := fake.ServiceClient()

	_, err := load_balancers.IDFromName(client, "Load Balancer 1")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
