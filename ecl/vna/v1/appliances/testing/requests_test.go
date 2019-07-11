package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
	"github.com/nttcom/eclcloud/pagination"
	"github.com/nttcom/eclcloud/testhelper/client"

	th "github.com/nttcom/eclcloud/testhelper"
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

			fmt.Fprintf(w, listResponse)
		})

	client := ServiceClient()
	count := 0

	appliances.List(client, appliances.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := appliances.ExtractAppliances(page)
		if err != nil {
			t.Errorf("Failed to extract virtual network appliances: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, expectedAppliancesSlice, actual)

		return true, nil
	})

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

		fmt.Fprintf(w, getResponse)
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
			fmt.Fprintf(w, createResponse)
		})

	createOpts := appliances.CreateOpts{
		Name:                          "appliance_1",
		Description:                   "appliance_1_description",
		DefaultGateway:                "192.168.1.1",
		AvailabilityZone:              "zone1-groupb",
		VirtualNetworkAppliancePlanID: idVirtualNetworkAppliancePlan,
		Tags:                          map[string]string{"k1": "v1"},
		Interfaces: appliances.CreateOptsInterfaces{
			Interface1: appliances.CreateOptsInterface{
				Name:        "interface_1",
				Description: "interface_1_description",
				NetworkID:   "dummyNetworkID",
				Tags:        map[string]string{},
				FixedIPs: [1]appliances.CreateOptsFixedIP{
					appliances.CreateOptsFixedIP{
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

		fmt.Fprintf(w, updateMetadataResponse)
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
		Interfaces: appliances.UpdateMetadataInterfaces{
			Interface1: updateOptsInterface1,
			Interface2: interface{}(nil),
			Interface3: interface{}(nil),
			Interface4: interface{}(nil),
			Interface5: interface{}(nil),
			Interface6: interface{}(nil),
			Interface7: interface{}(nil),
			Interface8: interface{}(nil),
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

		fmt.Fprintf(w, updateNetworkIDAndFixedIPResponse)
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
		Interfaces: appliances.UpdateFixedIPInterfaces{
			Interface1: updateOptsInterface1,
			Interface2: interface{}(nil),
			Interface3: interface{}(nil),
			Interface4: interface{}(nil),
			Interface5: interface{}(nil),
			Interface6: interface{}(nil),
			Interface7: interface{}(nil),
			Interface8: interface{}(nil),
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
