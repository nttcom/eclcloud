package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v3/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_group_rules"
	"github.com/nttcom/eclcloud/v3/pagination"
	th "github.com/nttcom/eclcloud/v3/testhelper"
)

func TestListSecurityGroupRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	security_group_rules.List(client, security_group_rules.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := security_group_rules.ExtractSecurityGroupRules(page)
		if err != nil {
			t.Errorf("Failed to extract security group rules: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedSecurityGroupRuleSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetSecurityGroupRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-group-rules/2bc0accf-312e-429a-956e-e4407625eb62", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	rule, err := security_group_rules.Get(fake.ServiceClient(), "2bc0accf-312e-429a-956e-e4407625eb62").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SecurityGroupRule1, rule)
}

func TestCreateSecurityGroupRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	portRangeMax := 22
	portRangeMin := 22
	remoteIPPrefix := "0.0.0.0/0"

	options := &security_group_rules.CreateOpts{
		Description:     "Allow SSH",
		Direction:       "ingress",
		Ethertype:       "IPv4",
		PortRangeMax:    &portRangeMax,
		PortRangeMin:    &portRangeMin,
		Protocol:        "tcp",
		RemoteIPPrefix:  &remoteIPPrefix,
		SecurityGroupID: "a7734e61-b545-452d-a3cd-0189cbd9747a",
		TenantID:        "e4f50856753b4dc6afee5fa6b9b6c550",
	}
	rule, err := security_group_rules.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &SecurityGroupRule1, rule)
}

func TestCreateSecurityGroupRuleWithRemoteGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-group-rules", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequestWithRemoteGroup)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponseWithRemoteGroup)
	})

	remoteGroupID := "a7734e61-b545-452d-a3cd-0189cbd9747a"

	options := &security_group_rules.CreateOpts{
		Description:     "Allow from same group",
		Direction:       "ingress",
		Ethertype:       "IPv4",
		Protocol:        "tcp",
		RemoteGroupID:   &remoteGroupID,
		SecurityGroupID: "a7734e61-b545-452d-a3cd-0189cbd9747a",
	}
	rule, err := security_group_rules.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	expected := &security_group_rules.SecurityGroupRule{
		Description:     "Allow from same group",
		Direction:       "ingress",
		Ethertype:       "IPv4",
		ID:              "rule-with-remote-group-id",
		PortRangeMax:    nil,
		PortRangeMin:    nil,
		Protocol:        "tcp",
		RemoteGroupID:   &remoteGroupID,
		RemoteIPPrefix:  nil,
		SecurityGroupID: "a7734e61-b545-452d-a3cd-0189cbd9747a",
		TenantID:        "e4f50856753b4dc6afee5fa6b9b6c550",
	}
	th.AssertDeepEquals(t, expected, rule)
}

func TestRequiredCreateOptsSecurityGroupRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	// Test missing direction
	res := security_group_rules.Create(fake.ServiceClient(), security_group_rules.CreateOpts{
		SecurityGroupID: "a7734e61-b545-452d-a3cd-0189cbd9747a",
	})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}

	// Test missing security_group_id
	res = security_group_rules.Create(fake.ServiceClient(), security_group_rules.CreateOpts{
		Direction: "ingress",
	})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestDeleteSecurityGroupRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-group-rules/2bc0accf-312e-429a-956e-e4407625eb62", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := security_group_rules.Delete(fake.ServiceClient(), "2bc0accf-312e-429a-956e-e4407625eb62")
	th.AssertNoErr(t, res.Err)
}
