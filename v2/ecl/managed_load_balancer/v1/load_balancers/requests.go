/*
Generated by https://github.com/tamac-io/openapi-to-eclcloud-rb
*/
package load_balancers

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

/*
List Load Balancers
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the load balancer attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	Name string `q:"name"`

	// - Description of the resource
	Description string `q:"description"`

	// - Configuration status of the resource
	ConfigurationStatus string `q:"configuration_status"`

	// - Monitoring status of the load balancer
	MonitoringStatus string `q:"monitoring_status"`

	// - Operation status of the resource
	OperationStatus string `q:"operation_status"`

	// - The zone / group where the primary virtual server of load balancer is deployed
	PrimaryAvailabilityZone string `q:"primary_availability_zone"`

	// - The zone / group where the secondary virtual server of load balancer is deployed
	SecondaryAvailabilityZone string `q:"secondary_availability_zone"`

	// - Primary or secondary availability zone where the load balancer is currently running
	ActiveAvailabilityZone string `q:"active_availability_zone"`

	// - Revision of the load balancer
	Revision int `q:"revision"`

	// - ID of the plan
	PlanID string `q:"plan_id"`

	// - ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToLoadBalancerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToLoadBalancerListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToLoadBalancerListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of load balancers.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToLoadBalancerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return LoadBalancerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Create Load Balancer
*/

// CreateOptsReservedFixedIP represents reserved_fixed_ip information in the load balancer creation.
type CreateOptsReservedFixedIP struct {

	// - The IP address assign to this interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	// - Must be specified the unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
	// - Must not be specified the network IP address and broadcast IP address
	IPAddress string `json:"ip_address"`
}

// CreateOptsInterface represents interface information in the load balancer creation.
type CreateOptsInterface struct {

	// - ID of the network that this interface belongs to
	// - Must be specified the unique network ID in `interfaces`
	// - Must be specified the network that plane is data
	// - Must not be specified the network that uses ISP shared address (RFC 6598)
	NetworkID string `json:"network_id"`

	// - Virtual IP address of the interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	// - Must be specified the unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
	// - Must not be specified the network IP address and broadcast IP address
	// - Must not be specified for existing interfaces
	VirtualIPAddress string `json:"virtual_ip_address"`

	// - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
	// - Must not be specified for existing interfaces
	ReservedFixedIPs *[]CreateOptsReservedFixedIP `json:"reserved_fixed_ips"`
}

// CreateOptsSyslogServer represents syslog_server information in the load balancer creation.
type CreateOptsSyslogServer struct {

	// - IP address of the syslog server
	// - The load balancer sends ICMP to this IP address for health check purpose
	IPAddress string `json:"ip_address"`

	// - Port number of the syslog server
	Port int `json:"port,omitempty"`

	// - Protocol of the syslog server
	// - Must be specified same protocol in all syslog servers belongs to the same load balancer
	Protocol string `json:"protocol,omitempty"`
}

// CreateOpts represents options used to create a new load balancer.
type CreateOpts struct {

	// - Name of the load balancer
	Name string `json:"name,omitempty"`

	// - Description of the load balancer
	Description string `json:"description,omitempty"`

	// - Tags of the load balancer
	// - Must be specified as JSON object
	Tags map[string]string `json:"tags,omitempty"`

	// - ID of the plan
	PlanID string `json:"plan_id,omitempty"`

	// - Syslog servers to which access logs are transferred
	// - Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
	//   - When `protocol` of `syslog_servers` is `"tcp"`
	//     - Access logs are transferred to all healthy syslog servers specified in `syslog_servers`
	//   - When `protocol` of `syslog_servers` is `"udp"`
	//     - Access logs are transferred to the syslog server specified first in `syslog_servers` as long as it is healthy
	//     - Access logs are transferred to the syslog server specified second (last) in `syslog_servers` when the first syslog server is not healthy
	SyslogServers *[]CreateOptsSyslogServer `json:"syslog_servers,omitempty"`

	// - Interfaces that attached to the load balancer
	// - `virtual_ip_address` and `reserved_fixed_ips` can not be changed once attached
	//   - To change `virtual_ip_address` and `reserved_fixed_ips` , recreating the interface is needed
	Interfaces *[]CreateOptsInterface `json:"interfaces,omitempty"`
}

// ToLoadBalancerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToLoadBalancerCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "load_balancer")
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToLoadBalancerCreateMap() (map[string]interface{}, error)
}

// Create accepts a CreateOpts struct and creates a new load balancer using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToLoadBalancerCreateMap()
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
Show Load Balancer
*/

// ShowOpts represents options used to show a load balancer.
type ShowOpts struct {

	// - When `true` is specified, `current` and `staged` are returned in response body
	Changes bool `q:"changes"`
}

// ToLoadBalancerShowQuery formats a ShowOpts into a query string.
func (opts ShowOpts) ToLoadBalancerShowQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ShowOptsBuilder allows extensions to add additional parameters to the Show request.
type ShowOptsBuilder interface {
	ToLoadBalancerShowQuery() (string, error)
}

