package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2/ecl/sss/v1/users"
	"github.com/nttcom/eclcloud/v2/pagination"

	th "github.com/nttcom/eclcloud/v2/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v2/testhelper/client"
)

func TestListUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := users.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := users.ExtractUsers(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedUsersSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListUserAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	allPages, err := users.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allZones, err := users.ExtractUsers(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allZones))
}

func TestGetUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/users/%s", idUser1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := users.Get(fakeclient.ServiceClient(), idUser1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &getResponseStruct, actual)
}

func TestCreateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	createOpts := users.CreateOpts{
		LoginID:        "login_id_1",
		MailAddress:    "user1@example.com",
		NotifyPassword: "false",
		Password:       "Passw0rd",
	}

	// clone FirstTenant into createdUser(Used as assertion target)
	// and initialize StartTime
	// createdUser := firstUser
	// createdUser.StartTime = time.Time{}

	actual, err := users.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createdUser, actual)
}

func TestUpdateUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/users/%s", idUser1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, updateRequest)

		w.WriteHeader(http.StatusNoContent)
	})

	loginID := "login_id_1_update"
	mailAddress := "user1_update@example.com"
	newPassword := "NewPassw0rd"

	updateOpts := users.UpdateOpts{
		LoginID:     &loginID,
		MailAddress: &mailAddress,
		NewPassword: &newPassword,
	}

	// In ECL2.0 user update API returns
	// - StatusNoContent
	// - No response body as PUT response
	res := users.Update(fakeclient.ServiceClient(), idUser1, updateOpts)
	th.AssertNoErr(t, res.Err)
}

func TestDeleteUser(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/users/%s", idUser1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	res := users.Delete(fakeclient.ServiceClient(), idUser1)
	th.AssertNoErr(t, res.Err)
}
