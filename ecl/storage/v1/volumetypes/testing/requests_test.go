package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/ecl/storage/v1/volumetypes"
	"github.com/nttcom/eclcloud/pagination"

	th "github.com/nttcom/eclcloud/testhelper"
	fakeclient "github.com/nttcom/eclcloud/testhelper/client"
)

func TestListVolumeType(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc(
		"/volume_types/detail",
		func(w http.ResponseWriter, r *http.Request) {
			th.TestMethod(t, r, "GET")
			th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			fmt.Fprintf(w, ListResponse)
		})

	client := fakeclient.ServiceClient()
	count := 0

	volumetypes.List(client, volumetypes.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := volumetypes.ExtractVolumeTypes(page)
		if err != nil {
			t.Errorf("Failed to extract volume types: %v", err)
			return false, err
		}

		th.CheckDeepEquals(t, getExpectedVolumeTypesSlice(), actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}

}

func TestGetVolumeType(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volume_types/%s", idVolumeType1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	vtActual, err := volumetypes.Get(
		fakeclient.ServiceClient(), idVolumeType1).Extract()
	th.AssertNoErr(t, err)
	vtExpected := getExpectedVolumeTypesSlice()[0]
	th.CheckDeepEquals(t, &vtExpected, vtActual)
}
