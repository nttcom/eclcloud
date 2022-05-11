package testing

import (
	"fmt"
	"net/http"
	"testing"

	az "github.com/nttcom/eclcloud/v2/ecl/compute/v2/extensions/availabilityzones"
	th "github.com/nttcom/eclcloud/v2/testhelper"

	fakeclient "github.com/nttcom/eclcloud/v2/testhelper/client"
)

func TestListAvailabilityZone(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/os-availability-zone", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	allPages, err := az.List(fakeclient.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)

	actual, err := az.ExtractAvailabilityZones(allPages)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, azResult, actual)
}
