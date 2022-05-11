package testing

import (
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/compute/v2/extensions/startstop"
	th "github.com/nttcom/eclcloud/v3/testhelper"
	"github.com/nttcom/eclcloud/v3/testhelper/client"
)

const serverID = "645b787e-7fbb-4111-a217-63a2882930f2"

func TestServerStart(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/servers/" + serverID + "/action"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"os-start": null}`)
		w.WriteHeader(http.StatusAccepted)
	})

	err := startstop.Start(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestServerStop(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/servers/" + serverID + "/action"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, `{"os-stop": null}`)
		w.WriteHeader(http.StatusAccepted)
	})

	err := startstop.Stop(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
