package load_balancer_actions

import (
	"github.com/nttcom/eclcloud/v2"
)

type commonResult struct {
	eclcloud.Result
}

// ExtractResetPassword is a function that accepts a result and extracts a result of reset_password.
func (r commonResult) ExtractResetPassword() (*Password, error) {
	var s Password
	err := r.ExtractInto(&s)
	return &s, err
}

// RebootResult represents the result of a reboot operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type RebootResult struct {
	eclcloud.ErrResult
}

// ResetPasswordResult represents the result of a Reset Password operation. Call its ExtractResetPassword
// method to interpret it as an action's result.
type ResetPasswordResult struct {
	commonResult
}

// Password represents a detail of a Reset Password operation.
type Password struct {

	// new password
	NewPassword string `json:"new_password"`

	// username
	Username string `json:"username"`
}
