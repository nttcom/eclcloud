/*
Generated by https://github.com/tamac-io/openapi-to-eclcloud-rb
*/
package testing

import (
	"fmt"
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

var applyConfigurationsRequest = fmt.Sprintf(`
{
    "apply-configurations": null
}`)

var cancelConfigurationsRequest = fmt.Sprintf(`
{
    "cancel-configurations": null
}`)

var systemUpdateRequest = fmt.Sprintf(`
{
    "system-update": {
        "system_update_id": "31746df7-92f9-4b5e-ad05-59f6684a54eb"
    }
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
