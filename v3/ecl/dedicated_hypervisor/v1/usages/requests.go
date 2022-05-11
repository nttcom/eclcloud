package usages

import (
	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToResourceListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	From        string `q:"from"`
	To          string `q:"to"`
	LicenseType string `q:"license_type"`
}

// ToResourceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// List retrieves a list of Usages.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToResourceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return UsagePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// GetHistoriesOpts provides options to filter the GetHistories results.
type GetHistoriesOpts struct {
	From string `q:"from"`
	To   string `q:"to"`
}

// ToResourceListQuery formats a GetHistoriesOpts into a query string.
func (opts GetHistoriesOpts) ToResourceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	if err != nil {
		return "", err
	}
	return q.String(), nil
}

// GetHistories retrieves details of usage histories.
func GetHistories(client *eclcloud.ServiceClient, usageID string, opts ListOptsBuilder) (r GetHistoriesResult) {
	url := getHistoriesURL(client, usageID)
	if opts != nil {
		query, err := opts.ToResourceListQuery()
		if err != nil {
			r.Err = err
			return
		}
		url += query
	}
	_, r.Err = client.Get(url, &r.Body, nil)
	return
}
