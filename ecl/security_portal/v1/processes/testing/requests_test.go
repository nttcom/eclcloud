package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/ecl/security_portal/v1/processes"

	th "github.com/nttcom/eclcloud/testhelper"
	fakeclient "github.com/nttcom/eclcloud/testhelper/client"
)

func TestGetProcess(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/ecl-api/process/%s/status", processID)
	fmt.Println(url)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := processes.Get(fakeclient.ServiceClient(), processID, nil).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &expectedProcess, actual)
}
