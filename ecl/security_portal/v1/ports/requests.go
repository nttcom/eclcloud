package ports

import (
	"github.com/nttcom/eclcloud"
)

// // ListOptsBuilder allows extensions to add additional parameters to
// // the List request
// type ListOptsBuilder interface {
// 	ToSingleDeviceQuery() (string, error)
// }

// // ListOpts enables filtering of a list request.
// // Currently SSS User API does not support any of query parameters.
// type ListOpts struct {
// 	TenantID string `q:"tenant_id"`
// 	Locale   string `q:"locale"`
// }

// // ToSingleDeviceQuery formats a ListOpts into a query string.
// func (opts ListOpts) ToSingleDeviceQuery() (string, error) {
// 	q, err := eclcloud.BuildQueryString(opts)
// 	return q.String(), err
// }

// // List enumerates the Users to which the current token has access.
// func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
// 	url := listURL(client)
// 	if opts != nil {
// 		query, err := opts.ToSingleDeviceQuery()
// 		if err != nil {
// 			return pagination.Pager{Err: err}
// 		}
// 		url += query
// 	}
// 	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
// 		return SingleFirewallUTMPage{pagination.LinkedPageBase{PageResult: r}}
// 	})
// }

// // Get retrieves details on a single user, by ID.
// func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
// 	_, r.Err = client.Get(getURL(client), &r.Body, nil)
// 	return
// }

// // CreateOptsBuilder allows extensions to add additional parameters to
// // the Create request.
// type CreateOptsBuilder interface {
// 	ToSingleDeviceCreateMap() (map[string]interface{}, error)
// }

// // GtHostInCreate represents parameters used to create a Single Firewall/UTM.
// type GtHostInCreate struct {
// 	OperatingMode string `json:"operatingmode" required:"true"`
// 	LicenseKind   string `json:"licensekind" required:"true"`
// 	AZGroup       string `json:"azgroup" required:"true"`
// }

// // CreateOpts represents parameters used to create a user.
// type CreateOpts struct {
// 	SOKind   string            `json:"sokind" required:"true"`
// 	TenantID string            `json:"tenant_id" required:"true"`
// 	Locale   string            `json:"locale,omitempty"`
// 	GtHost   [1]GtHostInCreate `json:"gt_host" required:"true"`
// }

// // ToSingleDeviceCreateMap formats a CreateOpts into a create request.
// func (opts CreateOpts) ToSingleDeviceCreateMap() (map[string]interface{}, error) {
// 	return eclcloud.BuildRequestBody(opts, "")
// }

// // Create creates a new user.
// func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
// 	b, err := opts.ToSingleDeviceCreateMap()
// 	if err != nil {
// 		r.Err = err
// 		return
// 	}
// 	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
// 		OkCodes: []int{200},
// 	})
// 	return
// }

// type DeleteOptsBuilder interface {
// 	ToSingleDeviceDeleteMap() (map[string]interface{}, error)
// }

// // GtHostInDelete represents parameters used to create a Single Firewall/UTM.
// type GtHostInDelete struct {
// 	HostName string `json:"hostname" required:"true"`
// }

// // DeleteOpts represents parameters used to create a user.
// type DeleteOpts struct {
// 	SOKind   string            `json:"sokind" required:"true"`
// 	TenantID string            `json:"tenant_id" required:"true"`
// 	GtHost   [1]GtHostInDelete `json:"gt_host" required:"true"`
// }

// // ToSingleDeviceDeleteMap formats a CreateOpts into a create request.
// func (opts DeleteOpts) ToSingleDeviceDeleteMap() (map[string]interface{}, error) {
// 	return eclcloud.BuildRequestBody(opts, "")
// }

// // Delete deletes a user.
// func Delete(client *eclcloud.ServiceClient, opts DeleteOptsBuilder) (r DeleteResult) {
// 	// _, r.Err = client.Delete(deleteURL(client), nil)
// 	// return
// 	b, err := opts.ToSingleDeviceDeleteMap()
// 	if err != nil {
// 		r.Err = err
// 		return
// 	}
// 	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
// 		OkCodes: []int{200},
// 	})
// 	return

// }

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToPortUpdateMap() (map[string]interface{}, error)
}

// GtHostInUpdate represents parameters used to create a Single Firewall/UTM.
// type GtHostInUpdate struct {
// 	OperatingMode string `json:"operatingmode" required:"true"`
// 	LicenseKind   string `json:"licensekind" required:"true"`
// 	HostName      string `json:"hostname" required:"true"`
// }

// EachPort represents parameters to update a Single Firewall/UTM.
type SinglePort struct {
	EnablePort string `json:"enable_port" required:"true"`
	IPAddress  string `json:"ip_address,omitempty"`
	NetworkID  string `json:"network_id,omitempty"`
	SubnetID   string `json:"subnet_id,omitempty"`
	Comment    string `json:"comment,omitempty"`
}

type UpdateOpts struct {
	Port [7]SinglePort `json:"port" required:"true"`
}

// ToPortUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToPortUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

type UpdateQueryOptsBuilder interface {
	ToUpdateQuery() (string, error)
}

type UpdateQueryOpts struct {
	TenantID  string `q:"tenant_id"`
	UserToken string `q:"usertoken"`
}

func (opts UpdateQueryOpts) ToUpdateQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Update modifies the attributes of a user.
func Update(client *eclcloud.ServiceClient,
	serviceType string, opts UpdateOptsBuilder,
	qOpts UpdateQueryOptsBuilder) (r UpdateResult) {
	b, err := opts.ToPortUpdateMap()
	if err != nil {
		r.Err = err
		return
	}

	url := updateURL(client, serviceType)
	if qOpts != nil {
		query, _ := qOpts.ToUpdateQuery()
		url += query
	}

	_, r.Err = client.Post(url, &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}
