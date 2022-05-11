package testing

import (
	"testing"

	th "github.com/nttcom/eclcloud/v3/testhelper"
)

func TestBootFromNewVolume(t *testing.T) {

	actual, err := newVolumeRequest.ToServerCreateMap()
	th.AssertNoErr(t, err)
	th.CheckJSONEquals(t, expectedNewVolumeRequest, actual)
}

func TestBootFromExistingVolume(t *testing.T) {
	actual, err := existingVolumeRequest.ToServerCreateMap()
	th.AssertNoErr(t, err)
	th.CheckJSONEquals(t, expectedExistingVolumeRequest, actual)
}

func TestBootFromImage(t *testing.T) {
	actual, err := imageRequest.ToServerCreateMap()
	th.AssertNoErr(t, err)
	th.CheckJSONEquals(t, expectedImageRequest, actual)
}

func TestCreateMultiEphemeralOpts(t *testing.T) {
	actual, err := multiEphemeralRequest.ToServerCreateMap()
	th.AssertNoErr(t, err)
	th.CheckJSONEquals(t, expectedMultiEphemeralRequest, actual)
}

func TestAttachNewVolume(t *testing.T) {
	actual, err := imageAndNewVolumeRequest.ToServerCreateMap()
	th.AssertNoErr(t, err)
	th.CheckJSONEquals(t, expectedImageAndNewVolumeRequest, actual)
}

func TestAttachExistingVolume(t *testing.T) {
	actual, err := imageAndExistingVolumeRequest.ToServerCreateMap()
	th.AssertNoErr(t, err)
	th.CheckJSONEquals(t, expectedImageAndExistingVolumeRequest, actual)
}
