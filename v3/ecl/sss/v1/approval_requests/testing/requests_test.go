package testing

import (
	"fmt"
	"net/http"
	"testing"

	ar "github.com/nttcom/eclcloud/v3/ecl/sss/v1/approval_requests"
	"github.com/nttcom/eclcloud/v3/pagination"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestListApprovalRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/approval-requests", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := ar.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := ar.ExtractApprovalRequests(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedApprovalRequestsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListApprovalRequestAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/approval-requests", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := ar.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allRequests, err := ar.ExtractApprovalRequests(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allRequests))
}

func TestGetApprovalRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/approval-requests/%s", idApprovalRequest1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := ar.Get(fakeclient.ServiceClient(), idApprovalRequest1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &firstApprovalRequest, actual)
}

func TestUpdateApprovalRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/approval-requests/%s", idApprovalRequest1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusNoContent)
	})

	updateOpts := ar.UpdateOpts{
		Status: "approved",
	}

	res := ar.Update(fakeclient.ServiceClient(), idApprovalRequest1, updateOpts)
	th.AssertNoErr(t, res.Err)
}
