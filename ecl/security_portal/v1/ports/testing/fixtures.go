package testing

import (
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/ports"
)

const updateRequest = `{
	"port": [
		{
			"comment": "port 0 comment",
			"enable_port":"true",
			"ip_address": "192.168.1.50/24",
			"mtu":"1500",
			"network_id": "32314bd2-3583-4fb9-b622-9b121e04e007",
			"subnet_id": "7fd77711-abae-4828-93f1-f3d682a8771f"
		}
	]
}`

const updateResponse = `{
    "message": "The process launch request has been accepted",
	"processId": 85385
}`

var expectedResult = ports.UpdateProcess{
	Message:   "The process launch request has been accepted",
	ProcessID: 85385,
	ID:        "85385",
}
