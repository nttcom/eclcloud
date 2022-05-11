package testing

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/v2/ecl/dedicated_hypervisor/v1/servers"

	th "github.com/nttcom/eclcloud/v2/testhelper"
	"github.com/nttcom/eclcloud/v2/testhelper/client"
)

// ListResult provides a single page of Server results.
const ListResult = `
{
	"servers": [
		{
			"id": "194573e4-8f53-4ee4-806f-d9b2db74a380",
			"name": "GP2v1",
			"links": [
				{
					"href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//v2/1bc271e7a8af4d988ff91612f5b122f8/servers/194573e4-8f53-4ee4-806f-d9b2db74a380",
					"rel": "self"
				},
				{
					"href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//1bc271e7a8af4d988ff91612f5b122f8/servers/194573e4-8f53-4ee4-806f-d9b2db74a380",
					"rel": "bookmark"
				}
			],
			"baremetal_server": {
				"id": "621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
				"links": [
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
						"rel": "self"
					},
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
						"rel": "bookmark"
					}
				],
				"name": "GP2v1"
			}
		},
		{
			"id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
			"name": "test",
			"links": [
				{
					"href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//v2/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
					"rel": "self"
				},
				{
					"href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
					"rel": "bookmark"
				}
			],
			"baremetal_server": {
				"id": "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
				"links": [
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
						"rel": "self"
					},
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
						"rel": "bookmark"
					}
				],
				"name": "test"
			}
		}
	]
}
`

