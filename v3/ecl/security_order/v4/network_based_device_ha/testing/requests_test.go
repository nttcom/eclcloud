package testing

import (
	"fmt"
	"net/http"
	"testing"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v3/network_based_device_ha"
	"github.com/nttcom/eclcloud/v3/pagination"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestListDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/API/ScreenEventFGHADeviceGet", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := security.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := security.ExtractHADevices(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedDevicesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListDeviceZoneAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/API/ScreenEventFGHADeviceGet", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := security.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allDevices, err := security.ExtractHADevices(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allDevices))
}

func TestCreateDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/API/SoEntryFGHA", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	gtHost1 := security.GtHostInCreate{
		AZGroup:          "zone1-groupa",
		LicenseKind:      "02",
		OperatingMode:    "FW_HA",
		HALink1IPAddress: "192.168.1.3",
		HALink1NetworkID: "c5b1b0a8-45a3-4c99-b808-84e7c13e557f",
		HALink1SubnetID:  "9a2116e2-52be-439c-9587-506a1a5d288d",
		HALink2IPAddress: "192.168.2.3",
		HALink2NetworkID: "a8df4d5f-8752-4574-a255-dc749acd458f",
		HALink2SubnetID:  "a2ff5669-8422-421c-bb85-a6d691ecf223",
	}

	gtHost2 := security.GtHostInCreate{
		AZGroup:          "zone1-groupb",
		LicenseKind:      "02",
		OperatingMode:    "FW_HA",
		HALink1IPAddress: "192.168.1.4",
		HALink1NetworkID: "c5b1b0a8-45a3-4c99-b808-84e7c13e557f",
		HALink1SubnetID:  "9a2116e2-52be-439c-9587-506a1a5d288d",
		HALink2IPAddress: "192.168.2.4",
		HALink2NetworkID: "a8df4d5f-8752-4574-a255-dc749acd458f",
		HALink2SubnetID:  "a2ff5669-8422-421c-bb85-a6d691ecf223",
	}

	createOpts := security.CreateOpts{
		SOKind:   "AH",
		Locale:   "ja",
		TenantID: "9ee80f2a926c49f88f166af47df4e9f5",
		GtHost:   [2]security.GtHostInCreate{gtHost1, gtHost2},
	}

	actual, err := security.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createResult, actual)
}

func TestUpdateDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryFGHA"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, updateResponse)
	})

	gtHost1 := security.GtHostInUpdate{
		OperatingMode: "UTM_HA",
		LicenseKind:   "08",
		HostName:      "CES11811",
	}

	gtHost2 := security.GtHostInUpdate{
		OperatingMode: "UTM_HA",
		LicenseKind:   "08",
		HostName:      "CES11812",
	}

	updateOpts := security.UpdateOpts{
		SOKind:   "MH",
		Locale:   "en",
		TenantID: "9ee80f2a926c49f88f166af47df4e9f5",
		GtHost:   [2]security.GtHostInUpdate{gtHost1, gtHost2},
	}

	actual, err := security.Update(fakeclient.ServiceClient(), updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &updateResult, actual)
}

func TestDeleteDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryFGHA"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, deleteRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, deleteResponse)
	})

	gtHost1 := security.GtHostInDelete{
		HostName: "CES11811",
	}

	gtHost2 := security.GtHostInDelete{
		HostName: "CES11812",
	}

	deleteOpts := security.DeleteOpts{
		SOKind:   "DH",
		TenantID: "9ee80f2a926c49f88f166af47df4e9f5",
		GtHost:   [2]security.GtHostInDelete{gtHost1, gtHost2},
	}

	actual, err := security.Delete(fakeclient.ServiceClient(), "CES11811", deleteOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &deleteResult, actual)
}
