package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connection_requests"
	th "github.com/nttcom/eclcloud/testhelper"
	"github.com/nttcom/eclcloud/testhelper/client"
)

// ListResult provides a single page of tenant_connection_request results.
const ListResult = `
{
  "tenant_connection_requests": [
    {
      "id": "5fbcc350-bd33-11e7-afb6-0050569c850d",
      "name": "test_name1",
      "description": "test_desc1",
      "tags": {
		"test_tags1": "test1"
	},
      "tenant_id": "c7f3a68a73e845d4ba6a42fb80fce03f",
      "name_other": "",
      "description_other": "",
      "tags_other": {},
      "tenant_id_other": "7e91b19b9baa423793ee74a8e1ff2be1",
      "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
      "status": "registering",
      "approval_request_id": "req0000010454"
    },
    {
      "id": "90381138-b572-11e7-9391-0050569c850d",
      "name": "created_name",
      "description": "created_desc",
      "tags": {
		"test_tags2": "test2"
	},
      "tenant_id": "7e91b19b9baa423793ee74a8e1ff2be1",
      "name_other": "",
      "description_other": "",
      "tags_other": {
		"test_tags_other2": "test2"
	},
      "tenant_id_other": "c7f3a68a73e845d4ba6a42fb80fce03f",
      "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
      "status": "registered",
      "approval_request_id": "req0000010363"
    }
  ]
}
`

