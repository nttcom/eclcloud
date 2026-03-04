package testing

import (
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/security_groups"
)

const ListResponse = `
{
	"security_groups": [
	  {
		"description": "default security group",
		"id": "85cc3048-abc3-43cc-89b3-377341426ac5",
		"name": "default",
		"security_group_rules": [],
		"status": "ACTIVE",
		"tags": {},
		"tenant_id": "6f70656e737461636b20342065766572"
	  },
	  {
		"description": "Test security group",
		"id": "c0e1482e-2e3c-497e-8964-e4f818071700",
		"name": "test-sg",
		"security_group_rules": [
		  {
			"description": "Allow SSH",
			"direction": "ingress",
			"ethertype": "IPv4",
			"id": "rule-id-1",
			"port_range_max": 22,
			"port_range_min": 22,
			"protocol": "tcp",
			"remote_group_id": null,
			"remote_ip_prefix": "0.0.0.0/0",
			"security_group_id": "c0e1482e-2e3c-497e-8964-e4f818071700",
			"tenant_id": "6f70656e737461636b20342065766572"
		  }
		],
		"status": "ACTIVE",
		"tags": {
		  "env": "test"
		},
		"tenant_id": "6f70656e737461636b20342065766572"
	  }
	]
}`

const GetResponse = `
{
	"security_group": {
	  "description": "Test security group",
	  "id": "c0e1482e-2e3c-497e-8964-e4f818071700",
	  "name": "test-sg",
	  "security_group_rules": [
		{
		  "description": "Allow SSH",
		  "direction": "ingress",
		  "ethertype": "IPv4",
		  "id": "rule-id-1",
		  "port_range_max": 22,
		  "port_range_min": 22,
		  "protocol": "tcp",
		  "remote_group_id": null,
		  "remote_ip_prefix": "0.0.0.0/0",
		  "security_group_id": "c0e1482e-2e3c-497e-8964-e4f818071700",
		  "tenant_id": "6f70656e737461636b20342065766572"
		}
	  ],
	  "status": "ACTIVE",
	  "tags": {
		"env": "test"
	  },
	  "tenant_id": "6f70656e737461636b20342065766572"
	}
}`

const CreateRequest = `
{
	"security_group": {
	  "description": "Test security group",
	  "name": "test-sg",
	  "tags": {
		"env": "test"
	  },
	  "tenant_id": "6f70656e737461636b20342065766572"
	}
}`

const CreateResponse = `
{
	"security_group": {
	  "description": "Test security group",
	  "id": "c0e1482e-2e3c-497e-8964-e4f818071700",
	  "name": "test-sg",
	  "security_group_rules": [],
	  "status": "ACTIVE",
	  "tags": {
		"env": "test"
	  },
	  "tenant_id": "6f70656e737461636b20342065766572"
	}
}`

const UpdateRequest = `
{
	"security_group": {
	  "description": "Updated security group",
	  "name": "updated-sg",
	  "tags": {
		"env": "production"
	  }
	}
}`

const UpdateResponse = `
{
	"security_group": {
	  "description": "Updated security group",
	  "id": "c0e1482e-2e3c-497e-8964-e4f818071700",
	  "name": "updated-sg",
	  "security_group_rules": [
		{
		  "description": "Allow SSH",
		  "direction": "ingress",
		  "ethertype": "IPv4",
		  "id": "rule-id-1",
		  "port_range_max": 22,
		  "port_range_min": 22,
		  "protocol": "tcp",
		  "remote_group_id": null,
		  "remote_ip_prefix": "0.0.0.0/0",
		  "security_group_id": "c0e1482e-2e3c-497e-8964-e4f818071700",
		  "tenant_id": "6f70656e737461636b20342065766572"
		}
	  ],
	  "status": "ACTIVE",
	  "tags": {
		"env": "production"
	  },
	  "tenant_id": "6f70656e737461636b20342065766572"
	}
}`

var port22 = 22

var SecurityGroup1 = security_groups.SecurityGroup{
	Description:        "default security group",
	ID:                 "85cc3048-abc3-43cc-89b3-377341426ac5",
	Name:               "default",
	SecurityGroupRules: []security_groups.SecurityGroupRule{},
	Status:             "ACTIVE",
	Tags:               map[string]string{},
	TenantID:           "6f70656e737461636b20342065766572",
}

var SecurityGroup2 = security_groups.SecurityGroup{
	Description: "Test security group",
	ID:          "c0e1482e-2e3c-497e-8964-e4f818071700",
	Name:        "test-sg",
	SecurityGroupRules: []security_groups.SecurityGroupRule{
		{
			Description:     "Allow SSH",
			Direction:       "ingress",
			Ethertype:       "IPv4",
			ID:              "rule-id-1",
			PortRangeMax:    &port22,
			PortRangeMin:    &port22,
			Protocol:        "tcp",
			RemoteGroupID:   nil,
			RemoteIPPrefix:  strPtr("0.0.0.0/0"),
			SecurityGroupID: "c0e1482e-2e3c-497e-8964-e4f818071700",
			TenantID:        "6f70656e737461636b20342065766572",
		},
	},
	Status: "ACTIVE",
	Tags: map[string]string{
		"env": "test",
	},
	TenantID: "6f70656e737461636b20342065766572",
}

var ExpectedSecurityGroupSlice = []security_groups.SecurityGroup{SecurityGroup1, SecurityGroup2}

func strPtr(s string) *string {
	return &s
}
