package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/provider_connectivity/v2/tenant_connection_requests"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
	"github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListTenantConnectionRequests(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantConnectionRequestsSuccessfully(t)

	count := 0
	err := tenant_connection_requests.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := tenant_connection_requests.ExtractTenantConnectionRequests(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedTenantConnectionRequestsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListTenantConnectionRequestsAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantConnectionRequestsSuccessfully(t)

	allPages, err := tenant_connection_requests.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := tenant_connection_requests.ExtractTenantConnectionRequests(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedTenantConnectionRequestsSlice, actual)
}

func TestGetTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetTenantConnectionRequestSuccessfully(t)

	actual, err := tenant_connection_requests.Get(client.ServiceClient(), SecondTenantConnectionRequest.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnectionRequest, *actual)
}

func TestCreateTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateTenantConnectionRequestSuccessfully(t)

	createOpts := tenant_connection_requests.CreateOpts{
		TenantIDOther: "7e91b19b9baa423793ee74a8e1ff2be1",
		NetworkID:     "77cfc6b0-d032-4e5a-b6fb-4cce2537f4d1",
		Name:          "test_name1",
		Description:   "test_desc1",
		Tags:          map[string]string{"test_tags1": "test1"},
	}

	actual, err := tenant_connection_requests.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &FirstTenantConnectionRequest, actual)
}

func TestDeleteTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteTenantConnectionRequestSuccessfully(t)

	res := tenant_connection_requests.Delete(client.ServiceClient(), FirstTenantConnectionRequest.ID)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateTenantConnectionRequestSuccessfully(t)

	name := "updated_name"
	description := "updated_desc"
	tags := map[string]string{"k2": "v2"}

	updateOpts := tenant_connection_requests.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := tenant_connection_requests.Update(client.ServiceClient(), SecondTenantConnectionRequest.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnectionRequestUpdated, *actual)
}

func TestUpdateOtherMetadataTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateOtherMetadataTenantConnectionRequestSuccessfully(t)

	nameOther := "updated_name_other"
	descriptionOther := "updated_desc_other"
	tagsOther := map[string]string{"k3": "v3"}

	updateOpts := tenant_connection_requests.UpdateOpts{
		NameOther:        &nameOther,
		DescriptionOther: &descriptionOther,
		TagsOther:        &tagsOther,
	}

	actual, err := tenant_connection_requests.Update(client.ServiceClient(), SecondTenantConnectionRequest.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnectionRequestOtherMetadataUpdated, *actual)
}

func TestBlankUpdateTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleBlankUpdateTenantConnectionRequestSuccessfully(t)

	name := ""
	description := ""
	tags := map[string]string{}

	updateOpts := tenant_connection_requests.UpdateOpts{
		Name:        &name,
		Description: &description,
		Tags:        &tags,
	}

	actual, err := tenant_connection_requests.Update(client.ServiceClient(), SecondTenantConnectionRequest.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnectionRequestBlankUpdated, *actual)
}

func TestNilUpdateTenantConnectionRequest(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleNilUpdateTenantConnectionRequestSuccessfully(t)

	name := "nilupdate"

	updateOpts := tenant_connection_requests.UpdateOpts{
		Name: &name,
	}

	actual, err := tenant_connection_requests.Update(client.ServiceClient(), SecondTenantConnectionRequest.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnectionRequestNilUpdated, *actual)
}
