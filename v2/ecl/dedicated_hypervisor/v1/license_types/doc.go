/*
Package license_types manages and retrieves license type in the Enterprise Cloud Dedicated Hypervisor Service.

Example to List License types

	allPages, err := license_types.List(dhClient).AllPages()
	if err != nil {
		panic(err)
	}

	allLicenseTypes, err := license_types.ExtractLicenseTypes(allPages)
	if err != nil {
		panic(err)
	}

	for _, licenseType := range allLicenseTypes {
		fmt.Printf("%+v\n", licenseType)
	}
*/
package license_types