// GetResult provides a Get result.
const GetResult = `
{
  "tenant_connection_request": {
    "id": "90381138-b572-11e7-9391-0050569c850d",
    "name": "created_name",
    "description": "created_desc",
    "tags": {
		"test_tags2":"test2"
	},
    "tenant_id": "7e91b19b9baa423793ee74a8e1ff2be1",
    "name_other": "",
    "description_other": "",
    "tags_other": {
		"test_tags_other2":"test2"
	},
    "tenant_id_other": "c7f3a68a73e845d4ba6a42fb80fce03f",
    "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
    "status": "registered",
    "approval_request_id": "req0000010363"
  }
}
`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
  "tenant_connection_request": {
    "tenant_id_other": "7e91b19b9baa423793ee74a8e1ff2be1",
    "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
    "name": "test_name1",
    "description": "test_desc1",
    "tags": {
		"test_tags1": "test1"
	}
  }
}
`

// CreateResponse provides the output from a Create request.
const CreateResponse = `
{
  "tenant_connection_request": {
    "id": "5fbcc350-bd33-11e7-afb6-0050569c850d",
    "name": "test_name1",
    "description": "test_desc1",
    "tags": {
		"test_tags1": "test1"
	},
    "tenant_id": "c7f3a68a73e845d4ba6a42fb80fce03f",
    "name_other": "",
    "description_other": "",
    "tenant_id_other": "7e91b19b9baa423793ee74a8e1ff2be1",
    "tags_other": {},
    "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
    "status": "registering",
    "approval_request_id": "req0000010454"
  }
}
`

// UpdateRequest provides the input to as Update request.
const UpdateRequest = `
{
  "tenant_connection_request":{
    "name": "updated_name",
    "description": "updated_desc",
    "tags": {
		"k2":"v2"
	},
    "name_other": "updeted_name_other",
    "description_other": "updated_desc_other",
    "tags_other": {
		"k3":"v3"
	}
  }
}
`

// UpdateResult provides an update result.
const UpdateResult = `
{
  "tenant_connection_request": {
    "id": "90381138-b572-11e7-9391-0050569c850d",
    "name": "updated_name",
    "description": "updated_desc",
    "tags": {
		"k2": "v2"
	},
    "tenant_id": "7e91b19b9baa423793ee74a8e1ff2be1",
    "name_other": "updated_name_other",
    "description_other": "updated_desc_other",
    "tenant_id_other": "c7f3a68a73e845d4ba6a42fb80fce03f",
    "tags_other": {
		"k3": "v3"
	},
    "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
    "status": "registered",
    "approval_request_id": "req0000010363"
  }
}
`

// FirstTenantConnectionRequest is the first tenant_connection_request in the List request.
var FirstTenantConnectionRequest = tenant_connection_requests.TenantConnectionRequest{
	ID:                "5fbcc350-bd33-11e7-afb6-0050569c850d",
	Status:            "registering",
	KeystoneUserID:    "",
	Name:              "test_name1",
	Description:       "test_desc1",
	Tags:              map[string]string{"test_tags1": "test1"},
	TenantID:          "c7f3a68a73e845d4ba6a42fb80fce03f",
	NameOther:         "",
	DescriptionOther:  "",
	TagsOther:         map[string]string{},
	TenantIDOther:     "7e91b19b9baa423793ee74a8e1ff2be1",
	NetworkID:         "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
	ApprovalRequestID: "req0000010454",
}

// SecondTenantConnectionRequest is the second tenant_connection_request in the List request.
var SecondTenantConnectionRequest = tenant_connection_requests.TenantConnectionRequest{
	ID:                "90381138-b572-11e7-9391-0050569c850d",
	Status:            "registered",
	KeystoneUserID:    "",
	Name:              "created_name",
	Description:       "created_desc",
	Tags:              map[string]string{"test_tags2": "test2"},
	TenantID:          "7e91b19b9baa423793ee74a8e1ff2be1",
	NameOther:         "",
	DescriptionOther:  "",
	TagsOther:         map[string]string{"test_tags_other2": "test2"},
	TenantIDOther:     "c7f3a68a73e845d4ba6a42fb80fce03f",
	NetworkID:         "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
	ApprovalRequestID: "req0000010363",
}

// SecondTenantConnectionRequestUpdated is how second tenant_connection_request should look after an Update.
var SecondTenantConnectionRequestUpdated = tenant_connection_requests.TenantConnectionRequest{
	ID:                "90381138-b572-11e7-9391-0050569c850d",
	Status:            "registered",
	KeystoneUserID:    "",
	Name:              "updated_name",
	Description:       "updated_desc",
	Tags:              map[string]string{"k2": "v2"},
	TenantID:          "7e91b19b9baa423793ee74a8e1ff2be1",
	NameOther:         "updated_name_other",
	DescriptionOther:  "updated_desc_other",
	TagsOther:         map[string]string{"k3": "v3"},
	TenantIDOther:     "c7f3a68a73e845d4ba6a42fb80fce03f",
	NetworkID:         "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
	ApprovalRequestID: "req0000010363",
}

// ExpectedTenantConnectionRequestsSlice is the slice of tenant_connection_request expected to be returned from ListResult.
var ExpectedTenantConnectionRequestsSlice = []tenant_connection_requests.TenantConnectionRequest{FirstTenantConnectionRequest, SecondTenantConnectionRequest}

// HandleListTenantConnectionRequestsSuccessfully creates an HTTP handler at `/tenant_connection_requests` on the
// test handler mux that responds with a list of two tenant_connection_requests.
func HandleListTenantConnectionRequestsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/tenant_connection_requests", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, ListResult)
	})
}

// HandleGetTenantConnectionRequestSuccessfully creates an HTTP handler at `/tenant_connection_requests` on the
// test handler mux that responds with a single tenant_connection_request.
func HandleGetTenantConnectionRequestSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/tenant_connection_requests/%s", SecondTenantConnectionRequest.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, GetResult)
	})
}

// HandleCreateTenantConnectionRequestSuccessfully creates an HTTP handler at `/tenant_connection_requests` on the
// test handler mux that tests tenant_connection_request creation.
func HandleCreateTenantConnectionRequestSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/tenant_connection_requests", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, CreateResponse)
	})
}

// HandleDeleteTenantConnectionRequestSuccessfully creates an HTTP handler at `/tenant_connection_requests` on the
// test handler mux that tests tenant_connection_request deletion.
func HandleDeleteTenantConnectionRequestSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/tenant_connection_requests/%s", FirstTenantConnectionRequest.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleUpdateTenantConnectionRequestSuccessfully creates an HTTP handler at `/tenant_connection_requests` on the
// test handler mux that tests tenant_connection_request update.
func HandleUpdateTenantConnectionRequestSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/tenant_connection_requests/%s", SecondTenantConnectionRequest.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, UpdateResult)
	})
}
