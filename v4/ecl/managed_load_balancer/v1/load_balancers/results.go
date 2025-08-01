package load_balancers

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// CreateResult represents the result of a Create operation.
// Call its Extract method to interpret it as a LoadBalancer.
type CreateResult struct {
	commonResult
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a LoadBalancer.
type ShowResult struct {
	commonResult
}

// UpdateResult represents the result of a Update operation.
// Call its Extract method to interpret it as a LoadBalancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a Delete operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// ActionResult represents the result of a Action operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type ActionResult struct {
	eclcloud.ErrResult
}

// CreateStagedResult represents the result of a CreateStaged operation.
// Call its Extract method to interpret it as a LoadBalancer.
type CreateStagedResult struct {
	commonResult
}

// ShowStagedResult represents the result of a ShowStaged operation.
// Call its Extract method to interpret it as a LoadBalancer.
type ShowStagedResult struct {
	commonResult
}

// UpdateStagedResult represents the result of a UpdateStaged operation.
// Call its Extract method to interpret it as a LoadBalancer.
type UpdateStagedResult struct {
	commonResult
}

// CancelStagedResult represents the result of a CancelStaged operation.
// Call its ExtractErr method to determine if the request succeeded or failed.
type CancelStagedResult struct {
	eclcloud.ErrResult
}

// ReservedFixedIPInResponse represents a reserved fixed ip in a load balancer.
type ReservedFixedIPInResponse struct {

	// - The IP address assign to this interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	IPAddress string `json:"ip_address"`
}

// ConfigurationInResponse represents a configuration in a load balancer.
type ConfigurationInResponse struct {

	// - Syslog servers to which access logs are transferred
	// - The facility code of syslog is 0 (kern), and the severity level is 6 (info)
	// - Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
	//   - If `protocol` of `syslog_servers` is `"tcp"`
	//     - Access logs are transferred to all healthy syslog servers set in `syslog_servers`
	//   - If `protocol` of `syslog_servers` is `"udp"`
	//     - Access logs are transferred to the syslog server set first in `syslog_servers` as long as it is healthy
	//     - Access logs are transferred to the syslog server set second (last) in `syslog_servers` if the first syslog server is not healthy
	SyslogServers []SyslogServerInResponse `json:"syslog_servers,omitempty"`

	// - Interfaces that attached to the load balancer
	Interfaces []InterfaceInResponse `json:"interfaces,omitempty"`
}

// InterfaceInResponse represents a interface in a load balancer.
type InterfaceInResponse struct {

	// - ID of the network that this interface belongs to
	NetworkID string `json:"network_id"`

	// - Virtual IP address of the interface within subnet
	// - Do not use this IP address at the interface of other devices, allowed address pairs, etc
	VirtualIPAddress string `json:"virtual_ip_address"`

	// - IP addresses that are pre-reserved for applying configurations of load balancer to be performed without losing redundancy
	ReservedFixedIPs []ReservedFixedIPInResponse `json:"reserved_fixed_ips"`
}

// SyslogServerInResponse represents a syslog server in a load balancer.
type SyslogServerInResponse struct {

	// - IP address of the syslog server
	// - The load balancer sends ICMP to this IP address for health check purpose
	IPAddress string `json:"ip_address"`

	// - Port number of the syslog server
	Port int `json:"port"`

	// - Protocol of the syslog server
	Protocol string `json:"protocol"`
}

