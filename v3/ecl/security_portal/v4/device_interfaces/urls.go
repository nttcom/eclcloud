package device_interfaces

import (
	"fmt"

	"github.com/nttcom/eclcloud/v3"
)

func listURL(client *eclcloud.ServiceClient, serverUUID string) string {
	url := fmt.Sprintf("ecl-api/v4/devices/%s/interfaces", serverUUID)
	return client.ServiceURL(url)
}
