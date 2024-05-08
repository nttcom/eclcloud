package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/storage/v1/volumes"
	"github.com/nttcom/eclcloud/v3/pagination"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestListVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/volumes/detail",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, ListResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	volumes.List(client, volumes.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := volumes.ExtractVolumes(page)
		if err != nil {
			t.Errorf("Failed to extract volumes: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, getExpectedVolumesSlice(), actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}

}

func TestGetVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volumes/%s", idVolume1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	volActual, err := volumes.Get(
		fakeclient.ServiceClient(), idVolume1).Extract()
	th.AssertNoErr(t, err)
	volExpected := getExpectedVolumesSlice()[0]
	th.CheckDeepEquals(t, &volExpected, volActual)
}

func TestCreateBlockStorageTypeVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, CreateRequestBlock)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted) // 202

			fmt.Fprintf(w, CreateResponseBlock)
		})

	createOpts := volumes.CreateOpts{
		VirtualStorageID: idVirtualStorage,
		Name:             nameVolume1,
		Description:      descriptionVolume1,
		Size:             100,
		IOPSPerGB:        "2",
		InitiatorIQNs:    []string{IQN1},
		AvailabilityZone: "zone1_groupa",
	}
	volActual, err := volumes.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, volActual.Status, "creating")

	volExpected := getExpectedCreateBlockStorageTypeVolume()
	th.AssertDeepEquals(t, &volExpected, volActual)
}

func TestCreateFileStorageTypeVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "POST")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, CreateRequestFile)
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted) // 202

			fmt.Fprintf(w, CreateResponseFile)
		})

	createOpts := volumes.CreateOpts{
		VirtualStorageID: idVirtualStorage,
		Name:             nameVolume1,
		Description:      descriptionVolume1,
		Size:             256,
		Throughput:       "50",
		AvailabilityZone: "zone1_groupa",
	}
	volActual, err := volumes.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, volActual.Status, "creating")

	volExpected := getExpectedCreateFileStorageTypeVolume()
	th.AssertDeepEquals(t, &volExpected, volActual)
}

// TestUpdateBlockStorageTypeVolume covers file storage type's codes
// So, contrary to creation tests, tests for file storage type updating is not implemented
func TestUpdateBlockStorageTypeVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(fmt.Sprintf("/volumes/%s", idVolume1),
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "PUT")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
			th.TestHeader(t, r, "Content-Type", "application/json")
			th.TestHeader(t, r, "Accept", "application/json")
			th.TestJSONRequest(t, r, UpdateRequest)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)

			fmt.Fprintf(w, UpdateResponse)
		})

	name := nameVolume1Update
	description := descriptionVolume1Update
	initiatorIQNs := []string{IQN1, IQN2}

	updateOpts := volumes.UpdateOpts{
		Name:          &name,
		Description:   &description,
		InitiatorIQNs: &initiatorIQNs,
	}

	volActual, err := volumes.Update(
		fakeclient.ServiceClient(), idVolume1, updateOpts).Extract()

	th.AssertNoErr(t, err)

	th.AssertEquals(t, volActual.Name, nameVolume1Update)
	th.AssertEquals(t, volActual.Description, descriptionVolume1Update)

	th.AssertEquals(t, volActual.InitiatorIQNs[0], IQN1)
	th.AssertEquals(t, volActual.InitiatorIQNs[1], IQN2)
}

func TestDeleteVirtualStorage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volumes/%s", idVolume1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.WriteHeader(http.StatusOK)
	})

	res := volumes.Delete(fakeclient.ServiceClient(), idVolume1)
	th.AssertNoErr(t, res.Err)
}
