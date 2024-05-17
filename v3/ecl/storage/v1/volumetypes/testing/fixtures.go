package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v3/ecl/storage/v1/volumetypes"
)

// Define parameters which are used in assertion.
// Additionally, kind of IDs are defined here.
const idVolumeType1 = "6328d234-7939-4d61-9216-736de66d15f9"
const idVolumeType2 = "bf33db2a-d13e-11e5-8949-005056ab5d30"
const idVolumeType3 = "704db6e5-8a93-41a5-850d-405913600341"

// ListResponse is mocked response of volumetypes.List
var ListResponse = fmt.Sprintf(`
{
	"volume_types": [
		{
			"extra_specs": {
				"available_volume_size": [
					100,
					250,
					500,
					1000,
					2000,
					4000,
					8000,
					12000
				],
				"available_iops_per_gb": [
					"2",
					"4"
				]
			},
			"id": "%s",
			"name": "piops_iscsi_na"
		},
		{
			"extra_specs": {
				"available_volume_size": [
					256,
					512
				],
				"available_volume_throughput": [
					"50",
					"100",
					"250",
					"400"
				]
			},
			"id": "%s",
			"name": "pre_nfs_na"
		},
		{
			"extra_specs": {
				"available_volume_size": [
					1024,
					2048,
					3072,
					4096,
					5120,
					10240,
					15360,
					20480,
					25600,
					30720,
					35840,
					40960,
					46080,
					51200,
					56320,
					61440,
					66560,
					71680,
					76800,
					81920,
					87040,
					92160,
					97280,
					102400
				]
			},
			"id": "%s",
			"name": "standard_nfs_na"
		}
	]
}`,
	idVolumeType1,
	idVolumeType2,
	idVolumeType3,
)

// GetResponse is mocked format of volumetypes.Get
var GetResponse = fmt.Sprintf(`
{
	"volume_type": {
		"extra_specs": {
			"available_volume_size": [
				100,
				250,
				500,
				1000,
				2000,
				4000,
				8000,
				12000
			],
			"available_iops_per_gb": [
				"2",
				"4"
			]
		},
		"id": "%s",
		"name": "piops_iscsi_na"
	}
}`, idVolumeType1,
)

func getExpectedVolumeTypesSlice() []volumetypes.VolumeType {

	// For Block Storage Type
	var volumetype1 = volumetypes.VolumeType{
		ID:   idVolumeType1,
		Name: "piops_iscsi_na",
		ExtraSpecs: volumetypes.ExtraSpec{
			AvailableVolumeSize: []int{
				100, 250, 500, 1000, 2000, 4000, 8000, 12000,
			},
			AvailableIOPSPerGB: []string{"2", "4"},
		},
	}

	// For File Storage(Premium) Type
	var volumetype2 = volumetypes.VolumeType{
		ID:   idVolumeType2,
		Name: "pre_nfs_na",
		ExtraSpecs: volumetypes.ExtraSpec{
			AvailableVolumeSize: []int{
				256, 512,
			},
			AvailableVolumeThroughput: []string{
				"50", "100", "250", "400",
			},
		},
	}

	// For File Storage(Standard) Type
	var volumetype3 = volumetypes.VolumeType{
		ID:   idVolumeType3,
		Name: "standard_nfs_na",
		ExtraSpecs: volumetypes.ExtraSpec{
			AvailableVolumeSize: []int{
				1024, 2048, 3072, 4096, 5120, 10240,
				15360, 20480, 25600, 30720, 35840, 40960,
				46080, 51200, 56320, 61440, 66560, 71680,
				76800, 81920, 87040, 92160, 97280, 102400,
			},
		},
	}

	// ExpectedVolumeTypesSlice is expected assertion target
	ExpectedVolumeTypesSlice := []volumetypes.VolumeType{
		volumetype1,
		volumetype2,
		volumetype3,
	}

	return ExpectedVolumeTypesSlice
}
