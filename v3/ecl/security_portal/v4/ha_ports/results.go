package ha_ports

import (
	"encoding/json"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract is a function that accepts a result
// and extracts a Port resource.
func (r commonResult) Extract() (*UpdateProcess, error) {
	var p UpdateProcess
	err := r.ExtractInto(&p)
	return &p, err
}

// Extract interprets any commonResult as a Port if possible.
func (r commonResult) ExtractInto(v interface{}) error {
	return r.Result.ExtractIntoStructPtr(v, "")
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Port.
type UpdateResult struct {
	commonResult
}

// UpdateProcess represents the result of a each element in
// response of port api result.
type UpdateProcess struct {
	Message   string `json:"message"`
	ProcessID string `json:"processId"`
	ID        string `json:"-"`
}

// ProcessPage is the page returned by a pager
// when traversing over a collection of Single Port.
type ProcessPage struct {
	pagination.LinkedPageBase
}

// UnmarshalJSON function overrides original functionality,
// to parse processId as unique identifier of process.
// Note:
// ID parameter in each struct must be string.
// As of MSS v4, processId is returned as a UUID string (previously numeric),
// so ID is simply copied from ProcessID.
func (r *UpdateProcess) UnmarshalJSON(b []byte) error {
	type tmp UpdateProcess
	var s struct {
		tmp
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = UpdateProcess(s.tmp)

	r.ID = r.ProcessID

	return err
}
