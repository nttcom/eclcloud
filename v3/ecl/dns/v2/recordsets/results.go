package recordsets

import (
	"encoding/json"
	"fmt"
	// "log"
	"time"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type commonResult struct {
	eclcloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a RecordSet.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*RecordSet, error) {
	var s *RecordSet
	err := r.ExtractInto(&s)
	return s, err
}

func (r commonResult) ExtractCreatedRecordSet() (*RecordSet, error) {
	var sl []*RecordSet
	err := r.ExtractIntoSlicePtr(&sl, "recordsets")
	if err != nil {
		return nil, fmt.Errorf("[Error] Error in parsing result of recordset create  %s", err)
	}
	return sl[0], nil
}

// CreateResult is the result of a Create operation. Call its Extract method to
// interpret the result as a RecordSet.
type CreateResult struct {
	commonResult
}

// GetResult is the result of a Get operation. Call its Extract method to
// interpret the result as a RecordSet.
type GetResult struct {
	commonResult
}

// RecordSetPage is a single page of RecordSet results.
type RecordSetPage struct {
	pagination.LinkedPageBase
}

// UpdateResult is result of an Update operation. Call its Extract method to
// interpret the result as a RecordSet.
type UpdateResult struct {
	commonResult
}

// DeleteResult is result of a Delete operation. Call its ExtractErr method to
// determine if the operation succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// IsEmpty returns true if the page contains no results.
func (r RecordSetPage) IsEmpty() (bool, error) {
	s, err := ExtractRecordSets(r)
	return len(s) == 0, err
}

// ExtractRecordSets extracts a slice of RecordSets from a List result.
func ExtractRecordSets(r pagination.Page) ([]RecordSet, error) {
	var s struct {
		RecordSets []RecordSet `json:"recordsets"`
	}
	err := (r.(RecordSetPage)).ExtractInto(&s)
	return s.RecordSets, err
}

// RecordSet represents a DNS Record Set.
type RecordSet struct {
	// ID is the unique ID of the recordset
	ID string `json:"id"`

	// ZoneID is the ID of the zone the recordset belongs to.
	ZoneID string `json:"zone_id"`

	// ProjectID is the ID of the project that owns the recordset.
	// ProjectID string `json:"project_id"`

	// Name is the name of the recordset.
	Name string `json:"name"`

	// Type is the RRTYPE of the recordset.
	Type string `json:"type"`

	// Records are the DNS records of the recordset.
	// This is original code.
	// Records []string `json:"records"`
	//
	// But in ECL2.0, record set will be returned as simple string
	// e.g.
	// Usual response(like creation) reccordset: "[10.0.0.1]"
	// Update response(like creation) reccordset: "10.0.0.1]"
	Records interface{} `json:"records"`

	// TTL is the time to live of the recordset.
	TTL int `json:"ttl"`

	// Description is the description of the recordset.
	Description string `json:"description"`

	// Version is the revision of the recordset.
	Version int `json:"version"`

	// CreatedAt is the date when the recordset was created.
	CreatedAt time.Time `json:"-"`

	// UpdatedAt is the date when the recordset was updated.
	UpdatedAt time.Time `json:"-"`

	// Status is the current status of recordset.
	Status string `json:"status"`

	// Current action in progress on the resource.
	// This parameter is not currently supported. it always return an empty.
	Action string `json:"action"`

	// Links includes HTTP references to the itself,
	// useful for passing along to other APIs that might want a recordset
	// reference.
	Links []eclcloud.Link `json:"-"`
}

func (r *RecordSet) UnmarshalJSON(b []byte) error {
	type tmp RecordSet
	var s struct {
		tmp
		CreatedAt eclcloud.JSONRFC3339MilliNoZ `json:"created_at"`
		UpdatedAt eclcloud.JSONRFC3339MilliNoZ `json:"updated_at"`
		Links     map[string]interface{}       `json:"links"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = RecordSet(s.tmp)

	r.CreatedAt = time.Time(s.CreatedAt)
	r.UpdatedAt = time.Time(s.UpdatedAt)

	if s.Links != nil {
		for rel, href := range s.Links {
			if v, ok := href.(string); ok {
				link := eclcloud.Link{
					Rel:  rel,
					Href: v,
				}
				r.Links = append(r.Links, link)
			}
		}
	}

	return err
}
