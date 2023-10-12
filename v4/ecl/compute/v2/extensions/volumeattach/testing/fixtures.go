package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/compute/v2/extensions/volumeattach"
)

const serverID = "4d8c3732-a248-40ed-bebc-539a6ffd25c0"
const volumeID = "a26887c6-c47b-4654-abb5-dfadf7d3f803"
const attachID = volumeID

var listResponse = fmt.Sprintf(`
{
  "volumeAttachments": [
    {
      "device": "/dev/vdd",
      "id": "%s",
      "serverId": "%s",
      "volumeId": "%s"
    },
    {
      "device": "/dev/vdc",
      "id": "%s",
      "serverId": "%s",
      "volumeId": "%s"
    }
  ]
}`,
	volumeID, serverID, volumeID,
	volumeID, serverID, volumeID,
)

var expectedVolumeAttachmentSlice = []volumeattach.VolumeAttachment{
	firstVolumeAttachment,
	secondVolumeAttachment,
}

var firstVolumeAttachment = volumeattach.VolumeAttachment{
	Device:   "/dev/vdd",
	ID:       volumeID,
	ServerID: serverID,
	VolumeID: volumeID,
}

var secondVolumeAttachment = volumeattach.VolumeAttachment{
	Device:   "/dev/vdc",
	ID:       volumeID,
	ServerID: serverID,
	VolumeID: volumeID,
}

var getResponse = fmt.Sprintf(`
{
  "volumeAttachment": {
    "device": "/dev/vdc",
    "id": "%s",
    "serverId": "%s",
    "volumeId": "%s"
  }
}`, volumeID, serverID, volumeID)

var createRequest = fmt.Sprintf(`
{
  "volumeAttachment": {
    "volumeId": "%s",
    "device": "/dev/vdc"
  }
}`, volumeID)

var createResponse = fmt.Sprintf(`
{
  "volumeAttachment": {
    "device": "/dev/vdc",
    "id": "%s",
    "serverId": "%s",
    "volumeId": "%s"
  }
}`, volumeID, serverID, volumeID)

var createdVolumeAttachment = volumeattach.VolumeAttachment{
	Device:   "/dev/vdc",
	ID:       volumeID,
	ServerID: serverID,
	VolumeID: volumeID,
}
