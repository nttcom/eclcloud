package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/ecl/security_portal/v1/devices"
	"github.com/nttcom/eclcloud/pagination"

	th "github.com/nttcom/eclcloud/testhelper"
	fakeclient "github.com/nttcom/eclcloud/testhelper/client"
)

func TestListDevices(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/ecl-api/devices", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := devices.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := devices.ExtractDevices(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedDevicesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListDeviceAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/ecl-api/devices", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := devices.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allDevices, err := devices.ExtractDevices(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allDevices))
}
