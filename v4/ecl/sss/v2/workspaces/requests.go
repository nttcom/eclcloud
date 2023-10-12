package workspaces

import (
	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToWorkspaceListQuery() (string, error)
}

// ListOpts enables filtering of a list request.
// Currently SSS Workspace API does not support any of query parameters.
type ListOpts struct {
	ContractID string `q:"contract_id"`
}

// ToWorkspaceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToWorkspaceListQuery() (string, error) {
	q, err := eclcloud.BuildQueryString(opts)
	return q.String(), err
}

// List enumerates the Workspaces to which the current token has access.
func List(client *eclcloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	url := listURL(client)
	if opts != nil {
		query, err := opts.ToWorkspaceListQuery()
		if err != nil {
			return pagination.Pager{Err: err}
		}
		url += query
	}
	return pagination.NewPager(client, url, func(r pagination.PageResult) pagination.Page {
		return WorkspacePage{pagination.LinkedPageBase{PageResult: r}}
	})
}

// Get retrieves details on a single workspace, by ID.
func Get(client *eclcloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}

// CreateOptsBuilder allows extensions to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToWorkspaceCreateMap() (map[string]interface{}, error)
}

// CreateOpts represents parameters used to create a workspace.
type CreateOpts struct {
	// Workspace Name.
	WorkspaceName string `json:"workspace_name" required:"true"`
	// Workspace description.
	Description string `json:"description,omitempty"`
	// ContractID to be associated with the workspace.
	ContractID string `json:"contract_id,omitempty"`
}

// ToWorkspaceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToWorkspaceCreateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Create creates a new workspace.
func Create(client *eclcloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	b, err := opts.ToWorkspaceCreateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Post(createURL(client), &b, &r.Body, nil)
	return
}

// Delete deletes a workspace.
func Delete(client *eclcloud.ServiceClient, workspaceID string) (r DeleteResult) {
	_, r.Err = client.Delete(deleteURL(client, workspaceID), nil)
	return
}

// UpdateOptsBuilder allows extensions to add additional parameters to the Update request.
type UpdateOptsBuilder interface {
	ToWorkspaceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a workspace.
type UpdateOpts struct {
	// Workspace description.
	Description *string `json:"description" required:"true"`
}

// ToWorkspaceUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToWorkspaceUpdateMap() (map[string]interface{}, error) {
	return eclcloud.BuildRequestBody(opts, "")
}

// Update modifies the attributes of a workspace.
// SSS Workspace PUT API does not have response body, so set JSONResponse option as nil.
func Update(client *eclcloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	b, err := opts.ToWorkspaceUpdateMap()
	if err != nil {
		r.Err = err
		return
	}
	_, r.Err = client.Put(
		updateURL(client, id),
		b,
		nil,
		&eclcloud.RequestOpts{
			OkCodes: []int{204},
		},
	)
	return
}
