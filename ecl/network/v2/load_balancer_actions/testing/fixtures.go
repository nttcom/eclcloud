package testing

import (
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_actions"
)

const RebootRequest = `
{
  "type": "HARD"
}
`
const ResetPasswordResponse = `
{
  "new_password": "ABCDabcd4321",
  "username": "user-read"
}
`
const ResetPasswordRequest = `
{
  "username": "user-read"
}
`

var ResetPasswordDetail = load_balancer_actions.ResetPasswordDetail{
	NewPassword: "ABCDabcd4321",
	Username:    "user-read",
}
