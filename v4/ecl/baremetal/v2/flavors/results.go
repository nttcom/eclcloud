package flavors

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// GetResult is the result of Get operations. Call its Extract method to
// interpret it as a Flavor.
type GetResult struct {
	commonResult
}

// Extract provides access to the individual Flavor returned by
// the Get and functions.
func (r commonResult) Extract() (*Flavor, error) {
	var s struct {
		Flavor *Flavor `json:"flavor"`
	}
	err := r.ExtractInto(&s)
	return s.Flavor, err
}

// Flavor represent (virtual) hardware configurations for server resources
// in a region.
type Flavor struct {
	// ID is the flavor's unique ID.
	ID string `json:"id"`

	// Name is the name of the flavor.
	Name string `json:"name"`

	// Disk is the amount of root disk, measured in GB.
	Disk int `json:"disk"`

	// RAM is the amount of memory, measured in MB.
	RAM int `json:"ram"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor.
	VCPUs int `json:"vcpus"`
}

// FlavorPage contains a single page of all flavors from a ListDetails call.
type FlavorPage struct {
	pagination.LinkedPageBase
}

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (page FlavorPage) NextPageURL() (string, error) {
	var s struct {
		Links []eclcloud.Link `json:"flavors_links"`
	}
	err := page.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return eclcloud.ExtractNextURL(s.Links)
}

// IsEmpty determines if a FlavorPage contains any results.
func (page FlavorPage) IsEmpty() (bool, error) {
	flavors, err := ExtractFlavors(page)
	return len(flavors) == 0, err
}

// ExtractFlavors provides access to the list of flavors in a page acquired
// from the ListDetail operation.
func ExtractFlavors(r pagination.Page) ([]Flavor, error) {
	var s struct {
		Flavors []Flavor `json:"flavors"`
	}
	err := (r.(FlavorPage)).ExtractInto(&s)
	return s.Flavors, err
}
