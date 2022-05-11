package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/computevolume/v2/volumes"
	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestListVolumeAll(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, listResponse)
	})

	allPages, err := volumes.List(fakeclient.ServiceClient(), &volumes.ListOpts{}).AllPages()
	th.AssertNoErr(t, err)
	actual, err := volumes.ExtractVolumes(allPages)
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, expectedVolumesSlice, actual)

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
		fmt.Fprintf(w, getResponse)
	})

	v, err := volumes.Get(fakeclient.ServiceClient(), idVolume1).Extract()
	th.AssertNoErr(t, err)

	th.CheckDeepEquals(t, &expectedVolumesSlice[0], v)
}

func TestCreateVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/volumes", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, createResponse)
	})

	options := &volumes.CreateOpts{
		Size:             15,
		AvailabilityZone: az,
		Description:      descriptionVolume1,
		Name:             nameVolume1,
		ImageID:          "dummyimage",
	}
	v, err := volumes.Create(fakeclient.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, v.AvailabilityZone, az)
	th.AssertEquals(t, v.Size, 15)
	th.AssertEquals(t, v.ID, idVolume1)
	th.AssertEquals(t, v.Description, descriptionVolume1)
}

func TestUpdatVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volumes/%s", idVolume1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, updateResponse)
	})

	name := nameVolume1Update
	description := descriptionVolume1Update
	metadata := map[string]string{}

	updateOpts := volumes.UpdateOpts{
		Name:        &name,
		Description: &description,
		Metadata:    &metadata,
	}

	v, err := volumes.Update(fakeclient.ServiceClient(), idVolume1, updateOpts).Extract()

	blankMeta := map[string]string{}
	th.AssertNoErr(t, err)
	th.CheckEquals(t, nameVolume1Update, v.Name)
	th.CheckEquals(t, descriptionVolume1Update, v.Description)
	th.CheckDeepEquals(t, &blankMeta, &v.Metadata)
}

func TestDeleteVolume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volumes/%s", idVolume1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.WriteHeader(http.StatusAccepted)
	})

	res := volumes.Delete(fakeclient.ServiceClient(), idVolume1)
	th.AssertNoErr(t, res.Err)
}