// ListDetailsResult provides a single page of Server results in details.
const ListDetailsResult = `
{
	"servers": [
		{
			"id": "194573e4-8f53-4ee4-806f-d9b2db74a380",
			"name": "GP2v1",
			"imageRef": "293063f6-8986-4b79-becd-7a6d28794bb8",
			"description": null,
			"status": "ACTIVE",
			"hypervisor_type": "vsphere_esxi",
			"baremetal_server": {
				"OS-EXT-STS:power_state": "RUNNING",
				"OS-EXT-STS:task_state": "None",
				"OS-EXT-STS:vm_state": "ACTIVE",
				"OS-EXT-AZ:availability_zone": "groupb",
				"progress": 100,
				"created": "2019-10-18T07:42:35Z",
				"flavor": {
					"id": "303b4993-cf29-4301-abd0-99512b5413a5",
					"links": [
						{
							"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/303b4993-cf29-4301-abd0-99512b5413a5",
							"rel": "bookmark"
						}
					]
				},
				"id": "621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
				"image": {
					"id": "02441adc-0d9a-4e9d-b359-ce23413e7ea7",
					"links": [
						{
							"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/02441adc-0d9a-4e9d-b359-ce23413e7ea7",
							"rel": "bookmark"
						}
					]
				},
				"links": [
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
						"rel": "self"
					},
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
						"rel": "bookmark"
					}
				],
				"metadata": {},
				"name": "GP2v1",
				"status": "ACTIVE",
				"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
				"updated": "2019-10-18T07:44:18Z",
				"user_id": "55891ce6a3cb4bb0833514667d67288c",
				"raid_arrays": [
					{
						"primary_storage": true,
						"partitions": null,
						"raid_card_hardware_id": "24184dcf-dc76-4ea2-a34e-bccc6c11d5be",
						"disk_hardware_ids": [
							"4de7d3df-8e2f-4193-98dc-145f78df29a2",
							"de158fab-7bc2-49e3-98d1-9db5451d43e3",
							"97087307-cab2-40cb-a84c-86d98da0f393"
						]
					}
				],
				"lvm_volume_groups": null,
				"filesystems": null,
				"nic_physical_ports": [
					{
						"id": "b49c0624-e89a-469f-8c90-7a27ee1f61cb",
						"mac_addr": "8C:DC:D4:B7:41:48",
						"plane": "DATA",
						"network_physical_port_id": "4cfbe3b2-a502-485f-82fa-a0949396e567",
						"hardware_id": "8e1e2fe0-60a7-4211-a891-80e808426708",
						"attached_ports": [
							{
								"network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
								"port_id": "8808acc2-d930-40fb-b382-ce7074baef83",
								"fixed_ips": [
									{
										"subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
										"ip_address": "169.254.0.11"
									}
								]
							},
							{
								"network_id": "722f9e4f-39f8-406a-b98c-5fbd5689b89a",
								"port_id": "9d3baa16-e0e5-4e50-9677-08dd338e0c14",
								"fixed_ips": [
									{
										"subnet_id": "dc84c9dc-0b4d-40fc-8605-e518af7cdd30",
										"ip_address": "192.168.4.3"
									}
								]
							}
						]
					},
					{
						"id": "8bea93c4-721b-480c-8713-f2a4b6e5dbad",
						"mac_addr": "8C:DC:D4:B7:41:49",
						"plane": "STORAGE",
						"network_physical_port_id": "ab38075d-128f-4f3d-a16a-c6426375a380",
						"hardware_id": "8e1e2fe0-60a7-4211-a891-80e808426708",
						"attached_ports": []
					},
					{
						"id": "a7fbca5e-ff49-4ddb-8659-02ec462f98ec",
						"mac_addr": "8C:DC:D4:B7:45:89",
						"plane": "STORAGE",
						"network_physical_port_id": "9ef62803-7848-460c-9dae-17fd02606a26",
						"hardware_id": "1aa1c2a4-2608-41e6-b4f5-87679d1aea43",
						"attached_ports": []
					},
					{
						"id": "71af4de1-0c9b-4870-8c95-c7b9b4115bb8",
						"mac_addr": "8C:DC:D4:B7:45:88",
						"plane": "DATA",
						"network_physical_port_id": "ad92f33a-eac4-408d-bc27-22d91eccd465",
						"hardware_id": "1aa1c2a4-2608-41e6-b4f5-87679d1aea43",
						"attached_ports": [
							{
								"network_id": "722f9e4f-39f8-406a-b98c-5fbd5689b89a",
								"port_id": "705bde94-7189-40b1-b8a2-188e6cc3c546",
								"fixed_ips": [
									{
										"subnet_id": "dc84c9dc-0b4d-40fc-8605-e518af7cdd30",
										"ip_address": "192.168.4.4"
									}
								]
							},
							{
								"network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
								"port_id": "7b860eb4-0eb6-4c2a-873e-ccefd6029d97",
								"fixed_ips": [
									{
										"subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
										"ip_address": "169.254.0.12"
									}
								]
							}
						]
					}
				],
				"chassis-status": {
					"chassis-power": true,
					"power-supply": true,
					"cpu": true,
					"memory": true,
					"fan": true,
					"disk": 0,
					"nic": true,
					"system-board": true,
					"etc": true,
					"console": true
				},
				"media_attachments": [],
				"managed_by_service": "dedicated-hypervisor",
				"managed_service_resource_id": "194573e4-8f53-4ee4-806f-d9b2db74a380"
			}
		},
		{
			"id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
			"name": "test",
			"imageRef": "dfd25820-b368-4012-997b-29a6d0cf8518",
			"description": "test",
			"status": "ACTIVE",
			"hypervisor_type": "vsphere_esxi",
			"baremetal_server": {
				"OS-EXT-STS:power_state": "RUNNING",
				"OS-EXT-STS:task_state": "None",
				"OS-EXT-STS:vm_state": "ACTIVE",
				"OS-EXT-AZ:availability_zone": "groupb",
				"progress": 100,
				"created": "2019-10-10T04:11:41Z",
				"flavor": {
					"id": "a830b61c-3155-4a61-b7ed-c450862845e6",
					"links": [
						{
							"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/a830b61c-3155-4a61-b7ed-c450862845e6",
							"rel": "bookmark"
						}
					]
				},
				"id": "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
				"image": {
					"id": "112a26a0-ff25-4513-afe1-407e41b0a48b",
					"links": [
						{
							"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/112a26a0-ff25-4513-afe1-407e41b0a48b",
							"rel": "bookmark"
						}
					]
				},
				"links": [
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
						"rel": "self"
					},
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
						"rel": "bookmark"
					}
				],
				"metadata": {},
				"name": "test",
				"status": "ACTIVE",
				"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
				"updated": "2019-10-10T04:14:08Z",
				"user_id": "55891ce6a3cb4bb0833514667d67288c",
				"raid_arrays": [
					{
						"primary_storage": true,
						"partitions": null,
						"raid_card_hardware_id": "bdfb75d1-194d-426d-b288-f588dfa5ac49",
						"disk_hardware_ids": [
							"76649053-863e-4533-86e3-f194a79485a6",
							"a25827e3-67da-47be-ba96-849ab4685a1d"
						]
					}
				],
				"lvm_volume_groups": null,
				"filesystems": null,
				"nic_physical_ports": [
					{
						"id": "a2f63380-6c77-4cd5-8868-e3556ffd35ce",
						"mac_addr": "48:DF:37:90:B4:58",
						"plane": "DATA",
						"network_physical_port_id": "d8e40a51-f1e2-4681-8953-9fe1e9992c42",
						"hardware_id": "be2d30d6-f891-4200-b827-95f229fb8c6b",
						"attached_ports": [
							{
								"network_id": "94055904-6b2c-4839-a14a-c61c93a8bc48",
								"port_id": "30fc1c27-fb5f-4955-94d0-a56cd28d09e8",
								"fixed_ips": [
									{
										"subnet_id": "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
										"ip_address": "2.1.1.10"
									}
								]
							},
							{
								"network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
								"port_id": "aa6c61f4-db8a-44c7-a91c-7e636dac1dc6",
								"fixed_ips": [
									{
										"subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
										"ip_address": "169.254.0.9"
									}
								]
							}
						]
					},
					{
						"id": "b01dfdb0-f247-47d8-8224-c257aa3265e9",
						"mac_addr": "48:DF:37:90:B4:50",
						"plane": "STORAGE",
						"network_physical_port_id": "00dfea92-5c5b-4860-aa05-efef6c2bb2af",
						"hardware_id": "be2d30d6-f891-4200-b827-95f229fb8c6b",
						"attached_ports": []
					},
					{
						"id": "f4355e8e-39fc-48bd-a283-a2dbef8a2e32",
						"mac_addr": "48:DF:37:82:B0:A0",
						"plane": "STORAGE",
						"network_physical_port_id": "cf798cc0-c869-45d5-a5a7-bcc578a300b0",
						"hardware_id": "84c74a86-7045-4284-80f9-0e7aff5d27ad",
						"attached_ports": []
					},
					{
						"id": "5ef177fd-888c-4fae-9925-a8920beb07cb",
						"mac_addr": "48:DF:37:82:B0:A8",
						"plane": "DATA",
						"network_physical_port_id": "2bbbb516-c75a-42b2-8a46-9cb5f26c219e",
						"hardware_id": "84c74a86-7045-4284-80f9-0e7aff5d27ad",
						"attached_ports": [
							{
								"network_id": "94055904-6b2c-4839-a14a-c61c93a8bc48",
								"port_id": "4e329a01-2cf4-4028-9259-03b7aa145cb6",
								"fixed_ips": [
									{
										"subnet_id": "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
										"ip_address": "2.1.1.20"
									}
								]
							},
							{
								"network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
								"port_id": "a256b4a1-3ae3-4102-a14e-987ae1610f97",
								"fixed_ips": [
									{
										"subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
										"ip_address": "169.254.0.10"
									}
								]
							}
						]
					}
				],
				"chassis-status": {
					"chassis-power": true,
					"power-supply": true,
					"cpu": true,
					"memory": true,
					"fan": true,
					"disk": 0,
					"nic": true,
					"system-board": true,
					"etc": true,
					"console": true
				},
				"media_attachments": [],
				"managed_by_service": "dedicated-hypervisor",
				"managed_service_resource_id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5"
			}
		}
	]
}
`

