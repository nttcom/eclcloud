package health_monitors

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a health monitor resource.
func (r commonResult) Extract() (*HealthMonitor, error) {
	var healthMonitor HealthMonitor
	err := r.ExtractInto(&healthMonitor)
	return &healthMonitor, err
}

// Extract interprets any commonResult as a Health Monitor, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "health_monitor")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Health Monitor.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Health Monitor.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Health Monitor.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// Configuration represents current and staged configurations.
type Configuration struct {

	// Port number of the health monitor for healthchecking
	// Returns 0 when protocol is "icmp"
	Port int `json:"port"`

	// Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol"`

	// Interval of healthchecking (in seconds)
	Interval int `json:"interval"`

	// Retry count of healthchecking
	// Initial monitoring is not included
	// Retry is executed at the interval specified by interval
	Retry int `json:"retry"`

	// Timeout of healthchecking (in seconds)
	Timeout int `json:"timeout"`
}

// HealthMonitor represents a health monitor.
type HealthMonitor struct {
	// ID of the health monitor
	ID string `json:"id"`

	// Name of the health monitor
	Name string `json:"name"`

	// Description of the health monitor
	Description string `json:"description"`

	// Tags of the health monitor (JSON object format)
	Tags map[string]string `json:"tags"`

	// Configuration status of the health monitor
	ConfigurationStatus string `json:"configuration_status"`

	// Operation status of the load balancer which the health monitor belongs to
	OperationStatus string `json:"operation_status"`

	// ID of the load balancer which the health monitor belongs to
	LoadBalancerID string `json:"load_balancer_id"`

	// ID of the owner tenant of the health monitor
	TenantID string `json:"tenant_id"`

	// Port number of the health monitor for healthchecking
	// Returns 0 when protocol is "icmp"
	Port int `json:"port"`

	// Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol"`

	// Interval of healthchecking (in seconds)
	Interval int `json:"interval"`

	// Retry count of healthchecking
	// Initial monitoring is not included
	// Retry is executed at the interval specified by interval
	Retry int `json:"retry"`

	// Timeout of healthchecking (in seconds)
	Timeout int `json:"timeout"`

	// Running configurations of the health monitor
	// Return object when changes is true
	// Return null when current configuretion does not exist
	Current *Configuration `json:"current,omitempty"`

	// Added or changed configurations of the health monitor that waiting to be applied
	// Return object when changes is true
	// Return null when staged configuretion does not exist
	Staged *Configuration `json:"staged,omitempty"`
}

// HealthMonitorPage is the page returned by a pager
// when traversing over a collection of health monitor.
type HealthMonitorPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a HealthMonitorPage struct is empty.
func (r HealthMonitorPage) IsEmpty() (bool, error) {
	is, err := ExtractHealthMonitors(r)
	return len(is) == 0, err
}

// ExtractHealthMonitors accepts a Page struct,
// specifically a HealthMonitorPage struct, and extracts the elements
// into a slice of Health Monitor structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractHealthMonitors(r pagination.Page) ([]HealthMonitor, error) {
	var s struct {
		HealthMonitors []HealthMonitor `json:"health_monitors"`
	}
	err := (r.(HealthMonitorPage)).ExtractInto(&s)
	return s.HealthMonitors, err
}
