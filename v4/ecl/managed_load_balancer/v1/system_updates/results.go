package system_updates

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// ShowResult represents the result of a Show operation.
// Call its Extract method to interpret it as a SystemUpdate.
type ShowResult struct {
	commonResult
}

// SystemUpdate represents a system update.
type SystemUpdate struct {

	// - ID of the system update
	ID string `json:"id"`

	// - Name of the system update
	Name string `json:"name"`

	// - Description of the system update
	Description string `json:"description"`

	// - URL of announcement for the system update (for example, Knowledge Center news)
	Href string `json:"href"`

	// - The time when the system update has been announced
	// - Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
	PublishDatetime string `json:"publish_datetime"`

	// - The deadline for applying the system update to the load balancer at any time
	//   - **For load balancers that have not been applied the system update even after the deadline, the provider will automatically apply it in the maintenance window of each region**
	// - Format: `"%Y-%m-%d %H:%M:%S"` (UTC)
	LimitDatetime string `json:"limit_datetime"`

	// - Current revision for the system update
	// - The system update can be applied to the load balancers that is this revision
	CurrentRevision int `json:"current_revision"`

	// - Next revision for the system update
	// - The load balancer to which the system update is applied will be this revision
	NextRevision int `json:"next_revision"`

	// - Whether the system update can be applied to the load balancer
	Applicable bool `json:"applicable"`
}

// ExtractInto interprets any commonResult as a system update, if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "system_update")
}

// Extract is a function that accepts a result and extracts a SystemUpdate resource.
func (r commonResult) Extract() (*SystemUpdate, error) {
	var systemUpdate SystemUpdate

	err := r.ExtractInto(&systemUpdate)

	return &systemUpdate, err
}

// SystemUpdatePage is the page returned by a pager when traversing over a collection of system update.
type SystemUpdatePage struct {
	pagination.LinkedPageBase
}

// IsEmpty checks whether a SystemUpdatePage struct is empty.
func (r SystemUpdatePage) IsEmpty() (bool, error) {
	is, err := ExtractSystemUpdates(r)

	return len(is) == 0, err
}

// ExtractSystemUpdatesInto interprets the results of a single page from a List() call, producing a slice of system update entities.
func ExtractSystemUpdatesInto(r pagination.Page, v interface{}) error {
	return r.(SystemUpdatePage).Result.ExtractIntoSlicePtr(v, "system_updates")
}

// ExtractSystemUpdates accepts a Page struct, specifically a NetworkPage struct, and extracts the elements into a slice of SystemUpdate structs.
// In other words, a generic collection is mapped into a relevant slice.
func ExtractSystemUpdates(r pagination.Page) ([]SystemUpdate, error) {
	var s []SystemUpdate

	err := ExtractSystemUpdatesInto(r, &s)

	return s, err
}
