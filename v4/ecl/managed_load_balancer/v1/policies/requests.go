package policies

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

/*
List Policies
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the policy attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts single-byte characters only
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts single-byte characters only
	Description string `q:"description"`

	// - Configuration status of the resource
	ConfigurationStatus string `q:"configuration_status"`

	// - Operation status of the resource
	OperationStatus string `q:"operation_status"`

	// - Load balancing algorithm (method) of the policy
	Algorithm string `q:"algorithm"`

	// - Persistence setting of the policy
	Persistence string `q:"persistence"`

	// - The duration (in seconds) during which a session is allowed to remain inactive
	IdleTimeout int `q:"idle_timeout"`

	// - URL of the sorry page to which accesses are redirected if all members in the target group are down
	SorryPageUrl string `q:"sorry_page_url"`

	// - Source NAT setting of the policy
	SourceNat string `q:"source_nat"`

	// - ID of the certificate that assigned to the policy
	CertificateID string `q:"certificate_id"`

	// - ID of the health monitor that assigned to the policy
	HealthMonitorID string `q:"health_monitor_id"`

	// - ID of the listener that assigned to the policy
	ListenerID string `q:"listener_id"`

	// - ID of the default target group that assigned to the policy
	DefaultTargetGroupID string `q:"default_target_group_id"`

	// - ID of the TLS policy that assigned to the policy
	TLSPolicyID string `q:"tls_policy_id"`

	// - ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of policies.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return PolicyPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Policy
*/

// CreateOpts represents options used to create a new policy.
type CreateOpts struct {

	// - Name of the policy
	// - This field accepts single-byte characters only
	Name string `json:"name,omitempty"`

	// - Description of the policy
	// - This field accepts single-byte characters only
	Description string `json:"description,omitempty"`

	// - Tags of the policy
	// - Set JSON object up to 32,768 characters
	//   - Nested structure is permitted
	// - This field accepts single-byte characters only
	Tags map[string]interface{} `json:"tags,omitempty"`

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
	// - If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
	// - If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
	SorryPageUrl string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat string `json:"source_nat,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - You can set a ID of the certificate in which `ca_cert.status`, `ssl_cert.status` and `ssl_key.status` are all `"UPLOADED"`
	// - If `listener.protocol` is `"https"`, set `certificate.id`
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	CertificateID string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	// - Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
	HealthMonitorID string `json:"health_monitor_id"`

	// - ID of the listener that assigned to the policy
	// - Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
	// - Must not set ID of the listener that already assigned to the other policy
	ListenerID string `json:"listener_id"`

	// - ID of the default target group that assigned to the policy
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	DefaultTargetGroupID string `json:"default_target_group_id"`

	// - ID of the TLS policy that assigned to the policy
	// - If `listener.protocol` is `"https"`, you can set this parameter explicitly
	//   - If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	TLSPolicyID string `json:"tls_policy_id,omitempty"`

	// - ID of the load balancer which the policy belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToPolicyCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new policy using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToPolicyCreateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createURL(c), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Policy
*/

// ShowOpts represents options used to show a policy.
type ShowOpts struct {

	// - If `true` is set, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToPolicyShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToPolicyShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToPolicyShowQuery() (string, error)
}

// Show retrieves a specific policy based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToPolicyShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Policy Attributes
*/

// UpdateOpts represents options used to update a existing policy.
type UpdateOpts struct {

	// - Name of the policy
	// - This field accepts single-byte characters only
	Name *string `json:"name,omitempty"`

	// - Description of the policy
	// - This field accepts single-byte characters only
	Description *string `json:"description,omitempty"`

	// - Tags of the policy
	// - Set JSON object up to 32,768 characters
	//   - Nested structure is permitted
	// - This field accepts single-byte characters only
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// ToPolicyUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing policy using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPolicyUpdateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Delete Policy
*/

// Delete accepts a unique ID and deletes the policy associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Policy Configurations
*/

// CreateStagedOpts represents options used to create new policy configurations.
type CreateStagedOpts struct {

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
	// - If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
	// - If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"http"` or `"https"` to others, set `""`
	SorryPageUrl string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat string `json:"source_nat,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - You can set a ID of the certificate in which `ca_cert.status`, `ssl_cert.status` and `ssl_key.status` are all `"UPLOADED"`
	// - If `listener.protocol` is `"https"`, set `certificate.id`
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	CertificateID string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	// - Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
	HealthMonitorID string `json:"health_monitor_id,omitempty"`

	// - ID of the listener that assigned to the policy
	// - Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
	// - Must not set ID of the listener that already assigned to the other policy
	ListenerID string `json:"listener_id,omitempty"`

	// - ID of the default target group that assigned to the policy
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	DefaultTargetGroupID string `json:"default_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If `listener.protocol` is `"https"`, you can set this parameter explicitly
	//   - If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	TLSPolicyID string `json:"tls_policy_id,omitempty"`
}

// ToPolicyCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToPolicyCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToPolicyCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new policy configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToPolicyCreateStagedMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(createStagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Show Staged Policy Configurations
*/

// ShowStaged retrieves specific policy configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Policy Configurations
*/

// UpdateStagedOpts represents options used to update existing Policy configurations.
type UpdateStagedOpts struct {

	// - Load balancing algorithm (method) of the policy
	Algorithm *string `json:"algorithm,omitempty"`

	// - Persistence setting of the policy
	// - If `listener.protocol` is `"http"` or `"https"`, `"cookie"` is available
	Persistence *string `json:"persistence,omitempty"`

	// - The duration (in seconds) during which a session is allowed to remain inactive
	// - There may be a time difference up to 60 seconds, between the set value and the actual timeout
	// - If `listener.protocol` is `"tcp"` or `"udp"`
	//   - Default value is 120
	// - If `listener.protocol` is `"http"` or `"https"`
	//   - Default value is 600
	//   - On session timeout, the load balancer sends TCP RST packets to both the client and the real server
	IdleTimeout *int `json:"idle_timeout,omitempty"`

	// - URL of the sorry page to which accesses are redirected if all members in the target group are down
	// - If `listener.protocol` is `"http"` or `"https"`, this parameter can be set
	// - If `listener.protocol` is neither `"http"` nor `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"http"` or `"https"` to others, set `""`
	SorryPageUrl *string `json:"sorry_page_url,omitempty"`

	// - Source NAT setting of the policy
	// - If `source_nat` is `"enable"` and `listener.protocol` is `"http"` or `"https"`
	//   - The source IP address of the request is replaced with `virtual_ip_address` which is assigned to the interface from which the request was sent
	//   - `X-Forwarded-For` header with the IP address of the client is added
	SourceNat *string `json:"source_nat,omitempty"`

	// - ID of the certificate that assigned to the policy
	// - You can set a ID of the certificate in which `ca_cert.status`, `ssl_cert.status` and `ssl_key.status` are all `"UPLOADED"`
	// - If `listener.protocol` is `"https"`, set `certificate.id`
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	CertificateID *string `json:"certificate_id,omitempty"`

	// - ID of the health monitor that assigned to the policy
	// - Must not set ID of the health monitor that `configuration_status` is `"DELETE_STAGED"`
	HealthMonitorID *string `json:"health_monitor_id,omitempty"`

	// - ID of the listener that assigned to the policy
	// - Must not set ID of the listener that `configuration_status` is `"DELETE_STAGED"`
	// - Must not set ID of the listener that already assigned to the other policy
	ListenerID *string `json:"listener_id,omitempty"`

	// - ID of the default target group that assigned to the policy
	// - Must not set ID of the target group that `configuration_status` is `"DELETE_STAGED"`
	DefaultTargetGroupID *string `json:"default_target_group_id,omitempty"`

	// - ID of the TLS policy that assigned to the policy
	// - If `listener.protocol` is `"https"`, you can set this parameter explicitly
	//   - If not set this parameter, the ID of the `tls_policy` with `default: true` will be automatically set
	// - If `listener.protocol` is not `"https"`, must not set this parameter or set `""`
	//   - If you change `listener.protocol` from `"https"` to others, set `""`
	TLSPolicyID *string `json:"tls_policy_id,omitempty"`
}

// ToPolicyUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToPolicyUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "policy")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToPolicyUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Policy configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToPolicyUpdateStagedMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Patch(updateStagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Cancel Staged Policy Configurations
*/

// CancelStaged accepts a unique ID and deletes policy configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
