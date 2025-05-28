package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/policies"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "policies": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "policy",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "configuration_status": "ACTIVE",
            "operation_status": "COMPLETE",
            "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "algorithm": "round-robin",
            "persistence": "cookie",
            "persistence_timeout": 525600,
            "idle_timeout": 600,
            "sorry_page_url": "https://example.com/sorry",
            "source_nat": "enable",
            "server_name_indications": [
                {
                    "server_name": "*.example.com",
                    "input_type": "fixed",
                    "priority": 1,
                    "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
                }
            ],
            "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
            "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
            "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
            "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
            "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
            "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        }
    ]
}`)

func listResult() []policies.Policy {
	var policy1 policies.Policy

	serverNameIndication11 := policies.ServerNameIndicationInResponse{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	policy1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policy1.Name = "policy"
	policy1.Description = "description"
	policy1.Tags = tags1
	policy1.ConfigurationStatus = "ACTIVE"
	policy1.OperationStatus = "COMPLETE"
	policy1.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	policy1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	policy1.Algorithm = "round-robin"
	policy1.Persistence = "cookie"
	policy1.PersistenceTimeout = 525600
	policy1.IdleTimeout = 600
	policy1.SorryPageUrl = "https://example.com/sorry"
	policy1.SourceNat = "enable"
	policy1.ServerNameIndications = []policies.ServerNameIndicationInResponse{serverNameIndication11}
	policy1.CertificateID = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	policy1.HealthMonitorID = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	policy1.ListenerID = "68633f4f-f52a-402f-8572-b8173418904f"
	policy1.DefaultTargetGroupID = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	policy1.BackupTargetGroupID = "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	policy1.TLSPolicyID = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"

	return []policies.Policy{policy1}
}

var createRequest = fmt.Sprintf(`
{
    "policy": {
        "name": "policy",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040"
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "persistence_timeout": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "server_name_indications": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "backup_target_group_id": null,
        "tls_policy_id": null
    }
}`)

func createResult() *policies.Policy {
	var policy policies.Policy

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	policy.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policy.Name = "policy"
	policy.Description = "description"
	policy.Tags = tags
	policy.ConfigurationStatus = "CREATE_STAGED"
	policy.OperationStatus = "NONE"
	policy.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	policy.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	policy.Algorithm = ""
	policy.Persistence = ""
	policy.PersistenceTimeout = 0
	policy.IdleTimeout = 0
	policy.SorryPageUrl = ""
	policy.SourceNat = ""
	policy.ServerNameIndications = nil
	policy.CertificateID = ""
	policy.HealthMonitorID = ""
	policy.ListenerID = ""
	policy.DefaultTargetGroupID = ""
	policy.BackupTargetGroupID = ""
	policy.TLSPolicyID = ""

	return &policy
}

var showResponse = fmt.Sprintf(`
{
    "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "ACTIVE",
        "operation_status": "COMPLETE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
        "current": {
            "algorithm": "round-robin",
            "persistence": "cookie",
            "persistence_timeout": 525600,
            "idle_timeout": 600,
            "sorry_page_url": "https://example.com/sorry",
            "source_nat": "enable",
            "server_name_indications": [
                {
                    "server_name": "*.example.com",
                    "input_type": "fixed",
                    "priority": 1,
                    "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
                }
            ],
            "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
            "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
            "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
            "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
            "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
            "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
        },
        "staged": null
    }
}`)

func showResult() *policies.Policy {
	var policy policies.Policy

	serverNameIndication1 := policies.ServerNameIndicationInResponse{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}
	var staged policies.ConfigurationInResponse
	current := policies.ConfigurationInResponse{
		Algorithm:             "round-robin",
		Persistence:           "cookie",
		PersistenceTimeout:    525600,
		IdleTimeout:           600,
		SorryPageUrl:          "https://example.com/sorry",
		SourceNat:             "enable",
		ServerNameIndications: []policies.ServerNameIndicationInResponse{serverNameIndication1},
		CertificateID:         "f57a98fe-d63e-4048-93a0-51fe163f30d7",
		HealthMonitorID:       "dd7a96d6-4e66-4666-baca-a8555f0c472c",
		ListenerID:            "68633f4f-f52a-402f-8572-b8173418904f",
		DefaultTargetGroupID:  "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
		BackupTargetGroupID:   "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
		TLSPolicyID:           "4ba79662-f2a1-41a4-a3d9-595799bbcd86",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	policy.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policy.Name = "policy"
	policy.Description = "description"
	policy.Tags = tags
	policy.ConfigurationStatus = "ACTIVE"
	policy.OperationStatus = "COMPLETE"
	policy.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	policy.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	policy.Algorithm = "round-robin"
	policy.Persistence = "cookie"
	policy.PersistenceTimeout = 525600
	policy.IdleTimeout = 600
	policy.SorryPageUrl = "https://example.com/sorry"
	policy.SourceNat = "enable"
	policy.ServerNameIndications = []policies.ServerNameIndicationInResponse{serverNameIndication1}
	policy.CertificateID = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	policy.HealthMonitorID = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	policy.ListenerID = "68633f4f-f52a-402f-8572-b8173418904f"
	policy.DefaultTargetGroupID = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	policy.BackupTargetGroupID = "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	policy.TLSPolicyID = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
	policy.Current = current
	policy.Staged = staged

	return &policy
}

var updateRequest = fmt.Sprintf(`
{
    "policy": {
        "name": "policy",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "policy",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "configuration_status": "CREATE_STAGED",
        "operation_status": "NONE",
        "load_balancer_id": "67fea379-cff0-4191-9175-de7d6941a040",
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "algorithm": null,
        "persistence": null,
        "persistence_timeout": null,
        "idle_timeout": null,
        "sorry_page_url": null,
        "source_nat": null,
        "server_name_indications": null,
        "certificate_id": null,
        "health_monitor_id": null,
        "listener_id": null,
        "default_target_group_id": null,
        "backup_target_group_id": null,
        "tls_policy_id": null
    }
}`)

func updateResult() *policies.Policy {
	var policy policies.Policy

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	policy.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	policy.Name = "policy"
	policy.Description = "description"
	policy.Tags = tags
	policy.ConfigurationStatus = "CREATE_STAGED"
	policy.OperationStatus = "NONE"
	policy.LoadBalancerID = "67fea379-cff0-4191-9175-de7d6941a040"
	policy.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	policy.Algorithm = ""
	policy.Persistence = ""
	policy.PersistenceTimeout = 0
	policy.IdleTimeout = 0
	policy.SorryPageUrl = ""
	policy.SourceNat = ""
	policy.ServerNameIndications = nil
	policy.CertificateID = ""
	policy.HealthMonitorID = ""
	policy.ListenerID = ""
	policy.DefaultTargetGroupID = ""
	policy.BackupTargetGroupID = ""
	policy.TLSPolicyID = ""

	return &policy
}

var createStagedRequest = fmt.Sprintf(`
{
    "policy": {
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
    }
}`)

var createStagedResponse = fmt.Sprintf(`
{
    "policy": {
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
    }
}`)

func createStagedResult() *policies.Policy {
	var policy policies.Policy

	serverNameIndication1 := policies.ServerNameIndicationInResponse{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}

	policy.Algorithm = "round-robin"
	policy.Persistence = "cookie"
	policy.PersistenceTimeout = 525600
	policy.IdleTimeout = 600
	policy.SorryPageUrl = "https://example.com/sorry"
	policy.SourceNat = "enable"
	policy.ServerNameIndications = []policies.ServerNameIndicationInResponse{serverNameIndication1}
	policy.CertificateID = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	policy.HealthMonitorID = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	policy.ListenerID = "68633f4f-f52a-402f-8572-b8173418904f"
	policy.DefaultTargetGroupID = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	policy.BackupTargetGroupID = "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	policy.TLSPolicyID = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"

	return &policy
}

var showStagedResponse = fmt.Sprintf(`
{
    "policy": {
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
    }
}`)

func showStagedResult() *policies.Policy {
	var policy policies.Policy

	serverNameIndication1 := policies.ServerNameIndicationInResponse{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}

	policy.Algorithm = "round-robin"
	policy.Persistence = "cookie"
	policy.PersistenceTimeout = 525600
	policy.IdleTimeout = 600
	policy.SorryPageUrl = "https://example.com/sorry"
	policy.SourceNat = "enable"
	policy.ServerNameIndications = []policies.ServerNameIndicationInResponse{serverNameIndication1}
	policy.CertificateID = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	policy.HealthMonitorID = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	policy.ListenerID = "68633f4f-f52a-402f-8572-b8173418904f"
	policy.DefaultTargetGroupID = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	policy.BackupTargetGroupID = "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	policy.TLSPolicyID = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"

	return &policy
}

var updateStagedRequest = fmt.Sprintf(`
{
    "policy": {
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
    }
}`)

var updateStagedResponse = fmt.Sprintf(`
{
    "policy": {
        "algorithm": "round-robin",
        "persistence": "cookie",
        "persistence_timeout": 525600,
        "idle_timeout": 600,
        "sorry_page_url": "https://example.com/sorry",
        "source_nat": "enable",
        "server_name_indications": [
            {
                "server_name": "*.example.com",
                "input_type": "fixed",
                "priority": 1,
                "certificate_id": "fdfed344-e8ab-4f20-bd62-a4039453a389"
            }
        ],
        "certificate_id": "f57a98fe-d63e-4048-93a0-51fe163f30d7",
        "health_monitor_id": "dd7a96d6-4e66-4666-baca-a8555f0c472c",
        "listener_id": "68633f4f-f52a-402f-8572-b8173418904f",
        "default_target_group_id": "a44c4072-ed90-4b50-a33a-6b38fb10c7db",
        "backup_target_group_id": "f1a117f1-f8df-ce07-6c8c-4bbf103059b6",
        "tls_policy_id": "4ba79662-f2a1-41a4-a3d9-595799bbcd86"
    }
}`)

func updateStagedResult() *policies.Policy {
	var policy policies.Policy

	serverNameIndication1 := policies.ServerNameIndicationInResponse{
		ServerName:    "*.example.com",
		InputType:     "fixed",
		Priority:      1,
		CertificateID: "fdfed344-e8ab-4f20-bd62-a4039453a389",
	}

	policy.Algorithm = "round-robin"
	policy.Persistence = "cookie"
	policy.PersistenceTimeout = 525600
	policy.IdleTimeout = 600
	policy.SorryPageUrl = "https://example.com/sorry"
	policy.SourceNat = "enable"
	policy.ServerNameIndications = []policies.ServerNameIndicationInResponse{serverNameIndication1}
	policy.CertificateID = "f57a98fe-d63e-4048-93a0-51fe163f30d7"
	policy.HealthMonitorID = "dd7a96d6-4e66-4666-baca-a8555f0c472c"
	policy.ListenerID = "68633f4f-f52a-402f-8572-b8173418904f"
	policy.DefaultTargetGroupID = "a44c4072-ed90-4b50-a33a-6b38fb10c7db"
	policy.BackupTargetGroupID = "f1a117f1-f8df-ce07-6c8c-4bbf103059b6"
	policy.TLSPolicyID = "4ba79662-f2a1-41a4-a3d9-595799bbcd86"

	return &policy
}
