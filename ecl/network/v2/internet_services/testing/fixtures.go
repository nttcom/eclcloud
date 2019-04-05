package testing

import (
	"github.com/nttcom/eclcloud/ecl/network/v2/internet_services"
)

const ListResponse = `
{
    "internet_services": [
        {
            "description": "Example internet_service 1 description",
            "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
            "minimal_submask_length": 26,
            "name": "Internet-Service-01",
            "zone": "jp1-zone1"
        },
        {
            "description": "Example internet_service 2 description.",
            "id": "5d6eaf32-8c42-4187-973b-dcee142dcb9d",
            "minimal_submask_length": 26,
            "name": "Internet-Service-01",
            "zone": "jp2-zone1"
        }
    ]
}`

const GetResponse = `{
    "internet_service": {
        "description": "Example internet_service 1 description",
        "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
        "minimal_submask_length": 26,
        "name": "Internet-Service-01",
        "zone": "jp1-zone1"
    }
}`

var InternetService1 = internet_services.InternetService{
	Description:          "Example internet_service 1 description",
	ID:                   "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
	MinimalSubmaskLength: 26,
	Name:                 "Internet-Service-01",
	Zone:                 "jp1-zone1",
}

var InternetService2 = internet_services.InternetService{
	Description:          "Example internet_service 2 description.",
	ID:                   "5d6eaf32-8c42-4187-973b-dcee142dcb9d",
	MinimalSubmaskLength: 26,
	Name:                 "Internet-Service-01",
	Zone:                 "jp2-zone1",
}

var ExpectedInternetServiceSlice = []internet_services.InternetService{InternetService1, InternetService2}
