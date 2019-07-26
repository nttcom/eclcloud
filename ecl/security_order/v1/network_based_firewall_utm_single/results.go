package network_based_firewall_utm_single

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Single Firewall/UTM resource.
func (r commonResult) Extract() (*SingleFirewallUTMOrder, error) {
	var sdo SingleFirewallUTMOrder
	err := r.ExtractInto(&sdo)
	return &sdo, err
}

// Extract interprets any commonResult as a Single Firewall/UTM if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Single Firewall/UTM.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Single Firewall/UTM.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Single Firewall/UTM.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	commonResult
}

// SingleFirewallUTM represents the result of a each element in
// response of single device api result.
type SingleFirewallUTM struct {
	ID   int      `json:"id"`
	Cell []string `json:"cell"`
}

// SingleFirewallUTMOrder represents a Single Firewall/UTM's each order.
type SingleFirewallUTMOrder struct {
	ID      string `json:"soId"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

// SingleFirewallUTMPage is the page returned by a pager
// when traversing over a collection of Single Firewall/UTM.
type SingleFirewallUTMPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of Single Firewall/UTM
//  has reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r SingleFirewallUTMPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"single_firewall_utm_links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty checks whether a SingleFirewallUTMPage struct is empty.
func (r SingleFirewallUTMPage) IsEmpty() (bool, error) {
	is, err := ExtractSingleFirewallUTMs(r)
	return len(is) == 0, err
}

// ExtractSingleFirewallUTMs accepts a Page struct,
// specifically a NetworkPage struct, and extracts the elements
// into a slice of Single Firewall/UTM structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractSingleFirewallUTMs(r pagination.Page) ([]SingleFirewallUTM, error) {
	var s []SingleFirewallUTM
	err := ExtractSingleFirewallUTMsInto(r, &s)
	return s, err
}

// ExtractSingleFirewallUTMsInto interprets the results of a single page from a List() call,
// producing a slice of Server entities.
func ExtractSingleFirewallUTMsInto(r pagination.Page, v interface{}) error {
	return r.(SingleFirewallUTMPage).Result.ExtractIntoSlicePtr(v, "rows")
}
