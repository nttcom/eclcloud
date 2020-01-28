/*
Package load_balancer_actions contains functionality for working with
ECL Load Balancer/Actions resources.

Example to reboot a Load Balancer

	loadBalancerID := "9ab7ab3c-38a6-417c-926b-93772c4eb2f9"

	rebootOpts := load_balancer_actions.RebootOpts{
		Type: "HARD",
	}

	err := load_balancer_actions.Reboot(networkClient, loadBalancerID, rebootOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to reset password of Load Balancer

	loadBalancerID := "9ab7ab3c-38a6-417c-926b-93772c4eb2f9"

	resetPasswordOpts := load_balancer_actions.ResetPasswordOpts{
		Username: "user-read",
	}

	resetPasswordResult, err := load_balancer_actions.ResetPassword(networkClient, loadBalancerID, resetPasswordOpts).ExtractResetPassword()
	if err != nil {
		panic(err)
	}

*/
package load_balancer_actions
