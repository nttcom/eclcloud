package load_balancer_plans

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result and extracts a Load Balancer Plan resource.
func (r commonResult) Extract() (*LoadBalancerPlan, error) {
	var s struct {
		LoadBalancerPlan *LoadBalancerPlan `json:"load_balancer_plan"`
	}
	err := r.ExtractInto(&s)
	return s.LoadBalancerPlan, err
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Load Balancer Plan.
type GetResult struct {
	commonResult
}

// Model of Load Balancer.
type Model struct {
	Edition     string `json:"edition"`
	Size        string `json:"size"`
}

// LoadBalancerPlan represents a Load Balancer Plan. See package documentation for a top-level
// description of what this is.
type LoadBalancerPlan struct {

	// Description is description
	Description string `json:"description"`

	// Is user allowed to create new load balancers with this plan.
	Enabled bool `json:"enabled"`

	// UUID representing the Load Balancer Plan.
	ID string `json:"id"`

	// Maximum number of syslog servers
	MaximumSyslogServers int `json:"maximum_syslog_servers"`

	// Model of load balancer
	Model Model `json:"model"`

	// Name of the Load Balancer Plan
	Name string `json:"name"`

	// Load Balancer Type
	Vendor string `json:"vendor"`

	// Version name
	Version string `json:"version"`
}

// LoadBalancerPage is the page returned by a pager when traversing over a collection
// of load balancers.
type LoadBalancerPlanPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of load balancer Plans has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r LoadBalancerPlanPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"load_balancer_plans_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a LoadBalancerPlanPage struct is empty.
func (r LoadBalancerPlanPage) IsEmpty() (bool, error) {
	is, err := ExtractLoadBalancerPlans(r)
	return len(is) == 0, err
}

// ExtractLoadBalancerPlans accepts a Page struct, specifically a LoadBalancerPage struct,
// and extracts the elements into a slice of Load Balancer Plan structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractLoadBalancerPlans(r pagination.Page) ([]LoadBalancerPlan, error) {
	var s struct {
		LoadBalancerPlans []LoadBalancerPlan `json:"load_balancer_plans"`
	}
	err := (r.(LoadBalancerPlanPage)).ExtractInto(&s)
	return s.LoadBalancerPlans, err
}
