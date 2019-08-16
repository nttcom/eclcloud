package host_based

import (
	"github.com/nttcom/eclcloud"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Host Based Security resource.
func (r commonResult) Extract() (*HostBasedOrder, error) {
	var hbo HostBasedOrder
	err := r.ExtractInto(&hbo)
	return &hbo, err
}

// Extract interprets any commonResult as a Host Based Security if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Host Based Security.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Host Based Security.
type GetResult struct {
	commonResult
}

// HostBasedSecurity represents a Host Based Security's each order.
type HostBasedSecurity struct {
	Code                string `json:"code"`
	Message             string `json:"message"`
	Region              string `json:"region"`
	TenantName          string `json:"tenant_name"`
	TenantDescription   string `json:"tenant_description"`
	ContractID          string `json:"contract_id"`
	ServiceOrderService string `json:"service_order_service"`
	MaxAgentValue       string `json:"max_agent_value"`
	TimeZone            string `json:"time_zone"`
	CustomerName        string `json:"customer_name"`
	MailAddress         string `json:"mailaddress"`
	DSMLang             string `json:"dsm_lang"`
	TenantFlg           bool   `json:"tenant_flg"`
	Status              int    `json:"status"`
}

// Extract is a function that accepts a result
// and extracts a Host Based Security resource.
func (r GetResult) Extract() (*HostBasedSecurity, error) {
	var h HostBasedSecurity
	err := r.ExtractInto(&h)
	return &h, err
}

// ExtractInto interprets any commonResult as a Host Based Security if possible.
func (r GetResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Host Based Security.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

// HostBasedOrder represents a Host Based Security's each order.
type HostBasedOrder struct {
	ID      string `json:"soId"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// // HostBasedPage is the page returned by a pager
// // when traversing over a collection of Host Based Security.
// type HostBasedPage struct {
// 	pagination.LinkedPageBase
// }

// // NextPageURL is invoked when a paginated collection of Host Based Security
// // has reached the end of a page and the pager seeks to traverse over a new one.
// // In order to do this, it needs to construct the next page's URL.
// func (r HostBasedPage) NextPageURL() (string, error) {
// 	var s struct {
// 		Links []eclcloud.Link `json:"host_based_security_links"`
// 	}
// 	err := r.ExtractInto(&s)
// 	if err != nil {
// 		return "", err
// 	}
// 	return eclcloud.ExtractNextURL(s.Links)
// }

// // IsEmpty checks whether a HostBasedPage struct is empty.
// func (r HostBasedPage) IsEmpty() (bool, error) {
// 	is, err := ExtractHostBased(r)
// 	return len(is) == 0, err
// }

// // ExtractHostBased accepts a Page struct,
// // specifically a HostBasedPage struct, and extracts the elements
// // into a slice of Host Based Security structs.
// // In other words, a generic collection is mapped into a relevant slice.
// func ExtractHostBased(r pagination.Page) ([]HADevice, error) {
// 	var s []HADevice
// 	err := ExtractHostBasedInto(r, &s)
// 	return s, err
// }

// // ExtractHostBasedInto interprets the results of a single page from a List() call,
// // producing a slice of Device entities.
// func ExtractHostBasedInto(r pagination.Page, v interface{}) error {
// 	return r.(HostBasedPage).Result.ExtractIntoSlicePtr(v, "")
// }
