package testing

import (
	security "github.com/nttcom/eclcloud/v3/ecl/security_order/v3/network_based_device_ha"
)

const listResponse = `
{
    "status": 1,
    "code": "FOV-01",
    "message": "Successful completion",
    "records": 2,
    "rows": [
        {
            "cell": ["false", "1", "1902F60E", "CES12085", "FW_HA", "02", "ha", "zone1-groupa", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.3", "dummyNetworkID2", "dummySubnetID2", "192.168.2.3"],
            "id": 1
        }, 
        {
            "cell": ["false", "2", "1902F60E", "CES12086", "FW_HA", "02", "ha", "zone1-groupb", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.4", "dummyNetworkID2", "dummySubnetID2", "192.168.2.4"],
            "id": 2
        }
    ]
}
`

var expectedDevicesSlice = []security.HADevice{firstDevice, secondDevice}

var firstDevice = security.HADevice{
	ID: 1,
	Cell: []string{
		"false", "1", "1902F60E", "CES12085", "FW_HA", "02", "ha", "zone1-groupa", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.3", "dummyNetworkID2", "dummySubnetID2", "192.168.2.3",
	},
}

var secondDevice = security.HADevice{
	ID: 2,
	Cell: []string{
		"false", "2", "1902F60E", "CES12086", "FW_HA", "02", "ha", "zone1-groupb", "jp4_zone1", "dummyNetworkID1", "dummySubnetID1", "192.168.1.4", "dummyNetworkID2", "dummySubnetID2", "192.168.2.4",
	},
}

var createRequest = `
{
    "gt_host": [
        {
            "azgroup": "zone1-groupa",
            "licensekind": "02",
            "operatingmode": "FW_HA",
            "halink1ipaddress": "192.168.1.3",
            "halink1networkid": "c5b1b0a8-45a3-4c99-b808-84e7c13e557f",
            "halink1subnetid": "9a2116e2-52be-439c-9587-506a1a5d288d",
            "halink2ipaddress": "192.168.2.3",
            "halink2networkid": "a8df4d5f-8752-4574-a255-dc749acd458f",
            "halink2subnetid": "a2ff5669-8422-421c-bb85-a6d691ecf223"
        },
        {
            "azgroup": "zone1-groupb",
            "licensekind": "02",
            "operatingmode": "FW_HA",
            "halink1ipaddress": "192.168.1.4",
            "halink1networkid": "c5b1b0a8-45a3-4c99-b808-84e7c13e557f",
            "halink1subnetid": "9a2116e2-52be-439c-9587-506a1a5d288d",
            "halink2ipaddress": "192.168.2.4",
            "halink2networkid": "a8df4d5f-8752-4574-a255-dc749acd458f",
            "halink2subnetid": "a2ff5669-8422-421c-bb85-a6d691ecf223"
        }
    ],
    "locale": "ja",
    "sokind": "AH",
    "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5"
}`

var createResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var createResult = security.HADeviceOrder{
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
			"operatingmode": "UTM_HA"
		},
		{
			"hostname": "CES11812",
			"licensekind": "08",
			"operatingmode": "UTM_HA"
		}
	],
	"locale": "en",
	"sokind": "MH",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5"
}`

var updateResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var updateResult = security.HADeviceOrder{
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
		},
		{
			"hostname": "CES11812"
		}
	],
	"sokind": "DH",
	"tenant_id": "9ee80f2a926c49f88f166af47df4e9f5"
}`

var deleteResponse = `
{
    "status": 1,
    "code": "FOV-02",
    "message": "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
    "soId": "FGS_3B6A7602ACD04E16B6EBEF215AE8E642"
}`

var deleteResult = security.HADeviceOrder{
	Status:  1,
	Code:    "FOV-02",
	Message: "オーダーを受け付けました。ProgressRateにて状況を確認できます。",
	ID:      "FGS_3B6A7602ACD04E16B6EBEF215AE8E642",
}
