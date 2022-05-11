package testing

import (
	"fmt"
	"net/http"
	"testing"

	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v2/host_based"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestGetHostBasedSecurity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/ScreenEventHBSOrderInfoGet"
	fmt.Println(url)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := security.Get(fakeclient.ServiceClient(), nil).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &expectedResult, actual)
}

func TestCreateHostBasedSecurity(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/API/SoEntryHBS", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	createOpts := security.CreateOpts{
		SOKind:              "N",
		TenantID:            "9ee80f2a926c49f88f166af47df4e9f5",
		Locale:              "ja",
		ServiceOrderService: "Managed Anti-Virus",
		MaxAgentValue:       1,
		MailAddress:         "terraform@example.com",
		DSMLang:             "ja",
		TimeZone:            "Asia/Tokyo",
	}

	actual, err := security.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createResult, actual)
}

func TestUpdateHostBasedSecurityTypeM1(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryHBS"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequestM1)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, updateResponseM1)
	})

	updateOpts := security.UpdateOpts{
		SOKind:      "M1",
		TenantID:    "9ee80f2a926c49f88f166af47df4e9f5",
		Locale:      "ja",
		MailAddress: "terraform@example.com",
	}

	serviceOrderService := "Managed Anti-Virus"
	updateOpts.ServiceOrderService = &serviceOrderService
	actual, err := security.Update(fakeclient.ServiceClient(), updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &updateResultM1, actual)
}

func TestUpdateHostBasedSecurityTypeM2(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryHBS"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequestM2)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, updateResponseM2)
	})

	updateOpts := security.UpdateOpts{
		SOKind:      "M2",
		TenantID:    "9ee80f2a926c49f88f166af47df4e9f5",
		Locale:      "ja",
		MailAddress: "terraform@example.com",
	}

	maxAgentValue := 10
	updateOpts.MaxAgentValue = &maxAgentValue
	actual, err := security.Update(fakeclient.ServiceClient(), updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &updateResultM2, actual)
}

func TestDeleteDevice(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := "/API/SoEntryHBS"
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, deleteRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, deleteResponse)
	})

	deleteOpts := security.DeleteOpts{
		SOKind:      "C",
		TenantID:    "9ee80f2a926c49f88f166af47df4e9f5",
		Locale:      "ja",
		MailAddress: "terraform@example.com",
	}

	actual, err := security.Delete(fakeclient.ServiceClient(), deleteOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &deleteResult, actual)
}
