package testing

import (
	"fmt"
	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connections"
	th "github.com/nttcom/eclcloud/testhelper"
	"github.com/nttcom/eclcloud/testhelper/client"
	"net/http"
	"testing"
)

// ListResult provides a single page of tenant_connection results.
const ListResult = `
{
  "tenant_connections": [
    {
      "id": "2a23e5a6-bd34-11e7-afb6-0050569c850d",
      "tenant_id": "7e91b19b9baa423793ee74a8e1ff2be1",
      "tenant_id_other": "c7f3a68a73e845d4ba6a42fb80fce03f",
      "tenant_connection_request_id": "5fbcc350-bd33-11e7-afb6-0050569c850d",
      "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
      "device_type": "ECL::Compute::Server",
      "device_id": "8c235a3b-8dee-41a1-b81a-64e06edc0986",
      "device_interface_id": "",
      "port_id": "b404ed73-9438-41a1-91ed-49d0e403be64",
      "status": "creating",
      "name": "test_name_1",
      "description": "test_desc_1",
      "tags": {
        "test_tags1": "test1"
    },
      "name_other": "",
      "description_other": "",
      "tags_other": {}
    },
    {
      "id": "ea5d975c-bd31-11e7-bcac-0050569c850d",
      "tenant_id": "c7f3a68a73e845d4ba6a42fb80fce03f",
      "tenant_id_other": "7e91b19b9baa423793ee74a8e1ff2be1",
      "tenant_connection_request_id": "90381138-b572-11e7-9391-0050569c850d",
      "network_id": "c4d5fc41-b7e8-4f19-96f4-85299e54373c",
      "device_type": "ECL::Compute::Server",
      "device_id": "7cc34d4b-a345-4e51-b3d9-62540faca7bf",
      "device_interface_id": "",
      "port_id": "c9c3de44-0720-4acd-87c1-9c76f0f77cac",
      "status": "down",
      "name": "test_name_2",
      "description": "test_desc_2",
      "tags": {
        "test_tags2": "test2"
    },
      "name_other": "test_name_other_2",
      "description_other": "test_desc_other_2",
      "tags_other": {
        "test_tags_other2": "test2"
    }
    }
  ]
}
`

// GetResult provides a Get result.
const GetResult = `
{
  "tenant_connection": {
    "id": "ea5d975c-bd31-11e7-bcac-0050569c850d",
    "tenant_id": "c7f3a68a73e845d4ba6a42fb80fce03f",
    "tenant_id_other": "7e91b19b9baa423793ee74a8e1ff2be1",
    "tenant_connection_request_id": "90381138-b572-11e7-9391-0050569c850d",
    "network_id": "c4d5fc41-b7e8-4f19-96f4-85299e54373c",
    "device_type": "ECL::Compute::Server",
    "device_id": "7cc34d4b-a345-4e51-b3d9-62540faca7bf",
    "device_interface_id": "",
    "port_id": "c9c3de44-0720-4acd-87c1-9c76f0f77cac",
    "status": "down",
    "name": "test_name_2",
    "description": "test_desc_2",
    "tags": {
        "test_tags2": "test2"
    },
    "name_other": "test_name_other_2",
    "description_other": "test_desc_other_2",
    "tags_other": {
        "test_tags_other2": "test2"
    }
  }
}
`

// CreateAttachServerRequest provides the input to a Create request.
const CreateAttachServerRequest = `
{
    "tenant_connection": {
        "name": "test_name_1",
        "description": "test_desc_1",
		"tags": {
			"test_tags1": "test1"
		},
		"tenant_connection_request_id": "21b344d8-be11-11e7-bf3c-0050569c850d",
		"device_type": "ECL::Compute::Server",
		"device_id": "8c235a3b-8dee-41a1-b81a-64e06edc0986",
		"attachment_opts": {
			"fixed_ips": [
				{
					"ip_address": "192.168.1.1",
					"subnet_id": "1f424165-2202-4022-ad70-0fa6f9ec99e1"
				}
			],
			"allowed_address_pairs": [
				{
					"ip_address": "192.168.1.2",
					"mac_address": "11:22:33:aa:bb:cc"
				}
			]
		}
    }
}
`

// CreateAttachServerResponse provides the output from a Create request.
const CreateAttachServerResponse = `
{
    "tenant_connection":{
        "id": "2a23e5a6-bd34-11e7-afb6-0050569c850d",
        "tenant_connection_request_id": "5fbcc350-bd33-11e7-afb6-0050569c850d",
        "name": "test_name_1",
        "description": "test_desc_1",
        "tags": {
			"test_tags1": "test1"
		},
        "tenant_id": "7e91b19b9baa423793ee74a8e1ff2be1",
        "name_other": "",
        "description_other": "",
        "tags_other": {},
        "tenant_id_other": "c7f3a68a73e845d4ba6a42fb80fce03f",
        "network_id": "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
        "device_type": "ECL::Compute::Server",
        "device_id": "8c235a3b-8dee-41a1-b81a-64e06edc0986",
        "device_interface_id": "",
        "port_id": "b404ed73-9438-41a1-91ed-49d0e403be64",
        "status": "creating"
    }
}
`

