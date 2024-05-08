package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/v3/ecl/computevolume/extensions/volumeactions"
	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

func TestVolumeUploadImage(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volumes/%s/action", volumeID)

	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, uploadImageRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)

		fmt.Fprintf(w, uploadImageResponse)
	})

	options := &volumeactions.UploadImageOpts{
		ContainerFormat: "bare",
		Force:           true,
		ImageName:       "imagetest",
		DiskFormat:      "raw",
	}

	actual, err := volumeactions.UploadImage(fakeclient.ServiceClient(), volumeID, options).Extract()
	th.AssertNoErr(t, err)

	expected := volumeactions.VolumeImage{
		Status:    "uploading",
		ImageID:   "49d7efe7-975e-46d7-af0a-fd94fe8e62bf",
		ImageName: "imagetest",
		VolumeType: volumeactions.ImageVolumeType{
			Name:       "nfsdriver",
			QosSpecsID: "",
			Deleted:    false,
			CreatedAt:  time.Date(2018, 6, 4, 8, 5, 9, 0, time.UTC),
			UpdatedAt:  time.Time{},
			DeletedAt:  time.Time{},
			ID:         "1f02ea8f-3823-4e69-a232-695adc39f5e0",
		},
		ContainerFormat: "bare",
		Size:            40,
		DiskFormat:      "raw",
		VolumeID:        volumeID,
		Description:     "test volume 2update",
		UpdatedAt:       time.Date(2019, 2, 6, 22, 6, 27, 0, time.UTC),
	}
	th.AssertDeepEquals(t, expected, actual)
}

func TestVolumeExtendSize(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/volumes/%s/action", volumeID)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, extendRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})

	options := &volumeactions.ExtendSizeOpts{
		NewSize: 40,
	}

	err := volumeactions.ExtendSize(fakeclient.ServiceClient(), volumeID, options).ExtractErr()
	th.AssertNoErr(t, err)
}
