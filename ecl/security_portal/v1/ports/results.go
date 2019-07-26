package ports

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Single Firewall/UTM resource.
func (r commonResult) Extract() (*UpdateProcess, error) {
	var p UpdateProcess
	err := r.ExtractInto(&p)
	return &p, err
}

// Extract interprets any commonResult as a Single Firewall/UTM if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Single Firewall/UTM.
// type CreateResult struct {
// 	commonResult
// }

// // GetResult represents the result of a get operation. Call its Extract
// // method to interpret it as a Single Firewall/UTM.
// type GetResult struct {
// 	commonResult
// }

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Single Firewall/UTM.
type UpdateResult struct {
	commonResult
}

// // DeleteResult represents the result of a delete operation. Call its
// // ExtractErr method to determine if the request succeeded or failed.
// type DeleteResult struct {
// 	commonResult
// }

// UpdateProcess represents the result of a each element in
// response of single device api result.
type UpdateProcess struct {
	Message string `json:"message"`
	ID      string `json:"processId"`
}

// ProcessPage is the page returned by a pager
// when traversing over a collection of Single Firewall/UTM.
type ProcessPage struct {
	pagination.LinkedPageBase
}

// // NextPageURL is invoked when a paginated collection of Single Firewall/UTM
// //  has reached the end of a page and the pager seeks to traverse over a new one.
// // In order to do this, it needs to construct the next page's URL.
// func (r ProcessPage) NextPageURL() (string, error) {
// 	var s struct {
// 		Links []eclcloud.Link `json:"process_links"`
// 	}
// 	err := r.ExtractInto(&s)
// 	if err != nil {
// 		return "", err
// 	}
// 	return eclcloud.ExtractNextURL(s.Links)
// }

// // IsEmpty checks whether a ProcessPage struct is empty.
// func (r ProcessPage) IsEmpty() (bool, error) {
// 	is, err := ExtractProcesses(r)
// 	return len(is) == 0, err
// }

// // ExtractProcesses accepts a Page struct,
// // specifically a NetworkPage struct, and extracts the elements
// // into a slice of Single Firewall/UTM structs.
// // In other words, a generic collection is mapped into a relevant slice.
// func ExtractProcesses(r pagination.Page) ([]Process, error) {
// 	var s []SingleFirewallUTM
// 	err := ExtractProcessesInto(r, &s)
// 	return s, err
// }

// // ExtractProcessesInto interprets the results of a single page from a List() call,
// // producing a slice of Server entities.
// func ExtractProcessesInto(r pagination.Page, v interface{}) error {
// 	return r.(ProcessPage).Result.ExtractIntoSlicePtr(v, "")
// }