// Show retrieves a specific load balancer based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string, opts ShowOptsBuilder) (r ShowResult) {
	url := showURL(c, id)

	if opts != nil {
		query, _ := opts.ToLoadBalancerShowQuery()
		url += query
	}

	_, r.Err = c.Get(url, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Load Balancer Attributes
*/

// UpdateOpts represents options used to update a existing load balancer.
type UpdateOpts struct {

	// - Name of the load balancer
	Name *string `json:"name,omitempty"`

	// - Description of the load balancer
	Description *string `json:"description,omitempty"`

	// - Tags of the load balancer
	// - Must be specified as JSON object
	Tags *map[string]string `json:"tags,omitempty"`
}

// ToLoadBalancerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToLoadBalancerUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "load_balancer")
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToLoadBalancerUpdateMap() (map[string]interface{}, error)
}

// Update accepts a UpdateOpts struct and updates a existing load balancer using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToLoadBalancerUpdateMap()
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
Delete Load Balancer
*/

// Delete accepts a unique ID and deletes the load balancer associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Action Load Balancer
*/

// ApplyConfigurations performs action on a existing load balancer.
func ApplyConfigurations(c *eclcloud.ServiceClient, id string) (r ActionResult) {
	b := map[string]interface{}{"apply-configurations": nil}
	_, r.Err = c.Post(actionURL(c, id), b, nil, &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

// SystemUpdateOpts represents options used to perform action on a existing load balancer.
type SystemUpdateOpts struct {

	// - ID of the system update that will be applied to the load balancer
	SystemUpdateID string `json:"system_update_id"`
}

// ToLoadBalancerSystemUpdateMap builds a request body from SystemUpdateOpts.
func (opts SystemUpdateOpts) ToLoadBalancerSystemUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "system-update")
}

// SystemUpdateOptsBuilder allows extensions to add additional parameters to the SystemUpdate request.
type SystemUpdateOptsBuilder interface {
	ToLoadBalancerSystemUpdateMap() (map[string]interface{}, error)
}

