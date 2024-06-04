package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/target_groups"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "target_groups": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "target_group",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "members": [
                {
                    "ip_address": "192.168.0.7",
                    "port": 80,
                    "weight": 1
                }
            ]
        }
    ]
}`)

func listResult() []target_groups.TargetGroup {
	var targetGroup1 target_groups.TargetGroup

	member11 := target_groups.MemberInResponse{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	targetGroup1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroup1.Name = "target_group"
	targetGroup1.Description = "description"
	targetGroup1.Tags = tags1
	targetGroup1.ConfigurationStatus = "ACTIVE"
	targetGroup1.OperationStatus = "COMPLETE"
	targetGroup1.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	targetGroup1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	targetGroup1.Members = []target_groups.MemberInResponse{member11}

	return []target_groups.TargetGroup{targetGroup1}
}

var createRequest = fmt.Sprintf(`
{
    "target_group": {
        "name": "target_group",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ]
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null
    }
}`)

func createResult() *target_groups.TargetGroup {
	var targetGroup target_groups.TargetGroup

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	targetGroup.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroup.Name = "target_group"
	targetGroup.Description = "description"
	targetGroup.Tags = tags
	targetGroup.ConfigurationStatus = "CREATE_STAGED"
	targetGroup.OperationStatus = "NONE"
	targetGroup.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	targetGroup.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	targetGroup.Members = nil

	return &targetGroup
}

var showResponse = fmt.Sprintf(`
{
    "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ],
        "current": {
            "members": [
                {
                    "ip_address": "192.168.0.7",
                    "port": 80,
                    "weight": 1
                }
            ]
        },
        "staged": null
    }
}`)

func showResult() *target_groups.TargetGroup {
	var targetGroup target_groups.TargetGroup

	member1 := target_groups.MemberInResponse{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}
	var staged target_groups.ConfigurationInResponse
	current := target_groups.ConfigurationInResponse{
		Members: []target_groups.MemberInResponse{member1},
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	targetGroup.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroup.Name = "target_group"
	targetGroup.Description = "description"
	targetGroup.Tags = tags
	targetGroup.ConfigurationStatus = "ACTIVE"
	targetGroup.OperationStatus = "COMPLETE"
	targetGroup.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	targetGroup.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	targetGroup.Members = []target_groups.MemberInResponse{member1}
	targetGroup.Current = current
	targetGroup.Staged = staged

	return &targetGroup
}

var updateRequest = fmt.Sprintf(`
{
    "target_group": {
        "name": "target_group",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "target_group": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "target_group",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "members": null
    }
}`)

func updateResult() *target_groups.TargetGroup {
	var targetGroup target_groups.TargetGroup

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	targetGroup.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	targetGroup.Name = "target_group"
	targetGroup.Description = "description"
	targetGroup.Tags = tags
	targetGroup.ConfigurationStatus = "CREATE_STAGED"
	targetGroup.OperationStatus = "NONE"
	targetGroup.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	targetGroup.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	targetGroup.Members = nil

	return &targetGroup
}

var createStagedRequest = fmt.Sprintf(`
{
    "target_group": {
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ]
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "target_group": {
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ]
    }
}`)

func createStagedResult() *target_groups.TargetGroup {
	var targetGroup target_groups.TargetGroup

	member1 := target_groups.MemberInResponse{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}

	targetGroup.Members = []target_groups.MemberInResponse{member1}

	return &targetGroup
}

var showStagedResponse = fmt.Sprintf(`
{
    "target_group": {
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ]
    }
}`)

func showStagedResult() *target_groups.TargetGroup {
	var targetGroup target_groups.TargetGroup

	member1 := target_groups.MemberInResponse{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}

	targetGroup.Members = []target_groups.MemberInResponse{member1}

	return &targetGroup
}

var updateStagedRequest = fmt.Sprintf(`
{
    "target_group": {
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ]
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "target_group": {
        "members": [
            {
                "ip_address": "192.168.0.7",
                "port": 80,
                "weight": 1
            }
        ]
    }
}`)

func updateStagedResult() *target_groups.TargetGroup {
	var targetGroup target_groups.TargetGroup

	member1 := target_groups.MemberInResponse{
		IPAddress: "192.168.0.7",
		Port:      80,
		Weight:    1,
	}

	targetGroup.Members = []target_groups.MemberInResponse{member1}

	return &targetGroup
}
