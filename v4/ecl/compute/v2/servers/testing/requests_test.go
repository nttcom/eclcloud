package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/compute/v2/servers"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListServers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServersDetailsSuccessfully(t)

	count := 0
	err := servers.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := servers.ExtractServers(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, expectedServers, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListServersAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListServersDetailsSuccessfully(t)

	allPages, err := servers.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := servers.ExtractServers(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedServers, actual)
}

func TestGetServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetServerSuccessfully(t)

	actual, err := servers.Get(client.ServiceClient(), expectedServer2.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectedServer2, *actual)
}

func TestCreateServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateServerSuccessfully(t)

	configDrive := true
	createOpts := servers.CreateOpts{
		Name:             "Test Server1",
		ImageRef:         "c11a6d55-70e9-4d04-a086-4451f07da0d7",
		FlavorRef:        "1CPU-4GB",
		UserData:         []byte("user_data"),
		AvailabilityZone: "zone1-groupa",
		Networks: []servers.Network{
			{
				UUID: "4d98b876-b5d1-4861-8650-b5a53024486a",
			},
		},
		Metadata: map[string]string{
			"foo": "bar",
		},
		ConfigDrive: &configDrive,
	}

	actual, err := servers.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedServer2.ID, actual.ID)
	th.AssertDeepEquals(t, expectedServer2.Links, actual.Links)
	th.AssertEquals(t, "aabbccddeeff", actual.AdminPass)
}

func TestDeleteServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteServerSuccessfully(t)

	res := servers.Delete(client.ServiceClient(), expectedServer1.ID)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateServerSuccessfully(t)

	name := "Update Name"
	updateOpts := servers.UpdateOpts{Name: &name}

	actual, err := servers.Update(client.ServiceClient(), expectedServer2.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, serverNameUpdated, *actual)
}

func TestGetMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetMetadataSuccessfully(t)

	actual, err := servers.Metadata(client.ServiceClient(), expectedServer2.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectMetadata, actual)
}

func TestGetMetadatum(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetMetadatumSuccessfully(t)

	actual, err := servers.Metadatum(client.ServiceClient(), expectedServer2.ID, "vmha").Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectMetadatum, actual)
}

func TestCreateMetadatum(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateMetadatumSuccessfully(t)

	createOpts := servers.MetadatumOpts{"key1": "val1"}

	actual, err := servers.CreateMetadatum(client.ServiceClient(), expectedServer2.ID, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectCreateMetadatum, actual)
}

func TestDeleteMetadatum(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteMetadatumSuccessfully(t)

	res := servers.DeleteMetadatum(client.ServiceClient(), expectedServer1.ID, "vmha")
	th.AssertNoErr(t, res.Err)
}

func TestUpdateMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateMetadataSuccessfully(t)

	updateOpts := servers.MetadataOpts{"key1": "update_val"}

	actual, err := servers.UpdateMetadata(client.ServiceClient(), expectedServer2.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ecpectUpdateMetadata, actual)
}

func TestResetMetadata(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleResetMetadataSuccessfully(t)

	createOpts := servers.MetadataOpts{
		"key1": "val1",
		"key2": "val2",
	}

	actual, err := servers.ResetMetadata(client.ServiceClient(), expectedServer2.ID, createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, expectResetMetadata, actual)
}

func TestResizeServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleResizeServerSuccessfully(t)

	resizeOpts := servers.ResizeOpts{FlavorRef: "2CPU-8GB"}

	err := servers.Resize(client.ServiceClient(), expectedServer2.ID, resizeOpts).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestCreateImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateImageSuccessfully(t)

	snapshotOpts := servers.CreateImageOpts{
		Name:     "Test Create Image",
		Metadata: map[string]string{"key": "create_image"},
	}

	result := servers.CreateImage(client.ServiceClient(), expectedServer2.ID, snapshotOpts)
	th.AssertNoErr(t, result.Err)
}
