package testing

import (
	"testing"

	"github.com/nttcom/eclcloud/ecl/rca/v1/users"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
	"github.com/nttcom/eclcloud/testhelper/client"
)

func TestListUsers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUsersSuccessfully(t)

	count := 0
	err := users.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := users.ExtractUsers(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedUsersSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListUsersAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListUsersSuccessfully(t)

	allPages, err := users.List(client.ServiceClient()).AllPages()
	th.AssertNoErr(t, err)
	actual, err := users.ExtractUsers(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedUsersSlice, actual)
}

func TestGetUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetUserSuccessfully(t)

	actual, err := users.Get(client.ServiceClient(), SecondUser.Name).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondUser, *actual)
}

func TestCreateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateUserSuccessfully(t)

	createOpts := users.CreateOpts{
		Password: password,
	}

	actual, err := users.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, SecondUser.Name, actual.Name)
	th.AssertEquals(t, password, actual.Password)
	th.AssertDeepEquals(t, SecondUser.VPNEndpoints, actual.VPNEndpoints)
}

func TestDeleteUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteUserSuccessfully(t)

	res := users.Delete(client.ServiceClient(), FirstUser.Name)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateUserSuccessfully(t)

	updateOpts := users.UpdateOpts{
		Password: passwordUpdated,
	}

	actual, err := users.Update(client.ServiceClient(), SecondUser.Name, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondUserUpdated, *actual)
}
