package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/dedicated_hypervisor/v1/license_types"

	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListLicenseTypes(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLicenseTypesSuccessfully(t)

	count := 0
	err := license_types.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := license_types.ExtractLicenseTypes(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedLicenseTypesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListLicenseTypesAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLicenseTypesSuccessfully(t)

	allPages, err := license_types.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)
	actual, err := license_types.ExtractLicenseTypes(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedLicenseTypesSlice, actual)
}
