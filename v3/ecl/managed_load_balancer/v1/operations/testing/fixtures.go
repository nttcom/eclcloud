package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/operations"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "operations": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "resource_id": "4d5215ed-38bb-48ed-879a-fdb9ca58522f",
            "resource_type": "ECL::ManagedLoadBalancer::LoadBalancer",
            "request_id": "",
            "request_types": [
                "Action::apply-configurations"
            ],
            "status": "COMPLETE",
            "reception_datetime": "2019-08-24 14:15:22",
            "commit_datetime": "2019-08-24 14:30:44",
            "warning": "",
            "error": "",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0"
        }
    ]
}`)

func listResult() []operations.Operation {
	var operation1 operations.Operation

	operation1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	operation1.ResourceID = "4d5215ed-38bb-48ed-879a-fdb9ca58522f"
	operation1.ResourceType = "ECL::ManagedLoadBalancer::LoadBalancer"
	operation1.RequestID = ""
	operation1.RequestTypes = []string{"Action::apply-configurations"}
	operation1.Status = "COMPLETE"
	operation1.ReceptionDatetime = "2019-08-24 14:15:22"
	operation1.CommitDatetime = "2019-08-24 14:30:44"
	operation1.Warning = ""
	operation1.Error = ""
	operation1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"

	return []operations.Operation{operation1}
}

var showResponse = fmt.Sprintf(`
{
    "operation": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "resource_id": "4d5215ed-38bb-48ed-879a-fdb9ca58522f",
        "resource_type": "ECL::ManagedLoadBalancer::LoadBalancer",
        "request_id": "",
        "request_types": [
            "Action::apply-configurations"
        ],
        "request_body": {
            "apply-configurations": {
                "load_balancer": {
                    "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
                    "configuration_status": "CREATE_STAGED",
                    "current": null,
                    "staged": {
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
                },
                "health_monitors": [
                    {
                        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
                        "configuration_status": "CREATE_STAGED",
                        "current": null,
                        "staged": {
                            "port": 0,
                            "protocol": "icmp",
                            "interval": 5,
                            "retry": 3,
                            "timeout": 5
                        }
                    }
                ],
                "listeners": [
                    {
                        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
                        "configuration_status": "CREATE_STAGED",
                        "current": null,
                        "staged": {
                            "ip_address": "10.0.0.1",
                            "port": 80,
                            "protocol": "tcp"
                        }
                    }
                ],
                "policies": [
                    {
                        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
                        "configuration_status": "CREATE_STAGED",
                        "current": null,
                        "staged": {
                            "algorithm": "round-robin",
                            "persistence": "none",
                            "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
                            "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
                            "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
                        }
                    }
                ],
                "routes": [
                    {
                        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
                        "configuration_status": "CREATE_STAGED",
                        "current": null,
                        "staged": {
                            "next_hop_ip_address": "192.168.0.254"
                        }
                    }
                ],
                "target_groups": [
                    {
                        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
                        "configuration_status": "CREATE_STAGED",
                        "current": null,
                        "staged": {
                            "members": [
                                {
                                    "ip_address": "192.168.0.6",
                                    "port": 80,
                                    "weight": 1
                                }
                            ]
                        }
                    }
                ]
            }
        },
        "status": "COMPLETE",
        "reception_datetime": "2019-08-24 14:15:22",
        "commit_datetime": "2019-08-24 14:30:44",
        "warning": "",
        "error": "",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0"
    }
}`)

func showResult() *operations.Operation {
	var operation operations.Operation

	var requestBody map[string]interface{}
	requestBodyJson := `{"apply-configurations":{"load_balancer":{"id":"497f6eca-6276-4993-bfeb-53cbbbba6f08","configuration` +
		`_status":"CREATE_STAGED","current":null,"staged":{"interfaces":[{"network_id":"d6797cf4-42b9-4cad-85` +
		`91-9dd91c3f0fc3","virtual_ip_address":"192.168.0.1","reserved_fixed_ips":[{"ip_address":"192.168.0.2` +
		`"},{"ip_address":"192.168.0.3"},{"ip_address":"192.168.0.4"},{"ip_address":"192.168.0.5"}]}]}},"heal` +
		`th_monitors":[{"id":"497f6eca-6276-4993-bfeb-53cbbbba6f08","configuration_status":"CREATE_STAGED","c` +
		`urrent":null,"staged":{"port":0,"protocol":"icmp","interval":5,"retry":3,"timeout":5}}],"listeners":` +
		`[{"id":"497f6eca-6276-4993-bfeb-53cbbbba6f08","configuration_status":"CREATE_STAGED","current":null,` +
		`"staged":{"ip_address":"10.0.0.1","port":80,"protocol":"tcp"}}],"policies":[{"id":"497f6eca-6276-499` +
		`3-bfeb-53cbbbba6f08","configuration_status":"CREATE_STAGED","current":null,"staged":{"algorithm":"ro` +
		`und-robin","persistence":"none","health_monitor_id":"dd7a96d6-4e66-4666-baca-a8555f0c472c","listener` +
		`_id":"68633f4f-f52a-402f-8572-b8173418904f","default_target_group_id":"a44c4072-ed90-4b50-a33a-6b38f` +
		`b10c7db"}}],"routes":[{"id":"497f6eca-6276-4993-bfeb-53cbbbba6f08","configuration_status":"CREATE_ST` +
		`AGED","current":null,"staged":{"next_hop_ip_address":"192.168.0.254"}}],"target_groups":[{"id":"497f` +
		`6eca-6276-4993-bfeb-53cbbbba6f08","configuration_status":"CREATE_STAGED","current":null,"staged":{"m` +
		`embers":[{"ip_address":"192.168.0.6","port":80,"weight":1}]}}]}}`
	err := json.Unmarshal([]byte(requestBodyJson), &requestBody)
	if err != nil {
		panic(err)
	}

	operation.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	operation.ResourceID = "4d5215ed-38bb-48ed-879a-fdb9ca58522f"
	operation.ResourceType = "ECL::ManagedLoadBalancer::LoadBalancer"
	operation.RequestID = ""
	operation.RequestTypes = []string{"Action::apply-configurations"}
	operation.RequestBody = requestBody
	operation.Status = "COMPLETE"
	operation.ReceptionDatetime = "2019-08-24 14:15:22"
	operation.CommitDatetime = "2019-08-24 14:30:44"
	operation.Warning = ""
	operation.Error = ""
	operation.TenantID = "34f5c98ef430457ba81292637d0c6fd0"

	return &operation
}
