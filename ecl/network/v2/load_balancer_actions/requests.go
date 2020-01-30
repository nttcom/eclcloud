package load_balancer_actions

import (
	"github.com/nttcom/eclcloud"
)

// RebootOpts represents the attributes used when rebooting a Load Balancer.
type RebootOpts struct {

	// should syslog record acl info
	Type string `json:"type" required:"true"`
}

// ToLoadBalancerActionRebootMap builds a request body from RebootOpts.
func (opts RebootOpts) ToLoadBalancerActionRebootMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Reboot accepts a RebootOpts struct and reboots an existing Load Balancer using the
// values provided.
func Reboot(c *eclcloud.ServiceClient, id string, opts RebootOpts) (r RebootResult) {
	b, err := opts.ToLoadBalancerActionRebootMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(rebootURL(c, id), b, nil, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// ResetPasswordOpts represents the attributes used when resetting password of load_balancer instance.
type ResetPasswordOpts struct {

	// should syslog record acl info
	Username string `json:"username" required:"true"`
}

// ToLoadBalancerActionResetPasswordMap builds a request body from ResetPasswordOpts.
func (opts ResetPasswordOpts) ToLoadBalancerActionResetPasswordMap() (map[string]interface{}, error) {
	b, err := eclcloud.BuildRequestBody(opts, "")
	if err != nil {
		return nil, err
	}

	return b, nil
}

// ResetPassword accepts a ResetPasswordOpts struct and resets an existing Load Balancer password using the
// values provided.
func ResetPassword(c *eclcloud.ServiceClient, id string, opts ResetPasswordOpts) (r ResetPasswordResult) {
	b, err := opts.ToLoadBalancerActionResetPasswordMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(resetPasswordURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
