package policies

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a Policy.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a Policy.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a Policy.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a Policy.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a Policy.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a Policy.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ConfigurationInResponse represents a configuration in a policy.
type ConfigurationInResponse struct {

	// - Load balancing algorithm (method) of the policy
	Algorithm string `json:"algorithm,omitempty"`

	// - Persistence setting of the policy
	// - If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
	Persistence string `json:"persistence,omitempty"`

	// - The duration (in seconds) during which a session is allowed to remain inactive
	// - There may be a time difference up to 60 seconds, between the set value and the actual timeout
	// - If `listener.protocol` is `"tcp"` or `"udp"`
	//   - Default value is 120
	// - If `listener.protocol` is `"http"` or `"https"`
	//   - Default value is 600
	//   - On session timeout, the load balancer sends TCP RST packets to both the client and the real server
	IdleTimeout int `json:"idle_timeout,omitempty"`

	// - URL of the sorry page to which accesses are redirected if all members in the target group are down
	// - If protocol is not `"http"` or `"https"`, returns `""`
	SorryPageUrl string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"` ,
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat string `json:"source_nat,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - If protocol is not `"https"`, returns `""`
	CertificateID string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	HealthMonitorID string `json:"health_monitor_id,omitempty"`

	// - ID of the listener that assigned to the policy
	ListenerID string `json:"listener_id,omitempty"`

	// - ID of the default target group that assigned to the policy
	DefaultTargetGroupID string `json:"default_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If protocol is not `"https"`, returns `""`
	TLSPolicyID string `json:"tls_policy_id,omitempty"`
}

// Policy represents a policy.
type Policy struct {

	// - ID of the policy
	ID string `json:"id"`

	// - Name of the policy
	Name string `json:"name"`

	// - Description of the policy
	Description string `json:"description"`

	// - Tags of the policy (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the policy
	//   - `"ACTIVE"`
	//     - There are no configurations of the policy that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The policy has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the policy exists that waiting to be applied
	//   - `"DELETE_STAGED"`
	//     - The policy has been removed and waiting to be applied
	ConfigurationStatus string `json:"configuration_status"`

	// - Operation status of the load balancer which the policy belongs to
	//   - `"NONE"` :
	//     - There are no operations of the load balancer
	//     - The load balancer and related resources can be operated
	//   - `"PROCESSING"`
	//     - The latest operation of the load balancer is processing
	//     - The load balancer and related resources cannot be operated
	//   - `"COMPLETE"`
	//     - The latest operation of the load balancer has been succeeded
	//     - The load balancer and related resources can be operated
	//   - `"STUCK"`
	//     - The latest operation of the load balancer has been stopped
	//     - Operators of NTT Communications will investigate the operation
	//     - The load balancer and related resources cannot be operated
	//   - `"ERROR"`
	//     - The latest operation of the load balancer has been failed
	//     - The operation was roll backed normally
	//     - The load balancer and related resources can be operated
	OperationStatus string `json:"operation_status"`

	// - ID of the load balancer which the policy belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// - ID of the owner tenant of the policy
	TenantID string `json:"tenant_id"`

	// - Load balancing algorithm (method) of the policy
	Algorithm string `json:"algorithm,omitempty"`

	// - Persistence setting of the policy
	// - If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
	Persistence string `json:"persistence,omitempty"`

	// - The duration (in seconds) during which a session is allowed to remain inactive
	// - There may be a time difference up to 60 seconds, between the set value and the actual timeout
	// - If `listener.protocol` is `"tcp"` or `"udp"`
	//   - Default value is 120
	// - If `listener.protocol` is `"http"` or `"https"`
	//   - Default value is 600
	//   - On session timeout, the load balancer sends TCP RST packets to both the client and the real server
	IdleTimeout int `json:"idle_timeout,omitempty"`

	// - URL of the sorry page to which accesses are redirected if all members in the target group are down
	// - If protocol is not `"http"` or `"https"`, returns `""`
	SorryPageUrl string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"` ,
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat string `json:"source_nat,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - If protocol is not `"https"`, returns `""`
	CertificateID string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	HealthMonitorID string `json:"health_monitor_id,omitempty"`

	// - ID of the listener that assigned to the policy
	ListenerID string `json:"listener_id,omitempty"`

	// - ID of the default target group that assigned to the policy
	DefaultTargetGroupID string `json:"default_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If protocol is not `"https"`, returns `""`
	TLSPolicyID string `json:"tls_policy_id,omitempty"`

	// - Running configurations of the policy
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the policy that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a policy, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "policy")
}

// Extract is a function that accepts a result and extracts a Policy resource.
func (r commonResult) Extract() (*Policy, error) {
	var policy Policy

	err := r.ExtractInto(&policy)

	return &policy, err
}

// PolicyPage is the page returned by a pager when traversing over a collection of policy.
type PolicyPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a PolicyPage struct is empty.
func (r PolicyPage) IsEmpty() (bool, error) {
	is, err := ExtractPolicies(r)

	return len(is) == 0, err
}

// ExtractPoliciesInto interprets the results of a single page from a List() call, producing a slice of policy entities.
func ExtractPoliciesInto(r pagination.Page, v interface{}) error {
	return r.(PolicyPage).Result.ExtractIntoSlicePtr(v, "policies")
}

// ExtractPolicies accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of Policy structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractPolicies(r pagination.Page) ([]Policy, error) {
	var s []Policy

	err := ExtractPoliciesInto(r, &s)

	return s, err
}
