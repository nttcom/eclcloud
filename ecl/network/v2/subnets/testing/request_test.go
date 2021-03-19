package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v2/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v2/ecl/network/v2/subnets"
	"github.com/nttcom/eclcloud/v2/pagination"
	th "github.com/nttcom/eclcloud/v2/testhelper"
)

func TestListSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	subnets.List(client, subnets.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := subnets.ExtractSubnets(page)
		if err != nil {
			t.Errorf("Failed to extrace ports: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedSubnetSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnets/ab49eb24-667f-4a4e-9421-b4d915bff416", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := subnets.Get(fake.ServiceClient(), "ab49eb24-667f-4a4e-9421-b4d915bff416").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &Subnet1, s)
}

func TestCreateSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnets", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := &subnets.CreateOpts{
		CIDR:      "192.168.10.0/24",
		NetworkID: "8f36b88a-443f-4d97-9751-34d34af9e782",
	}
	s, err := subnets.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &Subnet2, s)
}

func TestRequiredCreateOptsSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := subnets.Create(fake.ServiceClient(), subnets.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestUpdateSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnets/ab49eb24-667f-4a4e-9421-b4d915bff416", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	description := "UPDATED"
	dnsNameservers := []string{
		"0.0.0.0",
		"1.1.1.1",
	}
	enableDHCP := false
	gatewayIP := "192.168.10.1"
	hostRoutes := []subnets.HostRoute{{
		DestinationCIDR: "10.2.0.0/24",
		NextHop:         "10.1.0.10",
	}}
	name := "UPDATED"
	ntpServers := []string{
		"2.2.2.2",
	}
	tags := map[string]string{
		"updated": "true",
	}

	options := subnets.UpdateOpts{
		Description:    &description,
		DNSNameservers: dnsNameservers,
		EnableDHCP:     &enableDHCP,
		GatewayIP:      &gatewayIP,
		HostRoutes:     &hostRoutes,
		Name:           &name,
		NTPServers:     &ntpServers,
		Tags:           &tags,
	}

	s, err := subnets.Update(fake.ServiceClient(), "ab49eb24-667f-4a4e-9421-b4d915bff416", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckEquals(t, description, s.Description)
	th.CheckDeepEquals(t, dnsNameservers, s.DNSNameservers)
	th.CheckEquals(t, enableDHCP, s.EnableDHCP)
	th.CheckEquals(t, gatewayIP, s.GatewayIP)
	th.CheckDeepEquals(t, hostRoutes, s.HostRoutes)
	th.CheckEquals(t, name, s.Name)
	th.CheckDeepEquals(t, ntpServers, s.NTPServers)
	th.CheckDeepEquals(t, tags, s.Tags)
}

func TestDeleteSubnet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/subnets/ab49eb24-667f-4a4e-9421-b4d915bff416", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := subnets.Delete(fake.ServiceClient(), "ab49eb24-667f-4a4e-9421-b4d915bff416")
	th.AssertNoErr(t, res.Err)
}
