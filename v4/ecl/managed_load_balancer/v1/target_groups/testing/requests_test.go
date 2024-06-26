package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/target_groups"
	"github.com/nttcom/eclcloud/v4/pagination"
	"github.com/nttcom/eclcloud/v4/testhelper/client"

	th "github.com/nttcom/eclcloud/v4/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"

	return sc
}

func TestListTargetGroups(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/target_groups",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := target_groups.ListOpts{}

	err := target_groups.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := target_groups.ExtractTargetGroups(page)
		if err != nil {
			t.Errorf("Failed to extract target groups: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestCreateTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/target_groups",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, createResponse)
		})

	cli := ServiceClient()
	member1 := target_groups.CreateOptsMember{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	createOpts := target_groups.CreateOpts{
		Name:           "target_group",
		Description:    "description",
		Tags:           tags,
		LoadBalancerID: "67fea379-cff0-4191-9175-de7d6941a040",
		Members:        &[]target_groups.CreateOptsMember{member1},
	}

	actual, err := target_groups.Create(cli, createOpts).Extract()

	th.CheckDeepEquals(t, createResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()
	showOpts := target_groups.ShowOpts{}

	actual, err := target_groups.Show(cli, id, showOpts).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, updateResponse)
		})

	cli := ServiceClient()

	name := "target_group"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	updateOpts := target_groups.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := target_groups.Update(cli, id, updateOpts).Extract()

	th.CheckDeepEquals(t, updateResult(), actual)
	th.AssertNoErr(t, err)
}

func TestDeleteTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := target_groups.Delete(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCreateStagedTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, createStagedRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, createStagedResponse)
		})

	cli := ServiceClient()
	member1 := target_groups.CreateStagedOptsMember{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}
	createStagedOpts := target_groups.CreateStagedOpts{
		Members: &[]target_groups.CreateStagedOptsMember{member1},
	}

	actual, err := target_groups.CreateStaged(cli, id, createStagedOpts).Extract()

	th.CheckDeepEquals(t, createStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowStagedTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showStagedResponse)
		})

	cli := ServiceClient()
	actual, err := target_groups.ShowStaged(cli, id).Extract()

	th.CheckDeepEquals(t, showStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateStagedTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PATCH")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, updateStagedRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, updateStagedResponse)
		})

	cli := ServiceClient()

	member1IPAddress := "192.168.0.7"
	member1Port := 80
	member1Weight := 1
	member1 := target_groups.UpdateStagedOptsMember{
		IPAddress: &member1IPAddress,
		Port:      &member1Port,
		Weight:    &member1Weight,
	}

	updateStagedOpts := target_groups.UpdateStagedOpts{
		Members: &[]target_groups.UpdateStagedOptsMember{member1},
	}

	actual, err := target_groups.UpdateStaged(cli, id, updateStagedOpts).Extract()

	th.CheckDeepEquals(t, updateStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestCancelStagedTargetGroup(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/target_groups/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := target_groups.CancelStaged(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}
