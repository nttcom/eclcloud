package service_order_status

import (
	"github.com/nttcom/eclcloud"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a common function gateway resource.
func (r commonResult) Extract() (*OrderProgress, error) {
	var sd OrderProgress
	err := r.ExtractInto(&sd)
	return &sd, err
}

// Extract interprets any commonResult as a Common Function Gateway, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Common Function Gateway.
type GetResult struct {
	commonResult
}

// OrderProgress represents, well, a common function gateway.
// SingleDevice represents an ECL SSS User.
type OrderProgress struct {
	Status       string `json:"status"`
	Code         string `json:"code"`
	Message      string `json:"message"`
	ProgressRate string `json:"progressRate"`
}

// // OrderProgressPage is the page returned by a pager
// // when traversing over a collection of common function gateway.
// type OrderProgressPage struct {
// 	pagination.LinkedPageBase
// }

// // NextPageURL is invoked when a paginated collection of common function gateway
// //  has reached the end of a page and the pager seeks to traverse over a new one.
// // In order to do this, it needs to construct the next page's URL.
// func (r OrderProgressPage) NextPageURL() (string, error) {
// 	var s struct {
// 		Links []eclcloud.Link `json:"order_progress_links"`
// 	}
// 	err := r.ExtractInto(&s)
// 	if err != nil {
// 		return "", err
// 	}
// 	return eclcloud.ExtractNextURL(s.Links)
// }

// // IsEmpty checks whether a SingleDevicePage struct is empty.
// func (r OrderProgressPage) IsEmpty() (bool, error) {
// 	is, err := ExtractOrderProgress(r)
// 	return len(is) == 0, err
// }

// // ExtractSingleDevices accepts a Page struct,
// // specifically a NetworkPage struct, and extracts the elements
// // into a slice of Common Function Gateway structs.
// // In other words, a generic collection is mapped into a relevant slice.
// func ExtractSingleDevices(r pagination.Page) ([]SingleDevice, error) {
// 	var s []SingleDevice
// 	err := ExtractSingleDevicesInto(r, &s)
// 	return s, err
// }

// // ExtractSingleDevicesInto interprets the results of a single page from a List() call,
// // producing a slice of Server entities.
// func ExtractSingleDevicesInto(r pagination.Page, v interface{}) error {
// 	return r.(SingleDevicePage).Result.ExtractIntoSlicePtr(v, "")
// }
