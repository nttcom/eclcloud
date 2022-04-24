/*
Package listeners contains functionality for working with
ECL Managed Load Balancer resources.

Example to List Listener

	listOpts := listeners.ListOpts{
		TenantID: "a99e9b4e620e4db09a2dfb6e42a01e66",
	}

	allPages, err := listeners.List(mlbClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allListeners, err := listeners.ExtractListeners(allPages)
	if err != nil {
		panic(err)
	}

	for _, listener := range allListeners {
		fmt.Printf("%+v", listener)
	}

Example to Create a Listener

  loadBalancerID := "67fea379-cff0-4191-9175-de7d6941a040"

	createOpts := listeners.CreateOpts{
		Name:           "listener_1",
		IPAddress:      "10.0.0.1",
		Port:           80,
		Protocol:       "tcp",
		LoadBalancerID: loadBalancerID,
	}

	listener, err := listeners.Create(mlbClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Show a Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	listener, err := listeners.Get(mlbClient, listenerID, listeners.GetOpts{}).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	updateOpts := listeners.UpdateOpts{
		Name: "new_name",
	}

	listener, err := listeners.Update(mlbClient, listenerID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := listeners.Delete(mlbClient, listenerID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Create a Staged Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	createStagedOpts := listeners.CreateStagedOpts{
		Port:     443,
		Protocol: "tcp",
	}

	err := listeners.CreateStaged(mlbClient, listenerID, createStagedOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Show a Staged Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	err := listeners.GetStaged(mlbClient, listenerID).ExtractErr()
	if err != nil {
		panic(err)
	}


Example to Update a Staged Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	updateStagedOpts := listeners.UpdateStagedOpts{
		Port:     8080,
		Protocol: "tcp",
	}

	err := listeners.UpdateStaged(mlbClient, listenerID, updateStagedOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Delete a Staged Listener

	listenerID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	err := listeners.DeleteStaged(mlbClient, listenerID).ExtractErr()
	if err != nil {
		panic(err)
	}

*/
package listeners
