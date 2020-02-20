package testing

import (
	"encoding/json"
	"fmt"

	ar "github.com/nttcom/eclcloud/ecl/sss/v1/approval_requests"
)

const idApprovalRequest1 = "9a76dca6-d8cd-4391-aac6f-2ea052f10f4"
const idApprovalRequest2 = "fc578e8b-dea2-4f8c-aa7e-9026fa173632"
const stringBody = "{\n\t\"firewall\": {\n\t\t\"availability_zone\": \"zone1-groupa\",\n\t\t\"default_gateway\": \"\",\n\t\t\"description\": \"abcdefghijklmnopqrstuvwxyz\",\n\t\t\"firewall_plan_id\": \"bd12784a-c66e-4f13-9f72-5143d64762b6\",\n\t\t\"name\": \"abcdefghijklmnopqrstuvwxyz\",\n\t\t\"tenant_id\": \"6a156ddf2ecd497ca786ff2da6df5aa8\"\n\t}\n}"

var actionsBody = json.RawMessage(fmt.Sprint(
	`{`,
	`"firewall":`,
	`{`,
	`"availability_zone":"zone1-groupa",`,
	`"default_gateway":"",`,
	`"description":"abcdefghijklmnopqrstuvwxyz",`,
	`"firewall_plan_id":"bd12784a-c66e-4f13-9f72-5143d64762b6",`,
	`"name":"abcdefghijklmnopqrstuvwxyz",`,
	`"tenant_id":"6a156ddf2ecd497ca786ff2da6df5aa8"`,
	`}`,
	`}`,
))

var listResponse = fmt.Sprintf(`
{
	"approval_requests": [
		{
			"request_id": "%s",
			"external_request_id": "test007",
			"approver_type":"tenant_owner",
			"approver_id":"11a98bf9cb144af5a204c9da566d2bd0",
			"request_user_id":"ecid9999888881",
			"service":"network",
			"actions" : [
				{
					"service": "network",
					"region": "jp1",
					"api_path": "/network/v1/firewall",
					"method": "POST",
					"body": %s
				}
			],
			"descriptions": [
				{
					"lang": "en",
					"text": "approval resquest test"
				}
			],
			"request_user": false,
			"approver": true,
			"approval_deadline": "2017-02-05 09:45:22",
			"approval_expire": null,
			"registered_time": "2017-01-31 07:43:13",
			"updated_time": null,
			"status": "registered"
		},
		{
			"request_id": "%s",
			"external_request_id": "test006",
			"approver_type":"tenant_owner",
			"approver_id":"66a98bf9cb1238192a204c9da566dbd0",
			"request_user_id":"ecid9999888882",
			"service":"network",
			"actions" : [
				{
					"service": "network",
					"region": "jp1",
					"api_path": "/network/v1/firewall",
					"method": "POST",
					"body": %s
				}
			],
			"descriptions": [
				{
					"lang": "en",
					"text": "approval resquest test"
				}
			],
			"request_user": false,
			"approver": true,
			"approval_deadline": "2016-12-25 09:45:22",
			"approval_expire": null,
			"registered_time": "2016-12-13 02:20:21",
			"updated_time": null,
			"status": "expired"
		}
	]
}
`,
	idApprovalRequest1,
	stringBody,
	idApprovalRequest2,
	stringBody,
)

var expectedApprovalRequestsSlice = []ar.ApprovalRequest{
	firstApprovalRequest,
	secondApprovalRequest,
}

var firstApprovalRequest = ar.ApprovalRequest{
	RequestID:         idApprovalRequest1,
	ExternalRequestID: "test007",
	ApproverType:      "tenant_owner",
	ApproverID:        "11a98bf9cb144af5a204c9da566d2bd0",
	RequestUserID:     "ecid9999888881",
	Service:           "network",
	Actions: []ar.Action{
		{
			Service: "network",
			Region:  "jp1",
			APIPath: "/network/v1/firewall",
			Method:  "POST",
			Body:    actionsBody,
		},
	},
	Descriptions: []ar.Description{
		{
			Lang: "en",
			Text: "approval resquest test",
		},
	},
	RequestUser:      false,
	Approver:         true,
	ApprovalDeadLine: interface{}("2017-02-05 09:45:22"),
	ApprovalExpire:   interface{}(nil),
	RegisteredTime:   interface{}("2017-01-31 07:43:13"),
	UpdatedTime:      interface{}(nil),
	Status:           "registered",
}

var secondApprovalRequest = ar.ApprovalRequest{
	RequestID:         idApprovalRequest2,
	ExternalRequestID: "test006",
	ApproverType:      "tenant_owner",
	ApproverID:        "66a98bf9cb1238192a204c9da566dbd0",
	RequestUserID:     "ecid9999888882",
	Service:           "network",
	Actions: []ar.Action{
		{
			Service: "network",
			Region:  "jp1",
			APIPath: "/network/v1/firewall",
			Method:  "POST",
			Body:    actionsBody,
		},
	},
	Descriptions: []ar.Description{
		{
			Lang: "en",
			Text: "approval resquest test",
		},
	},
	RequestUser:      false,
	Approver:         true,
	ApprovalDeadLine: interface{}("2016-12-25 09:45:22"),
	ApprovalExpire:   interface{}(nil),
	RegisteredTime:   interface{}("2016-12-13 02:20:21"),
	UpdatedTime:      interface{}(nil),
	Status:           "expired",
}

var getResponse = fmt.Sprintf(`
	{
		"request_id": "%s",
		"external_request_id": "test007",
		"approver_type":"tenant_owner",
		"approver_id":"11a98bf9cb144af5a204c9da566d2bd0",
		"request_user_id":"ecid9999888881",
		"service":"network",
		"actions" : [
			{
				"service": "network",
				"region": "jp1",
				"api_path": "/network/v1/firewall",
				"method": "POST",
				"body": {
				"firewall": {
					"availability_zone": "zone1-groupa",
					"default_gateway":   "",
					"description":       "abcdefghijklmnopqrstuvwxyz",
					"firewall_plan_id":  "bd12784a-c66e-4f13-9f72-5143d64762b6",
					"name":              "abcdefghijklmnopqrstuvwxyz",
					"tenant_id":         "6a156ddf2ecd497ca786ff2da6df5aa8"
				}
			}
			}
		],
		"descriptions": [
			{
				"lang": "en",
				"text": "approval resquest test"
			}
		],
		"request_user": false,
		"approver": true,
		"approval_deadline": "2017-02-05 09:45:22",
		"approval_expire": null,
		"registered_time": "2017-01-31 07:43:13",
		"updated_time": null,
		"status": "registered"
	}
`,
	idApprovalRequest1,
)

const updateRequest = `
{
	"status": "approved"
}
`
