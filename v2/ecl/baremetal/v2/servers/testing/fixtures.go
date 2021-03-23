package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud"
	"github.com/nttcom/eclcloud/ecl/baremetal/v2/servers"
)

var listResponse = fmt.Sprintf(`
{
	"servers": [
	  {
		"OS-EXT-STS:power_state": "RUNNING",
		"OS-EXT-STS:task_state": "None",
		"OS-EXT-STS:vm_state": "ACTIVE",
		"OS-EXT-AZ:availability_zone": "zone1-groupa",
		"created": "2012-09-07T16:56:37Z",
		"flavor": {
		  "id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
		  "links": [
			{
			  "href": "http://openstack.example.com/openstack/flavors/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			  "rel": "bookmark"
			}
		  ]
		},
		"id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
		"image": {
		  "id": "70a599e0-31e7-49b7-b260-868f441e862b",
		  "links": [
			{
			  "href": "http://openstack.example.com/openstack/images/70a599e0-31e7-49b7-b260-868f441e862b",
			  "rel": "bookmark"
			}
		  ]
		},
		"links": [
		  {
			"href": "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			"rel": "self"
		  },
		  {
			"href": "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			"rel": "bookmark"
		  }
		],
		"metadata": {
		  "My Server Name": "Apache1"
		},
		"name": "Test Server1",
		"progress": 0,
		"status": "ACTIVE",
		"tenant_id": "openstack",
		"updated": "2012-09-07T16:56:37Z",
		"user_id": "fake",
		"raid_arrays": [
		  {
			"primary_storage": true,
			"raid_card_hardware_id": "raid_card_uuid",
			"disk_hardware_ids": [
			  "disk0_uuid",
			  "disk1_uuid",
			  "disk2_uuid",
			  "disk3_uuid"
			],
			"partitions": [
			  {
				"lvm": true,
				"partition_label": "primary-part1"
			  },
			  {
				"lvm": false,
				"size": 100,
				"partition_label": "var"
			  }
			]
		  },
		  {
			"primary_storage": false,
			"raid_card_hardware_id": "raid_card_uuid",
			"internal_disk_ids": [
			  "disk4_uuid",
			  "disk5_uuid",
			  "disk6_uuid",
			  "disk7_uuid"
			],
			"raid_level": 10,
			"partitions": [
			  {
				"lvm": true,
				"partition_label": "secondary-part1"
			  }
			]
		  }
		],
		"lvm_volume_groups": [
		  {
			"vg_label": "VG_root",
			"physical_volume_partition_labels": [
			  "primary-part1",
			  "secondary-part1"
			],
			"logical_volumes": [
			  {
				"lv_label": "LV_root"
			  },
			  {
				"size": 2,
				"lv_label": "LV_swap"
			  }
			]
		  }
		],
		"filesystems": [
		  {
			"label": "LV_root",
			"mount_point": "/",
			"fs_type": "xfs"
		  },
		  {
			"label": "var",
			"mount_point": "/var",
			"fs_type": "xfs"
		  },
		  {
			"label": "LV_swap",
			"fs_type": "swap"
		  }
		],
		"nic_physical_ports": [
		  {
			"id": "39285bf9-12fb-4064-b98b-a552efc51cfc",
			"mac_addr": "0a:31:c1:d5:6d:9c",
			"network_physical_port_id": "38268d94-584a-4f14-96ff-732a68aa7301",
			"plane": "data",
			"attached_ports": [
			  {
				"port_id": "61b7da1e-9571-4d63-b779-e003a56b8105",
				"network_id": "9aa93722-1ec4-4912-b813-b975c21460a5",
				"fixed_ips": [
				  {
					"subnet_id": "0419bbde-2b82-4107-9d8a-6bba76e364af",
					"ip_address": "192.168.10.2"
				  }
				]
			  }
			],
			"hardware_id": "063468e8-61ab-4afd-be38-c937254aeb9a"
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
		  "etc": true
		},
		"media_attachments": [],
		"personality": [
		  {
			"path": "/home/big/banner.txt",
			"contents": "ZWNobyAiS3VtYSBQZXJzb25hbGl0eSIgPj4gL2hvbWUvYmlnL3BlcnNvbmFsaXR5"
		  }
		]
	  },
	  {
		"OS-EXT-STS:power_state": "RUNNING",
		"OS-EXT-STS:task_state": "None",
		"OS-EXT-STS:vm_state": "ACTIVE",
		"OS-EXT-AZ:availability_zone": "zone1-groupa",
		"created": "2012-09-07T16:56:37Z",
		"flavor": {
		  "id": "05184ba3-00ba-4fbc-b7a2-03b62b884932",
		  "links": [
			{
			  "href": "http://openstack.example.com/openstack/flavors/1",
			  "rel": "bookmark"
			}
		  ]
		},
		"hostId": "16d193736a5cfdb60c697ca27ad071d6126fa13baeb670fc9d10645e",
		"id": "05184ba3-00ba-4fbc-b7a2-03b62b884932",
		"image": {
		  "id": "70a599e0-31e7-49b7-b260-868f441e862b",
		  "links": [
			{
			  "href": "http://openstack.example.com/openstack/images/70a599e0-31e7-49b7-b260-868f441e862b",
			  "rel": "bookmark"
			}
		  ]
		},
		"links": [
		  {
			"href": "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			"rel": "self"
		  },
		  {
			"href": "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			"rel": "bookmark"
		  }
		],
		"metadata": {
		  "My Server Name": "Apache1"
		},
		"name": "Test Server2",
		"progress": 0,
		"status": "ACTIVE",
		"tenant_id": "openstack",
		"updated": "2012-09-07T16:56:37Z",
		"user_id": "fake",
		"raid_arrays": [
		  {
			"primary_storage": true,
			"raid_card_hardware_id": "raid_card_uuid",
			"disk_hardware_ids": [
			  "disk0_uuid",
			  "disk1_uuid",
			  "disk2_uuid",
			  "disk3_uuid"
			],
			"partitions": [
			  {
				"lvm": true,
				"partition_label": "primary-part1"
			  },
			  {
				"lvm": false,
				"size": 100,
				"partition_label": "var"
			  }
			]
		  },
		  {
			"primary_storage": false,
			"raid_card_hardware_id": "raid_card_uuid",
			"internal_disk_ids": [
			  "disk4_uuid",
			  "disk5_uuid",
			  "disk6_uuid",
			  "disk7_uuid"
			],
			"raid_level": 10,
			"partitions": [
			  {
				"lvm": true,
				"partition_label": "secondary-part1"
			  }
			]
		  }
		],
		"lvm_volume_groups": [
		  {
			"vg_label": "VG_root",
			"physical_volume_partition_labels": [
			  "primary-part1",
			  "secondary-part1"
			],
			"logical_volumes": [
			  {
				"lv_label": "LV_root"
			  },
			  {
				"size": 2,
				"lv_label": "LV_swap"
			  }
			]
		  }
		],
		"filesystems": [
		  {
			"label": "LV_root",
			"mount_point": "/",
			"fs_type": "xfs"
		  },
		  {
			"label": "var",
			"mount_point": "/var",
			"fs_type": "xfs"
		  },
		  {
			"label": "LV_swap",
			"fs_type": "swap"
		  }
		],
		"nic_physical_ports": [
		  {
			"id": "f4732cd9-31f7-408e-9f27-cc9b0ee17457",
			"mac_addr": "0a:31:c1:d5:6d:9d",
			"network_physical_port_id": "ab17a82d-e9a5-4e95-9b18-de3f8a47670f",
			"plane": "storage",
			"attached_ports": [
			  {
				"port_id": "6fb0d979-f05b-466c-b50c-64d5ae4c4ef6",
				"network_id": "99babdfc-79eb-470a-b0d4-df02482cc509",
				"fixed_ips": [
				  {
					"subnet_id": "9632ce5d-8750-40bf-871d-968aa3324367",
					"ip_address": "192.168.10.8"
				  }
				]
			  }
			],
			"hardware_id": "ab36f541-b854-46c3-8891-e9484a1ba1ac"
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
		  "etc": true
		},
		"media_attachments": [
		  {
			"image": {
			  "id": "3339fd5f-ec06-4ef8-9337-c1c70218a748",
			  "links": [
				{
				  "href": "http://openstack.example.com/openstack/images/3339fd5f-ec06-4ef8-9337-c1c70218a748",
				  "rel": "bookmark"
				}
			  ]
			}
		  }
		]
	  }
	]
}`)

