package load_balancers

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result and extracts a Load Balancer resource.
func (r commonResult) Extract() (*LoadBalancer, error) {
	var s struct {
		LoadBalancer *LoadBalancer `json:"load_balancer"`
	}
	err := r.ExtractInto(&s)
	return s.LoadBalancer, err
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Load Balancer.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Load Balancer.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Load Balancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// Interface represents a attached interfaces by Load Balancer.
type Interface struct {
	ID                  string `json:"id"`
	IPAddress           string `json:"ip_address"`
	Name                string `json:"name"`
	NetworkID           string `json:"network_id"`
	SlotNumber          int    `json:"slot_number"`
	Status              string `json:"status"`
	Type                string `json:"type"`
	VirtualIPAddress    string `json:"virtual_ip_address"`
	VirtualIPProperties string `json:"virtual_ip_properties"`
}

// Interface represents a attached interfaces by Load Balancer.
type SyslogServer struct {
	ID          string `json:"id"`
	IPAddress   string `json:"ip_address"`
	LogFacility string `json:"log_facility"`
	LogLevel    string `json:"log_level"`
	Name        string `json:"name"`
	PortNumber  int    `json:"port_number"`
	Status      string `json:"status"`
}

// LoadBalancer represents a Load Balancer. See package documentation for a top-level
// description of what this is.
type LoadBalancer struct {

	// AdminPassword is admin's password
	AdminPassword string `json:"admin_password"`

	// AdminUsername is admin's username
	AdminUsername string `json:"admin_username"`

	// AvailabilityZone is one of the Virtual Server (Nova)’s availability zone.
	AvailabilityZone string `json:"availability_zone"`

	// Description is description
	DefaultGateway string `json:"default_gateway"`

	// Description is description
	Description string `json:"description"`

	// UUID representing the Load Balancer.
	ID string `json:"id"`

	// Attached interfaces by Load Balancer.
	Interfaces []Interface `json:"interfaces"`

	// LoadBalancerPlanID is the UUID of Load Balancer Plan.
	LoadBalancerPlanID string `json:"load_balancer_plan_id"`

	// Name of the Load Balancer.
	Name string `json:"name"`

	// The Load Balancer status.
	Status string `json:"status"`

	// Connected syslog servers.
	SyslogServers []SyslogServer `json:"syslog_servers"`

	// TenantID is the project owner of the Load Balancer.
	TenantID string `json:"tenant_id"`

	// User's password placeholder.
	UserPassword string `json:"user_password"`

	// User's username placeholder.
	UserUsername string `json:"user_username"`
}

// LoadBalancerPage is the page returned by a pager when traversing over a collection
// of load balancers.
type LoadBalancerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of load balancers has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r LoadBalancerPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"load_balancers_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a LoadBalancerPage struct is empty.
func (r LoadBalancerPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancers(r)
	return len(is) == 0, err
}

// ExtractLoadBalancers accepts a Page struct, specifically a LoadBalancerPage struct,
// and extracts the elements into a slice of Load Balancer structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) {
	var s struct {
		LoadBalancers []LoadBalancer `json:"load_balancers"`
	}
	err := (r.(LoadBalancerPage)).ExtractInto(&s)
	return s.LoadBalancers, err
}
