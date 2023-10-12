package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/dedicated_hypervisor/v1/usages"

	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListUsages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUsagesSuccessfully(t)

	count := 0
	err := usages.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := usages.ExtractUsages(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedUsagesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListUsagesAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUsagesSuccessfully(t)

	allPages, err := usages.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := usages.ExtractUsages(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedUsagesSlice, actual)
}

func TestGetUsageHistories(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetHistoriesSuccessfully(t)

	result := usages.GetHistories(client.ServiceClient(), usageID, nil)
	th.AssertNoErr(t, result.Err)
	actual, err := result.ExtractHistories()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedHistories, actual)
}
