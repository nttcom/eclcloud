package listeners

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a listener resource.
func (r commonResult) Extract() (*Listener, error) {
	var listener Listener
	err := r.ExtractInto(&listener)
	return &listener, err
}

// Extract interprets any commonResult as a Listener, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "listener")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Listener.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Listener.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Listener.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// Listener represents a listener.
type Listener struct {
	// ID of the listener
	ID string `json:"id,omitempty"`

	// Name of the listener
	Name string `json:"name,omitempty"`

	// Description of the listener
	Description string `json:"description,omitempty"`

	// Tags of the listener (JSON object format)
	Tags map[string]string `json:"tags,omitempty"`

	// Configuration status of the listener
	ConfigurationStatus string `json:"configuration_status,omitempty"`

	// Operation status of the load balancer which the listener belongs to
	OperationStatus string `json:"operation_status,omitempty"`

	// ID of the load balancer which the listener belongs to
	LoadBalancerID string `json:"load_balancer_id,omitempty"`

	// ID of the owner tenant of the listener
	TenantID string `json:"tenant_id,omitempty"`

	// IP address of the listener for listening
	IPAddress string `json:"ip_address"`

	// Port number of the listener for listening
	Port int `json:"port"`

	// Protocol of the listener for listening
	Protocol string `json:"protocol"`

	// Running configurations of the listener
	// Return object when changes is true
	// Return null when current configuretion does not exist
	Current *Listener `json:"current,omitempty"`

	// Added or changed configurations of the Listener that waiting to be applied
	// Return object when changes is true
	// Return null when staged configuretion does not exist
	Staged *Listener `json:"staged,omitempty"`
}

// ListenerPage is the page returned by a pager
// when traversing over a collection of listener.
type ListenerPage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a ListenerPage struct is empty.
func (r ListenerPage) IsEmpty() (bool, error) {
	is, err := ExtractListeners(r)
	return len(is) == 0, err
}

// ExtractListeners accepts a Page struct,
// specifically a ListenerPage struct, and extracts the elements
// into a slice of Listener structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractListeners(r pagination.Page) ([]Listener, error) {
	var s struct {
		Listeners []Listener `json:"listeners"`
	}
	err := (r.(ListenerPage)).ExtractInto(&s)
	return s.Listeners, err
}
