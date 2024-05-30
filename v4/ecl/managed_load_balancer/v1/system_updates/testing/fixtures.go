package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/system_updates"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "system_updates": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "security_update_202210",
            "description": "description",
            "href": "https://sdpf.ntt.com/news/2022100301/",
            "publish_datetime": "2022-10-03 00:00:00",
            "limit_datetime": "2022-10-11 12:59:59",
            "current_revision": 1,
            "next_revision": 2,
            "applicable": true
        }
    ]
}`)

func listResult() []system_updates.SystemUpdate {
	var systemUpdate1 system_updates.SystemUpdate

	systemUpdate1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	systemUpdate1.Name = "security_update_202210"
	systemUpdate1.Description = "description"
	systemUpdate1.Href = "https://sdpf.ntt.com/news/2022100301/"
	systemUpdate1.PublishDatetime = "2022-10-03 00:00:00"
	systemUpdate1.LimitDatetime = "2022-10-11 12:59:59"
	systemUpdate1.CurrentRevision = 1
	systemUpdate1.NextRevision = 2
	systemUpdate1.Applicable = true

	return []system_updates.SystemUpdate{systemUpdate1}
}

var showResponse = fmt.Sprintf(`
{
    "system_update": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "security_update_202210",
        "description": "description",
        "href": "https://sdpf.ntt.com/news/2022100301/",
        "publish_datetime": "2022-10-03 00:00:00",
        "limit_datetime": "2022-10-11 12:59:59",
        "current_revision": 1,
        "next_revision": 2,
        "applicable": true
    }
}`)

func showResult() *system_updates.SystemUpdate {
	var systemUpdate system_updates.SystemUpdate

	systemUpdate.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	systemUpdate.Name = "security_update_202210"
	systemUpdate.Description = "description"
	systemUpdate.Href = "https://sdpf.ntt.com/news/2022100301/"
	systemUpdate.PublishDatetime = "2022-10-03 00:00:00"
	systemUpdate.LimitDatetime = "2022-10-11 12:59:59"
	systemUpdate.CurrentRevision = 1
	systemUpdate.NextRevision = 2
	systemUpdate.Applicable = true

	return &systemUpdate
}
