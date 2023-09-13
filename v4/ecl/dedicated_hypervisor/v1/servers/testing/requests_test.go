package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/dedicated_hypervisor/v1/servers"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListServers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServersSuccessfully(t)

	count := 0
	err := servers.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := servers.ExtractServers(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedServersSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListServersAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServersSuccessfully(t)

	allPages, err := servers.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := servers.ExtractServers(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedServersSlice, actual)
}

func TestListServersDetails(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServersDetailsSuccessfully(t)

	count := 0
	err := servers.ListDetails(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := servers.ExtractServers(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedServersDetailsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListServersDetailsAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServersDetailsSuccessfully(t)

	allPages, err := servers.ListDetails(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := servers.ExtractServers(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedServersDetailsSlice, actual)
}

func TestGetServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetServerSuccessfully(t)

	actual, err := servers.Get(client.ServiceClient(), SecondServer.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondServerDetail, *actual)
}

func TestCreateServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateServerSuccessfully(t)

	createOpts := servers.CreateOpts{
		Name: "test",
		Networks: []servers.Network{
			{
				UUID:           "94055904-6b2c-4839-a14a-c61c93a8bc48",
				Plane:          "data",
				SegmentationID: 6,
			},
			{
				UUID:           "94055904-6b2c-4839-a14a-c61c93a8bc48",
				Plane:          "data",
				SegmentationID: 6,
			},
		},
		ImageRef:  "dfd25820-b368-4012-997b-29a6d0cf8518",
		FlavorRef: "a830b61c-3155-4a61-b7ed-c450862845e6",
	}

	actual, err := servers.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, SecondServer.ID, actual.ID)
	th.AssertDeepEquals(t, SecondServer.Links, actual.Links)
	th.AssertEquals(t, "aabbccddeeff", actual.AdminPass)
}

func TestDeleteResource(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteServerSuccessfully(t)

	res := servers.Delete(client.ServiceClient(), FirstServer.ID)
	th.AssertNoErr(t, res.Err)
}

func TestAddLicense(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleAddLicenseSuccessfully(t)

	addLicenseOpts := servers.AddLicenseOpts{
		VmName: "Alice",
		LicenseTypes: []string{
			"Windows Server",
			"SQL Server Standard 2014",
		},
	}

	actual, err := servers.AddLicense(client.ServiceClient(), SecondServer.ID, addLicenseOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, AddLicenseJob.JobID, actual.JobID)
}

func TestGetAddLicenseResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetAddLicenseResultSuccessfully(t)

	getAddLicenseResultOpts := servers.GetAddLicenseResultOpts{
		JobID: AddLicenseJob.JobID,
	}

	actual, err := servers.GetAddLicenseResult(client.ServiceClient(), SecondServer.ID, getAddLicenseResultOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, AddLicenseJob, *actual)
}
