package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/listeners"
)

const idListener1 = "497f6eca-6276-4993-bfeb-53cbbbba6f01"
const idListener2 = "497f6eca-6276-4993-bfeb-53cbbbba6f02"
const idListener3 = "497f6eca-6276-4993-bfeb-53cbbbba6f03"

const idLoadBalancer = "67fea379-cff0-4191-9175-de7d6941a040"

var listResponse = fmt.Sprintf(`
{
  "listeners": [
    {
      "id": "%s",
      "name": "listener_1",
      "description": "listener_1_description",
      "tags": {
        "key1": "value1"
      },
      "configuration_status": "ACTIVE",
      "operation_status": "COMPLETE",
      "load_balancer_id": "%s",
      "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
      "ip_address": "10.0.0.1",
      "port": 80,
      "protocol": "tcp"
    },
    {
      "id": "%s",
      "name": "listener_2",
      "description": "listener_2_description",
      "tags": {
        "key2": "value2"
      },
      "configuration_status": "ACTIVE",
      "operation_status": "COMPLETE",
      "load_balancer_id": "%s",
      "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
      "ip_address": "10.0.0.2",
      "port": 123,
      "protocol": "udp"
    }
  ]
}
`,
	// for listener1
	idListener1,
	idLoadBalancer,
	// for listener2
	idListener2,
	idLoadBalancer,
)

var listener1 = listeners.Listener{
	ID:                  idListener1,
	Name:                "listener_1",
	Description:         "listener_1_description",
	Tags:                map[string]string{"key1": "value1"},
	ConfigurationStatus: "ACTIVE",
	OperationStatus:     "COMPLETE",
	LoadBalancerID:      idLoadBalancer,
	TenantID:            "34f5c98ef430457ba81292637d0c6fd0",
	IPAddress:           "10.0.0.1",
	Port:                80,
	Protocol:            "tcp",
}

var listener2 = listeners.Listener{
	ID:                  idListener2,
	Name:                "listener_2",
	Description:         "listener_2_description",
	Tags:                map[string]string{"key2": "value2"},
	ConfigurationStatus: "ACTIVE",
	OperationStatus:     "COMPLETE",
	LoadBalancerID:      idLoadBalancer,
	TenantID:            "34f5c98ef430457ba81292637d0c6fd0",
	IPAddress:           "10.0.0.2",
	Port:                123,
	Protocol:            "udp",
}

var expectedListenersSlice = []listeners.Listener{
	listener1,
	listener2,
}

var getResponse = fmt.Sprintf(`
{
  "listener": {
    "id": "%s",
    "name": "listener_1",
    "description": "listener_1_description",
    "tags": {
      "key1": "value1"
    },
    "configuration_status": "ACTIVE",
    "operation_status": "COMPLETE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "ip_address": "10.0.0.1",
    "port": 80,
    "protocol": "tcp"
  }
}
`,
	idListener1,
	idLoadBalancer,
)

var getChangesResponse = fmt.Sprintf(`
{
  "listener": {
    "id": "%s",
    "name": "listener_1",
    "description": "listener_1_description",
    "tags": {
      "key1": "value1"
    },
    "configuration_status": "ACTIVE",
    "operation_status": "COMPLETE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "ip_address": "10.0.0.1",
    "port": 80,
    "protocol": "tcp",
    "current": {
      "ip_address": "10.0.0.1",
      "port": 80,
      "protocol": "tcp"
    },
    "staged": null
  }
}
`,
	idListener1,
	idLoadBalancer,
)

var listener3 = listeners.Listener{
	ID:                  idListener3,
	Name:                "listener_3",
	Description:         "listener_3_description",
	Tags:                map[string]string{"key3": "value3"},
	ConfigurationStatus: "CREATE_STAGED",
	OperationStatus:     "NONE",
	LoadBalancerID:      idLoadBalancer,
	TenantID:            "34f5c98ef430457ba81292637d0c6fd0",
	IPAddress:           "",
	Port:                0,
	Protocol:            "",
}

var createRequest = fmt.Sprintf(`
{
  "listener": {
    "name": "listener_3",
    "description": "listener_3_description",
    "tags": {
      "key3": "value3"
    },
    "ip_address": "10.0.0.3",
    "port": 443,
    "protocol": "tcp",
    "load_balancer_id": "%s"
  }
}
`,
	idLoadBalancer,
)

var createResponse = fmt.Sprintf(`
{
  "listener": {
    "id": "%s",
    "name": "listener_3",
    "description": "listener_3_description",
    "tags": {
      "key3": "value3"
    },
    "configuration_status": "CREATE_STAGED",
    "operation_status": "NONE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "ip_address": null,
    "port": null,
    "protocol": null
  }
}`,
	idListener3,
	idLoadBalancer,
)

var updateRequest = `
{
  "listener": {
    "name": "listener_1-update",
    "description": "listener_1_description-update"
  }
}
`

var updateResponse = fmt.Sprintf(`
{
  "listener": {
    "id": "%s",
    "name": "listener_1-update",
    "description": "listener_1_description-update",
    "tags": {
      "key1": "value1"
    },
    "configuration_status": "ACTIVE",
    "operation_status": "COMPLETE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "ip_address": "10.0.0.1",
    "port": 80,
    "protocol": "tcp"
  }
}
`,
	idListener1,
	idLoadBalancer,
)

var stagedListener = listeners.Listener{
	IPAddress: "10.0.1.1",
	Port:      443,
	Protocol:  "tcp",
}
var createStagedRequest = `
{
  "listener": {
    "ip_address": "10.0.1.1",
    "port": 443,
    "protocol": "tcp"
  }
}
`
var createStagedResponse = `
{
  "listener": {
    "ip_address": "10.0.1.1",
    "port": 443,
    "protocol": "tcp"
  }
}
`
var getStagedResponse = `
{
  "listener": {
    "ip_address": "10.0.1.1",
    "port": 443,
    "protocol": "tcp"
  }
}
`

var updateStagedListener = listeners.Listener{
	IPAddress: "10.0.1.1",
	Port:      8080,
	Protocol:  "tcp",
}

var updateStagedRequest = `
{
  "listener": {
    "port": 8080
  }
}
`
var updateStagedResponse = `
{
  "listener": {
    "ip_address": "10.0.1.1",
    "port": 8080,
    "protocol": "tcp"
  }
}
`
