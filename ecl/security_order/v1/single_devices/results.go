package single_devices

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a common function gateway resource.
func (r commonResult) Extract() (*SingleDeviceOrder, error) {
	var sdo SingleDeviceOrder
	err := r.ExtractInto(&sdo)
	return &sdo, err
}

// Extract interprets any commonResult as a Common Function Gateway, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Common Function Gateway.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Common Function Gateway.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Common Function Gateway.
// type UpdateResult struct {
// 	commonResult
// }

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

type SingleDevice struct {
	ID   int      `json:"id"`
	Cell []string `json:"cell"`
}

// SingleDeviceOrder represents, well, a common function gateway.
// SingleDevice represents an ECL SSS User.
type SingleDeviceOrder struct {
	ID      string `json:"soId"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// SingleDevicePage is the page returned by a pager
// when traversing over a collection of common function gateway.
type SingleDevicePage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of common function gateway
//  has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r SingleDevicePage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"single_device_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SingleDevicePage struct is empty.
func (r SingleDevicePage) IsEmpty() (bool, error) {
	is, err := ExtractSingleDevices(r)
	return len(is) == 0, err
}

// ExtractSingleDevices accepts a Page struct,
// specifically a NetworkPage struct, and extracts the elements
// into a slice of Common Function Gateway structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractSingleDevices(r pagination.Page) ([]SingleDevice, error) {
	var s []SingleDevice
	err := ExtractSingleDevicesInto(r, &s)
	return s, err
}

// ExtractSingleDevicesInto interprets the results of a single page from a List() call,
// producing a slice of Server entities.
func ExtractSingleDevicesInto(r pagination.Page, v interface{}) error {
	return r.(SingleDevicePage).Result.ExtractIntoSlicePtr(v, "rows")
}
