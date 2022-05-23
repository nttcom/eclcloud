/*
Package workspace_roles contains workspace-role management functionality on SSS.

Example to Create a workspace-role

	workspaceID := "ws00000000001"
	userID := "ecid0000000001"

	createOpts := workspace_roles.CreateOpts{
		UserID:      userID,
		WorkspaceID: workspaceID,
	}

	role, err := workspace_roles.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", role)

Example to Delete a workspace-role

	workspaceID := "ws00000000001"
	userID := "ecid00000000001"
	result := workspace_roles.Delete(client, workspaceID, userID)
	if result.Err != nil {
		panic(result.Err)
	}

*/
package workspace_roles
