package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v4/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/security_groups"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
)

func TestListSecurityGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	security_groups.List(client, security_groups.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := security_groups.ExtractSecurityGroups(page)
		if err != nil {
			t.Errorf("Failed to extract security groups: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedSecurityGroupSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetSecurityGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-groups/c0e1482e-2e3c-497e-8964-e4f818071700", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	sg, err := security_groups.Get(fake.ServiceClient(), "c0e1482e-2e3c-497e-8964-e4f818071700").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SecurityGroup2, sg)
}

func TestCreateSecurityGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-groups", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	options := &security_groups.CreateOpts{
		Description: "Test security group",
		Name:        "test-sg",
		Tags: map[string]string{
			"env": "test",
		},
		TenantID: "6f70656e737461636b20342065766572",
	}
	sg, err := security_groups.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	expected := &security_groups.SecurityGroup{
		Description:        "Test security group",
		ID:                 "c0e1482e-2e3c-497e-8964-e4f818071700",
		Name:               "test-sg",
		SecurityGroupRules: []security_groups.SecurityGroupRule{},
		Status:             "ACTIVE",
		Tags: map[string]string{
			"env": "test",
		},
		TenantID: "6f70656e737461636b20342065766572",
	}
	th.AssertDeepEquals(t, expected, sg)
}

func TestUpdateSecurityGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-groups/c0e1482e-2e3c-497e-8964-e4f818071700", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	description := "Updated security group"
	name := "updated-sg"
	tags := map[string]string{
		"env": "production",
	}

	options := &security_groups.UpdateOpts{
		Description: &description,
		Name:        &name,
		Tags:        &tags,
	}
	sg, err := security_groups.Update(fake.ServiceClient(), "c0e1482e-2e3c-497e-8964-e4f818071700", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckEquals(t, description, sg.Description)
	th.CheckEquals(t, name, sg.Name)
	th.CheckDeepEquals(t, tags, sg.Tags)
}

func TestDeleteSecurityGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/security-groups/c0e1482e-2e3c-497e-8964-e4f818071700", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := security_groups.Delete(fake.ServiceClient(), "c0e1482e-2e3c-497e-8964-e4f818071700")
	th.AssertNoErr(t, res.Err)
}
