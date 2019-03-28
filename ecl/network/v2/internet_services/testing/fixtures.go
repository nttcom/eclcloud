package testing

import(
	"github.com/nttcom/eclcloud/ecl/network/v2/internet_services"
)

const ListResponse = `
{
	"internet_services": [
        {
            "description": "lab3 internet connectivity provider",
            "id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
            "minimal_submask_length": 26,
            "name": "Lab3-Internet-Service-Provider-01",
            "zone": "Lab3"
        },
        {
            "description": "Example internet_service 2 description.",
            "id": "5d6eaf32-8c42-4187-973b-dcee142dcb9d",
            "minimal_submask_length": 26,
            "name": "Example internet_service 2",
            "zone": "Lab3"
        }
    ]
}`

const GetResponse = `{
	"internet_service": {
		"description": "lab3 internet connectivity provider",
		"id": "a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
		"minimal_submask_length": 26,
		"name": "Lab3-Internet-Service-Provider-01",
		"zone": "Lab3"
	}
}`


var InternetService1 = internet_services.InternetService{
	Description:			"lab3 internet connectivity provider",
	ID:						"a7791c79-19b0-4eb6-9a8f-ea739b44e8d5",
	MinimalSubmaskLength:	26,
	Name:					"Lab3-Internet-Service-Provider-01",
	Zone:					"Lab3",
}

var InternetService2 = internet_services.InternetService{
	Description:			"Example internet_service 2 description.",
	ID:						"5d6eaf32-8c42-4187-973b-dcee142dcb9d",
	MinimalSubmaskLength:	26,
	Name:					"Example internet_service 2",
	Zone:					"Lab3",
}

var ExpectedInternetServiceSlice = []internet_services.InternetService{InternetService1, InternetService2}