// LoadBalancer represents a load balancer.
type LoadBalancer struct {

	// - ID of the load balancer
	ID string `json:"id"`

	// - Name of the load balancer
	Name string `json:"name"`

	// - Description of the load balancer
	Description string `json:"description"`

	// - Tags of the load balancer (JSON object format)
	Tags map[string]interface{} `json:"tags"`

	// - Configuration status of the load balancer
	//   - `"ACTIVE"`
	//     - There are no configurations of the load balancer that waiting to be applied
	//   - `"CREATE_STAGED"`
	//     - The load balancer has been added and waiting to be applied
	//   - `"UPDATE_STAGED"`
	//     - Changed configurations of the load balancer exists that waiting to be applied
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	ConfigurationStatus string `json:"configuration_status"`

	// - Monitoring status of the load balancer
	//   - `"ACTIVE"`
	//     - The load balancer is operating normally
	//   - `"INITIAL"`
	//     - The load balancer is not deployed and does not monitored
	//   - `"UNAVAILABLE"`
	//     - The load balancer is not operating normally
	MonitoringStatus string `json:"monitoring_status"`

	// - Operation status of the load balancer
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
	//     - The operators will investigate the operation
	//     - The load balancer and related resources cannot be operated
	//   - `"ERROR"`
	//     - The latest operation of the load balancer has been failed
	//     - The operation was roll backed normally
	//     - The load balancer and related resources can be operated
	// - For detail, refer to the API reference appendix
	//     - https://sdpf.ntt.com/services/docs/managed-lb/service-descriptions/api_reference_appendix.html
	OperationStatus string `json:"operation_status"`

	// - The zone / group where the primary virtual server of load balancer is deployed
	PrimaryAvailabilityZone string `json:"primary_availability_zone,omitempty"`

	// - The zone / group where the secondary virtual server of load balancer is deployed
	SecondaryAvailabilityZone string `json:"secondary_availability_zone,omitempty"`

	// - Primary or secondary availability zone where the load balancer is currently running
	// - If can not define active availability zone, returns `"UNDEFINED"`
	ActiveAvailabilityZone string `json:"active_availability_zone"`

	// - Revision of the load balancer
	Revision int `json:"revision"`

	// - ID of the plan
	PlanID string `json:"plan_id"`

	// - Name of the plan
	PlanName string `json:"plan_name"`

	// - ID of the owner tenant of the load balancer
	TenantID string `json:"tenant_id"`

	// - Syslog servers to which access logs are transferred
	// - The facility code of syslog is 0 (kern), and the severity level is 6 (info)
	// - Only access logs to listeners which `protocol` is either `"http"` or `"https"` are transferred
	//   - If `protocol` of `syslog_servers` is `"tcp"`
	//     - Access logs are transferred to all healthy syslog servers set in `syslog_servers`
	//   - If `protocol` of `syslog_servers` is `"udp"`
	//     - Access logs are transferred to the syslog server set first in `syslog_servers` as long as it is healthy
	//     - Access logs are transferred to the syslog server set second (last) in `syslog_servers` if the first syslog server is not healthy
	SyslogServers []SyslogServerInResponse `json:"syslog_servers,omitempty"`

	// - Interfaces that attached to the load balancer
	Interfaces []InterfaceInResponse `json:"interfaces,omitempty"`

	// - Running configurations of the load balancer
	// - If `changes` is `true`, return object
	// - If current configuration does not exist, return `null`
	Current ConfigurationInResponse `json:"current,omitempty"`

	// - Added or changed configurations of the load balancer that waiting to be applied
	// - If `changes` is `true`, return object
	// - If staged configuration does not exist, return `null`
	Staged ConfigurationInResponse `json:"staged,omitempty"`
}

// ExtractInto interprets any commonResult as a load balancer, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "load_balancer")
}

// Extract is a function that accepts a result and extracts a LoadBalancer resource.
func (r commonResult) Extract() (*LoadBalancer, error) {
	var loadBalancer LoadBalancer

	err := r.ExtractInto(&loadBalancer)

	return &loadBalancer, err
}

// LoadBalancerPage is the page returned by a pager when traversing over a collection of load balancer.
type LoadBalancerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a LoadBalancerPage struct is empty.
func (r LoadBalancerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancers(r)

	return len(is) == 0, err
}

// ExtractLoadBalancersInto interprets the results of a single page from a List() call, producing a slice of load balancer entities.
func ExtractLoadBalancersInto(r pagination.Page, v interface{}) error {
	return r.(LoadBalancerPage).Result.ExtractIntoSlicePtr(v, "load_balancers")
}

// ExtractLoadBalancers accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of LoadBalancer structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) {
	var s []LoadBalancer

	err := ExtractLoadBalancersInto(r, &s)

	return s, err
}
