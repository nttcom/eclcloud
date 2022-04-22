package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/health_monitors"
)

const idHealthMonitor1 = "497f6eca-6276-4993-bfeb-53cbbbba6f01"
const idHealthMonitor2 = "497f6eca-6276-4993-bfeb-53cbbbba6f02"

const idLoadBalancer = "67fea379-cff0-4191-9175-de7d6941a040"

var listResponse = fmt.Sprintf(`
{
  "health_monitors": [
    {
      "id": "%s",
      "name": "health_monitor_1",
      "description": "health_monitor_1_description",
      "tags": {
        "key1": "value1"
      },
      "configuration_status": "CREATE_STAGED",
      "operation_status": "NONE",
      "load_balancer_id": "%s",
      "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
      "port": 0,
      "protocol": "icmp",
      "interval": 5,
      "retry": 3,
      "timeout": 5
    },
    {
      "id": "%s",
      "name": "health_monitor_2",
      "description": "health_monitor_2_description",
      "tags": {
        "key2": "value2"
      },
      "configuration_status": "ACTIVE",
      "operation_status": "COMPLETE",
      "load_balancer_id": "%s",
      "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
      "port": 80,
      "protocol": "tcp",
      "interval": 6,
      "retry": 4,
      "timeout": 6
    }
  ]
}
`,
	// for health_monitor1
	idHealthMonitor1,
	idLoadBalancer,
	// for appliance2
	idHealthMonitor2,
	idLoadBalancer,
)

var healthMonitor1 = health_monitors.HealthMonitor{
	ID:                  idHealthMonitor1,
	Name:                "health_monitor_1",
	Description:         "health_monitor_1_description",
	Tags:                map[string]string{"key1": "value1"},
	ConfigurationStatus: "CREATE_STAGED",
	OperationStatus:     "NONE",
	LoadBalancerID:      idLoadBalancer,
	TenantID:            "34f5c98ef430457ba81292637d0c6fd0",
	Port:                0,
	Protocol:            "icmp",
	Interval:            5,
	Retry:               3,
	Timeout:             5,
}

var current = health_monitors.Configuration{
	Port:                0,
	Protocol:            "icmp",
	Interval:            5,
	Retry:               3,
	Timeout:             5,
}

var changesHealthMonitor1 = health_monitors.HealthMonitor{
	ID:                  idHealthMonitor1,
	Name:                "health_monitor_1",
	Description:         "health_monitor_1_description",
	Tags:                map[string]string{"key1": "value1"},
	ConfigurationStatus: "CREATE_STAGED",
	OperationStatus:     "NONE",
	LoadBalancerID:      idLoadBalancer,
	TenantID:            "34f5c98ef430457ba81292637d0c6fd0",
	Port:                0,
	Protocol:            "icmp",
	Interval:            5,
	Retry:               3,
	Timeout:             5,
  Current:             &current,
}

var healthMonitor2 = health_monitors.HealthMonitor{
	ID:                  idHealthMonitor2,
	Name:                "health_monitor_2",
	Description:         "health_monitor_2_description",
	Tags:                map[string]string{"key2": "value2"},
	ConfigurationStatus: "ACTIVE",
	OperationStatus:     "COMPLETE",
	LoadBalancerID:      idLoadBalancer,
	TenantID:            "34f5c98ef430457ba81292637d0c6fd0",
	Port:                80,
	Protocol:            "tcp",
	Interval:            6,
	Retry:               4,
	Timeout:             6,
}

var expectedHealthMonitorsSlice = []health_monitors.HealthMonitor{
	healthMonitor1,
	healthMonitor2,
}

var getResponse = fmt.Sprintf(`
{
  "health_monitor": {
    "id": "%s",
    "name": "health_monitor_1",
    "description": "health_monitor_1_description",
    "tags": {
      "key1": "value1"
    },
    "configuration_status": "CREATE_STAGED",
    "operation_status": "NONE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "port": 0,
    "protocol": "icmp",
    "interval": 5,
    "retry": 3,
    "timeout": 5
  }
}
`,
	idHealthMonitor1,
	idLoadBalancer,
)

var getChangesResponse = fmt.Sprintf(`
{
  "health_monitor": {
    "id": "%s",
    "name": "health_monitor_1",
    "description": "health_monitor_1_description",
    "tags": {
      "key1": "value1"
    },
    "configuration_status": "CREATE_STAGED",
    "operation_status": "NONE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "port": 0,
    "protocol": "icmp",
    "interval": 5,
    "retry": 3,
    "timeout": 5,
    "current": {
      "port": 0,
      "protocol": "icmp",
      "interval": 5,
      "retry": 3,
      "timeout": 5
    },
    "staged": null
  }
}
`,
	idHealthMonitor1,
	idLoadBalancer,
)

var createRequest = fmt.Sprintf(`
{
  "health_monitor": {
    "name": "health_monitor_1",
    "description": "health_monitor_1_description",
    "tags": {
      "key1": "value1"
    },
    "port": 0,
    "protocol": "icmp",
    "interval": 5,
    "retry": 3,
    "timeout": 5,
    "load_balancer_id": "%s"
  }
}
`,
	idLoadBalancer,
)

var createResponse = fmt.Sprintf(`
{
  "health_monitor": {
    "id": "%s",
    "name": "health_monitor_1",
    "description": "health_monitor_1_description",
    "tags": {
      "key1": "value1"
    },
    "configuration_status": "CREATE_STAGED",
    "operation_status": "NONE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "port": 0,
    "protocol": "icmp",
    "interval": 5,
    "retry": 3,
    "timeout": 5
  }
}`,
	idHealthMonitor1,
	idLoadBalancer,
)

var updateRequest = fmt.Sprintf(`
{
  "health_monitor": {
    "name": "health_monitor_1-update",
    "description": "health_monitor_1_description-update",
    "tags": {
      "key1": "value1",
      "key2": "value2"
    }
  }
}
`,
)

var updateResponse = fmt.Sprintf(`
{
  "health_monitor": {
    "id": "%s",
    "name": "health_monitor_1-update",
    "description": "health_monitor_1_description-update",
    "tags": {
      "key1": "value1",
      "key2": "value2"
    },
    "configuration_status": "CREATE_STAGED",
    "operation_status": "NONE",
    "load_balancer_id": "%s",
    "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
    "port": 0,
    "protocol": "icmp",
    "interval": 5,
    "retry": 3,
    "timeout": 5
  }
}
`,
	idHealthMonitor1,
	idLoadBalancer,
)