// SystemUpdate accepts a SystemUpdateOpts struct and performs action on a existing load balancer using the values provided.
func SystemUpdate(c *eclcloud.ServiceClient, id string, opts SystemUpdateOptsBuilder) (r ActionResult) {
	b, err := opts.ToLoadBalancerSystemUpdateMap()
	if err != nil {
		r.Err = err

		return
	}

	_, r.Err = c.Post(actionURL(c, id), b, nil, &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

// CancelConfigurations performs action on a existing load balancer.
func CancelConfigurations(c *eclcloud.ServiceClient, id string) (r ActionResult) {
	b := map[string]interface{}{"cancel-configurations": nil}
	_, r.Err = c.Post(actionURL(c, id), b, nil, &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}

/*
Create Staged Load Balancer Configurations
*/

// CreateStagedOptsReservedFixedIP represents reserved_fixed_ip information in the load balancer configurations creation.
type CreateStagedOptsReservedFixedIP struct {

	// - The IP address assign to this interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	// - Must be specified the unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
	// - Must not be specified the network IP address and broadcast IP address
	IPAddress string `json:"ip_address"`
}

// CreateStagedOptsInterface represents interface information in the load balancer configurations creation.
type CreateStagedOptsInterface struct {

	// - ID of the network that this interface belongs to
	// - Must be specified the unique network ID in `interfaces`
	// - Must be specified the network that plane is data
	// - Must not be specified the network that uses ISP shared address (RFC 6598)
	NetworkID string `json:"network_id"`

	// - Virtual IP address of the interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	// - Must be specified the unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
	// - Must not be specified the network IP address and broadcast IP address
	// - Must not be specified for existing interfaces
	VirtualIPAddress string `json:"virtual_ip_address"`

	// - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
	// - Must not be specified for existing interfaces
	ReservedFixedIPs *[]CreateStagedOptsReservedFixedIP `json:"reserved_fixed_ips"`
}

// CreateStagedOptsSyslogServer represents syslog_server information in the load balancer configurations creation.
type CreateStagedOptsSyslogServer struct {

	// - IP address of the syslog server
	// - The load balancer sends ICMP to this IP address for health check purpose
	IPAddress string `json:"ip_address"`

	// - Port number of the syslog server
	Port int `json:"port,omitempty"`

	// - Protocol of the syslog server
	// - Must be specified same protocol in all syslog servers belongs to the same load balancer
	Protocol string `json:"protocol,omitempty"`
}

// CreateStagedOpts represents options used to create new load balancer configurations.
type CreateStagedOpts struct {

	// - Syslog servers to which access logs are transferred
	// - Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
	//   - When `protocol` of `syslog_servers` is `"tcp"`
	//     - Access logs are transferred to all healthy syslog servers specified in `syslog_servers`
	//   - When `protocol` of `syslog_servers` is `"udp"`
	//     - Access logs are transferred to the syslog server specified first in `syslog_servers` as long as it is healthy
	//     - Access logs are transferred to the syslog server specified second (last) in `syslog_servers` when the first syslog server is not healthy
	SyslogServers *[]CreateStagedOptsSyslogServer `json:"syslog_servers,omitempty"`

	// - Interfaces that attached to the load balancer
	// - `virtual_ip_address` and `reserved_fixed_ips` can not be changed once attached
	//   - To change `virtual_ip_address` and `reserved_fixed_ips` , recreating the interface is needed
	Interfaces *[]CreateStagedOptsInterface `json:"interfaces,omitempty"`
}

// ToLoadBalancerCreateStagedMap builds a request body from CreateStagedOpts.
func (opts CreateStagedOpts) ToLoadBalancerCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "load_balancer")
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToLoadBalancerCreateStagedMap() (map[string]interface{}, error)
}

// CreateStaged accepts a CreateStagedOpts struct and creates new load balancer configurations using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateStagedResult) {
	b, err := opts.ToLoadBalancerCreateStagedMap()
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
Show Staged Load Balancer Configurations
*/

// ShowStaged retrieves specific load balancer configurations based on its unique ID.
func ShowStaged(c *eclcloud.ServiceClient, id string) (r ShowStagedResult) {
	_, r.Err = c.Get(showStagedURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}

/*
Update Staged Load Balancer Configurations
*/

// UpdateStagedOptsReservedFixedIP represents reserved_fixed_ip information in load balancer configurations updation.
type UpdateStagedOptsReservedFixedIP struct {

	// - The IP address assign to this interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	// - Must be specified the unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
	// - Must not be specified the network IP address and broadcast IP address
	IPAddress *string `json:"ip_address"`
}

// UpdateStagedOptsInterface represents interface information in load balancer configurations updation.
type UpdateStagedOptsInterface struct {

	// - ID of the network that this interface belongs to
	// - Must be specified the unique network ID in `interfaces`
	// - Must be specified the network that plane is data
	// - Must not be specified the network that uses ISP shared address (RFC 6598)
	NetworkID *string `json:"network_id"`

	// - Virtual IP address of the interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	// - Must be specified the unique IP address in `virtual_ip_address` and `reserved_fixed_ips`
	// - Must not be specified the network IP address and broadcast IP address
	// - Must not be specified for existing interfaces
	VirtualIPAddress *string `json:"virtual_ip_address"`

	// - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
	// - Must not be specified for existing interfaces
	ReservedFixedIPs *[]UpdateStagedOptsReservedFixedIP `json:"reserved_fixed_ips"`
}

// UpdateStagedOptsSyslogServer represents syslog_server information in load balancer configurations updation.
type UpdateStagedOptsSyslogServer struct {

	// - IP address of the syslog server
	// - The load balancer sends ICMP to this IP address for health check purpose
	IPAddress *string `json:"ip_address"`

	// - Port number of the syslog server
	Port *int `json:"port,omitempty"`

	// - Protocol of the syslog server
	// - Must be specified same protocol in all syslog servers belongs to the same load balancer
	Protocol *string `json:"protocol,omitempty"`
}

// UpdateStagedOpts represents options used to update existing Load Balancer configurations.
type UpdateStagedOpts struct {

	// - Syslog servers to which access logs are transferred
	// - Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
	//   - When `protocol` of `syslog_servers` is `"tcp"`
	//     - Access logs are transferred to all healthy syslog servers specified in `syslog_servers`
	//   - When `protocol` of `syslog_servers` is `"udp"`
	//     - Access logs are transferred to the syslog server specified first in `syslog_servers` as long as it is healthy
	//     - Access logs are transferred to the syslog server specified second (last) in `syslog_servers` when the first syslog server is not healthy
	SyslogServers *[]UpdateStagedOptsSyslogServer `json:"syslog_servers,omitempty"`

	// - Interfaces that attached to the load balancer
	// - `virtual_ip_address` and `reserved_fixed_ips` can not be changed once attached
	//   - To change `virtual_ip_address` and `reserved_fixed_ips` , recreating the interface is needed
	Interfaces *[]UpdateStagedOptsInterface `json:"interfaces,omitempty"`
}

// ToLoadBalancerUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToLoadBalancerUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "load_balancer")
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToLoadBalancerUpdateStagedMap() (map[string]interface{}, error)
}

// UpdateStaged accepts a UpdateStagedOpts struct and updates existing Load Balancer configurations using the values provided.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateStagedResult) {
	b, err := opts.ToLoadBalancerUpdateStagedMap()
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
Cancel Staged Load Balancer Configurations
*/

// CancelStaged accepts a unique ID and deletes load balancer configurations associated with it.
func CancelStaged(c *eclcloud.ServiceClient, id string) (r CancelStagedResult) {
	_, r.Err = c.Delete(cancelStagedURL(c, id), &eclcloud.RequestOpts{
		OkCodes: []int{204},
	})

	return
}
