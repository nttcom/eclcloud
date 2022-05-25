package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/sss/v2/workspace_roles"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestCreateWorkspaceRole(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/workspace-roles", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	createOpts := workspace_roles.CreateOpts{
		UserID:      "ecid1234567891",
		WorkspaceID: "ws0000000001",
	}

	actual, err := workspace_roles.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createdWorkspaceRole, actual)
}

func TestDeleteWorkspace(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/workspace-roles/workspaces/%s/users/%s", workspaceID, userID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	res := workspace_roles.Delete(fakeclient.ServiceClient(), workspaceID, userID)
	th.AssertNoErr(t, res.Err)
}
