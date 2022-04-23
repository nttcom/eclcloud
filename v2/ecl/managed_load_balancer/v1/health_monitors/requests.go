package health_monitors

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
	ToHealthMonitorListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the health monitor attributes you want to see returned.
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

	// Port number of the resource for healthchecking or listening
	Port int `q:"port"`

	// Protocol of the resource for healthchecking or listening
	Protocol string `q:"protocol"`

	// Interval of healthchecking (in seconds)
	Interval int `q:"interval"`

	// Retry count of healthchecking
	Retry int `q:"retry"`

	// Timeout of healthchecking (in seconds)
	Timeout int `q:"timeout"`

	// ID of the load balancer which the resource belongs to
	LoadBalancerID string `q:"load_balancer_id"`

	// ID of the owner tenant of the resource
	TenantID string `q:"tenant_id"`
}

// ToHealthMonitorListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToHealthMonitorListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List returns a Pager which allows you to iterate over a collection of health monitors.
// It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)
	if opts != nil {
		query, err := opts.ToHealthMonitorListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return HealthMonitorPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Parameters for Get
*/

// GetOptsBuilder allows extensions to add additional parameters to
// the health monitor API request
type GetOptsBuilder interface {
	ToHealthMonitorGetQuery() (string, error)
}

// GetOpts represents options used to show a health monitor.
type GetOpts struct {
	Changes bool `q:"changes"`
}

// ToHealthMonitorGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToHealthMonitorGetQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves a specific health monitor based on its unique ID.
func Get(c *eclcloud.ServiceClient, id string, opts GetOptsBuilder) (r GetResult) {
	url := getURL(c, id)
	if opts != nil {
		query, _ := opts.ToHealthMonitorGetQuery()
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
	ToHealthMonitorCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents options used to create a health monitor.
type CreateOpts struct {

	// Name of the health monitor
	Name string `json:"name,omitempty"`

	// Description of the health monitor
	Description string `json:"description,omitempty"`

	// Tags of the health monitor
	// Must be specified as JSON object
	Tags map[string]string `json:"tags,omitempty"`

	// Port number of the health monitor for healthchecking
	// Must be specified 0 when protocol is "icmp"
	Port int `json:"port"`

	// Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol"`

	// Interval of healthchecking (in seconds)
	Interval int `json:"interval,omitempty"`

	// Retry count of healthchecking
	// Initial monitoring is not included
	// Retry is executed at the interval specified by interval
	Retry int `json:"retry,omitempty"`

	// Timeout of healthchecking (in seconds)
	// Must be specified a number less than or equal to interval
	Timeout int `json:"timeout,omitempty"`

	// ID of the load balancer which the health monitor belongs to
	LoadBalancerID string `json:"load_balancer_id"`
}

// ToHealthMonitorCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToHealthMonitorCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// Create accepts a CreateOpts struct and creates a new health monitor
// using the values provided.
func Create(c *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToHealthMonitorCreateMap()
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
	ToHealthMonitorUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents options used to update metadata of exisiting health monitor.
type UpdateOpts struct {

	// Name of the health monitor
	Name string `json:"name,omitempty"`

	// Description of the health monitor
	Description string `json:"description,omitempty"`

	// Tags of the health monitor
	// Must be specified as JSON object
	Tags map[string]string `json:"tags,omitempty"`
}

// ToHealthMonitorUpdateMap builds a request body from UpdateMOpts.
func (opts UpdateOpts) ToHealthMonitorUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// Update accepts a UpdateOpts struct and updates an existing health monitor
// using the values provided.
func Update(c *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToHealthMonitorUpdateMap()
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

// Delete accepts a unique ID and deletes the health monitor associated with it.
func Delete(c *eclcloud.ServiceClient, id string) (r DeleteResult) {
	_, r.Err = c.Delete(deleteURL(c, id), nil)
	return
}

// IDFromName is a convenience function that returns a health monitor's
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

	all, err := ExtractHealthMonitors(pages)
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
		return "", eclcloud.ErrResourceNotFound{Name: name, ResourceType: "health_monitor"}
	case 1:
		return id, nil
	default:
		return "", eclcloud.ErrMultipleResourcesFound{Name: name, Count: count, ResourceType: "health_monitor"}
	}
}

/*
Parameters for Staged request
*/

// StagedOpts represents options used to create/update staged a health monitor.
type StagedOpts struct {

	// Port number of the health monitor for healthchecking
	// Must be specified 0 when protocol is "icmp"
	Port int `json:"port"`

	// Protocol of the health monitor for healthchecking
	Protocol string `json:"protocol,omitempty"`

	// Interval of healthchecking (in seconds)
	Interval int `json:"interval,omitempty"`

	// Retry count of healthchecking
	// Initial monitoring is not included
	// Retry is executed at the interval specified by interval
	Retry int `json:"retry,omitempty"`

	// Timeout of healthchecking (in seconds)
	// Must be specified a number less than or equal to interval
	Timeout int `json:"timeout,omitempty"`
}

// CreateStagedOptsBuilder allows extensions to add additional parameters to the
// CreateStaged request.
type CreateStagedOptsBuilder interface {
	ToHealthMonitorCreateStagedMap() (map[string]interface{}, error)
}

type CreateStagedOpts StagedOpts

// ToHealthMonitorCreateStagedMap builds a request body from CreateOpts.
func (opts CreateStagedOpts) ToHealthMonitorCreateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// CreateStaged accepts a CreateStagedOpts struct and creates a new health monitor
// using the values provided.
func CreateStaged(c *eclcloud.ServiceClient, id string, opts CreateStagedOptsBuilder) (r CreateResult) {
	b, err := opts.ToHealthMonitorCreateStagedMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Post(stagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// GetStaged retrieves a specific health monitor based on its unique ID.
func GetStaged(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Get(stagedURL(c, id), &r.Body, nil)
	return
}

// UpdateStagedOptsBuilder allows extensions to add additional parameters to the
// UpdateStaged request.
type UpdateStagedOptsBuilder interface {
	ToHealthMonitorUpdateStagedMap() (map[string]interface{}, error)
}

type UpdateStagedOpts StagedOpts

// ToHealthMonitorUpdateStagedMap builds a request body from UpdateStagedOpts.
func (opts UpdateStagedOpts) ToHealthMonitorUpdateStagedMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "health_monitor")
}

// UpdateStaged accepts a UpdateOpts struct and updates an existing health monitor
// using the values provided. For more information, see the Create function.
func UpdateStaged(c *eclcloud.ServiceClient, id string, opts UpdateStagedOptsBuilder) (r UpdateResult) {
	b, err := opts.ToHealthMonitorUpdateStagedMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = c.Patch(stagedURL(c, id), b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

// DeleteStaged retrieves a specific health monitor based on its unique ID.
func DeleteStaged(c *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = c.Delete(stagedURL(c, id), nil)
	return
}
