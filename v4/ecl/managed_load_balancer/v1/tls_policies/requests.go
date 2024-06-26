package tls_policies

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

/*
List TLS Policies
*/

// ListOpts allows the filtering and sorting of paginated collections through the API.
// Filtering is achieved by passing in struct field values that map to the tls policy attributes you want to see returned.
type ListOpts struct {

	// - ID of the resource
	ID string `q:"id"`

	// - Name of the resource
	// - This field accepts single-byte characters only
	Name string `q:"name"`

	// - Description of the resource
	// - This field accepts single-byte characters only
	Description string `q:"description"`

	// - Whether the TLS policy will be set `policy.tls_policy_id` when that is not specified
	Default bool `q:"default"`
}

// ToTLSPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTLSPolicyListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)

	return q.String(), err
}

// ListOptsBuilder allows extensions to add additional parameters to the List request.
type ListOptsBuilder interface {
	ToTLSPolicyListQuery() (string, error)
}

// List returns a Pager which allows you to iterate over a collection of tls policies.
// It accepts a ListOpts struct, which allows you to filter and sort the returned collection for greater efficiency.
func List(c *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(c)

	if opts != nil {
		query, err := opts.ToTLSPolicyListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}

		url += query
	}

	return pagination.NewPager(c, url, func(r pagination.PageResult) pagination.Page {
		return TLSPolicyPage{pagination.LinkedPageBase{PageResult: r}}
	})
}

/*
Show TLS Policy
*/

// Show retrieves a specific tls policy based on its unique ID.
func Show(c *eclcloud.ServiceClient, id string) (r ShowResult) {
	_, r.Err = c.Get(showURL(c, id), &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})

	return
}
