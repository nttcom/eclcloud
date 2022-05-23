/*
Package workspaces contains workspace management functionality on SSS.

Example to List workspaces

	listOpts := workspaces.ListOpts{ContractID: "econ0000000001"}

	allPages, err := workspaces.List(client, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allRequests, err := workspaces.ExtractWorkspaces(allPages)
	if err != nil {
		panic(err)
	}

	for _, request := range allRequests {
		fmt.Printf("%+v\n", request)
	}

Example to Get a workspace

	id := "ws0000000001"
	workspace, err := workspaces.Get(client, id).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", workspace)

Example to Create a workspace

	createOpts := workspaces.CreateOpts{
		WorkspaceName: "Example-Workspace",
		Description:   "Example Workspace",
		ContractID:    "econ0000000001",
	}

	workspace, err := workspaces.Create(client, createOpts).Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", workspace)

Example to Update a workspace

	workspaceID := "ws0000000001"
	description := "update description"
	updateOpts := workspaces.UpdateOpts{Description: &description}

	result := workspaces.Update(client, workspaceID, updateOpts)
	if result.Err != nil {
		panic(result.Err)
	}

Example to Delete a workspace

	workspaceID := "ws0000000001"
	res := workspaces.Delete(client, workspaceID)
	if res.Err != nil {
		panic(res.Err)
	}

*/
package workspaces
