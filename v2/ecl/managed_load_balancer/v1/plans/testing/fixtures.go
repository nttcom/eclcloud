package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/plans"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "plans": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "50M_HA_4IF",
            "description": "description",
            "bandwidth": "50M",
            "redundancy": "HA",
            "max_number_of_interfaces": 4,
            "max_number_of_health_monitors": 50,
            "max_number_of_listeners": 50,
            "max_number_of_policies": 50,
            "max_number_of_routes": 25,
            "max_number_of_target_groups": 50,
            "max_number_of_members": 100,
            "max_number_of_rules": 50,
            "max_number_of_conditions": 5,
            "enabled": true
        }
    ]
}`)

func listResult() []plans.Plan {
	var plan1 plans.Plan

	plan1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	plan1.Name = "50M_HA_4IF"
	plan1.Description = "description"
	plan1.Bandwidth = "50M"
	plan1.Redundancy = "HA"
	plan1.MaxNumberOfInterfaces = 4
	plan1.MaxNumberOfHealthMonitors = 50
	plan1.MaxNumberOfListeners = 50
	plan1.MaxNumberOfPolicies = 50
	plan1.MaxNumberOfRoutes = 25
	plan1.MaxNumberOfTargetGroups = 50
	plan1.MaxNumberOfMembers = 100
	plan1.MaxNumberOfRules = 50
	plan1.MaxNumberOfConditions = 5
	plan1.Enabled = true

	return []plans.Plan{plan1}
}

var showResponse = fmt.Sprintf(`
{
    "plan": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "50M_HA_4IF",
        "description": "description",
        "bandwidth": "50M",
        "redundancy": "HA",
        "max_number_of_interfaces": 4,
        "max_number_of_health_monitors": 50,
        "max_number_of_listeners": 50,
        "max_number_of_policies": 50,
        "max_number_of_routes": 25,
        "max_number_of_target_groups": 50,
        "max_number_of_members": 100,
        "max_number_of_rules": 50,
        "max_number_of_conditions": 5,
        "enabled": true
    }
}`)

func showResult() *plans.Plan {
	var plan plans.Plan

	plan.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	plan.Name = "50M_HA_4IF"
	plan.Description = "description"
	plan.Bandwidth = "50M"
	plan.Redundancy = "HA"
	plan.MaxNumberOfInterfaces = 4
	plan.MaxNumberOfHealthMonitors = 50
	plan.MaxNumberOfListeners = 50
	plan.MaxNumberOfPolicies = 50
	plan.MaxNumberOfRoutes = 25
	plan.MaxNumberOfTargetGroups = 50
	plan.MaxNumberOfMembers = 100
	plan.MaxNumberOfRules = 50
	plan.MaxNumberOfConditions = 5
	plan.Enabled = true

	return &plan
}
