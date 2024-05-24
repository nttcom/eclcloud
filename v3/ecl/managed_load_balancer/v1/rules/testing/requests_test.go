package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/rules"
	"github.com/nttcom/eclcloud/v3/pagination"
	"github.com/nttcom/eclcloud/v3/testhelper/client"

	th "github.com/nttcom/eclcloud/v3/testhelper"
)

const TokenID = client.TokenID

func ServiceClient() *eclcloud.ServiceClient {
	sc := client.ServiceClient()
	sc.ResourceBase = sc.Endpoint + "v1.0/"

	return sc
}

func TestListRules(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/rules",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, listResponse)
		})

	cli := ServiceClient()
	count := 0
	listOpts := rules.ListOpts{}

	err := rules.List(cli, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := rules.ExtractRules(page)
		if err != nil {
			t.Errorf("Failed to extract rules: %v", err)

			return false, err
		}

		th.CheckDeepEquals(t, listResult(), actual)

		return true, nil
	})

	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestCreateRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/v1.0/rules",
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
	condition := rules.CreateOptsCondition{
		PathPatterns: []string{"^/statics/"},
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	createOpts := rules.CreateOpts{
		Name:          "rule",
		Description:   "description",
		Tags:          tags,
		Priority:      1,
		TargetGroupID: "29527a3c-9e5d-48b7-868f-6442c7d21a95",
		PolicyID:      "fcb520e5-858d-4f9f-bc6c-7bd225fe7cf4",
		Conditions:    &condition,
	}

	actual, err := rules.Create(cli, createOpts).Extract()

	th.CheckDeepEquals(t, createResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showResponse)
		})

	cli := ServiceClient()
	showOpts := rules.ShowOpts{}

	actual, err := rules.Show(cli, id, showOpts).Extract()

	th.CheckDeepEquals(t, showResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s", id),
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

	name := "rule"
	description := "description"

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)

	th.AssertNoErr(t, err)

	updateOpts := rules.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := rules.Update(cli, id, updateOpts).Extract()

	th.CheckDeepEquals(t, updateResult(), actual)
	th.AssertNoErr(t, err)
}

func TestDeleteRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := rules.Delete(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}

func TestCreateStagedRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s/staged", id),
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
	condition := rules.CreateStagedOptsCondition{
		PathPatterns: []string{"^/statics/"},
	}
	createStagedOpts := rules.CreateStagedOpts{
		Priority:      1,
		TargetGroupID: "29527a3c-9e5d-48b7-868f-6442c7d21a95",
		Conditions:    &condition,
	}

	actual, err := rules.CreateStaged(cli, id, createStagedOpts).Extract()

	th.CheckDeepEquals(t, createStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestShowStagedRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprint(w, showStagedResponse)
		})

	cli := ServiceClient()
	actual, err := rules.ShowStaged(cli, id).Extract()

	th.CheckDeepEquals(t, showStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestUpdateStagedRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s/staged", id),
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

	condition := rules.UpdateStagedOptsCondition{
		PathPatterns: &[]string{"^/statics/"},
	}

	priority := 1
	targetGroupID := "29527a3c-9e5d-48b7-868f-6442c7d21a95"
	updateStagedOpts := rules.UpdateStagedOpts{
		Priority:      &priority,
		TargetGroupID: &targetGroupID,
		Conditions:    &condition,
	}

	actual, err := rules.UpdateStaged(cli, id, updateStagedOpts).Extract()

	th.CheckDeepEquals(t, updateStagedResult(), actual)
	th.AssertNoErr(t, err)
}

func TestCancelStagedRule(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		fmt.Sprintf("/v1.0/rules/%s/staged", id),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "DELETE")
			th.TestHeader(t, r, "X-Auth-Token", TokenID)

			w.WriteHeader(http.StatusNoContent)
		})

	cli := ServiceClient()

	err := rules.CancelStaged(cli, id).ExtractErr()

	th.AssertNoErr(t, err)
}
