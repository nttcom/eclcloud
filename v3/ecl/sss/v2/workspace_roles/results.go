package workspace_roles

import (
	"github.com/nttcom/eclcloud/v3"
)

type workspaceRoleResult struct {
	eclcloud.Result
}

// CreateResult is the result of a Create request. Call its Extract method to
// interpret it as a Workspace-Role.
type CreateResult struct {
	workspaceRoleResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteResult struct {
	eclcloud.ErrResult
}

type WorkspaceRole struct {
	UserID        string `json:"user_id"`
	WorkspaceID   string `json:"workspace_id"`
	WorkspaceName string `json:"workspace_name"`
}

// Extract interprets any projectResults as a Workspace-Role.
func (r workspaceRoleResult) Extract() (*WorkspaceRole, error) {
	var s WorkspaceRole
	err := r.ExtractInto(&s)
	return &s, err
}
