package service_order_status

import (
	"github.com/nttcom/eclcloud"
)

// GetOptsBuilder allows extensions to add additional parameters to
// the order API request
type GetOptsBuilder interface {
	ToServiceOrderQuery() (string, error)
}

// GetOpts represents result of order API response.
type GetOpts struct {
	TenantID string `q:"tenant_id"`
	Locale   string `q:"locale"`
	SoID     string `q:"soid"`
}

// ToServiceOrderQuery formats a GetOpts into a query string.
func (opts GetOpts) ToServiceOrderQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single order, by ID.
func Get(client *eclcloud.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	url := getURL(client)
	if opts != nil {
		query, _ := opts.ToServiceOrderQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}