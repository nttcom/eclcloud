package service_order_status

import (
	"github.com/nttcom/eclcloud"
)

// // ListOptsBuilder allows extensions to add additional parameters to
// // the List request
// type ListOptsBuilder interface {
// 	ToOrderProgressQuery() (string, error)
// }

// // ListOpts enables filtering of a list request.
// // Currently SSS User API does not support any of query parameters.
// type ListOpts struct {
// }

// // ToOrderProgressQuery formats a ListOpts into a query string.
// func (opts ListOpts) ToOrderProgressQuery() (string, error) {
// 	q, err := eclcloud.BuildQueryString(opts)
// 	return q.String(), err
// }

// // List enumerates the Users to which the current token has access.
// func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
// 	url := listURL(client)
// 	if opts != nil {
// 		query, err := opts.ToOrderProgressQuery()
// 		if err != nil {
// 			return pagination.Pager{Err: err}
// 		}
// 		url += query
// 	}
// 	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
// 		return OrderProgressPage{pagination.LinkedPageBase{PageResult: r}}
// 	})
// }

type GetOptsBuilder interface {
	ToServiceOrderQuery() (string, error)
}

type GetOpts struct {
	TenantID string `q:"tenant_id"`
	Locale   string `q:"locale"`
	SoID     string `q:"soid"`
}

// ToSingleDeviceQuery formats a ListOpts into a query string.
func (opts GetOpts) ToServiceOrderQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// Get retrieves details on a single user, by ID.
func Get(client *eclcloud.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	url := getURL(client)
	if opts != nil {
		query, _ := opts.ToServiceOrderQuery()
		url += query
	}

	_, r.Err = client.Get(url, &r.Body, nil)
	return
}

// // CreateOptsBuilder allows extensions to add additional parameters to
// // the Create request.
// type CreateOptsBuilder interface {
// 	ToUserCreateMap() (map[string]interface{}, error)
// }

// // GtHost represents parameters used to create a Single Device.
// type GtHost struct {
// 	OperatingMode string `json:"operatingmode" required:"true"`
// 	LicenseKind   string `json:"licensekind" required:"true"`
// 	AZGroup       string `json:"azgroup" required:"true"`
// }

// // CreateOpts represents parameters used to create a user.
// type CreateOpts struct {
// 	SOKind   string   `json:"sokind" required:"true"`
// 	TenantID string   `json:"tenant_id" required:"true"`
// 	Locale   string   `json:"locale,omitempty"`
// 	GtHost   []GtHost `json:"gt_host,omitempty"`
// }

// // ToUserCreateMap formats a CreateOpts into a create request.
// func (opts CreateOpts) ToUserCreateMap() (map[string]interface{}, error) {
// 	return eclcloud.BuildRequestBody(opts, "")
// }

// // Create creates a new user.
// func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
// 	b, err := opts.ToUserCreateMap()
// 	if err != nil {
// 		r.Err = err
// 		return
// 	}
// 	_, r.Err = client.Post(createURL(client), &b, &r.Body, nil)
// 	return
// }

// // Delete deletes a user.
// func Delete(client *eclcloud.ServiceClient, userID string) (r DeleteResult) {
// 	_, r.Err = client.Delete(deleteURL(client), nil)
// 	return
// }

// // UpdateOptsBuilder allows extensions to add additional parameters to
// // the Update request.
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
