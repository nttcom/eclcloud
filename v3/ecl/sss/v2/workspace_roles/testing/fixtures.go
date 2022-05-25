package testing

import "github.com/nttcom/eclcloud/v3/ecl/sss/v2/workspace_roles"

const workspaceID = "ws0000000001"

const userID = "ecid1234567891"

var createRequest = `
{
	"user_id": "ecid1234567891",
	"workspace_id": "ws0000000001"
}
`

var createResponse = `
{
	"user_id": "ecid1234567891",
	"workspace_id": "ws0000000001",
	"workspace_name": "testWorkspace001"
}
`

var createdWorkspaceRole = workspace_roles.WorkspaceRole{
	UserID:        userID,
	WorkspaceID:   workspaceID,
	WorkspaceName: "testWorkspace001",
}
