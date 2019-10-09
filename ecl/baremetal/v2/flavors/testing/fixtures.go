package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/ecl/baremetal/v2/flavors"
)

var listResponse = fmt.Sprintf(`
{
    "flavors": [
        {
            "id": "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
            "name": "General Purpose 1",
            "vcpus": 4,
            "ram": 32768,
            "disk": 550,
            "links": [
                {
                    "href": "https://baremetal-server.ntt/v2/1bc271e7a8af4d988ff91612f5b122f8/flavors/cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
                    "rel": "self"
                },
                {
                    "href": "https://baremetal-server.ntt/1bc271e7a8af4d988ff91612f5b122f8/flavors/cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
                    "rel": "bookmark"
                }
            ]
        },
        {
            "id": "303b4993-cf29-4301-abd0-99512b5413a5",
            "name": "General Purpose 2",
            "vcpus": 8,
            "ram": 262144,
            "disk": 3950,
            "links": [
                {
                    "href": "https://baremetal-server.ntt/v2/1bc271e7a8af4d988ff91612f5b122f8/flavors/303b4993-cf29-4301-abd0-99512b5413a5",
                    "rel": "self"
                },
                {
                    "href": "https://baremetal-server.ntt/1bc271e7a8af4d988ff91612f5b122f8/flavors/303b4993-cf29-4301-abd0-99512b5413a5",
                    "rel": "bookmark"
                }
			]
		}
	]
}`)

var getResponse = fmt.Sprintf(`
{
    "flavor": {
        "id": "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
        "links": [
            {
                "href": "https://baremetal-server.ntt/v2/1bc271e7a8af4d988ff91612f5b122f8/flavors/cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
                "rel": "self"
            },
            {
                "href": "https://baremetal-server.ntt/1bc271e7a8af4d988ff91612f5b122f8/flavors/cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
                "rel": "bookmark"
            }
        ],
        "name": "General Purpose 1",
        "vcpus": 4,
        "ram": 32768,
        "disk": 550
    }
}`)

var expectedFlavors = []flavors.Flavor{expectedFlavor1, expectedFlavor2}

var expectedFlavor1 = flavors.Flavor{
	ID:    "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79",
	Name:  "General Purpose 1",
	Disk:  550,
	RAM:   32768,
	VCPUs: 4,
}

var expectedFlavor2 = flavors.Flavor{
	ID:    "303b4993-cf29-4301-abd0-99512b5413a5",
	Name:  "General Purpose 2",
	Disk:  3950,
	RAM:   262144,
	VCPUs: 8,
}