var getResponse = fmt.Sprintf(`
{
	"server": {
	  "OS-EXT-STS:power_state": "RUNNING",
	  "OS-EXT-STS:task_state": "None",
	  "OS-EXT-STS:vm_state": "ACTIVE",
	  "OS-EXT-AZ:availability_zone": "zone1-groupa",
	  "created": "2012-09-07T16:56:37Z",
	  "flavor": {
		"id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
		"links": [
		  {
			"href": "http://openstack.example.com/openstack/flavors/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			"rel": "bookmark"
		  }
		]
	  },
	  "hostId": "16d193736a5cfdb60c697ca27ad071d6126fa13baeb670fc9d10645e",
	  "id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
	  "image": {
		"id": "70a599e0-31e7-49b7-b260-868f441e862b",
		"links": [
		  {
			"href": "http://openstack.example.com/openstack/images/70a599e0-31e7-49b7-b260-868f441e862b",
			"rel": "bookmark"
		  }
		]
	  },
	  "links": [
		{
		  "href": "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
		  "rel": "self"
		},
		{
		  "href": "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
		  "rel": "bookmark"
		}
	  ],
	  "metadata": {
		"My Server Name": "Apache1"
	  },
	  "name": "Test Server1",
	  "progress": 0,
	  "status": "ACTIVE",
	  "tenant_id": "openstack",
	  "updated": "2012-09-07T16:56:37Z",
	  "user_id": "fake",
	  "raid_arrays": [
		{
		  "primary_storage": true,
		  "raid_card_hardware_id": "raid_card_uuid",
		  "disk_hardware_ids": [
			"disk0_uuid",
			"disk1_uuid",
			"disk2_uuid",
			"disk3_uuid"
		  ],
		  "partitions": [
			{
			  "lvm": true,
			  "partition_label": "primary-part1"
			},
			{
			  "lvm": false,
			  "size": 100,
			  "partition_label": "var"
			}
		  ]
		},
		{
		  "primary_storage": false,
		  "raid_card_hardware_id": "raid_card_uuid",
		  "internal_disk_ids": [
			"disk4_uuid",
			"disk5_uuid",
			"disk6_uuid",
			"disk7_uuid"
		  ],
		  "raid_level": 10,
		  "partitions": [
			{
			  "lvm": true,
			  "partition_label": "secondary-part1"
			}
		  ]
		}
	  ],
	  "lvm_volume_groups": [
		{
		  "vg_label": "VG_root",
		  "physical_volume_partition_labels": [
			"primary-part1",
			"secondary-part1"
		  ],
		  "logical_volumes": [
			{
			  "lv_label": "LV_root"
			},
			{
			  "size": 2,
			  "lv_label": "LV_swap"
			}
		  ]
		}
	  ],
	  "filesystems": [
		{
		  "label": "LV_root",
		  "mount_point": "/",
		  "fs_type": "xfs"
		},
		{
		  "label": "var",
		  "mount_point": "/var",
		  "fs_type": "xfs"
		},
		{
		  "label": "LV_swap",
		  "fs_type": "swap"
		}
	  ],
	  "nic_physical_ports": [
		{
		  "id": "39285bf9-12fb-4064-b98b-a552efc51cfc",
		  "mac_addr": "0a:31:c1:d5:6d:9c",
		  "network_physical_port_id": "38268d94-584a-4f14-96ff-732a68aa7301",
		  "plane": "data",
		  "attached_ports": [
			{
			  "port_id": "61b7da1e-9571-4d63-b779-e003a56b8105",
			  "network_id": "9aa93722-1ec4-4912-b813-b975c21460a5",
			  "fixed_ips": [
				{
				  "subnet_id": "0419bbde-2b82-4107-9d8a-6bba76e364af",
				  "ip_address": "192.168.10.2"
				}
			  ]
			}
		  ],
		  "hardware_id": "063468e8-61ab-4afd-be38-c937254aeb9a"
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
		"etc": true
	  },
	  "media_attachments": [
		{
		  "image": {
			"id": "3339fd5f-ec06-4ef8-9337-c1c70218a748",
			"links": [
			  {
				"href": "http://openstack.example.com/openstack/images/3339fd5f-ec06-4ef8-9337-c1c70218a748",
				"rel": "bookmark"
			  }
			]
		  }
		}
	  ],
	  "personality": [
		{
		  "path": "/home/big/banner.txt",
		  "contents": "ZWNobyAiS3VtYSBQZXJzb25hbGl0eSIgPj4gL2hvbWUvYmlnL3BlcnNvbmFsaXR5"
		}
	  ]
	}
}`)

