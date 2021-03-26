package testing

import (
	order "github.com/nttcom/eclcloud/v2/ecl/security_order/v2/service_order_status"
)

const getResponse = `
{
    "status": 1,
    "code": "FOV-05",
    "message": "We accepted the order. Please wait",
    "progressRate": 45
}
`

var expectedResult = order.OrderProgress{
	Status:       1,
	Code:         "FOV-05",
	Message:      "We accepted the order. Please wait",
	ProgressRate: 45,
}
