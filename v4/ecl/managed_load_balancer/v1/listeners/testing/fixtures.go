package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/listeners"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "listeners": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "listener",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "ip_address": "10.0.0.1",
            "port": 443,
            "protocol": "https"
        }
    ]
}`)

func listResult() []listeners.Listener {
	var listener1 listeners.Listener

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	listener1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listener1.Name = "listener"
	listener1.Description = "description"
	listener1.Tags = tags1
	listener1.ConfigurationStatus = "ACTIVE"
	listener1.OperationStatus = "COMPLETE"
	listener1.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	listener1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	listener1.IPAddress = "10.0.0.1"
	listener1.Port = 443
	listener1.Protocol = "https"

	return []listeners.Listener{listener1}
}

var createRequest = fmt.Sprintf(`
{
    "listener": {
        "name": "listener",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040"
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null
    }
}`)

func createResult() *listeners.Listener {
	var listener listeners.Listener

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	listener.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listener.Name = "listener"
	listener.Description = "description"
	listener.Tags = tags
	listener.ConfigurationStatus = "CREATE_STAGED"
	listener.OperationStatus = "NONE"
	listener.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	listener.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	listener.IPAddress = ""
	listener.Port = 0
	listener.Protocol = ""

	return &listener
}

var showResponse = fmt.Sprintf(`
{
    "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https",
        "current": {
            "ip_address": "10.0.0.1",
            "port": 443,
            "protocol": "https"
        },
        "staged": null
    }
}`)

func showResult() *listeners.Listener {
	var listener listeners.Listener

	var staged listeners.ConfigurationInResponse
	current := listeners.ConfigurationInResponse{
		IPAddress: "10.0.0.1",
		Port:      443,
		Protocol:  "https",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	listener.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listener.Name = "listener"
	listener.Description = "description"
	listener.Tags = tags
	listener.ConfigurationStatus = "ACTIVE"
	listener.OperationStatus = "COMPLETE"
	listener.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	listener.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	listener.IPAddress = "10.0.0.1"
	listener.Port = 443
	listener.Protocol = "https"
	listener.Current = current
	listener.Staged = staged

	return &listener
}

var updateRequest = fmt.Sprintf(`
{
    "listener": {
        "name": "listener",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "listener": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "listener",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ip_address": null,
        "port": null,
        "protocol": null
    }
}`)

func updateResult() *listeners.Listener {
	var listener listeners.Listener

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	listener.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	listener.Name = "listener"
	listener.Description = "description"
	listener.Tags = tags
	listener.ConfigurationStatus = "CREATE_STAGED"
	listener.OperationStatus = "NONE"
	listener.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	listener.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	listener.IPAddress = ""
	listener.Port = 0
	listener.Protocol = ""

	return &listener
}

var createStagedRequest = fmt.Sprintf(`
{
    "listener": {
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "listener": {
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
    }
}`)

func createStagedResult() *listeners.Listener {
	var listener listeners.Listener

	listener.IPAddress = "10.0.0.1"
	listener.Port = 443
	listener.Protocol = "https"

	return &listener
}

var showStagedResponse = fmt.Sprintf(`
{
    "listener": {
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
    }
}`)

func showStagedResult() *listeners.Listener {
	var listener listeners.Listener

	listener.IPAddress = "10.0.0.1"
	listener.Port = 443
	listener.Protocol = "https"

	return &listener
}

var updateStagedRequest = fmt.Sprintf(`
{
    "listener": {
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "listener": {
        "ip_address": "10.0.0.1",
        "port": 443,
        "protocol": "https"
    }
}`)

func updateStagedResult() *listeners.Listener {
	var listener listeners.Listener

	listener.IPAddress = "10.0.0.1"
	listener.Port = 443
	listener.Protocol = "https"

	return &listener
}
