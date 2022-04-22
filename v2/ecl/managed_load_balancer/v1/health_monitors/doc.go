/*
Package appliances contains functionality for working with
ECL Commnon Function Gateway resources.

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

	createOpts := health_monitors.CreateOpts{
		Name:         "health_monitor_1",
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
*/
package health_monitors
