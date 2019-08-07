package testing

import (
	"fmt"
	"net/http"
	"testing"

	security "github.com/nttcom/eclcloud/ecl/security_order/v1/network_based_device_single"
	"github.com/nttcom/eclcloud/pagination"

	th "github.com/nttcom/eclcloud/testhelper"
	fakeclient "github.com/nttcom/eclcloud/testhelper/client"
)

func TestListDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/API/ScreenEventFGSDeviceGet", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := security.List(fakeclient.ServiceClient(), "UTM", nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := security.ExtractSingleDevices(page)
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

	th.Mux.HandleFunc("/API/ScreenEventFGSDeviceGet", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := security.List(fakeclient.ServiceClient(), "UTM", nil).AllPages()
	th.AssertNoErr(t, err)
	allDevices, err := security.ExtractSingleDevices(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allDevices))
}

func TestCreateDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/API/SoEntryFGS", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	gtHost := security.GtHostInCreate{
		AZGroup:       "zone1-groupb",
		LicenseKind:   "02",
		OperatingMode: "FW",
	}
	createOpts := security.CreateOpts{
		SOKind:   "A",
		Locale:   "ja",
		TenantID: "9ee80f2a926c49f88f166af47df4e9f5",
		GtHost:   [1]security.GtHostInCreate{gtHost},
	}

	actual, err := security.Create(fakeclient.ServiceClient(), "UTM", createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createResult, actual)
}

func TestUpdateDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryFGS"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, updateResponse)
	})

	gtHost := security.GtHostInUpdate{
		OperatingMode: "UTM",
		LicenseKind:   "08",
		HostName:      "CES11811",
	}
	updateOpts := security.UpdateOpts{
		SOKind:   "M",
		Locale:   "en",
		TenantID: "9ee80f2a926c49f88f166af47df4e9f5",
		GtHost:   [1]security.GtHostInUpdate{gtHost},
	}

	actual, err := security.Update(fakeclient.ServiceClient(), "CES11811", updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &updateResult, actual)
}

func TestDeleteDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryFGS"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, deleteRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, deleteResponse)
	})

	gtHost := security.GtHostInDelete{
		HostName: "CES11811",
	}
	deleteOpts := security.DeleteOpts{
		SOKind:   "D",
		TenantID: "9ee80f2a926c49f88f166af47df4e9f5",
		GtHost:   [1]security.GtHostInDelete{gtHost},
	}

	actual, err := security.Delete(fakeclient.ServiceClient(), "CES11811", deleteOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &updateResult, actual)
}
