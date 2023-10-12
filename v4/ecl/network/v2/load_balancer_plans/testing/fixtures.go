package testing

import (
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/load_balancer_plans"
)

const ListResponse = `
{
  "load_balancer_plans": [
    {
      "description": "Load Balancer Description 1",
      "enabled": true,
      "id": "58ab4df4-10f2-4fa0-b374-74b06dd648ee",
      "maximum_syslog_servers": 10,
      "model": {
        "edition": "Standard",
        "size": "50"
      },
      "name": "LB_Plan1",
      "vendor": "citrix",
      "version": "10.5-57.7"
    },
    {
      "description": "Load Balancer Description 2",
      "enabled": false,
      "id": "8b0cc5cc-b612-4810-ae45-7d6c5e806b3a",
      "maximum_syslog_servers": 10,
      "model": {
        "edition": "Standard",
        "size": "1000"
      },
      "name": "LB_Plan2",
      "vendor": "citrix",
      "version": "10.5-57.7"
    }
  ]
}
`
const GetResponse = `
{
  "load_balancer_plan": {
    "description": "Load Balance Plan Description",
    "enabled": true,
    "id": "6e5faf0c-9361-4b98-bfc4-670497c9bde3",
    "maximum_syslog_servers": 10,
    "model": {
      "edition": "Standard",
      "size": "50"
    },
    "name": "LB_Plan1",
    "vendor": "citrix",
    "version": "10.5-57.7"
  }
}
  `

var LoadBalancerPlan1 = load_balancer_plans.LoadBalancerPlan{
	Description:          "Load Balancer Description 1",
	Enabled:              true,
	ID:                   "58ab4df4-10f2-4fa0-b374-74b06dd648ee",
	MaximumSyslogServers: 10,
	Model: load_balancer_plans.Model{
		Edition: "Standard",
		Size:    "50",
	},
	Name:    "LB_Plan1",
	Vendor:  "citrix",
	Version: "10.5-57.7",
}

var LoadBalancerPlan2 = load_balancer_plans.LoadBalancerPlan{
	Description:          "Load Balancer Description 2",
	Enabled:              false,
	ID:                   "8b0cc5cc-b612-4810-ae45-7d6c5e806b3a",
	MaximumSyslogServers: 10,
	Model: load_balancer_plans.Model{
		Edition: "Standard",
		Size:    "1000",
	},
	Name:    "LB_Plan2",
	Vendor:  "citrix",
	Version: "10.5-57.7",
}

var LoadBalancerDetail = load_balancer_plans.LoadBalancerPlan{
	Description:          "Load Balance Plan Description",
	Enabled:              true,
	ID:                   "6e5faf0c-9361-4b98-bfc4-670497c9bde3",
	MaximumSyslogServers: 10,
	Model: load_balancer_plans.Model{
		Edition: "Standard",
		Size:    "50",
	},
	Name:    "LB_Plan1",
	Vendor:  "citrix",
	Version: "10.5-57.7",
}

var ExpectedLoadBalancerPlanSlice = []load_balancer_plans.LoadBalancerPlan{LoadBalancerPlan1, LoadBalancerPlan2}

const ListResponseDuplicatedNames = `
{
  "load_balancer_plans": [
    {
      "description": "Load Balancer Description 1",
      "enabled": true,
      "id": "58ab4df4-10f2-4fa0-b374-74b06dd648ee",
      "maximum_syslog_servers": 10,
      "model": {
        "edition": "Standard",
        "size": "50"
      },
      "name": "LB_Plan1",
      "vendor": "citrix",
      "version": "10.5-57.7"
    },
    {
      "description": "Load Balancer Description 2",
      "enabled": false,
      "id": "8b0cc5cc-b612-4810-ae45-7d6c5e806b3a",
      "maximum_syslog_servers": 10,
      "model": {
        "edition": "Standard",
        "size": "1000"
      },
      "name": "LB_Plan1",
      "vendor": "citrix",
      "version": "10.5-57.7"
    }
  ]
}
`