// CreateAttachVnaRequest provides the input to a Create request.
const CreateAttachVnaRequest = `
{
    "tenant_connection": {
        "name": "attach_vna_name",
        "description": "attach_vna_desc",
		"tags": {
			"test_tags1": "test1"
		},
		"tenant_connection_request_id": "67d76b00-3804-11ea-8088-525400060300",
		"device_type": "ECL::VirtualNetworkAppliance::VSRX",
		"device_interface_id": "interface_2",
		"device_id": "c291f4c4-a680-4db0-8b88-7e579f0aaa37",
		"attachment_opts": {
    		"fixed_ips": [
        		{
            		"ip_address": "192.168.1.3"
        		}
    		]
		}
    }
}
`

// CreateAttachVnaResponse provides the output from a Create request.
const CreateAttachVnaResponse = `
{
    "tenant_connection":{
        "id": "f6331886-3804-11ea-95a8-525400060400",
        "tenant_connection_request_id": "67d76b00-3804-11ea-8088-525400060300",
        "name": "attach_vna_name",
        "description": "attach_vna_desc",
        "tags": {
			"test_tags1": "test1"
		},
        "tenant_id": "7e91b19b9baa423793ee74a8e1ff2be1",
        "name_other": "",
        "description_other": "",
        "tags_other": {},
        "tenant_id_other": "c7f3a68a73e845d4ba6a42fb80fce03f",
        "network_id": "061dbaa9-a3e0-4343-b3fc-0a619db66854",
		"device_interface_id": "interface_2",
        "device_type": "ECL::VirtualNetworkAppliance::VSRX",
        "device_id": "c291f4c4-a680-4db0-8b88-7e579f0aaa37",
        "port_id": "",
        "status": "active"
    }
}
`

// UpdateRequest provides the input to as Update request.
const UpdateRequest = `
{
  "tenant_connection": {
    "name": "update_name",
    "description": "update_desc",
    "tags": {
      "update_tags": "update"
    },
	"name_other": "update_name_other",
    "description_other": "update_desc_other",
    "tags_other": {
      "test_tags_other": "update"
    }
  }
}
`

// UpdateResult provides an update result.
const UpdateResult = `
{
    "tenant_connection":{
        "id": "ea5d975c-bd31-11e7-bcac-0050569c850d",
        "tenant_connection_request_id": "90381138-b572-11e7-9391-0050569c850d",
        "name": "update_name",
        "description": "update_desc",
        "tags": {
			"update_tags": "update"
		},
        "tenant_id": "c7f3a68a73e845d4ba6a42fb80fce03f",
        "name_other": "update_name_other",
        "description_other": "update_desc_other",
        "tags_other": {
			"test_tags_other": "update"
		},
        "tenant_id_other": "7e91b19b9baa423793ee74a8e1ff2be1",
        "network_id": "c4d5fc41-b7e8-4f19-96f4-85299e54373c",
        "device_type": "ECL::Compute::Server",
        "device_id": "7cc34d4b-a345-4e51-b3d9-62540faca7bf",
        "device_interface_id": "",
        "port_id": "c9c3de44-0720-4acd-87c1-9c76f0f77cac",
        "status": "down"
    }
}
`

// FirstTenantConnection is the first tenant_connection in the List request.
var FirstTenantConnection = tenant_connections.TenantConnection{
	ID:                        "2a23e5a6-bd34-11e7-afb6-0050569c850d",
	TenantConnectionRequestID: "5fbcc350-bd33-11e7-afb6-0050569c850d",
	Name:                      "test_name_1",
	Description:               "test_desc_1",
	Tags: map[string]string{
		"test_tags1": "test1",
	},
	TenantID:          "7e91b19b9baa423793ee74a8e1ff2be1",
	NameOther:         "",
	DescriptionOther:  "",
	TagsOther:         map[string]string{},
	TenantIDOther:     "c7f3a68a73e845d4ba6a42fb80fce03f",
	NetworkID:         "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
	DeviceType:        "ECL::Compute::Server",
	DeviceID:          "8c235a3b-8dee-41a1-b81a-64e06edc0986",
	DeviceInterfaceID: "",
	PortID:            "b404ed73-9438-41a1-91ed-49d0e403be64",
	Status:            "creating",
}

