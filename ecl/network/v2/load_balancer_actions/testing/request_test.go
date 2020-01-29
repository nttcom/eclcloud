package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_actions"
	th "github.com/nttcom/eclcloud/testhelper"
)

func TestRebootLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers/6e9c7745-61f2-491f-9689-add8c5fc4b9a/reboot", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, RebootRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

	})

	options := load_balancer_actions.RebootOpts	{
		Type:                        "HARD",
	}
	res := load_balancer_actions.Reboot(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a", options)
	th.AssertNoErr(t, res.Err)
}

func TestRequiredRebootOptsLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := load_balancer_actions.Reboot(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a", load_balancer_actions.RebootOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestResetPasswordLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancers/6e9c7745-61f2-491f-9689-add8c5fc4b9a/reset_password", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, ResetPasswordRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ResetPasswordResponse)
	})

	options := load_balancer_actions.ResetPasswordOpts {
		Username: "user-read",
	}
	s, err := load_balancer_actions.ResetPassword(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a", options).ExtractResetPassword()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &LoadBalancerActionResetPassword, s)
}

func TestRequiredResetPasswordOptsLoadBalancer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := load_balancer_actions.ResetPassword(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a", load_balancer_actions.ResetPasswordOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

