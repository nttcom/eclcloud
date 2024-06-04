package testing

import (
	"github.com/nttcom/eclcloud/v3/ecl/vna/v1/appliance_plans"
)

const ListResponse = `
{
  "virtual_network_appliance_plans": [
    {
      "id": "37556569-87f2-4699-b5ff-bf38e7cbf8a7",
      "name": "appliance_plans_name",
      "description": "appliance_plans_description",
      "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
      "version": "",
      "flavor": "2CPU-8GB",
      "number_of_interfaces": 8,
      "enabled": true,
      "max_number_of_aap": 1,
      "licenses": [
        {
          "license_type": "STD"
        }
      ],
      "availability_zones": [
        {
          "availability_zone": "zone1_groupa",
          "available": true,
          "rank": 1
        },
        {
          "availability_zone": "zone1_groupb",
          "available": false,
          "rank": 2
        }
      ]
    }
  ]
}
`
const GetResponse = `
{
  "virtual_network_appliance_plan": {
    "id": "6589b37a-cf82-4918-96fe-255683f78e76",
    "name": "vSRX_15.1X49-D100_2CPU_4GB_8IF_STD",
    "description": "vSRX_15.1X49-D100_2CPU_4GB_8IF_STD",
    "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
    "version": "15.1X49-D100",
    "flavor": "VSRX-2CPU-4GB",
    "number_of_interfaces": 8,
    "enabled": true,
    "max_number_of_aap": 1,
    "licenses": [
      {
        "license_type": "STD"
      }
    ],
    "availability_zones": [
      {
        "availability_zone": "zone1_groupa",
        "available": true,
        "rank": 1
      },
      {
        "availability_zone": "zone1_groupb",
        "available": false,
        "rank": 2
      }
    ]
  }
}
`

var VirtualNetworkAppliancePlan1 = appliance_plans.VirtualNetworkAppliancePlan{
	ID:                 "37556569-87f2-4699-b5ff-bf38e7cbf8a7",
	Name:               "appliance_plans_name",
	Description:        "appliance_plans_description",
	ApplianceType:      "ECL::VirtualNetworkAppliance::VSRX",
	Version:            "",
	Flavor:             "2CPU-8GB",
	NumberOfInterfaces: 8,
	Enabled:            true,
	MaxNumberOfAap:     1,
	Licenses: []appliance_plans.License{
		{
			LicenseType: "STD",
		},
	},
	AvailabilityZones: []appliance_plans.AvailabilityZone{
		{
			AvailabilityZone: "zone1_groupa",
			Available:        true,
			Rank:             1,
		},
		{
			AvailabilityZone: "zone1_groupb",
			Available:        false,
			Rank:             2,
		},
	},
}

var VirtualNetworkApplianceDetail = appliance_plans.VirtualNetworkAppliancePlan{
	ID:                 "6589b37a-cf82-4918-96fe-255683f78e76",
	Name:               "vSRX_15.1X49-D100_2CPU_4GB_8IF_STD",
	Description:        "vSRX_15.1X49-D100_2CPU_4GB_8IF_STD",
	ApplianceType:      "ECL::VirtualNetworkAppliance::VSRX",
	Version:            "15.1X49-D100",
	Flavor:             "VSRX-2CPU-4GB",
	NumberOfInterfaces: 8,
	Enabled:            true,
	MaxNumberOfAap:     1,
	Licenses: []appliance_plans.License{
		{
			LicenseType: "STD",
		},
	},
	AvailabilityZones: []appliance_plans.AvailabilityZone{
		{
			AvailabilityZone: "zone1_groupa",
			Available:        true,
			Rank:             1,
		},
		{
			AvailabilityZone: "zone1_groupb",
			Available:        false,
			Rank:             2,
		},
	},
}

var ExpectedVirtualNetworkAppliancePlanSlice = []appliance_plans.VirtualNetworkAppliancePlan{VirtualNetworkAppliancePlan1}

const ListResponseDuplicatedNames = `
{
  "virtual_network_appliance_plans": [
    {
      "id": "37556569-87f2-4699-b5ff-bf38e7cbf8a7",
      "name": "appliance_plans_name",
      "description": "appliance_plans_description",
      "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
      "version": "",
      "flavor": "2CPU-8GB",
      "number_of_interfaces": 8,
      "enabled": true,
      "max_number_of_aap": 1,
      "licenses": [
        {
          "license_type": "STD"
        }
      ],
      "availability_zones": [
        {
          "availability_zone": "zone1_groupa",
          "available": true,
          "rank": 1
        },
        {
          "availability_zone": "zone1_groupb",
          "available": false,
          "rank": 2
        }
      ]
    },
    {
      "id": "6589b37a-cf82-4918-96fe-255683f78e76",
      "name": "appliance_plans_name",
      "description": "appliance_plans_description",
      "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
      "version": "",
      "flavor": "2CPU-8GB",
      "number_of_interfaces": 8,
      "enabled": true,
      "max_number_of_aap": 1,
      "licenses": [
        {
          "license_type": "STD"
        }
      ],
      "availability_zones": [
        {
          "availability_zone": "zone1_groupa",
          "available": true,
          "rank": 1
        },
        {
          "availability_zone": "zone1_groupb",
          "available": false,
          "rank": 2
        }
      ]
    }
  ]
}
`
