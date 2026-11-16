package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/security_portal/v3/device_interfaces"
	"github.com/nttcom/eclcloud/v3/pagination"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestListDeviceInterfaces(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/ecl-api/devices/%s/interfaces", deviceUUID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := device_interfaces.List(fakeclient.ServiceClient(), deviceUUID, nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := device_interfaces.ExtractDeviceInterfaces(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedDeviceInterfacesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListDeviceInterfaceAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/ecl-api/devices/%s/interfaces", deviceUUID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := device_interfaces.List(fakeclient.ServiceClient(), deviceUUID, nil).AllPages()
	th.AssertNoErr(t, err)
	allDeviceInterfaces, err := device_interfaces.ExtractDeviceInterfaces(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allDeviceInterfaces))
}
