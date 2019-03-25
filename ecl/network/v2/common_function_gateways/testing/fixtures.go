package testing

import (
	"fmt"
	"github.com/nttcom/eclcloud/ecl/network/v2/common_function_gateways"
)

// Define parameters which are used in assertion.
// Additionally, kind of IDs are defined here.
const idCommonFunctionGatway1 = "fb3efc23-ca8c-4eb5-b7f6-6fc66ff24f9c"
const idCommonFunctionGatway2 = "3535de20-192d-4f5a-a74a-cd1a9c1bf747"
const idCommonFunctionPool = "4f4971a5-899d-42b4-8442-24f17eac9683"

const nameCommonFunctionGateway1 = "common_function_gateway_name_1"
const descriptionCommonFunctionGateway1 = "common_function_gateway_description_1"

const nameCommonFunctionGateway1Update = "common_function_gateway_name_1-update"
const descriptionCommonFunctionGateway1Update = "common_function_gateway_description_1-update"

const tenantID = "2d5b878c-147a-4d7c-87fd-90a8be9d255f"

const networkID = "511f266e-a8bf-4547-ab2a-fc4d2bda9f81"
const subnetID = "9f3fd369-e4d4-4c3a-84f1-9c5ba7686297"

// ListResponse is mocked response of common_function_gateways.List
var ListResponse = fmt.Sprintf(`
{
    "common_function_gateways": [
        {
            "id": "%s",
            "common_function_pool_id": "%s",
            "name": "%s",
            "description": "%s",
            "tenant_id": "%s",
            "network_id": "%s",
            "subnet_id": "%s",
            "status": "ACTIVE"
        },
        {
            "id": "%s",
            "common_function_pool_id": "%s",
            "tenant_id": "%s",
            "name": "common_function_gateway_name_2",
            "description": "common_function_gateway_description_2",
            "network_id": "%s",
            "subnet_id": "%s",
            "status": "ACTIVE"
        }
    ]
}`,
	// for common function gateway1
	idCommonFunctionGatway1,
	idCommonFunctionPool,
	nameCommonFunctionGateway1,
	descriptionCommonFunctionGateway1,
	tenantID,
	networkID,
	subnetID,
	// for common function gateway2
	idCommonFunctionGatway2,
	idCommonFunctionPool,
	tenantID,
	networkID,
	subnetID)

// GetResponse is mocked format of common_function_gateways.Get
var GetResponse = fmt.Sprintf(`
{
    "common_function_gateway": {
        "id": "%s",
        "common_function_pool_id": "%s",
        "name": "%s",
        "description": "%s",
        "tenant_id": "%s",
        "network_id": "%s",
        "subnet_id": "%s",
        "status": "ACTIVE"
    }
}`, idCommonFunctionGatway1,
	idCommonFunctionPool,
	nameCommonFunctionGateway1,
	descriptionCommonFunctionGateway1,
	tenantID,
	networkID,
	subnetID)

// CreateRequest is mocked request for common_function_gateways.Create
var CreateRequest = fmt.Sprintf(`
{
    "common_function_gateway": {
        "name": "%s",
        "description": "%s",
        "common_function_pool_id": "%s",
        "tenant_id": "%s"
    }
}`, nameCommonFunctionGateway1,
	descriptionCommonFunctionGateway1,
	idCommonFunctionPool,
	tenantID)

// CreateResponse is mocked response of common_function_gateways.Create
var CreateResponse = fmt.Sprintf(`
{
    "common_function_gateway": {
        "id": "%s",
        "common_function_pool_id": "%s",
        "name": "%s",
        "description": "%s",
        "tenant_id": "%s",
        "network_id": "%s",
        "subnet_id": "%s",
        "status": "ACTIVE"
    }
}`, idCommonFunctionGatway1,
	idCommonFunctionPool,
	nameCommonFunctionGateway1,
	descriptionCommonFunctionGateway1,
	tenantID,
	networkID,
	subnetID)

// UpdateRequest is mocked request of common_function_gateways.Update
var UpdateRequest = fmt.Sprintf(`
{
    "common_function_gateway": {
        "name": "%s",
        "description": "%s"
    }
}`, nameCommonFunctionGateway1Update,
	descriptionCommonFunctionGateway1Update)

// UpdateResponse is mocked response of common_function_gateways.Update
var UpdateResponse = fmt.Sprintf(`
{
    "common_function_gateway": {
        "id": "%s",
        "common_function_pool_id": "%s",
        "name": "%s",
        "description": "%s",
        "tenant_id": "%s",
        "network_id": "%s",
        "subnet_id": "%s"
    }
}`, idCommonFunctionGatway1,
	idCommonFunctionPool,
	nameCommonFunctionGateway1Update,
	descriptionCommonFunctionGateway1Update,
	tenantID,
	networkID,
	subnetID)

var commonFunctionGateway1 = common_function_gateways.CommonFunctionGateway{
	ID:                   idCommonFunctionGatway1,
	CommonFunctionPoolID: idCommonFunctionPool,
	TenantID:             tenantID,
	Name:                 nameCommonFunctionGateway1,
	Description:          descriptionCommonFunctionGateway1,
	Status:               "ACTIVE",
	NetworkID:            networkID,
	SubnetID:             subnetID,
}

var commonFunctionGateway2 = common_function_gateways.CommonFunctionGateway{
	ID:                   idCommonFunctionGatway2,
	CommonFunctionPoolID: idCommonFunctionPool,
	TenantID:             tenantID,
	Name:                 "common_function_gateway_name_2",
	Description:          "common_function_gateway_description_2",
	Status:               "ACTIVE",
	NetworkID:            networkID,
	SubnetID:             subnetID,
}

// ExpectedCommonFunctionGatewaysSlice is expected assertion target
var ExpectedCommonFunctionGatewaysSlice = []common_function_gateways.CommonFunctionGateway{
	commonFunctionGateway1,
	commonFunctionGateway2,
}
