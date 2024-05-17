package testing

import (
	"fmt"
)

const volumeID = "ff2ac0fd-ea58-4e15-bd71-aec0bc58c469"
const instanceID = "ff2ac0fd-ea58-4e15-bd71-aec0bc58c469"

const uploadImageRequest = `{
	"os-volume_upload_image": {
		"container_format": "bare",
		"force": true,
		"image_name": "imagetest",
		"disk_format": "raw"
	}
}`

var uploadImageResponse = fmt.Sprintf(`{
	"os-volume_upload_image": {
		"status": "uploading",
		"image_id": "49d7efe7-975e-46d7-af0a-fd94fe8e62bf",
		"image_name": "imagetest",
		"volume_type": {
			"name": "nfsdriver",
			"qos_specs_id": null,
			"deleted": false,
			"created_at": "2018-06-04T08:05:09.000000",
			"updated_at": null,
			"deleted_at": null,
			"id": "1f02ea8f-3823-4e69-a232-695adc39f5e0"
		},
		"container_format": "bare",
		"size": 40,
		"disk_format": "raw",
		"id": "%s",
		"display_description": "test volume 2update",
		"updated_at": "2019-02-06T22:06:27.000000"
	}
}`,
	volumeID,
)

const extendRequest = `{
    "os-extend":
    {
        "new_size": 40
    }
}`
