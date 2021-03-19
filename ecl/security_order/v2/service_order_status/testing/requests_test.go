package testing

import (
	"fmt"
	"net/http"
	"testing"

	order "github.com/nttcom/eclcloud/v2/ecl/security_order/v2/service_order_status"

	th "github.com/nttcom/eclcloud/v2/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v2/testhelper/client"
)

func TestGetOrder(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/ScreenEventFGSOrderProgressRate"
	fmt.Println(url)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := order.Get(fakeclient.ServiceClient(), "UTM", nil).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &expectedResult, actual)
}
