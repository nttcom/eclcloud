package workspace_roles

import "github.com/nttcom/eclcloud/v3"

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToWorkspaceRoleCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents parameters used to create a workspace-role.
type CreateOpts struct {
	UserID      string `json:"user_id" required:"true"`
	WorkspaceID string `json:"workspace_id" required:"true"`
}

// ToWorkspaceRoleCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToWorkspaceRoleCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new workspace-role.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToWorkspaceRoleCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, nil)
	return
}

// Delete deletes a workspace-role.
func Delete(client *eclcloud.ServiceClient, workspaceID string, userID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, workspaceID, userID), nil)
	return
}