var createRequest = fmt.Sprintf(`
  {
    "server": {
        "name": "server-test-1",
        "adminPass": "aabbccddeeff",
        "imageRef": "b5660a6e-4b46-4be3-9707-6b47221b454f",
        "flavorRef": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
        "availability_zone": "zone1-groupa",
        "networks": [
            {
                "uuid": "d32019d3-bc6e-4319-9c1d-6722fc136a22",
                "fixed_ip": "10.0.0.100"
            }
        ],
        "raid_arrays": [
            {
                "primary_storage": true,
                "partitions": [
                    {
                        "lvm": true,
                        "partition_label": "primary-part1"
                    },
                    {
                        "size": "100G",
                        "partition_label": "var"
                    }
                ]
            },
            {
                "raid_card_hardware_id": "raid_card_uuid",
                "disk_hardware_ids": [
                    "disk1_uuid", "disk2_uuid", "disk3_uuid", "disk4_uuid"
                ],
                "partitions": [
                    {
                        "lvm": true,
                        "partition_label": "secondary-part1"
                    }
                ],
                "raid_level": 10
            }
        ],
        "lvm_volume_groups": [
            {
                "vg_label": "VG_root",
                "physical_volume_partition_labels": [
                    "primary-part1", "secondary-part1"
                ],
                "logical_volumes": [
                    {
                        "size": "300G",
                        "lv_label": "LV_root"
                    },
                    {
                        "size": "2G",
                        "lv_label": "LV_swap"
                    }
                ]
            }
        ],
        "filesystems": [
            {
                "label": "LV_root",
                "mount_point": "/",
                "fs_type": "xfs"
            },
            {
                "label": "var",
                "mount_point": "/var",
                "fs_type": "xfs"
            },
            {
                "label": "LV_swap",
                "fs_type": "swap"
            }
        ],
        "user_data": "dXNlcl9kYXRh",
        "metadata": {
            "foo": "bar"
        }
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "server": {
        "id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
        "links": [
            {
                "href": "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                "rel": "self"
            },
            {
                "href": "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
                "rel": "bookmark"
            }
        ],
        "adminPass": "aabbccddeeff"
    }
}`)