// SecondTenantConnection is the second tenant_connection in the List request.
var SecondTenantConnection = tenant_connections.TenantConnection{
	ID:                        "ea5d975c-bd31-11e7-bcac-0050569c850d",
	TenantConnectionRequestID: "90381138-b572-11e7-9391-0050569c850d",
	Name:                      "test_name_2",
	Description:               "test_desc_2",
	Tags: map[string]string{
		"test_tags2": "test2",
	},
	TenantID:         "c7f3a68a73e845d4ba6a42fb80fce03f",
	NameOther:        "test_name_other_2",
	DescriptionOther: "test_desc_other_2",
	TagsOther: map[string]string{
		"test_tags_other2": "test2",
	},
	TenantIDOther:     "7e91b19b9baa423793ee74a8e1ff2be1",
	NetworkID:         "c4d5fc41-b7e8-4f19-96f4-85299e54373c",
	DeviceType:        "ECL::Compute::Server",
	DeviceID:          "7cc34d4b-a345-4e51-b3d9-62540faca7bf",
	DeviceInterfaceID: "",
	PortID:            "c9c3de44-0720-4acd-87c1-9c76f0f77cac",
	Status:            "down",
}

// CreateTenantConnectionAttachVna is the tenant_connection in the Create Attach Vna request.
var CreateTenantConnectionAttachVna = tenant_connections.TenantConnection{
	ID:                        "f6331886-3804-11ea-95a8-525400060400",
	TenantConnectionRequestID: "67d76b00-3804-11ea-8088-525400060300",
	Name:                      "attach_vna_name",
	Description:               "attach_vna_desc",
	Tags: map[string]string{
		"test_tags1": "test1",
	},
	TenantID:          "7e91b19b9baa423793ee74a8e1ff2be1",
	NameOther:         "",
	DescriptionOther:  "",
	TagsOther:         map[string]string{},
	TenantIDOther:     "c7f3a68a73e845d4ba6a42fb80fce03f",
	NetworkID:         "061dbaa9-a3e0-4343-b3fc-0a619db66854",
	DeviceType:        "ECL::VirtualNetworkAppliance::VSRX",
	DeviceID:          "c291f4c4-a680-4db0-8b88-7e579f0aaa37",
	DeviceInterfaceID: "interface_2",
	PortID:            "",
	Status:            "active",
}

// SecondTenantConnectionUpdated is how second tenant_connection should look after an Update.
var SecondTenantConnectionUpdated = tenant_connections.TenantConnection{
	ID:                        "ea5d975c-bd31-11e7-bcac-0050569c850d",
	TenantConnectionRequestID: "90381138-b572-11e7-9391-0050569c850d",
	Name:                      "update_name",
	Description:               "update_desc",
	Tags: map[string]string{
		"update_tags": "update",
	},
	TenantID:         "c7f3a68a73e845d4ba6a42fb80fce03f",
	NameOther:        "update_name_other",
	DescriptionOther: "update_desc_other",
	TagsOther: map[string]string{
		"test_tags_other": "update",
	},
	TenantIDOther:     "7e91b19b9baa423793ee74a8e1ff2be1",
	NetworkID:         "c4d5fc41-b7e8-4f19-96f4-85299e54373c",
	DeviceType:        "ECL::Compute::Server",
	DeviceID:          "7cc34d4b-a345-4e51-b3d9-62540faca7bf",
	DeviceInterfaceID: "",
	PortID:            "c9c3de44-0720-4acd-87c1-9c76f0f77cac",
	Status:            "down",
}

// ExpectedTenantConnectionsSlice is the slice of tenant_connection expected to be returned from ListResult.
var ExpectedTenantConnectionsSlice = []tenant_connections.TenantConnection{FirstTenantConnection, SecondTenantConnection}

// HandleListTenantConnectionsSuccessfully creates an HTTP handler at `/tenant_connections` on the
// test handler mux that responds with a list of two tenant_connections.
func HandleListTenantConnectionsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/tenant_connections", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, ListResult)
	})
}

// HandleGetTenantConnectionSuccessfully creates an HTTP handler at `/tenant_connections` on the
// test handler mux that responds with a single tenant_connection.
func HandleGetTenantConnectionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/tenant_connections/%s", SecondTenantConnection.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, GetResult)
	})
}

// HandleCreateTenantConnectionAttachServerSuccessfully creates an HTTP handler at `/tenant_connections` on the
// test handler mux that tests creation of tenant_connection with Server attached.
func HandleCreateTenantConnectionAttachServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/tenant_connections", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateAttachServerRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, CreateAttachServerResponse)
	})
}

// HandleCreateTenantConnectionAttachVnaSuccessfully creates an HTTP handler at `/tenant_connections` on the
// test handler mux that that tests creation of tenant_connection with Vna attached.
func HandleCreateTenantConnectionAttachVnaSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/tenant_connections", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateAttachVnaRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, CreateAttachVnaResponse)
	})
}

// HandleDeleteTenantConnectionSuccessfully creates an HTTP handler at `/tenant_connections` on the
// test handler mux that tests tenant_connection deletion.
func HandleDeleteTenantConnectionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/tenant_connections/%s", FirstTenantConnection.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleUpdateTenantConnectionSuccessfully creates an HTTP handler at `/tenant_connections` on the
// test handler mux that tests tenant_connection update.
func HandleUpdateTenantConnectionSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/tenant_connections/%s", SecondTenantConnection.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, UpdateResult)
	})
}