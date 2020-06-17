package testing

import (
	security "github.com/nttcom/eclcloud/ecl/security_order/v2/host_based"
)

const getResponse = `
{
  "code": "DEP-01",
  "message": "Successful completion",
  "region": "jp4",
  "tenant_name": "Test Tenant",
  "tenant_description": "Test Tenant",
  "contract_id": "econ9999999999",
  "service_order_service": "Managed Anti-Virus",
  "max_agent_value": 1,
  "customer_name": "Customer",
  "time_zone": "Asia/Tokyo",
  "mailaddress": "terraform@example.com",
  "dsm_lang": "ja",
  "tenant_flg": true,
  "status": 1
}
`

var expectedResult = security.HostBasedSecurity{
	Code:                "DEP-01",
	Message:             "Successful completion",
	Region:              "jp4",
	TenantName:          "Test Tenant",
	TenantDescription:   "Test Tenant",
	ContractID:          "econ9999999999",
	ServiceOrderService: "Managed Anti-Virus",
	MaxAgentValue:       float64(1),
	TimeZone:            "Asia/Tokyo",
	CustomerName:        "Customer",
	MailAddress:         "terraform@example.com",
	DSMLang:             "ja",
	TenantFlg:           true,
	Status:              1,
}

var createRequest = `
{
    "sokind": "N",
    "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
    "locale": "ja",
    "service_order_service": "Managed Anti-Virus",
    "max_agent_value": 1,
    "mailaddress": "terraform@example.com",
    "dsm_lang": "ja",
    "time_zone": "Asia/Tokyo"
}`

var createResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var createResult = security.HostBasedOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}

var updateRequestM1 = `
{
	"sokind": "M1",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
    "locale": "ja",
    "mailaddress": "terraform@example.com",
    "service_order_service": "Managed Anti-Virus"
}`

var updateResponseM1 = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var updateResultM1 = security.HostBasedOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}

var updateRequestM2 = `
{
	"sokind": "M2",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
    "locale": "ja",
    "mailaddress": "terraform@example.com",
    "max_agent_value": 10
}`

var updateResponseM2 = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var updateResultM2 = security.HostBasedOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}

var deleteRequest = `
{
	"sokind": "C",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
    "locale": "ja",
    "mailaddress": "terraform@example.com"
}`

var deleteResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var deleteResult = security.HostBasedOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}