var expectedServers = []servers.Server{expectedServer1, expectedServer2}

var expectedCreated, _ = time.Parse(eclcloud.RFC3339Milli, "2012-09-07T16:56:37Z")
var expectedUpdated, _ = time.Parse(eclcloud.RFC3339Milli, "2012-09-07T16:56:37Z")

var expectedServer1 = servers.Server{
	ID:               "05184ba3-00ba-4fbc-b7a2-03b62b884931",
	TenantID:         "openstack",
	UserID:           "fake",
	Name:             "Test Server1",
	Updated:          expectedUpdated,
	Created:          expectedCreated,
	Status:           "ACTIVE",
	PowerState:       "RUNNING",
	TaskState:        "None",
	VMState:          "ACTIVE",
	AvailabilityZone: "zone1-groupa",
	Progress:         0,
	Image: map[string]interface{}{
		"id": "70a599e0-31e7-49b7-b260-868f441e862b",
		"links": []map[string]interface{}{
			{
				"href": "http://openstack.example.com/openstack/images/70a599e0-31e7-49b7-b260-868f441e862b",
				"rel":  "bookmark",
			},
		},
	},
	Flavor: map[string]interface{}{
		"id": "05184ba3-00ba-4fbc-b7a2-03b62b884931",
		"links": []map[string]interface{}{
			{
				"href": "http://openstack.example.com/openstack/flavors/05184ba3-00ba-4fbc-b7a2-03b62b884931",
				"rel":  "bookmark",
			},
		},
	},
	Metadata: map[string]string{
		"My Server Name": "Apache1",
	},
	Links: []eclcloud.Link{
		{
			Href: "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			Rel:  "self",
		},
		{
			Href: "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			Rel:  "bookmark",
		},
	},
	RaidArrays: []servers.RaidArray{
		{
			PrimaryStorage:     true,
			RaidCardHardwareID: "raid_card_uuid",
			DiskHardwareIDs: []string{
				"disk0_uuid",
				"disk1_uuid",
				"disk2_uuid",
				"disk3_uuid",
			},
			Partitions: []servers.Partition{
				{
					LVM:            true,
					PartitionLabel: "primary-part1",
				},
				{
					LVM:            false,
					Size:           100,
					PartitionLabel: "var",
				},
			},
		},
	},
	LVMVolumeGroups: []servers.LVMVolumeGroup{
		{
			VGLabel: "VG_root",
			PhysicalVolumePartitionLabels: []string{
				"primary-part1",
				"secondary-part1",
			},
			LogicalVolumes: []servers.LogicalVolume{
				{
					LVLabel: "LV_root",
				},
				{
					Size:    2,
					LVLabel: "LV_swap",
				},
			},
		},
	},
	Filesystems: []servers.Filesystem{
		{
			Label:      "LV_root",
			FSType:     "xfs",
			MountPoint: "/",
		},
		{
			Label:      "var",
			FSType:     "xfs",
			MountPoint: "/var",
		},
		{
			Label:  "LV_swap",
			FSType: "swap",
		},
	},
	NICPhysicalPorts: []servers.NICPhysicalPort{
		{
			ID:                    "39285bf9-12fb-4064-b98b-a552efc51cfc",
			MacAddr:               "0a:31:c1:d5:6d:9c",
			NetworkPhysicalPortID: "38268d94-584a-4f14-96ff-732a68aa7301",
			Plane:                 "data",
			AttachedPorts: []servers.AttachedPort{
				{
					PortID:    "61b7da1e-9571-4d63-b779-e003a56b8105",
					NetworkID: "9aa93722-1ec4-4912-b813-b975c21460a5",
					FixedIPs: []servers.FixedIP{
						{
							SubnetID:  "0419bbde-2b82-4107-9d8a-6bba76e364af",
							IPAddress: "192.168.10.2",
						},
					},
				},
			},
			HardwareID: "063468e8-61ab-4afd-be38-c937254aeb9a",
		},
	},
	ChassisStatus: servers.ChassisStatus{
		ChassisPower: true,
		PowerSupply:  true,
		CPU:          true,
		Memory:       true,
		Fan:          true,
		Disk:         0,
		NIC:          true,
		SystemBoard:  true,
		Etc:          true,
	},
	MediaAttachments: []map[string]interface{}{},
	Personality: []servers.Personality{
		{
			Path:     "/home/big/banner.txt",
			Contents: "ZWNobyAiS3VtYSBQZXJzb25hbGl0eSIgPj4gL2hvbWUvYmlnL3BlcnNvbmFsaXR5",
		},
	},
}

