/*
Package usages manages and retrieves usage in the Enterprise Cloud Dedicated Hypervisor Service.

Example to List Usages

	listOpts := usages.ListOpts{
		From:        "2019-10-10T00:00:00Z",
		To:          "2019-10-15T00:00:00Z",
		LicenseType: "dedicated-hypervisor.guest-image.vcenter-server-6-0-standard",
	}

	allPages, err := usages.List(dhClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allUsages, err := usages.ExtractUsages(allPages)
	if err != nil {
		panic(err)
	}

	for _, usage := range allUsages {
		fmt.Printf("%+v\n", usage)
	}

Example to Get Usage Histories

	getHistoriesOpts := usages.GetHistoriesOpts{
		From: "2019-10-10T00:00:00Z",
		To:   "2019-10-15T00:00:00Z",
	}

	usageID := "9ada4c06-a2a4-46d5-b969-72ac12433a79"
	histories, err := usages.GetHistories(client, usageID, getHistoriesOpts).ExtractHistories()
	if err != nil {
		panic(err)
	}
*/
package usages
