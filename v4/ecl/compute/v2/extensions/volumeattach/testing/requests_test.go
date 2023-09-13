package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v4/ecl/compute/v2/extensions/volumeattach"
	"github.com/nttcom/eclcloud/v4/pagination"

	th "github.com/nttcom/eclcloud/v4/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v4/testhelper/client"
)

func TestListVolumeAttachment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/servers/%s/os-volume_attachments", serverID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := volumeattach.List(fakeclient.ServiceClient(), serverID).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := volumeattach.ExtractVolumeAttachments(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, expectedVolumeAttachmentSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestCreateVolumeAttachment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/servers/%s/os-volume_attachments", serverID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, createResponse)
	})

	actual, err := volumeattach.Create(fakeclient.ServiceClient(), serverID, volumeattach.CreateOpts{
		Device:   "/dev/vdc",
		VolumeID: volumeID,
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &createdVolumeAttachment, actual)
}

func TestGetVolumeAttachment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/servers/%s/os-volume_attachments/%s", serverID, volumeID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := volumeattach.Get(fakeclient.ServiceClient(), serverID, attachID).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &secondVolumeAttachment, actual)
}

func TestDeleteVolumeAttachment(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/servers/%s/os-volume_attachments/%s", serverID, volumeID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusAccepted)
	})

	err := volumeattach.Delete(fakeclient.ServiceClient(), serverID, attachID).ExtractErr()
	th.AssertNoErr(t, err)
}
