package testing

import (
	"github.com/nttcom/eclcloud/v3/ecl/network/v2/security_group_rules"
)

const ListResponse = `
{
	"security_group_rules": [
	  {
		"description": "Allow SSH",
		"direction": "ingress",
		"ethertype": "IPv4",
		"id": "2bc0accf-312e-429a-956e-e4407625eb62",
		"port_range_max": 22,
		"port_range_min": 22,
		"protocol": "tcp",
		"remote_group_id": null,
		"remote_ip_prefix": "0.0.0.0/0",
		"security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
		"tenant_id": "e4f50856753b4dc6afee5fa6b9b6c550"
	  },
	  {
		"description": "Allow HTTPS",
		"direction": "ingress",
		"ethertype": "IPv4",
		"id": "c0e1482e-2e3c-497e-8964-e4f818071700",
		"port_range_max": 443,
		"port_range_min": 443,
		"protocol": "tcp",
		"remote_group_id": null,
		"remote_ip_prefix": "10.0.0.0/8",
		"security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
		"tenant_id": "e4f50856753b4dc6afee5fa6b9b6c550"
	  }
	]
}`

const GetResponse = `
{
	"security_group_rule": {
	  "description": "Allow SSH",
	  "direction": "ingress",
	  "ethertype": "IPv4",
	  "id": "2bc0accf-312e-429a-956e-e4407625eb62",
	  "port_range_max": 22,
	  "port_range_min": 22,
	  "protocol": "tcp",
	  "remote_group_id": null,
	  "remote_ip_prefix": "0.0.0.0/0",
	  "security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
	  "tenant_id": "e4f50856753b4dc6afee5fa6b9b6c550"
	}
}`

const CreateRequest = `
{
	"security_group_rule": {
	  "description": "Allow SSH",
	  "direction": "ingress",
	  "ethertype": "IPv4",
	  "port_range_max": 22,
	  "port_range_min": 22,
	  "protocol": "tcp",
	  "remote_ip_prefix": "0.0.0.0/0",
	  "security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
	  "tenant_id": "e4f50856753b4dc6afee5fa6b9b6c550"
	}
}`

const CreateResponse = `
{
	"security_group_rule": {
	  "description": "Allow SSH",
	  "direction": "ingress",
	  "ethertype": "IPv4",
	  "id": "2bc0accf-312e-429a-956e-e4407625eb62",
	  "port_range_max": 22,
	  "port_range_min": 22,
	  "protocol": "tcp",
	  "remote_group_id": null,
	  "remote_ip_prefix": "0.0.0.0/0",
	  "security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
	  "tenant_id": "e4f50856753b4dc6afee5fa6b9b6c550"
	}
}`

const CreateRequestWithRemoteGroup = `
{
	"security_group_rule": {
	  "description": "Allow from same group",
	  "direction": "ingress",
	  "ethertype": "IPv4",
	  "protocol": "tcp",
	  "remote_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
	  "security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a"
	}
}`

const CreateResponseWithRemoteGroup = `
{
	"security_group_rule": {
	  "description": "Allow from same group",
	  "direction": "ingress",
	  "ethertype": "IPv4",
	  "id": "rule-with-remote-group-id",
	  "port_range_max": null,
	  "port_range_min": null,
	  "protocol": "tcp",
	  "remote_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
	  "remote_ip_prefix": null,
	  "security_group_id": "a7734e61-b545-452d-a3cd-0189cbd9747a",
	  "tenant_id": "e4f50856753b4dc6afee5fa6b9b6c550"
	}
}`

var port22 = 22
var port443 = 443

var SecurityGroupRule1 = security_group_rules.SecurityGroupRule{
	Description:     "Allow SSH",
	Direction:       "ingress",
	Ethertype:       "IPv4",
	ID:              "2bc0accf-312e-429a-956e-e4407625eb62",
	PortRangeMax:    &port22,
	PortRangeMin:    &port22,
	Protocol:        "tcp",
	RemoteGroupID:   nil,
	RemoteIPPrefix:  strPtr("0.0.0.0/0"),
	SecurityGroupID: "a7734e61-b545-452d-a3cd-0189cbd9747a",
	TenantID:        "e4f50856753b4dc6afee5fa6b9b6c550",
}

var SecurityGroupRule2 = security_group_rules.SecurityGroupRule{
	Description:     "Allow HTTPS",
	Direction:       "ingress",
	Ethertype:       "IPv4",
	ID:              "c0e1482e-2e3c-497e-8964-e4f818071700",
	PortRangeMax:    &port443,
	PortRangeMin:    &port443,
	Protocol:        "tcp",
	RemoteGroupID:   nil,
	RemoteIPPrefix:  strPtr("10.0.0.0/8"),
	SecurityGroupID: "a7734e61-b545-452d-a3cd-0189cbd9747a",
	TenantID:        "e4f50856753b4dc6afee5fa6b9b6c550",
}

var ExpectedSecurityGroupRuleSlice = []security_group_rules.SecurityGroupRule{
	SecurityGroupRule1,
	SecurityGroupRule2,
}

func strPtr(s string) *string {
	return &s
}
