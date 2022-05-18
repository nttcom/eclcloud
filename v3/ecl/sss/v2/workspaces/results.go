package workspaces

import (
	"encoding/json"
	"time"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/pagination"
)

type workspaceResult struct {
	eclcloud.Result
}

// GetResult is the result of a Get request. Call its Extract method to
// interpret it as a Workspace.
type GetResult struct {
	workspaceResult
}

// CreateResult is the result of a Create request. Call its Extract method to
// interpret it as a Workspace.
type CreateResult struct {
	workspaceResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Workspace.
type UpdateResult struct {
	workspaceResult
}

type Workspace struct {
	ContractID    string    `json:"contract_id"`
	WorkspaceID   string    `json:"workspace_id"`
	WorkspaceName string    `json:"workspace_name"`
	Description   string    `json:"description"`
	StartTime     time.Time `json:"-"`
	Regions       []Region  `json:"regions"`
	Users         []User    `json:"users"`
}

type Region struct {
	RegionName string `json:"region_name"`
	TenantID   string `json:"tenant_id"`
}

type User struct {
	UserID        string `json:"user_id"`
	ContractID    string `json:"contract_id"`
	ContractOwner bool   `json:"contract_owner"`
}

// UnmarshalJSON creates JSON format of workspace
func (r *Workspace) UnmarshalJSON(b []byte) error {
	type tmp Workspace
	var s struct {
		tmp
		StartTime eclcloud.JSONRFC3339ZNoTNoZ `json:"start_time"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = Workspace(s.tmp)

	r.StartTime = time.Time(s.StartTime)

	return err
}

// WorkspacePage is a single page of Workspace results.
type WorkspacePage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Workspace contains any results.
func (r WorkspacePage) IsEmpty() (bool, error) {
	workspaces, err := ExtractWorkspaces(r)
	return len(workspaces) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r WorkspacePage) NextPageURL() (string, error) {
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

// ExtractWorkspaces returns a slice of Workspace contained in a single page of results.
func ExtractWorkspaces(r pagination.Page) ([]Workspace, error) {
	var s struct {
		ContractID string      `json:"contract_id"`
		Workspaces []Workspace `json:"workspaces"`
	}

	// In list response case, each json element does not have contract_id.
	// It is set at out layer of each element.
	// So following logic set contract_id into inside of workspaces slice forcibly.
	// In "show(get with ID of workspace)" case, this does not occur.
	err := (r.(WorkspacePage)).ExtractInto(&s)
	contractID := s.ContractID

	for i := 0; i < len(s.Workspaces); i++ {
		s.Workspaces[i].ContractID = contractID
	}
	return s.Workspaces, err
}

// Extract interprets any projectResults as a Workspace.
func (r workspaceResult) Extract() (*Workspace, error) {
	var s Workspace
	err := r.ExtractInto(&s)
	return &s, err
}
