package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v4/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/ports"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
)

func TestListPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	ports.List(client, ports.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ports.ExtractPorts(page)
		if err != nil {
			t.Errorf("Failed to extrace ports: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedPortSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/ports/ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	p, err := ports.Get(fake.ServiceClient(), "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &Port2, p)
}

func TestCreatePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/ports", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	asu := true

	options := &ports.CreateOpts{
		AdminStateUp: &asu,
		AllowedAddressPairs: []ports.AddressPair{{
			IPAddress:  "192.168.2.100",
			MACAddress: "00:00:5e:00:01:01",
		}},
		FixedIPs: []ports.IP{{
			IPAddress: "192.168.2.30",
			SubnetID:  "ab49eb24-667f-4a4e-9421-b4d915bff416",
		}},
		Name:             "port_12",
		NetworkID:        "8f36b88a-443f-4d97-9751-34d34af9e782",
		TenantID:         "dcb2d589c0c646d0bad45c0cf9f90cf1",
		SegmentationType: "flat",
	}
	p, err := ports.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &Port2, p)
}

func TestRequiredCreateOptsPort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := ports.Create(fake.ServiceClient(), ports.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}
func TestUpdatePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/ports/ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	aap := []ports.AddressPair{{
		IPAddress:  "192.168.2.100",
		MACAddress: "00:00:5e:00:01:01",
	}, {
		IPAddress:  "192.168.2.255",
		MACAddress: "26:8d:42:f6:c2:c4",
	}}
	description := "UPDATED"
	deviceID := "b269b8c0-1a42-4464-9314-4396e51e5107"
	deviceOwner := "UPDATED"
	fip := []ports.IP{{
		IPAddress: "192.168.2.30",
		SubnetID:  "ab49eb24-667f-4a4e-9421-b4d915bff416",
	}, {
		IPAddress: "192.168.2.31",
		SubnetID:  "ab49eb24-667f-4a4e-9421-b4d915bff417",
	}}
	name := "UPDATED"
	segmentationID := 2
	segmentationType := "vlan"
	tags := map[string]string{"some-key": "UPDATED"}

	options := ports.UpdateOpts{
		AllowedAddressPairs: &aap,
		Description:         &description,
		DeviceID:            &deviceID,
		DeviceOwner:         &deviceOwner,
		FixedIPs:            fip,
		Name:                &name,
		SegmentationID:      &segmentationID,
		SegmentationType:    &segmentationType,
		Tags:                &tags,
	}
	p, err := ports.Update(fake.ServiceClient(), "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, aap, p.AllowedAddressPairs)
	th.CheckEquals(t, description, p.Description)
	th.CheckEquals(t, deviceID, p.DeviceID)
	th.CheckEquals(t, deviceOwner, p.DeviceOwner)
	th.CheckDeepEquals(t, fip, p.FixedIPs)
	th.CheckEquals(t, name, p.Name)
	th.CheckEquals(t, segmentationID, p.SegmentationID)
	th.CheckEquals(t, segmentationType, p.SegmentationType)
	th.CheckDeepEquals(t, tags, p.Tags)
}

func TestRemoveAllowedAddressPairs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/ports/ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, RemoveAllowedAddressPairsRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, RemoveAllowedAddressPairsResponse)
	})

	name := "new_port_name"
	options := ports.UpdateOpts{
		Name:                &name,
		AllowedAddressPairs: &[]ports.AddressPair{},
	}

	s, err := ports.Update(fake.ServiceClient(), "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730", options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, s.Name, "new_port_name")
	th.AssertDeepEquals(t, s.AllowedAddressPairs, []ports.AddressPair{})
}

func TestDeletePort(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/ports/ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := ports.Delete(fake.ServiceClient(), "ac57c5c9-aaf4-4ffc-b8b8-f1ef84656730")
	th.AssertNoErr(t, res.Err)
}