// GetResult provides a Get result.
const GetResult = `
{
	"server": {
		"id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
		"name": "test",
		"imageRef": "dfd25820-b368-4012-997b-29a6d0cf8518",
		"description": "test",
		"status": "ACTIVE",
		"hypervisor_type": "vsphere_esxi",
		"baremetal_server": {
			"OS-EXT-STS:power_state": "RUNNING",
			"OS-EXT-STS:task_state": "None",
			"OS-EXT-STS:vm_state": "ACTIVE",
			"OS-EXT-AZ:availability_zone": "groupb",
			"progress": 100,
			"created": "2019-10-10T04:11:41Z",
			"flavor": {
				"id": "a830b61c-3155-4a61-b7ed-c450862845e6",
				"links": [
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/a830b61c-3155-4a61-b7ed-c450862845e6",
						"rel": "bookmark"
					}
				]
			},
			"id": "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
			"image": {
				"id": "112a26a0-ff25-4513-afe1-407e41b0a48b",
				"links": [
					{
						"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/112a26a0-ff25-4513-afe1-407e41b0a48b",
						"rel": "bookmark"
					}
				]
			},
			"links": [
				{
					"href": "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
					"rel": "self"
				},
				{
					"href": "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
					"rel": "bookmark"
				}
			],
			"metadata": {},
			"name": "test",
			"status": "ACTIVE",
			"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
			"updated": "2019-10-10T04:14:08Z",
			"user_id": "55891ce6a3cb4bb0833514667d67288c",
			"raid_arrays": [
				{
					"primary_storage": true,
					"partitions": null,
					"raid_card_hardware_id": "bdfb75d1-194d-426d-b288-f588dfa5ac49",
					"disk_hardware_ids": [
						"76649053-863e-4533-86e3-f194a79485a6",
						"a25827e3-67da-47be-ba96-849ab4685a1d"
					]
				}
			],
			"lvm_volume_groups": null,
			"filesystems": null,
			"nic_physical_ports": [
				{
					"id": "a2f63380-6c77-4cd5-8868-e3556ffd35ce",
					"mac_addr": "48:DF:37:90:B4:58",
					"plane": "DATA",
					"network_physical_port_id": "d8e40a51-f1e2-4681-8953-9fe1e9992c42",
					"hardware_id": "be2d30d6-f891-4200-b827-95f229fb8c6b",
					"attached_ports": [
						{
							"network_id": "94055904-6b2c-4839-a14a-c61c93a8bc48",
							"port_id": "30fc1c27-fb5f-4955-94d0-a56cd28d09e8",
							"fixed_ips": [
								{
									"subnet_id": "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
									"ip_address": "2.1.1.10"
								}
							]
						},
						{
							"network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
							"port_id": "aa6c61f4-db8a-44c7-a91c-7e636dac1dc6",
							"fixed_ips": [
								{
									"subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
									"ip_address": "169.254.0.9"
								}
							]
						}
					]
				},
				{
					"id": "b01dfdb0-f247-47d8-8224-c257aa3265e9",
					"mac_addr": "48:DF:37:90:B4:50",
					"plane": "STORAGE",
					"network_physical_port_id": "00dfea92-5c5b-4860-aa05-efef6c2bb2af",
					"hardware_id": "be2d30d6-f891-4200-b827-95f229fb8c6b",
					"attached_ports": []
				},
				{
					"id": "f4355e8e-39fc-48bd-a283-a2dbef8a2e32",
					"mac_addr": "48:DF:37:82:B0:A0",
					"plane": "STORAGE",
					"network_physical_port_id": "cf798cc0-c869-45d5-a5a7-bcc578a300b0",
					"hardware_id": "84c74a86-7045-4284-80f9-0e7aff5d27ad",
					"attached_ports": []
				},
				{
					"id": "5ef177fd-888c-4fae-9925-a8920beb07cb",
					"mac_addr": "48:DF:37:82:B0:A8",
					"plane": "DATA",
					"network_physical_port_id": "2bbbb516-c75a-42b2-8a46-9cb5f26c219e",
					"hardware_id": "84c74a86-7045-4284-80f9-0e7aff5d27ad",
					"attached_ports": [
						{
							"network_id": "94055904-6b2c-4839-a14a-c61c93a8bc48",
							"port_id": "4e329a01-2cf4-4028-9259-03b7aa145cb6",
							"fixed_ips": [
								{
									"subnet_id": "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
									"ip_address": "2.1.1.20"
								}
							]
						},
						{
							"network_id": "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
							"port_id": "a256b4a1-3ae3-4102-a14e-987ae1610f97",
							"fixed_ips": [
								{
									"subnet_id": "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
									"ip_address": "169.254.0.10"
								}
							]
						}
					]
				}
			],
			"chassis-status": {
				"chassis-power": true,
				"power-supply": true,
				"cpu": true,
				"memory": true,
				"fan": true,
				"disk": 0,
				"nic": true,
				"system-board": true,
				"etc": true,
				"console": true
			},
			"media_attachments": [],
			"managed_by_service": "dedicated-hypervisor",
			"managed_service_resource_id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5"
		}
	}
}
`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
	"server": {
		"imageRef": "dfd25820-b368-4012-997b-29a6d0cf8518",
		"name": "test",
		"networks": [
			{
				"segmentation_id": 6,
				"plane": "data",
				"uuid": "94055904-6b2c-4839-a14a-c61c93a8bc48"
			},
			{
				"segmentation_id": 6,
				"plane": "data",
				"uuid": "94055904-6b2c-4839-a14a-c61c93a8bc48"
			}
		],
		"flavorRef": "a830b61c-3155-4a61-b7ed-c450862845e6"
	}
}
`

const CreateResponse = `
{
	"server": {
		"id": "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
		"links": [
			{
				"href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//v2/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
				"rel": "self"
			},
			{
				"href": "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
				"rel": "bookmark"
			}
		],
		"adminPass": "aabbccddeeff"
	}
}
`

const AddLicenseRequest = `
{
	"add-license-to-vm": {
		"vm_name": "Alice",
		"license_types": [
			"Windows Server",
			"SQL Server Standard 2014"
		]
	}
}
`

const AddLicenseResponse = `
{
	"job_id": "b4f888dc2b9d4c41bb769cbd"
}
`

const GetAddLicenseResultRequest = `
{
	"get-result-for-add-license-to-vm": {
		"job_id": "b4f888dc2b9d4c41bb769cbd"
	}
}
`

const GetAddLicenseResultResponse = `
{
	"job_id": "b4f888dc2b9d4c41bb769cbd",
	"status": "COMPLETED",
	"requested_param": {
		"vm_name": "Alice",
		"license_types": ["Windows Server", "SQL Server Standard 2014"]
	}
}
`

// FirstServer is the first resource in the List request.
var FirstServer = servers.Server{
	ID:   "194573e4-8f53-4ee4-806f-d9b2db74a380",
	Name: "GP2v1",
	Links: []servers.Link{
		{
			Href: "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//v2/1bc271e7a8af4d988ff91612f5b122f8/servers/194573e4-8f53-4ee4-806f-d9b2db74a380",
			Rel:  "self",
		},
		{
			Href: "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//1bc271e7a8af4d988ff91612f5b122f8/servers/194573e4-8f53-4ee4-806f-d9b2db74a380",
			Rel:  "bookmark",
		},
	},
	BaremetalServer: servers.BaremetalServer{
		ID:   "621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
		Name: "GP2v1",
		Links: []servers.Link{
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
				Rel:  "self",
			},
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
				Rel:  "bookmark",
			},
		},
	},
}

// SecondServer is the second resource in the List request.
var SecondServer = servers.Server{
	ID:   "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
	Name: "test",
	Links: []servers.Link{
		{
			Href: "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//v2/1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
			Rel:  "self",
		},
		{
			Href: "https://dedicated-hypervisor-jp1-ecl.api.ntt.com/v1.0//1bc271e7a8af4d988ff91612f5b122f8/servers/f42dbc37-4642-4628-8b47-50bf95d8fdd5",
			Rel:  "bookmark",
		},
	},
	BaremetalServer: servers.BaremetalServer{
		ID:   "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
		Name: "test",
		Links: []servers.Link{
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
				Rel:  "self",
			},
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
				Rel:  "bookmark",
			},
		},
	},
}

// ExpectedServersSlice is the slice of resources expected to be returned from ListResult.
var ExpectedServersSlice = []servers.Server{FirstServer, SecondServer}

// FirstServerDetail is the first resource in the List details request.
var FirstServerDetail = servers.Server{
	ID:             "194573e4-8f53-4ee4-806f-d9b2db74a380",
	Name:           "GP2v1",
	ImageRef:       "293063f6-8986-4b79-becd-7a6d28794bb8",
	Description:    nil,
	Status:         "ACTIVE",
	HypervisorType: "vsphere_esxi",
	BaremetalServer: servers.BaremetalServer{
		PowerState:       "RUNNING",
		TaskState:        "None",
		VMState:          "ACTIVE",
		AvailabilityZone: "groupb",
		Created:          time.Date(2019, 10, 18, 7, 42, 35, 0, time.UTC),
		Flavor: servers.Flavor{
			ID: "303b4993-cf29-4301-abd0-99512b5413a5",
			Links: []servers.Link{
				{
					Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/303b4993-cf29-4301-abd0-99512b5413a5",
					Rel:  "bookmark",
				},
			},
		},
		ID: "621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
		Image: servers.Image{
			ID: "02441adc-0d9a-4e9d-b359-ce23413e7ea7",
			Links: []servers.Link{
				{
					Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/02441adc-0d9a-4e9d-b359-ce23413e7ea7",
					Rel:  "bookmark",
				},
			},
		},
		Metadata: map[string]string{},
		Name:     "GP2v1",
		Progress: 100,
		Status:   "ACTIVE",
		TenantID: "1bc271e7a8af4d988ff91612f5b122f8",
		Updated:  time.Date(2019, 10, 18, 7, 44, 18, 0, time.UTC),
		UserID:   "55891ce6a3cb4bb0833514667d67288c",
		NicPhysicalPorts: []servers.NicPhysicalPort{
			{
				ID:                    "b49c0624-e89a-469f-8c90-7a27ee1f61cb",
				MACAddr:               "8C:DC:D4:B7:41:48",
				Plane:                 "DATA",
				NetworkPhysicalPortID: "4cfbe3b2-a502-485f-82fa-a0949396e567",
				HardwareID:            "8e1e2fe0-60a7-4211-a891-80e808426708",
				AttachedPorts: []servers.Port{
					{
						NetworkID: "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
						PortID:    "8808acc2-d930-40fb-b382-ce7074baef83",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
								IPAddress: "169.254.0.11",
							},
						},
					},
					{
						NetworkID: "722f9e4f-39f8-406a-b98c-5fbd5689b89a",
						PortID:    "9d3baa16-e0e5-4e50-9677-08dd338e0c14",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "dc84c9dc-0b4d-40fc-8605-e518af7cdd30",
								IPAddress: "192.168.4.3",
							},
						},
					},
				},
			},
			{
				ID:                    "8bea93c4-721b-480c-8713-f2a4b6e5dbad",
				MACAddr:               "8C:DC:D4:B7:41:49",
				Plane:                 "STORAGE",
				NetworkPhysicalPortID: "ab38075d-128f-4f3d-a16a-c6426375a380",
				HardwareID:            "8e1e2fe0-60a7-4211-a891-80e808426708",
				AttachedPorts:         []servers.Port{},
			},
			{
				ID:                    "a7fbca5e-ff49-4ddb-8659-02ec462f98ec",
				MACAddr:               "8C:DC:D4:B7:45:89",
				Plane:                 "STORAGE",
				NetworkPhysicalPortID: "9ef62803-7848-460c-9dae-17fd02606a26",
				HardwareID:            "1aa1c2a4-2608-41e6-b4f5-87679d1aea43",
				AttachedPorts:         []servers.Port{},
			},
			{
				ID:                    "71af4de1-0c9b-4870-8c95-c7b9b4115bb8",
				MACAddr:               "8C:DC:D4:B7:45:88",
				Plane:                 "DATA",
				NetworkPhysicalPortID: "ad92f33a-eac4-408d-bc27-22d91eccd465",
				HardwareID:            "1aa1c2a4-2608-41e6-b4f5-87679d1aea43",
				AttachedPorts: []servers.Port{
					{
						NetworkID: "722f9e4f-39f8-406a-b98c-5fbd5689b89a",
						PortID:    "705bde94-7189-40b1-b8a2-188e6cc3c546",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "dc84c9dc-0b4d-40fc-8605-e518af7cdd30",
								IPAddress: "192.168.4.4",
							},
						},
					},
					{
						NetworkID: "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
						PortID:    "7b860eb4-0eb6-4c2a-873e-ccefd6029d97",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
								IPAddress: "169.254.0.12",
							},
						},
					},
				},
			},
		},
		ChassisStatus: servers.ChassisStatus{
			ChassisPower: true,
			PowerSupply:  true,
			CPU:          true,
			Memory:       true,
			Fan:          true,
			Disk:         0,
			Nic:          true,
			SystemBoard:  true,
			Etc:          true,
			Console:      true,
		},
		Links: []servers.Link{
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
				Rel:  "self",
			},
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/621b56e4-4aae-4de5-86a0-8ffeeda6a00b",
				Rel:  "bookmark",
			},
		},
		RaidArrays: []servers.RaidArray{
			{
				PrimaryStorage:     true,
				Partitions:         nil,
				RaidCardHardwareID: "24184dcf-dc76-4ea2-a34e-bccc6c11d5be",
				DiskHardwareIDs: []string{
					"4de7d3df-8e2f-4193-98dc-145f78df29a2",
					"de158fab-7bc2-49e3-98d1-9db5451d43e3",
					"97087307-cab2-40cb-a84c-86d98da0f393",
				},
			},
		},
		LvmVolumeGroups:          nil,
		Filesystems:              nil,
		MediaAttachments:         []servers.MediaAttachment{},
		ManagedByService:         "dedicated-hypervisor",
		ManagedServiceResourceID: "194573e4-8f53-4ee4-806f-d9b2db74a380",
	},
}

var SecondServerDescription = "test"

// SecondServerDetail is the second resource in the List detail request.
var SecondServerDetail = servers.Server{
	ID:             "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
	Name:           "test",
	ImageRef:       "dfd25820-b368-4012-997b-29a6d0cf8518",
	Description:    &SecondServerDescription,
	Status:         "ACTIVE",
	HypervisorType: "vsphere_esxi",
	BaremetalServer: servers.BaremetalServer{
		PowerState:       "RUNNING",
		TaskState:        "None",
		VMState:          "ACTIVE",
		AvailabilityZone: "groupb",
		Created:          time.Date(2019, 10, 10, 4, 11, 41, 0, time.UTC),
		Flavor: servers.Flavor{
			ID: "a830b61c-3155-4a61-b7ed-c450862845e6",
			Links: []servers.Link{
				{
					Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/a830b61c-3155-4a61-b7ed-c450862845e6",
					Rel:  "bookmark",
				},
			},
		},
		ID: "24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
		Image: servers.Image{
			ID: "112a26a0-ff25-4513-afe1-407e41b0a48b",
			Links: []servers.Link{
				{
					Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/112a26a0-ff25-4513-afe1-407e41b0a48b",
					Rel:  "bookmark",
				},
			},
		},
		Metadata: map[string]string{},
		Name:     "test",
		Progress: 100,
		Status:   "ACTIVE",
		TenantID: "1bc271e7a8af4d988ff91612f5b122f8",
		Updated:  time.Date(2019, 10, 10, 4, 14, 8, 0, time.UTC),
		UserID:   "55891ce6a3cb4bb0833514667d67288c",
		NicPhysicalPorts: []servers.NicPhysicalPort{
			{
				ID:                    "a2f63380-6c77-4cd5-8868-e3556ffd35ce",
				MACAddr:               "48:DF:37:90:B4:58",
				Plane:                 "DATA",
				NetworkPhysicalPortID: "d8e40a51-f1e2-4681-8953-9fe1e9992c42",
				HardwareID:            "be2d30d6-f891-4200-b827-95f229fb8c6b",
				AttachedPorts: []servers.Port{
					{
						NetworkID: "94055904-6b2c-4839-a14a-c61c93a8bc48",
						PortID:    "30fc1c27-fb5f-4955-94d0-a56cd28d09e8",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
								IPAddress: "2.1.1.10",
							},
						},
					},
					{
						NetworkID: "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
						PortID:    "aa6c61f4-db8a-44c7-a91c-7e636dac1dc6",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
								IPAddress: "169.254.0.9",
							},
						},
					},
				},
			},
			{
				ID:                    "b01dfdb0-f247-47d8-8224-c257aa3265e9",
				MACAddr:               "48:DF:37:90:B4:50",
				Plane:                 "STORAGE",
				NetworkPhysicalPortID: "00dfea92-5c5b-4860-aa05-efef6c2bb2af",
				HardwareID:            "be2d30d6-f891-4200-b827-95f229fb8c6b",
				AttachedPorts:         []servers.Port{},
			},
			{
				ID:                    "f4355e8e-39fc-48bd-a283-a2dbef8a2e32",
				MACAddr:               "48:DF:37:82:B0:A0",
				Plane:                 "STORAGE",
				NetworkPhysicalPortID: "cf798cc0-c869-45d5-a5a7-bcc578a300b0",
				HardwareID:            "84c74a86-7045-4284-80f9-0e7aff5d27ad",
				AttachedPorts:         []servers.Port{},
			},
			{
				ID:                    "5ef177fd-888c-4fae-9925-a8920beb07cb",
				MACAddr:               "48:DF:37:82:B0:A8",
				Plane:                 "DATA",
				NetworkPhysicalPortID: "2bbbb516-c75a-42b2-8a46-9cb5f26c219e",
				HardwareID:            "84c74a86-7045-4284-80f9-0e7aff5d27ad",
				AttachedPorts: []servers.Port{
					{
						NetworkID: "94055904-6b2c-4839-a14a-c61c93a8bc48",
						PortID:    "4e329a01-2cf4-4028-9259-03b7aa145cb6",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "acd41997-5ebb-4ff2-8cd2-22cae6cf2883",
								IPAddress: "2.1.1.20",
							},
						},
					},
					{
						NetworkID: "4a59f728-3920-4b71-ae54-d0d5c14ba04b",
						PortID:    "a256b4a1-3ae3-4102-a14e-987ae1610f97",
						FixedIPs: []servers.FixedIP{
							{
								SubnetID:  "b87d9c85-af5c-403d-a49a-55a6ab0a36d2",
								IPAddress: "169.254.0.10",
							},
						},
					},
				},
			},
		},
		ChassisStatus: servers.ChassisStatus{
			ChassisPower: true,
			PowerSupply:  true,
			CPU:          true,
			Memory:       true,
			Fan:          true,
			Disk:         0,
			Nic:          true,
			SystemBoard:  true,
			Etc:          true,
			Console:      true,
		},
		Links: []servers.Link{
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
				Rel:  "self",
			},
			{
				Href: "https://baremetal-server-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/24ebe7b8-ecfb-4d9f-a66b-c0120534fc90",
				Rel:  "bookmark",
			},
		},
		RaidArrays: []servers.RaidArray{
			{
				PrimaryStorage:     true,
				Partitions:         nil,
				RaidCardHardwareID: "bdfb75d1-194d-426d-b288-f588dfa5ac49",
				DiskHardwareIDs: []string{
					"76649053-863e-4533-86e3-f194a79485a6",
					"a25827e3-67da-47be-ba96-849ab4685a1d",
				},
			},
		},
		LvmVolumeGroups:          nil,
		Filesystems:              nil,
		MediaAttachments:         []servers.MediaAttachment{},
		ManagedByService:         "dedicated-hypervisor",
		ManagedServiceResourceID: "f42dbc37-4642-4628-8b47-50bf95d8fdd5",
	},
}

// ExpectedServersDetailsSlice is the slice of resources expected to be returned from ListDetailsResult.
var ExpectedServersDetailsSlice = []servers.Server{FirstServerDetail, SecondServerDetail}

var AddLicenseJob = servers.Job{
	JobID:  "b4f888dc2b9d4c41bb769cbd",
	Status: "COMPLETED",
	RequestedParam: servers.RequestedParam{
		VmName: "Alice",
		LicenseTypes: []string{
			"Windows Server",
			"SQL Server Standard 2014",
		},
	},
}

// HandleListServersSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that responds with a list of two servers.
func HandleListServersSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListResult)
	})
}

// HandleListServersDetailsSuccessfully creates an HTTP handler at `/servers/detail` on the
// test handler mux that responds with a list of two servers.
func HandleListServersDetailsSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/servers/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListDetailsResult)
	})
}

// HandleGetServerSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that responds with a single server.
func HandleGetServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s", SecondServer.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetResult)
	})
}

// HandleCreateServerSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server creation.
func HandleCreateServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateResponse)
	})
}

// HandleDeleteServerSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server deletion.
func HandleDeleteServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s", FirstServer.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
}

func HandleAddLicenseSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/action", SecondServer.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, AddLicenseRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, AddLicenseResponse)
	})
}

func HandleGetAddLicenseResultSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/action", SecondServer.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, GetAddLicenseResultRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetAddLicenseResultResponse)
	})
}
