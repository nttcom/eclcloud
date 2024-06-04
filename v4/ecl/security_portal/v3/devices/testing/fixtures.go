package testing

import (
	"github.com/nttcom/eclcloud/v4/ecl/security_portal/v3/devices"
)

const listResponse = `
{
  "devices": [
    {
      "msa_device_id": "CES11810",
      "os_server_id": "392a90bf-2c1b-45fd-8221-096894fff39d",
      "os_server_name": "UTM-CES11878",
      "os_availability_zone": "zone1-groupb",
      "os_admin_username": "jp4_sdp_mss_utm_admin",
      "msa_device_type": "FW",
      "os_server_status": "ACTIVE"
    },
    {
      "msa_device_id": "CES11811",
      "os_server_id": "12768064-e7c9-44d1-b01d-e66f138a278e",
      "os_server_name": "WAF-CES11816",
      "os_availability_zone": "zone1-groupb",
      "os_admin_username": "jp4_sdp_mss_utm_admin",
      "msa_device_type": "WAF",
      "os_server_status": "ACTIVE"
    }
  ]
}`

var expectedDevicesSlice = []devices.Device{firstDevice, secondDevice}

var firstDevice = devices.Device{
	MSADeviceID:        "CES11810",
	OSServerID:         "392a90bf-2c1b-45fd-8221-096894fff39d",
	OSServerName:       "UTM-CES11878",
	OSAvailabilityZone: "zone1-groupb",
	OSAdminUserName:    "jp4_sdp_mss_utm_admin",
	MSADeviceType:      "FW",
	OSServerStatus:     "ACTIVE",
}

var secondDevice = devices.Device{
	MSADeviceID:        "CES11811",
	OSServerID:         "12768064-e7c9-44d1-b01d-e66f138a278e",
	OSServerName:       "WAF-CES11816",
	OSAvailabilityZone: "zone1-groupb",
	OSAdminUserName:    "jp4_sdp_mss_utm_admin",
	MSADeviceType:      "WAF",
	OSServerStatus:     "ACTIVE",
}
