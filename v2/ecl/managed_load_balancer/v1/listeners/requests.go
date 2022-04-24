package listeners

import (
	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/pagination"
)

/*
Parameters for List
*/

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToListenerListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the listener attributes you want to see returned.
type ListOpts struct {

	// ID of the resource
	ID string `q:"id"`

	// Name of the resource
	Name string `q:"name"`

	// Description of the resource
	Description string `q:"description"`

	// Configuration status of the resource
	ConfigurationStatus string `q:"configuration_status"`

	// Operation status of the resource
	OperationStatus string `q:"operation_status"`

	// IP address of the resource for listening
	IPAddress string `q:"ip_address"`

	// Port number of the resource for healthchecking or listening
	Port int `q:"port"`

	// Protocol of the resource for healthchecking or listening
	Protocol string `q:"protocol"`

	// ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToListenerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToListenerListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of listeners.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToListenerListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return ListenerPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Parameters for Get
*/

// GetOptsBuilder allows extensions to add additional parameters to
// the listener API request
type GetOptsBuilder interface {
	ToListenerGetQuery() (string, error)
}

// GetOpts represents options used to show a listener.
type GetOpts struct {
	Changes bool `q:"changes"`
}

// ToListenerGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToListenerGetQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves a specific listener based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string, opts GetOptsBuilder) (r GetResult) {
	url := getURL(c, id)
	if opts != nil {
		query, _ := opts.ToListenerGetQuery()
		url += query
	}
	_, r.Err = c.Get(url, &r.Body, nil)
	return
}

/*
Parameters for Create
*/

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToListenerCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a listener.
type CreateOpts struct {

	// Name of the listener
	Name string `json:"name,omitempty"`

	// Description of the listener
	Description string `json:"description,omitempty"`

	// Tags of the listener
	// Must be specified as JSON object
	Tags map[string]string `json:"tags,omitempty"`

	// IP address of the listener for listening
	// Must be specified the unique combination of ip address and port
	// in all listeners belongs to the same load balancer
	// Must be specified the ip address
	// which is not included in subnet of load balancer interfaces that the listener belongs to
	IPAddress string `json:"ip_address"`

	// Port number of the listener for listening
	// Must be specified the unique combination of ip address and port
	// in all listeners belongs to the same load balancer
	Port int `json:"port"`

	// Protocol of the listener for listening
	Protocol string `json:"protocol"`

	// ID of the load balancer which the listener belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToListenerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToListenerCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// Create accepts a CreateOpts struct and creates a new listener
// using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateMap()
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
Parameters for Update
*/

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToListenerUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update metadata of exisiting listener.
type UpdateOpts struct {

	// Name of the listener
	Name string `json:"name,omitempty"`

	// Description of the listener
	Description string `json:"description,omitempty"`

	// Tags of the listener
	// Must be specified as JSON object
	Tags map[string]string `json:"tags,omitempty"`
}

// ToListenerUpdateMap builds a request body from UpdateMOpts.
func (opts UpdateOpts) ToListenerUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// Update accepts a UpdateOpts struct and updates an existing listener
// using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToListenerUpdateMap()
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
Parameters for Delete
*/

// Delete accepts a unique ID and deletes the listener associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}

// IDFromName is a convenience function that returns a listener's
// ID, given its name.
func IDFromName(client *eclcloud.ServiceClient, name string) (string, error) {
	count := 0
	id := ""

	listOpts := ListOpts{
		Name: name,
	}

	pages, err := List(client, listOpts).AllPages()
	if err != nil {
		return "", err
	}

	all, err := ExtractListeners(pages)
	if err != nil {
		return "", err
	}

	for _, s := range all {
		if s.Name == name {
			count++
			id = s.ID
		}
	}

	switch count {
	case 0:
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "listener"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "listener"}
	}
}

/*
Parameters for Staged request
*/

// StagedOpts represents options used to create/update staged a listener.
type StagedOpts struct {

	// IP address of the listener for listening
	// Must be specified the unique combination of ip address and port
	// in all listeners belongs to the same load balancer
	// Must be specified the ip address
	// which is not included in subnet of load balancer interfaces that the listener belongs to
	IPAddress string `json:"ip_address,omitempty"`

	// Port number of the listener for healthchecking
	// Must be specified 0 when protocol is "icmp"
	Port int `json:"port,omitempty"`

	// Protocol of the listener for healthchecking
	Protocol string `json:"protocol,omitempty"`
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the
// CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToListenerCreateStagedMap() (map[string]interface{}, error)
}

type CreateStagedOpts StagedOpts

// ToListenerCreateStagedMap builds a request body from CreateOpts.
func (opts CreateStagedOpts) ToListenerCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// CreateStaged accepts a CreateStagedOpts struct and creates a new listener
// using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateResult) {
	b, err := opts.ToListenerCreateStagedMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(stagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetStaged retrieves a specific listener based on its unique ID.
func GetStaged(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(stagedURL(c, id), &r.Body, nil)
	return
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the
// UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToListenerUpdateStagedMap() (map[string]interface{}, error)
}

type UpdateStagedOpts StagedOpts

// ToListenerUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToListenerUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "listener")
}

// UpdateStaged accepts a UpdateOpts struct and updates an existing listener
// using the values provided. For more information, see the Create function.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateResult) {
	b, err := opts.ToListenerUpdateStagedMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Patch(stagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteStaged retrieves a specific listener based on its unique ID.
func DeleteStaged(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Delete(stagedURL(c, id), nil)
	return
}
