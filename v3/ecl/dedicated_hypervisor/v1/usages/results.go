package usages

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v3"

	"github.com/nttcom/eclcloud/v3/pagination"
)

// Usage represents guest image usage information.
type Usage struct {
	ID            string `json:"id"`
	Type          string `json:"type"`
	Value         string `json:"value"`
	Unit          string `json:"unit"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	HasLicenseKey bool   `json:"has_license_key"`
	ResourceID    string `json:"resource_id"`
}

type UsageHistories struct {
	Unit        string    `json:"unit"`
	ResourceID  string    `json:"resource_id"`
	LicenseType string    `json:"license_type"`
	Histories   []History `json:"histories"`
}

type History struct {
	Time  time.Time `json:"-"`
	Value string    `json:"value"`
}

func (h *History) UnmarshalJSON(b []byte) error {
	type tmp History
	var s struct {
		tmp
		Time eclcloud.JSONRFC3339ZNoTNoZ `json:"time"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*h = History(s.tmp)

	h.Time = time.Time(s.Time)

	return err
}

type commonResult struct {
	eclcloud.Result
}

// GetHistoriesResult is the response from a Get operation. Call its Extract method
// to interpret it as usage histories.
type GetHistoriesResult struct {
	commonResult
}

// UsagePage is a single page of Usage results.
type UsagePage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Usages contains any results.
func (r UsagePage) IsEmpty() (bool, error) {
	usages, err := ExtractUsages(r)
	return len(usages) == 0, err
}

// ExtractUsages returns a slice of Usages contained in a single page of
// results.
func ExtractUsages(r pagination.Page) ([]Usage, error) {
	var s struct {
		Usages []Usage `json:"usages"`
	}
	err := (r.(UsagePage)).ExtractInto(&s)
	return s.Usages, err
}

// ExtractHistories interprets any commonResult as usage histories.
func (r commonResult) ExtractHistories() (*UsageHistories, error) {
	var s UsageHistories
	err := r.ExtractInto(&s)
	return &s, err
}
