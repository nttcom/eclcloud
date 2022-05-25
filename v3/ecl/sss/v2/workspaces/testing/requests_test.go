package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/sss/v2/workspaces"
	"github.com/nttcom/eclcloud/v3/pagination"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestListWorkspace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/workspaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := workspaces.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := workspaces.ExtractWorkspaces(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedWorkspacesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListWorkspaceAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/workspaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := workspaces.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allZones, err := workspaces.ExtractWorkspaces(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allZones))
}

func TestGetWorkspace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/workspaces/%s", workspaceID1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := workspaces.Get(fakeclient.ServiceClient(), workspaceID1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &getResponseStruct, actual)
}

func TestCreateWorkspace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/workspaces", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	createOpts := workspaces.CreateOpts{
		WorkspaceName: "sample_workspace",
		Description:   "sample workspace",
		ContractID:    "econ0000000001",
	}

	actual, err := workspaces.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createdWorkspace, actual)
}

func TestUpdateWorkspace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/workspaces/%s", workspaceID1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusNoContent)
	})

	description := "updated workspace"

	updateOpts := workspaces.UpdateOpts{
		Description: &description,
	}

	res := workspaces.Update(fakeclient.ServiceClient(), workspaceID1, updateOpts)
	th.AssertNoErr(t, res.Err)
}

func TestDeleteWorkspace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/workspaces/%s", workspaceID1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	res := workspaces.Delete(fakeclient.ServiceClient(), workspaceID1)
	th.AssertNoErr(t, res.Err)
}
