package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v2/ecl/managed_load_balancer/v1/load_balancers"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "load_balancers": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "load_balancer",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "monitoring_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "primary_availability_zone": "zone1_groupa",
            "secondary_availability_zone": "zone1_groupb",
            "active_availability_zone": "zone1_groupa",
            "revision": 1,
            "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
            "plan_name": "50M_HA_4IF",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "syslog_servers": [
                {
                    "ip_address": "192.168.0.6",
                    "port": 514,
                    "protocol": "udp"
                }
            ],
            "interfaces": [
                {
                    "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                    "virtual_ip_address": "192.168.0.1",
                    "reserved_fixed_ips": [
                        {
                            "ip_address": "192.168.0.2"
                        },
                        {
                            "ip_address": "192.168.0.3"
                        },
                        {
                            "ip_address": "192.168.0.4"
                        },
                        {
                            "ip_address": "192.168.0.5"
                        }
                    ]
                }
            ]
        }
    ]
}`)

func listResult() []load_balancers.LoadBalancer {
	var loadBalancer1 load_balancers.LoadBalancer

	reservedFixedIP11 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP12 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP13 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP14 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.5",
	}
	interface11 := load_balancers.InterfaceInResponse{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: []load_balancers.ReservedFixedIPInResponse{reservedFixedIP11, reservedFixedIP12, reservedFixedIP13, reservedFixedIP14},
	}
	syslogServer11 := load_balancers.SyslogServerInResponse{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	loadBalancer1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancer1.Name = "load_balancer"
	loadBalancer1.Description = "description"
	loadBalancer1.Tags = tags1
	loadBalancer1.ConfigurationStatus = "ACTIVE"
	loadBalancer1.MonitoringStatus = "ACTIVE"
	loadBalancer1.OperationStatus = "COMPLETE"
	loadBalancer1.PrimaryAvailabilityZone = "zone1_groupa"
	loadBalancer1.SecondaryAvailabilityZone = "zone1_groupb"
	loadBalancer1.ActiveAvailabilityZone = "zone1_groupa"
	loadBalancer1.Revision = 1
	loadBalancer1.PlanID = "00713021-9aea-41da-9a88-87760c08fa72"
	loadBalancer1.PlanName = "50M_HA_4IF"
	loadBalancer1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	loadBalancer1.SyslogServers = []load_balancers.SyslogServerInResponse{syslogServer11}
	loadBalancer1.Interfaces = []load_balancers.InterfaceInResponse{interface11}

	return []load_balancers.LoadBalancer{loadBalancer1}
}

var createRequest = fmt.Sprintf(`
{
    "load_balancer": {
        "name": "load_balancer",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ]
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null
    }
}`)

func createResult() *load_balancers.LoadBalancer {
	var loadBalancer load_balancers.LoadBalancer

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	loadBalancer.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancer.Name = "load_balancer"
	loadBalancer.Description = "description"
	loadBalancer.Tags = tags
	loadBalancer.ConfigurationStatus = "CREATE_STAGED"
	loadBalancer.MonitoringStatus = "INITIAL"
	loadBalancer.OperationStatus = "NONE"
	loadBalancer.PrimaryAvailabilityZone = ""
	loadBalancer.SecondaryAvailabilityZone = ""
	loadBalancer.ActiveAvailabilityZone = "UNDEFINED"
	loadBalancer.Revision = 1
	loadBalancer.PlanID = "00713021-9aea-41da-9a88-87760c08fa72"
	loadBalancer.PlanName = "50M_HA_4IF"
	loadBalancer.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	loadBalancer.SyslogServers = nil
	loadBalancer.Interfaces = nil

	return &loadBalancer
}

var showResponse = fmt.Sprintf(`
{
    "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "monitoring_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "primary_availability_zone": "zone1_groupa",
        "secondary_availability_zone": "zone1_groupb",
        "active_availability_zone": "zone1_groupa",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ],
        "current": {
            "syslog_servers": [
                {
                    "ip_address": "192.168.0.6",
                    "port": 514,
                    "protocol": "udp"
                }
            ],
            "interfaces": [
                {
                    "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                    "virtual_ip_address": "192.168.0.1",
                    "reserved_fixed_ips": [
                        {
                            "ip_address": "192.168.0.2"
                        },
                        {
                            "ip_address": "192.168.0.3"
                        },
                        {
                            "ip_address": "192.168.0.4"
                        },
                        {
                            "ip_address": "192.168.0.5"
                        }
                    ]
                }
            ]
        },
        "staged": null
    }
}`)

func showResult() *load_balancers.LoadBalancer {
	var loadBalancer load_balancers.LoadBalancer

	reservedFixedIP1 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.InterfaceInResponse{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: []load_balancers.ReservedFixedIPInResponse{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.SyslogServerInResponse{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}
	var staged load_balancers.ConfigurationInResponse
	current := load_balancers.ConfigurationInResponse{
		SyslogServers: []load_balancers.SyslogServerInResponse{syslogServer1},
		Interfaces:    []load_balancers.InterfaceInResponse{interface1},
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	loadBalancer.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancer.Name = "load_balancer"
	loadBalancer.Description = "description"
	loadBalancer.Tags = tags
	loadBalancer.ConfigurationStatus = "ACTIVE"
	loadBalancer.MonitoringStatus = "ACTIVE"
	loadBalancer.OperationStatus = "COMPLETE"
	loadBalancer.PrimaryAvailabilityZone = "zone1_groupa"
	loadBalancer.SecondaryAvailabilityZone = "zone1_groupb"
	loadBalancer.ActiveAvailabilityZone = "zone1_groupa"
	loadBalancer.Revision = 1
	loadBalancer.PlanID = "00713021-9aea-41da-9a88-87760c08fa72"
	loadBalancer.PlanName = "50M_HA_4IF"
	loadBalancer.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	loadBalancer.SyslogServers = []load_balancers.SyslogServerInResponse{syslogServer1}
	loadBalancer.Interfaces = []load_balancers.InterfaceInResponse{interface1}
	loadBalancer.Current = current
	loadBalancer.Staged = staged

	return &loadBalancer
}

var updateRequest = fmt.Sprintf(`
{
    "load_balancer": {
        "name": "load_balancer",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "load_balancer": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "load_balancer",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "monitoring_status": "INITIAL",
        "operation_status": "NONE",
        "primary_availability_zone": null,
        "secondary_availability_zone": null,
        "active_availability_zone": "UNDEFINED",
        "revision": 1,
        "plan_id": "00713021-9aea-41da-9a88-87760c08fa72",
        "plan_name": "50M_HA_4IF",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "syslog_servers": null,
        "interfaces": null
    }
}`)

func updateResult() *load_balancers.LoadBalancer {
	var loadBalancer load_balancers.LoadBalancer

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	loadBalancer.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	loadBalancer.Name = "load_balancer"
	loadBalancer.Description = "description"
	loadBalancer.Tags = tags
	loadBalancer.ConfigurationStatus = "CREATE_STAGED"
	loadBalancer.MonitoringStatus = "INITIAL"
	loadBalancer.OperationStatus = "NONE"
	loadBalancer.PrimaryAvailabilityZone = ""
	loadBalancer.SecondaryAvailabilityZone = ""
	loadBalancer.ActiveAvailabilityZone = "UNDEFINED"
	loadBalancer.Revision = 1
	loadBalancer.PlanID = "00713021-9aea-41da-9a88-87760c08fa72"
	loadBalancer.PlanName = "50M_HA_4IF"
	loadBalancer.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	loadBalancer.SyslogServers = nil
	loadBalancer.Interfaces = nil

	return &loadBalancer
}

var applyConfigurationsRequest = fmt.Sprintf(`
{
    "apply-configurations": null
}`)

var systemUpdateRequest = fmt.Sprintf(`
{
    "system-update": {
        "system_update_id": "31746df7-92f9-4b5e-ad05-59f6684a54eb"
    }
}`)

var applyConfigurationsAndSystemUpdateRequest = fmt.Sprintf(`
{
    "apply-configurations": null,
    "system-update": {
        "system_update_id": "31746df7-92f9-4b5e-ad05-59f6684a54eb"
    }
}`)

var cancelConfigurationsRequest = fmt.Sprintf(`
{
    "cancel-configurations": null
}`)

var createStagedRequest = fmt.Sprintf(`
{
    "load_balancer": {
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ]
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "load_balancer": {
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ]
    }
}`)

func createStagedResult() *load_balancers.LoadBalancer {
	var loadBalancer load_balancers.LoadBalancer

	reservedFixedIP1 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.InterfaceInResponse{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: []load_balancers.ReservedFixedIPInResponse{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.SyslogServerInResponse{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}

	loadBalancer.SyslogServers = []load_balancers.SyslogServerInResponse{syslogServer1}
	loadBalancer.Interfaces = []load_balancers.InterfaceInResponse{interface1}

	return &loadBalancer
}

var showStagedResponse = fmt.Sprintf(`
{
    "load_balancer": {
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ]
    }
}`)

func showStagedResult() *load_balancers.LoadBalancer {
	var loadBalancer load_balancers.LoadBalancer

	reservedFixedIP1 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.InterfaceInResponse{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: []load_balancers.ReservedFixedIPInResponse{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.SyslogServerInResponse{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}

	loadBalancer.SyslogServers = []load_balancers.SyslogServerInResponse{syslogServer1}
	loadBalancer.Interfaces = []load_balancers.InterfaceInResponse{interface1}

	return &loadBalancer
}

var updateStagedRequest = fmt.Sprintf(`
{
    "load_balancer": {
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ]
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "load_balancer": {
        "syslog_servers": [
            {
                "ip_address": "192.168.0.6",
                "port": 514,
                "protocol": "udp"
            }
        ],
        "interfaces": [
            {
                "network_id": "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
                "virtual_ip_address": "192.168.0.1",
                "reserved_fixed_ips": [
                    {
                        "ip_address": "192.168.0.2"
                    },
                    {
                        "ip_address": "192.168.0.3"
                    },
                    {
                        "ip_address": "192.168.0.4"
                    },
                    {
                        "ip_address": "192.168.0.5"
                    }
                ]
            }
        ]
    }
}`)

func updateStagedResult() *load_balancers.LoadBalancer {
	var loadBalancer load_balancers.LoadBalancer

	reservedFixedIP1 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.2",
	}
	reservedFixedIP2 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.3",
	}
	reservedFixedIP3 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.4",
	}
	reservedFixedIP4 := load_balancers.ReservedFixedIPInResponse{
		IPAddress: "192.168.0.5",
	}
	interface1 := load_balancers.InterfaceInResponse{
		NetworkID:        "d6797cf4-42b9-4cad-8591-9dd91c3f0fc3",
		VirtualIPAddress: "192.168.0.1",
		ReservedFixedIPs: []load_balancers.ReservedFixedIPInResponse{reservedFixedIP1, reservedFixedIP2, reservedFixedIP3, reservedFixedIP4},
	}
	syslogServer1 := load_balancers.SyslogServerInResponse{
		IPAddress: "192.168.0.6",
		Port:      514,
		Protocol:  "udp",
	}

	loadBalancer.SyslogServers = []load_balancers.SyslogServerInResponse{syslogServer1}
	loadBalancer.Interfaces = []load_balancers.InterfaceInResponse{interface1}

	return &loadBalancer
}
