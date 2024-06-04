package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/dedicated_hypervisor/v1/licenses"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListLicenses(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLicensesSuccessfully(t)

	count := 0
	err := licenses.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := licenses.ExtractLicenses(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedLicensesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListLicensesAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListLicensesSuccessfully(t)

	allPages, err := licenses.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := licenses.ExtractLicenses(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedLicensesSlice, actual)
}

func TestCreateLicense(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateLicenseSuccessfully(t)

	createOpts := licenses.CreateOpts{
		LicenseType: SecondLicense.LicenseType,
	}

	actual, err := licenses.Create(client.ServiceClient(), createOpts).ExtractLicenseInfo()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondLicense, *actual)
}

func TestDeleteLicense(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteLicenseSuccessfully(t)

	res := licenses.Delete(client.ServiceClient(), FirstLicense.ID)
	th.AssertNoErr(t, res.Err)
}
