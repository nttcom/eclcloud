/*
Package health_monitors contains functionality for working with
ECL Managed Load Balancer resources.

Example to List Health Monitor

	listOpts := health_monitors.ListOpts{
		TenantID: "a99e9b4e620e4db09a2dfb6e42a01e66",
	}

	allPages, err := health_monitors.List(mlbClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allHealthMonitors, err := health_monitors.ExtractHealthMonitors(allPages)
	if err != nil {
		panic(err)
	}

	for _, health_monitor := range allHealthMonitors {
		fmt.Printf("%+v", health_monitor)
	}

Example to Create a Health Monitor

  loadBalancerID := "67fea379-cff0-4191-9175-de7d6941a040"

	createOpts := health_monitors.CreateOpts{
		Name:           "health_monitor_1",
		Port:           80,
		Protocol:       "tcp",
		LoadBalancerID: loadBalancerID,
	}

	health_monitor, err := health_monitors.Create(mlbClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Show a Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	health_monitor, err := health_monitors.Get(mlbClient, healthMonitorID, health_monitors.GetOpts{}).Extract()
	if err != nil {
		panic(err)
	}

Example to Update a Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	updateOpts := health_monitors.UpdateOpts{
		Name: "new_name",
	}

	health_monitor, err := health_monitors.Update(mlbClient, healthMonitorID, updateOpts).Extract()
	if err != nil {
		panic(err)
	}

Example to Delete a Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	err := health_monitors.Delete(mlbClient, healthMonitorID).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Create a Staged Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	createStagedOpts := health_monitors.CreateStagedOpts{
		Port:     0,
		Protocol: "icmp",
	}

	err := health_monitors.CreateStaged(mlbClient, healthMonitorID, createStagedOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Show a Staged Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	err := health_monitors.GetStaged(mlbClient, healthMonitorID).ExtractErr()
	if err != nil {
		panic(err)
	}


Example to Update a Staged Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	updateStagedOpts := health_monitors.UpdateStagedOpts{
		Port:     0,
		Protocol: "icmp",
		Interval: 5,
		Retry:    3,
		Timeout:  5,
	}

	err := health_monitors.UpdateStaged(mlbClient, healthMonitorID, updateStagedOpts).ExtractErr()
	if err != nil {
		panic(err)
	}

Example to Delete a Staged Health Monitor

	healthMonitorID := "497f6eca-6276-4993-bfeb-53cbbbba6f08"

	err := health_monitors.DeleteStaged(mlbClient, healthMonitorID).ExtractErr()
	if err != nil {
		panic(err)
	}

*/
package health_monitors
