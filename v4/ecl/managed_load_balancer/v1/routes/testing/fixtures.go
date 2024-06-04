package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/routes"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "routes": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "route",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "destination_cidr": "172.16.0.0/24",
            "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "next_hop_ip_address": "192.168.0.254"
        }
    ]
}`)

func listResult() []routes.Route {
	var route1 routes.Route

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	route1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	route1.Name = "route"
	route1.Description = "description"
	route1.Tags = tags1
	route1.ConfigurationStatus = "ACTIVE"
	route1.OperationStatus = "COMPLETE"
	route1.DestinationCidr = "172.16.0.0/24"
	route1.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	route1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	route1.NextHopIPAddress = "192.168.0.254"

	return []routes.Route{route1}
}

var createRequest = fmt.Sprintf(`
{
    "route": {
        "name": "route",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "destination_cidr": "172.16.0.0/24",
        "next_hop_ip_address": "192.168.0.254",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040"
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null
    }
}`)

func createResult() *routes.Route {
	var route routes.Route

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	route.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	route.Name = "route"
	route.Description = "description"
	route.Tags = tags
	route.ConfigurationStatus = "CREATE_STAGED"
	route.OperationStatus = "NONE"
	route.DestinationCidr = "172.16.0.0/24"
	route.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	route.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	route.NextHopIPAddress = ""

	return &route
}

var showResponse = fmt.Sprintf(`
{
    "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": "192.168.0.254",
        "current": {
            "next_hop_ip_address": "192.168.0.254"
        },
        "staged": null
    }
}`)

func showResult() *routes.Route {
	var route routes.Route

	var staged routes.ConfigurationInResponse
	current := routes.ConfigurationInResponse{
		NextHopIPAddress: "192.168.0.254",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	route.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	route.Name = "route"
	route.Description = "description"
	route.Tags = tags
	route.ConfigurationStatus = "ACTIVE"
	route.OperationStatus = "COMPLETE"
	route.DestinationCidr = "172.16.0.0/24"
	route.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	route.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	route.NextHopIPAddress = "192.168.0.254"
	route.Current = current
	route.Staged = staged

	return &route
}

var updateRequest = fmt.Sprintf(`
{
    "route": {
        "name": "route",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "route": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "route",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "destination_cidr": "172.16.0.0/24",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "next_hop_ip_address": null
    }
}`)

func updateResult() *routes.Route {
	var route routes.Route

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	route.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	route.Name = "route"
	route.Description = "description"
	route.Tags = tags
	route.ConfigurationStatus = "CREATE_STAGED"
	route.OperationStatus = "NONE"
	route.DestinationCidr = "172.16.0.0/24"
	route.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	route.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	route.NextHopIPAddress = ""

	return &route
}

var createStagedRequest = fmt.Sprintf(`
{
    "route": {
        "next_hop_ip_address": "192.168.0.254"
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "route": {
        "next_hop_ip_address": "192.168.0.254"
    }
}`)

func createStagedResult() *routes.Route {
	var route routes.Route

	route.NextHopIPAddress = "192.168.0.254"

	return &route
}

var showStagedResponse = fmt.Sprintf(`
{
    "route": {
        "next_hop_ip_address": "192.168.0.254"
    }
}`)

func showStagedResult() *routes.Route {
	var route routes.Route

	route.NextHopIPAddress = "192.168.0.254"

	return &route
}

var updateStagedRequest = fmt.Sprintf(`
{
    "route": {
        "next_hop_ip_address": "192.168.0.254"
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "route": {
        "next_hop_ip_address": "192.168.0.254"
    }
}`)

func updateStagedResult() *routes.Route {
	var route routes.Route

	route.NextHopIPAddress = "192.168.0.254"

	return &route
}
