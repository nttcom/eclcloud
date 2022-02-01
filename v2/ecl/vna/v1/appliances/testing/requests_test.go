package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/vna/v1/appliances"
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

func TestListAppliances(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/virtual_network_appliances",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0

	err := appliances.List(cli, appliances.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := appliances.ExtractAppliances(page)
		if err != nil {
			t.Errorf("Failed to extract virtual network appliances: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedAppliancesSlice, actual)

		return true, nil
	})

	if err != nil {
		t.Errorf("Failed to get virtual network appliance list: %v", err)
	}

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetAppliance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/virtual_network_appliances/%s", idAppliance1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, getResponse)
	})

	ap, err := appliances.Get(ServiceClient(), idAppliance1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &appliance1, ap)
}

func TestCreateAppliance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v1.0/virtual_network_appliances",
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

	createOpts := appliances.CreateOpts{
		Name:                          "appliance_1",
		Description:                   "appliance_1_description",
		DefaultGateway:                "192.168.1.1",
		AvailabilityZone:              "zone1-groupb",
		VirtualNetworkAppliancePlanID: idVirtualNetworkAppliancePlan,
		Tags:                          map[string]string{"k1": "v1"},
		Interfaces: &appliances.CreateOptsInterfaces{
			Interface1: &appliances.CreateOptsInterface{
				Name:        "interface_1",
				Description: "interface_1_description",
				NetworkID:   "dummyNetworkID",
				Tags:        map[string]string{},
				FixedIPs: &[]appliances.CreateOptsFixedIP{
					{
						IPAddress: "192.168.1.51",
					},
				},
			},
		},
	}
	ap, err := appliances.Create(ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, ap.OperationStatus, "COMPLETE")
	th.AssertDeepEquals(t, &appliance1, ap)
}

func TestDeleteAppliance(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/virtual_network_appliances/%s", idAppliance1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := appliances.Delete(ServiceClient(), idAppliance1)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateApplianceMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/virtual_network_appliances/%s", idAppliance1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateMetadataRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, updateMetadataResponse)
	})

	name := "appliance_1-update"
	description := "appliance_1_description-update"
	tags := map[string]string{"k1": "v1", "k2": "v2"}

	interface1Name := "interface_1"
	interface1Description := "interface_1_description"
	interface1Tags := map[string]string{"k1": "v1", "k2": "v2"}

	updateOptsInterface1 := appliances.UpdateMetadataInterface{
		Name:        &interface1Name,
		Description: &interface1Description,
		Tags:        &interface1Tags,
	}
	updateOpts := appliances.UpdateMetadataOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
		Interfaces: &appliances.UpdateMetadataInterfaces{
			Interface1: &updateOptsInterface1,
		},
	}
	ap, err := appliances.Update(
		ServiceClient(), idAppliance1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, ap.Name, "appliance_1-update")
	th.AssertEquals(t, ap.Description, "appliance_1_description-update")
	th.AssertEquals(t, ap.ID, idAppliance1)
}

func TestUpdateApplianceNetworkIDAndFixedIP(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/virtual_network_appliances/%s", idAppliance1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateNetworkIDAndFixedIPRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, updateNetworkIDAndFixedIPResponse)
	})

	networkID := "dummyNetworkID2"

	updateAddressInfo1 := appliances.UpdateFixedIPAddressInfo{
		IPAddress: "192.168.1.51",
	}

	updateAddressInfo2 := appliances.UpdateFixedIPAddressInfo{
		IPAddress: "192.168.1.52",
	}
	updateFixedIPs := []appliances.UpdateFixedIPAddressInfo{
		updateAddressInfo1,
		updateAddressInfo2,
	}

	updateOptsInterface1 := appliances.UpdateFixedIPInterface{
		NetworkID: &networkID,
		FixedIPs:  &updateFixedIPs,
	}
	updateOpts := appliances.UpdateFixedIPOpts{
		Interfaces: &appliances.UpdateFixedIPInterfaces{
			Interface1: &updateOptsInterface1,
		},
	}
	ap, err := appliances.Update(
		ServiceClient(), idAppliance1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, ap.Interfaces.Interface1.NetworkID, "dummyNetworkID2")
	th.AssertEquals(t, ap.Interfaces.Interface1.FixedIPs[0].IPAddress, "192.168.1.51")
	th.AssertEquals(t, ap.Interfaces.Interface1.FixedIPs[1].IPAddress, "192.168.1.52")
	th.AssertEquals(t, ap.ID, idAppliance1)
}
func TestUpdateApplianceAllowedAddressPairs(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/v1.0/virtual_network_appliances/%s", idAppliance1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, updateAllowedAddressPairsRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprint(w, updateAllowedAddressPairsResponse)
	})

	mac1 := "aa:bb:cc:dd:ee:f1"
	type1 := "vrrp"

	var vrid1 interface{} = 123

	UpdateAllowedAddressPairAddressInfo1 := appliances.UpdateAllowedAddressPairAddressInfo{
		IPAddress:  "1.1.1.1",
		MACAddress: &mac1,
		Type:       &type1,
		VRID:       &vrid1,
	}

	mac2 := "aa:bb:cc:dd:ee:f2"
	type2 := ""

	var vrid2 interface{} = nil

	UpdateAllowedAddressPairAddressInfo2 := appliances.UpdateAllowedAddressPairAddressInfo{
		IPAddress:  "2.2.2.2",
		MACAddress: &mac2,
		Type:       &type2,
		VRID:       &vrid2,
	}

	updateAllowedAddressPairs := []appliances.UpdateAllowedAddressPairAddressInfo{
		UpdateAllowedAddressPairAddressInfo1,
		UpdateAllowedAddressPairAddressInfo2,
	}

	updateOptsInterface1 := appliances.UpdateAllowedAddressPairInterface{
		AllowedAddressPairs: &updateAllowedAddressPairs,
	}

	updateOpts := appliances.UpdateAllowedAddressPairOpts{
		Interfaces: &appliances.UpdateAllowedAddressPairInterfaces{
			Interface1: &updateOptsInterface1,
		},
	}
	ap, err := appliances.Update(
		ServiceClient(), idAppliance1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[0].IPAddress, "1.1.1.1")
	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[0].MACAddress, "aa:bb:cc:dd:ee:f1")
	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[0].Type, "vrrp")
	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[0].VRID, float64(123))

	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[1].IPAddress, "2.2.2.2")
	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[1].MACAddress, "aa:bb:cc:dd:ee:f2")
	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[1].Type, "")
	th.AssertEquals(t, ap.Interfaces.Interface1.AllowedAddressPairs[1].VRID, interface{}(nil))

	th.AssertEquals(t, ap.ID, idAppliance1)
}
