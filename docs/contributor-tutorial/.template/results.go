package RESOURCE

import (
	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/pagination"
)

// RESOURCE represents...
type Resource struct {
}

type commonResult struct {
	eclcloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a RESOURCE.
type GetResult struct {
	commonResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a RESOURCE.
type CreateResult struct {
	commonResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a RESOURCE.
type UpdateResult struct {
	commonResult
}

// SinglePage is a single page of RESOURCE results.
type SinglePage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of RESOURCES contains any results.
func (r SinglePage) IsEmpty() (bool, error) {
	resources, err := ExtractResources(r)
	return len(resources) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r SinglePage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

// ExtractResources returns a slice of Resources contained in a single page of
// results.
func ExtractResources(r pagination.Page) ([]Resource, error) {
	var s struct {
		Resources []Resource `json:"resources"`
	}
	err := (r.(SinglePage)).ExtractInto(&s)
	return s.Resources, err
}

// Extract interprets any commonResult as a Resource.
func (r commonResult) Extract() (*Resource, error) {
	var s struct {
		Resource *Resource `json:"resource"`
	}
	err := r.ExtractInto(&s)
	return s.Resource, err
}
