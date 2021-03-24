package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/sss/v1/tenants"
)

const contractID = "econ8000008888"

const idTenant1 = "9a76dca6d8cd4391aac6f2ea052f10f4"
const idTenant2 = "27a58d42769141ff8e94920a99aeb44b"

const nameTenant1 = "jp1_tenant01"

const descriptionTenant1 = "jp1 tenant01"
const descriptionTenant1Update = "jp1 tenant01-update"

const startTime = "2018-07-26 08:40:01"

// ListResponse is a sample response to a List call.
var ListResponse = fmt.Sprintf(`
{
	"contract_id": "%s",
	"tenants": [{
		"tenant_id": "%s",
		"tenant_name": "%s",
		"description": "%s",
		"region": "jp1",
		"start_time": "%s"
	}, {
		"tenant_id": "%s",
		"tenant_name": "jp2_tenant01",
		"description": "jp2 tenant",
		"region": "jp2",
		"start_time": "%s"
	}]
}
`,
	contractID,
	// fot tenant 1
	idTenant1,
	nameTenant1,
	descriptionTenant1,
	startTime,
	// for tenant 2
	idTenant2,
	startTime,
)

// ExpectedTenantsSlice is the slice of results that should be parsed
// from ListResponse in the expected order.
var ExpectedTenantsSlice = []tenants.Tenant{FirstTenant, SecondTenant}

// TenantStartTime is parsed tenant start time
var TenantStartTime, _ = time.Parse(eclcloud.RFC3339ZNoTNoZ, startTime)

// FirstTenant is the mock object of expected tenant-1
var FirstTenant = tenants.Tenant{
	ContractID:   contractID,
	TenantID:     idTenant1,
	TenantName:   nameTenant1,
	Description:  descriptionTenant1,
	TenantRegion: "jp1",
	StartTime:    TenantStartTime,
}

// SecondTenant is the mock object of expected tenant-2
var SecondTenant = tenants.Tenant{
	ContractID:   contractID,
	TenantID:     idTenant2,
	TenantName:   "jp2_tenant01",
	Description:  "jp2 tenant",
	TenantRegion: "jp2",
	StartTime:    TenantStartTime,
}

// GetResponse is a sample response to a Get call.
// This get result does not have action, attributes in ECL2.0
var GetResponse = fmt.Sprintf(`
{
	"tenant_id": "%s",
	"tenant_name": "%s",
	"description": "%s",
	"region": "jp1",
	"contract_id": "%s",
	"region_api_endpoint": "https://example.com:443/api",
	"start_time": "%s",
	"users": [{
		"user_id": "ecid8000008888",
		"contract_id": "%s",
		"contract_owner": true
	}],
	"brand_id": "ecl2"
}`, idTenant1,
	nameTenant1,
	descriptionTenant1,
	contractID,
	startTime,
	contractID,
)

// GetResponseStruct mocked actual tenant
var GetResponseStruct = tenants.Tenant{
	ContractID:   contractID,
	TenantID:     idTenant1,
	TenantName:   nameTenant1,
	Description:  descriptionTenant1,
	TenantRegion: "jp1",
	StartTime:    TenantStartTime,
}

// CreateRequest is a sample request to create a tenant.
var CreateRequest = fmt.Sprintf(`{
	"tenant_name": "%s",
	"region": "jp1",
	"description": "%s",
	"contract_id": "%s"
}`,
	nameTenant1,
	descriptionTenant1,
	contractID,
)

// CreateTenantResponse is a sample response to a create request.
var CreateResponse = fmt.Sprintf(`{
	"tenant_id": "%s",
	"tenant_name": "%s",
	"description": "%s",
	"region": "jp1",
	"contract_id": "%s"
}`, idTenant1,
	nameTenant1,
	descriptionTenant1,
	contractID,
)

// UpdateRequest is a sample request to update a zone.
var UpdateRequest = fmt.Sprintf(`
{
	"description": "%s"
}`,
	descriptionTenant1Update,
)
