package testing

import (
	security "github.com/nttcom/eclcloud/v2/ecl/security_order/v2/network_based_device_single"
)

const listResponse = `
{
    "status": 1,
    "code": "FOV-01",
    "message": "Successful completion",
    "records": 2,
    "rows": [
        {
            "id": 1,
            "cell": ["false", "1", "CES11810", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
        },
        {
            "id": 2,
            "cell": ["false", "1", "CES11811", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1"]
        }
    ]
}
`

var expectedDevicesSlice = []security.SingleDevice{firstDevice, secondDevice}

var firstDevice = security.SingleDevice{
	ID: 1,
	Cell: []string{
		"false", "1", "CES11810", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1",
	},
}

var secondDevice = security.SingleDevice{
	ID: 2,
	Cell: []string{
		"false", "1", "CES11811", "FW", "02", "standalone", "zone1-groupb", "jp4_zone1",
	},
}

var createRequest = `
{
	"gt_host":[
		{
			"azgroup": "zone1-groupb",
			"licensekind":"02",
			"operatingmode":"FW"
		}
	],
	"locale": "ja",
	"sokind": "A",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5"
}`

var createResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var createResult = security.SingleDeviceOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}

var updateRequest = `
{
	"gt_host": [
		{
			"hostname": "CES11811",
			"licensekind": "08",
			"operatingmode": "UTM"
		}
	],
	"locale": "en",
	"sokind": "M",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5"
}`

var updateResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var updateResult = security.SingleDeviceOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}

var deleteRequest = `
{
	"gt_host": [
		{
			"hostname": "CES11811"
		}
	],
	"sokind": "D",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5"
}`

var deleteResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var deleteResult = security.SingleDeviceOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}
