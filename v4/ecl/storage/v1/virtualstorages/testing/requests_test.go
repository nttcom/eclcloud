package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/storage/v1/virtualstorages"
	"github.com/nttcom/eclcloud/v4/pagination"

	th "github.com/nttcom/eclcloud/v4/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListVirtualStorage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/virtual_storages/detail",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, ListResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	virtualstorages.List(client, virtualstorages.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := virtualstorages.ExtractVirtualStorages(page)
		if err != nil {
			t.Errorf("Failed to extract virtual storages: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, getExpectedVirtualStoragesSlice(), actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}

}

func TestGetVirtualStorage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/virtual_storages/%s", idVirtualStorage1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	vsActual, err := virtualstorages.Get(
		fakeclient.ServiceClient(), idVirtualStorage1).Extract()
	th.AssertNoErr(t, err)
	vsExpected := getExpectedVirtualStoragesSlice()[0]
	th.CheckDeepEquals(t, &vsExpected, vsActual)
}

func TestCreate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/virtual_storages",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, CreateRequest)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted) // 202

			fmt.Fprintf(w, CreateResponse)
		})

	createOpts := virtualstorages.CreateOpts{
		VolumeTypeID: idVolumeType,
		Name:         nameVirtualStorage1,
		Description:  descriptionVirtualStorage1,
		NetworkID:    networkID,
		SubnetID:     subnetID,
		IPAddrPool:   getIPAddrPool(false),
		HostRoutes:   getHostRoutes(false),
	}
	vsActual, err := virtualstorages.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, vsActual.Status, "creating")

	vsExpected := getExpectedCreateVirtualStorage()
	th.AssertDeepEquals(t, &vsExpected, vsActual)
}

func TestUpdateVirtualStorage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/virtual_storages/%s", idVirtualStorage1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, UpdateResponse)
	})

	name := nameVirtualStorage1Update
	description := descriptionVirtualStorage1Update
	ipAddrPool := getIPAddrPool(true)
	hostRoutes := getHostRoutes(true)

	updateOpts := virtualstorages.UpdateOpts{
		Name:        &name,
		Description: &description,
		IPAddrPool:  &ipAddrPool,
		HostRoutes:  &hostRoutes,
	}
	vsActual, err := virtualstorages.Update(
		fakeclient.ServiceClient(), idVirtualStorage1, updateOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, vsActual.Name, nameVirtualStorage1Update)
	th.AssertEquals(t, vsActual.Description, descriptionVirtualStorage1Update)
	th.AssertEquals(t, vsActual.ID, idVirtualStorage1)

	th.AssertEquals(t, vsActual.IPAddrPool.Start, ipAddrPoolStartUpdate)
	th.AssertEquals(t, vsActual.IPAddrPool.End, ipAddrPoolEndUpdate)

	th.AssertEquals(t, vsActual.HostRoutes[2].Destination, hostRoute3Destination)
	th.AssertEquals(t, vsActual.HostRoutes[2].Nexthop, hostRoute3Nexthop)
}

func TestDeleteVirtualStorage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/virtual_storages/%s", idVirtualStorage1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	res := virtualstorages.Delete(fakeclient.ServiceClient(), idVirtualStorage1)
	th.AssertNoErr(t, res.Err)
}