var expectedServer2 = servers.Server{
	ID:               "05184ba3-00ba-4fbc-b7a2-03b62b884932",
	TenantID:         "openstack",
	UserID:           "fake",
	Name:             "Test Server2",
	Updated:          expectedUpdated,
	Created:          expectedCreated,
	Status:           "ACTIVE",
	PowerState:       "RUNNING",
	TaskState:        "None",
	VMState:          "ACTIVE",
	AvailabilityZone: "zone1-groupa",
	Progress:         0,
	Image: map[string]interface{}{
		"id": "70a599e0-31e7-49b7-b260-868f441e862b",
		"links": []map[string]interface{}{
			{
				"href": "http://openstack.example.com/openstack/images/70a599e0-31e7-49b7-b260-868f441e862b",
				"rel":  "bookmark",
			},
		},
	},
	Flavor: map[string]interface{}{
		"id": "05184ba3-00ba-4fbc-b7a2-03b62b884932",
		"links": []map[string]interface{}{
			{
				"href": "http://openstack.example.com/openstack/flavors/1",
				"rel":  "bookmark",
			},
		},
	},
	Metadata: map[string]string{
		"My Server Name": "Apache1",
	},
	Links: []eclcloud.Link{
		{
			Href: "http://openstack.example.com/v2/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			Rel:  "self",
		},
		{
			Href: "http://openstack.example.com/openstack/servers/05184ba3-00ba-4fbc-b7a2-03b62b884931",
			Rel:  "bookmark",
		},
	},
	RaidArrays: []servers.RaidArray{
		{
			PrimaryStorage:     true,
			RaidCardHardwareID: "raid_card_uuid",
			DiskHardwareIDs: []string{
				"disk0_uuid",
				"disk1_uuid",
				"disk2_uuid",
				"disk3_uuid",
			},
			Partitions: []servers.Partition{
				{
					LVM:            true,
					PartitionLabel: "primary-part1",
				},
				{
					LVM:            false,
					Size:           100,
					PartitionLabel: "var",
				},
			},
		},
	},
	LVMVolumeGroups: []servers.LVMVolumeGroup{
		{
			VGLabel: "VG_root",
			PhysicalVolumePartitionLabels: []string{
				"primary-part1",
				"secondary-part1",
			},
			LogicalVolumes: []servers.LogicalVolume{
				{
					LVLabel: "LV_root",
				},
				{
					Size:    2,
					LVLabel: "LV_swap",
				},
			},
		},
	},
	Filesystems: []servers.Filesystem{
		{
			Label:      "LV_root",
			FSType:     "xfs",
			MountPoint: "/",
		},
		{
			Label:      "var",
			FSType:     "xfs",
			MountPoint: "/var",
		},
		{
			Label:  "LV_swap",
			FSType: "swap",
		},
	},
	NICPhysicalPorts: []servers.NICPhysicalPort{
		{
			ID:                    "f4732cd9-31f7-408e-9f27-cc9b0ee17457",
			MacAddr:               "0a:31:c1:d5:6d:9d",
			NetworkPhysicalPortID: "ab17a82d-e9a5-4e95-9b18-de3f8a47670f",
			Plane:                 "storage",
			AttachedPorts: []servers.AttachedPort{
				{
					PortID:    "6fb0d979-f05b-466c-b50c-64d5ae4c4ef6",
					NetworkID: "99babdfc-79eb-470a-b0d4-df02482cc509",
					FixedIPs: []servers.FixedIP{
						{
							SubnetID:  "9632ce5d-8750-40bf-871d-968aa3324367",
							IPAddress: "192.168.10.8",
						},
					},
				},
			},
			HardwareID: "ab36f541-b854-46c3-8891-e9484a1ba1ac",
		},
	},
	ChassisStatus: servers.ChassisStatus{
		ChassisPower: true,
		PowerSupply:  true,
		CPU:          true,
		Memory:       true,
		Fan:          true,
		Disk:         0,
		NIC:          true,
		SystemBoard:  true,
		Etc:          true,
	},
	MediaAttachments: []map[string]interface{}{},
	Personality:      []servers.Personality(nil),
}
