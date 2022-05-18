package workspace_roles

import (
	"fmt"

	"github.com/nttcom/eclcloud/v3"
)

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("workspace-roles")
}

func deleteURL(client *eclcloud.ServiceClient, workspaceID string, userID string) string {
	url := fmt.Sprintf("workspace-roles/workspaces/%s/users/%s", workspaceID, userID)
	return client.ServiceURL(url)
}
