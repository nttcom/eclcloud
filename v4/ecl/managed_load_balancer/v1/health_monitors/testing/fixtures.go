package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/health_monitors"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "health_monitors": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "health_monitor",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "port": 80,
            "protocol": "http",
            "interval": 5,
            "retry": 3,
            "timeout": 5,
            "path": "/health",
            "http_status_code": "200-299"
        }
    ]
}`)

func listResult() []health_monitors.HealthMonitor {
	var healthMonitor1 health_monitors.HealthMonitor

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	healthMonitor1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitor1.Name = "health_monitor"
	healthMonitor1.Description = "description"
	healthMonitor1.Tags = tags1
	healthMonitor1.ConfigurationStatus = "ACTIVE"
	healthMonitor1.OperationStatus = "COMPLETE"
	healthMonitor1.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	healthMonitor1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	healthMonitor1.Port = 80
	healthMonitor1.Protocol = "http"
	healthMonitor1.Interval = 5
	healthMonitor1.Retry = 3
	healthMonitor1.Timeout = 5
	healthMonitor1.Path = "/health"
	healthMonitor1.HttpStatusCode = "200-299"

	return []health_monitors.HealthMonitor{healthMonitor1}
}

var createRequest = fmt.Sprintf(`
{
    "health_monitor": {
        "name": "health_monitor",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040"
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null
    }
}`)

func createResult() *health_monitors.HealthMonitor {
	var healthMonitor health_monitors.HealthMonitor

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	healthMonitor.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitor.Name = "health_monitor"
	healthMonitor.Description = "description"
	healthMonitor.Tags = tags
	healthMonitor.ConfigurationStatus = "CREATE_STAGED"
	healthMonitor.OperationStatus = "NONE"
	healthMonitor.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	healthMonitor.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	healthMonitor.Port = 0
	healthMonitor.Protocol = ""
	healthMonitor.Interval = 0
	healthMonitor.Retry = 0
	healthMonitor.Timeout = 0
	healthMonitor.Path = ""
	healthMonitor.HttpStatusCode = ""

	return &healthMonitor
}

var showResponse = fmt.Sprintf(`
{
    "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299",
        "current": {
            "port": 80,
            "protocol": "http",
            "interval": 5,
            "retry": 3,
            "timeout": 5,
            "path": "/health",
            "http_status_code": "200-299"
        },
        "staged": null
    }
}`)

func showResult() *health_monitors.HealthMonitor {
	var healthMonitor health_monitors.HealthMonitor

	var staged health_monitors.ConfigurationInResponse
	current := health_monitors.ConfigurationInResponse{
		Port:           80,
		Protocol:       "http",
		Interval:       5,
		Retry:          3,
		Timeout:        5,
		Path:           "/health",
		HttpStatusCode: "200-299",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	healthMonitor.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitor.Name = "health_monitor"
	healthMonitor.Description = "description"
	healthMonitor.Tags = tags
	healthMonitor.ConfigurationStatus = "ACTIVE"
	healthMonitor.OperationStatus = "COMPLETE"
	healthMonitor.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	healthMonitor.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	healthMonitor.Port = 80
	healthMonitor.Protocol = "http"
	healthMonitor.Interval = 5
	healthMonitor.Retry = 3
	healthMonitor.Timeout = 5
	healthMonitor.Path = "/health"
	healthMonitor.HttpStatusCode = "200-299"
	healthMonitor.Current = current
	healthMonitor.Staged = staged

	return &healthMonitor
}

var updateRequest = fmt.Sprintf(`
{
    "health_monitor": {
        "name": "health_monitor",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "health_monitor": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "health_monitor",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "port": null,
        "protocol": null,
        "interval": null,
        "retry": null,
        "timeout": null,
        "path": null,
        "http_status_code": null
    }
}`)

func updateResult() *health_monitors.HealthMonitor {
	var healthMonitor health_monitors.HealthMonitor

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	healthMonitor.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	healthMonitor.Name = "health_monitor"
	healthMonitor.Description = "description"
	healthMonitor.Tags = tags
	healthMonitor.ConfigurationStatus = "CREATE_STAGED"
	healthMonitor.OperationStatus = "NONE"
	healthMonitor.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	healthMonitor.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	healthMonitor.Port = 0
	healthMonitor.Protocol = ""
	healthMonitor.Interval = 0
	healthMonitor.Retry = 0
	healthMonitor.Timeout = 0
	healthMonitor.Path = ""
	healthMonitor.HttpStatusCode = ""

	return &healthMonitor
}

var createStagedRequest = fmt.Sprintf(`
{
    "health_monitor": {
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299"
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "health_monitor": {
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299"
    }
}`)

func createStagedResult() *health_monitors.HealthMonitor {
	var healthMonitor health_monitors.HealthMonitor

	healthMonitor.Port = 80
	healthMonitor.Protocol = "http"
	healthMonitor.Interval = 5
	healthMonitor.Retry = 3
	healthMonitor.Timeout = 5
	healthMonitor.Path = "/health"
	healthMonitor.HttpStatusCode = "200-299"

	return &healthMonitor
}

var showStagedResponse = fmt.Sprintf(`
{
    "health_monitor": {
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299"
    }
}`)

func showStagedResult() *health_monitors.HealthMonitor {
	var healthMonitor health_monitors.HealthMonitor

	healthMonitor.Port = 80
	healthMonitor.Protocol = "http"
	healthMonitor.Interval = 5
	healthMonitor.Retry = 3
	healthMonitor.Timeout = 5
	healthMonitor.Path = "/health"
	healthMonitor.HttpStatusCode = "200-299"

	return &healthMonitor
}

var updateStagedRequest = fmt.Sprintf(`
{
    "health_monitor": {
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299"
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "health_monitor": {
        "port": 80,
        "protocol": "http",
        "interval": 5,
        "retry": 3,
        "timeout": 5,
        "path": "/health",
        "http_status_code": "200-299"
    }
}`)

func updateStagedResult() *health_monitors.HealthMonitor {
	var healthMonitor health_monitors.HealthMonitor

	healthMonitor.Port = 80
	healthMonitor.Protocol = "http"
	healthMonitor.Interval = 5
	healthMonitor.Retry = 3
	healthMonitor.Timeout = 5
	healthMonitor.Path = "/health"
	healthMonitor.HttpStatusCode = "200-299"

	return &healthMonitor
}
