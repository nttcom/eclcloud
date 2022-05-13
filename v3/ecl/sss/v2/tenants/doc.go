/*
Package projects manages and retrieves Projects in the ECL SSS Service.

Example to List Tenants

	listOpts := tenants.ListOpts{
		Enabled: eclcloud.Enabled,
	}

	allPages, err := tenants.List(identityClient, listOpts).AllPages()
	if err != nil {
		panic(err)
	}

	allProjects, err := tenants.ExtractProjects(allPages)
	if err != nil {
		panic(err)
	}

	for _, tenant := range allTenants {
		fmt.Printf("%+v\n", tenant)
	}

Example to Create a Tenant

	createOpts := projects.CreateOpts{
		Name:        "tenant_name",
		Description: "Tenant Description"
	}

	tenant, err := tenants.Create(identityClient, createOpts).Extract()
	if err != nil {
		panic(err)
	}
*/
package tenants
