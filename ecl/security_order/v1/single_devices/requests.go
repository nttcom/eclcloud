package single_devices

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToSingleDeviceQuery() (string, error)
}

// ListOpts enables filtering of a list request.
// Currently SSS User API does not support any of query parameters.
type ListOpts struct {
	TenantID string `q:"tenant_id"`
	Locale   string `q:"locale"`
}

// ToSingleDeviceQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSingleDeviceQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the Users to which the current token has access.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToSingleDeviceQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return SingleDevicePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details on a single user, by ID.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToSingleDeviceCreateMap() (map[string]interface{}, error)
}

// GtHostInCreate represents parameters used to create a Single Device.
type GtHostInCreate struct {
	OperatingMode string `json:"operatingmode" required:"true"`
	LicenseKind   string `json:"licensekind" required:"true"`
	AZGroup       string `json:"azgroup" required:"true"`
}

// CreateOpts represents parameters used to create a user.
type CreateOpts struct {
	SOKind   string           `json:"sokind" required:"true"`
	TenantID string           `json:"tenant_id" required:"true"`
	Locale   string           `json:"locale,omitempty"`
	GtHost   []GtHostInCreate `json:"gt_host" required:"true"`
}

// ToSingleDeviceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToSingleDeviceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new user.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToSingleDeviceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return
}

type DeleteOptsBuilder interface {
	ToSingleDeviceDeleteMap() (map[string]interface{}, error)
}

// GtHostInDelete represents parameters used to create a Single Device.
type GtHostInDelete struct {
	HostName string `json:"hostname" required:"true"`
}

// DeleteOpts represents parameters used to create a user.
type DeleteOpts struct {
	SOKind   string            `json:"sokind" required:"true"`
	TenantID string            `json:"tenant_id" required:"true"`
	GtHost   [1]GtHostInDelete `json:"gt_host" required:"true"`
}

// ToSingleDeviceDeleteMap formats a CreateOpts into a create request.
func (opts DeleteOpts) ToSingleDeviceDeleteMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Delete deletes a user.
func Delete(client *eclcloud.ServiceClient, opts DeleteOptsBuilder) (r DeleteResult) {
	// _, r.Err = client.Delete(deleteURL(client), nil)
	// return
	b, err := opts.ToSingleDeviceDeleteMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, &eclcloud.RequestOpts{
		OkCodes: []int{200},
	})
	return

}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
// type UpdateOptsBuilder interface {
// 	ToUserUpdateMap() (map[string]interface{}, error)
// }

// // UpdateOpts represents parameters to update a user.
// type UpdateOpts struct {
// 	// New login id of the user.
// 	LoginID *string `json:"login_id" required:"true"`

// 	// New email address of the user
// 	MailAddress *string `json:"mail_address" required:"true"`

// 	// New password of the user
// 	NewPassword *string `json:"new_password" required:"true"`
// }

// // ToUserUpdateMap formats a UpdateOpts into an update request.
// func (opts UpdateOpts) ToUserUpdateMap() (map[string]interface{}, error) {
// 	return eclcloud.BuildRequestBody(opts, "")
// }

// // Update modifies the attributes of a user.
// // SSS User PUT API does not have response body,
// // so set JSONResponse option as nil.
// func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
// 	b, err := opts.ToUserUpdateMap()
// 	if err != nil {
// 		r.Err = err
// 		return
// 	}
// 	_, r.Err = client.Put(
// 		updateURL(client),
// 		b,
// 		nil,
// 		&eclcloud.RequestOpts{
// 			OkCodes: []int{204},
// 		},
// 	)
// 	return
// }
