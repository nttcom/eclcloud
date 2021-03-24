package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2/ecl/dns/v2/zones"
	"github.com/nttcom/eclcloud/v2/pagination"

	th "github.com/nttcom/eclcloud/v2/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v2/testhelper/client"
)

func TestListDNSZone(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, ListResponse)
	})

	count := 0
	err := zones.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := zones.ExtractZones(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedZonesSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListDNSZoneAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, ListResponse)
	})

	allPages, err := zones.List(fakeclient.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	allZones, err := zones.ExtractZones(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allZones))
}

func TestGetDNSZone(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s", idZone1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, GetResponse)
	})

	actual, err := zones.Get(fakeclient.ServiceClient(), idZone1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &GetResponseStruct, actual)
}

func TestCreateDNSZone(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/zones", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, CreateZoneRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, CreateZoneResponse)
	})

	createOpts := zones.CreateOpts{
		Description: descriptionZone1,
		Email:       "joe@example.org",
		Name:        nameZone1,
		TTL:         7200,
		Type:        "PRIMARY",
	}

	actual, err := zones.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedZone, actual)
}

func TestUpdateDNSZone(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s", idZone1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PATCH")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, UpdateZoneRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprint(w, UpdateZoneResponse)
	})

	description := descriptionZone1Update
	ttl := 600
	masters := make([]string, 0)
	email := ""

	updateOpts := zones.UpdateOpts{
		TTL:         &ttl,
		Description: &description,
		Masters:     &masters,
		Email:       &email,
	}

	actual, err := zones.Update(fakeclient.ServiceClient(), idZone1, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdatedZone, actual)
}

func TestDeleteDNSZone(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s", idZone1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusAccepted)
		w.Header().Add("Content-Type", "application/json")
	})

	res := zones.Delete(fakeclient.ServiceClient(), idZone1)
	th.AssertNoErr(t, res.Err)
}
