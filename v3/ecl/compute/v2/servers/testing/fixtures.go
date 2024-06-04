package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/compute/v2/servers"

	"github.com/nttcom/eclcloud/v3"
	th "github.com/nttcom/eclcloud/v3/testhelper"
	"github.com/nttcom/eclcloud/v3/testhelper/client"

	"time"
)

// ListResult provides a single page of Server results.
const ListResult = `
{
	"servers": [
		{
			"id": "707dbd55-b6bf-439d-804c-3002f49ac898",
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/707dbd55-b6bf-439d-804c-3002f49ac898",
					"rel": "self"
				},
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/707dbd55-b6bf-439d-804c-3002f49ac898",
					"rel": "bookmark"
				}
			],
			"name": "Test Server2"
		},
		{
			"id": "8e69a092-53f9-4225-bae6-57cfbb5d6857",
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
					"rel": "self"
				},
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
					"rel": "bookmark"
				}
			],
			"name": "Test Server1"
		}
	]
}
`

// ListDetailsResult provides a single page of Server results in details.
const ListDetailsResult = `
{
	"servers": [
		{
			"status": "ACTIVE",
			"updated": "2020-05-18T01:51:41Z",
			"hostId": "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
			"addresses": {
				"IF-4831": [
					{
						"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:f3:ed:05",
						"version": 4,
						"addr": "192.168.1.103",
						"OS-EXT-IPS:type": "fixed"
					}
				]
			},
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/707dbd55-b6bf-439d-804c-3002f49ac898",
					"rel": "self"
				},
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/707dbd55-b6bf-439d-804c-3002f49ac898",
					"rel": "bookmark"
				}
			],
			"key_name": null,
			"image": {
				"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
				"links": [
					{
						"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
						"rel": "bookmark"
					}
				]
			},
			"OS-EXT-STS:task_state": null,
			"OS-EXT-STS:vm_state": "active",
			"OS-SRV-USG:launched_at": "2020-05-11T06:25:56.000000",
			"flavor": {
				"id": "1CPU-4GB",
				"links": [
					{
						"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
						"rel": "bookmark"
					}
				]
			},
			"id": "707dbd55-b6bf-439d-804c-3002f49ac898",
			"OS-SRV-USG:terminated_at": null,
			"OS-EXT-AZ:availability_zone": "zone1_groupa",
			"user_id": "5e86848fbc63403daaeffc1b76b3a784",
			"name": "Test Server2",
			"created": "2020-05-11T06:25:53Z",
			"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
			"OS-DCF:diskConfig": "MANUAL",
			"os-extended-volumes:volumes_attached": [],
			"accessIPv4": "",
			"accessIPv6": "",
			"progress": 0,
			"OS-EXT-STS:power_state": 1,
			"config_drive": "True",
			"metadata": {
				"vmha": "false",
				"HA_Enabled": "false"
			}
		},
		{
			"status": "ACTIVE",
			"updated": "2020-05-18T01:51:41Z",
			"hostId": "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
			"addresses": {
				"IF-4831": [
					{
						"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:49:78:28",
						"version": 4,
						"addr": "192.168.1.101",
						"OS-EXT-IPS:type": "fixed"
					}
				]
			},
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
					"rel": "self"
				},
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
					"rel": "bookmark"
				}
			],
			"key_name": null,
			"image": {
				"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
				"links": [
					{
						"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
						"rel": "bookmark"
					}
				]
			},
			"OS-EXT-STS:task_state": null,
			"OS-EXT-STS:vm_state": "active",
			"OS-SRV-USG:launched_at": "2020-05-11T03:36:27.000000",
			"flavor": {
				"id": "1CPU-4GB",
				"links": [
					{
						"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
						"rel": "bookmark"
					}
				]
			},
			"id": "8e69a092-53f9-4225-bae6-57cfbb5d6857",
			"OS-SRV-USG:terminated_at": null,
			"OS-EXT-AZ:availability_zone": "zone1_groupa",
			"user_id": "5e86848fbc63403daaeffc1b76b3a784",
			"name": "Test Server1",
			"created": "2020-05-11T06:25:53Z",
			"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
			"OS-DCF:diskConfig": "MANUAL",
			"os-extended-volumes:volumes_attached": [],
			"accessIPv4": "",
			"accessIPv6": "",
			"progress": 0,
			"OS-EXT-STS:power_state": 1,
			"config_drive": "",
			"metadata": {
				"vmha": "false",
				"HA_Enabled": "false"
			}
		}
	]
}
`

// GetResult provides a Get result.
const GetResult = `
{
	"server": {
		"status": "ACTIVE",
		"updated": "2020-05-18T01:51:41Z",
		"hostId": "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
		"addresses": {
			"IF-4831": [
				{
					"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:49:78:28",
					"version": 4,
					"addr": "192.168.1.101",
					"OS-EXT-IPS:type": "fixed"
				}
			]
		},
		"links": [
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
				"rel": "self"
			},
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
				"rel": "bookmark"
			}
		],
		"key_name": null,
		"image": {
			"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
					"rel": "bookmark"
				}
			]
		},
		"OS-EXT-STS:task_state": null,
		"OS-EXT-STS:vm_state": "active",
		"OS-SRV-USG:launched_at": "2020-05-11T03:36:27.000000",
		"flavor": {
			"id": "1CPU-4GB",
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
					"rel": "bookmark"
				}
			]
		},
		"id": "8e69a092-53f9-4225-bae6-57cfbb5d6857",
		"OS-SRV-USG:terminated_at": null,
		"OS-EXT-AZ:availability_zone": "zone1_groupa",
		"user_id": "5e86848fbc63403daaeffc1b76b3a784",
		"name": "Test Server1",
		"created": "2020-05-11T06:25:53Z",
		"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
		"OS-DCF:diskConfig": "MANUAL",
		"os-extended-volumes:volumes_attached": [],
		"accessIPv4": "",
		"accessIPv6": "",
		"progress": 0,
		"OS-EXT-STS:power_state": 1,
		"config_drive": "",
		"metadata": {
			"vmha": "false",
			"HA_Enabled": "false"
		}
	}
}
`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
	"server": {
		"flavorRef": "1CPU-4GB",
		"imageRef": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
		"name": "Test Server1",
		"availability_zone": "zone1-groupa",
		"config_drive": true,
		"user_data": "dXNlcl9kYXRh",
		"metadata": {
			"foo": "bar"
		},
		"networks": [
			{
				"uuid": "4d98b876-b5d1-4861-8650-b5a53024486a"
			}
		]
	}
}
`

const CreateResponse = `
{
	"server": {
		"security_groups": [
			{
				"name": "default"
			}
		],
		"OS-DCF:diskConfig": "MANUAL",
		"id": "8e69a092-53f9-4225-bae6-57cfbb5d6857",
		"links": [
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
				"rel": "self"
			},
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
				"rel": "bookmark"
			}
		],
		"adminPass": "aabbccddeeff"
	}
}
`

const UpdateRequest = `
{
	"server": {
		"name": "Update Name"
	}
}
`

const UpdateResponse = `
{
	"server": {
		"status": "ACTIVE",
		"updated": "2020-05-18T01:51:41Z",
		"hostId": "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
		"addresses": {
			"IF-4831": [
				{
					"version": 4,
					"addr": "192.168.1.101"
				}
			]
		},
		"links": [
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
				"rel": "self"
			},
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
				"rel": "bookmark"
			}
		],
		"image": {
			"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
					"rel": "bookmark"
				}
			]
		},
		"flavor": {
			"id": "1CPU-4GB",
			"links": [
				{
					"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
					"rel": "bookmark"
				}
			]
		},
		"id": "8e69a092-53f9-4225-bae6-57cfbb5d6857",
		"user_id": "5e86848fbc63403daaeffc1b76b3a784",
		"name": "Update Name",
		"created": "2020-05-11T06:25:53Z",
		"tenant_id": "1bc271e7a8af4d988ff91612f5b122f8",
		"OS-DCF:diskConfig": "MANUAL",
		"accessIPv4": "",
		"accessIPv6": "",
		"progress": 0,
		"metadata": {
			"vmha": "false",
			"HA_Enabled": "false"
		}
	}
}
`

var MetadataResult = `
{
	"metadata": {
		"vmha": "false",
		"HA_Enabled": "false"
	}
}
`

var MetadatumResult = `
{
	"meta": {
		"vmha": "false"
	}
}
`

var CreateMetadatumRequest = `
{
	"meta": {
		"key1": "val1"
	}
}
`

var CreateMetadatumResponse = CreateMetadatumRequest

var UpdateMetadataRequest = `
{
	"metadata": {
		"key1": "update_val"
	}
}
`

var UpdateMetadataResponse = UpdateMetadataRequest

var ResetMetadataRequest = `
{
	"metadata": {
		"key1": "val1",
		"key2": "val2"
	}
}
`

var ResetMetadataResponse = ResetMetadataRequest

var ResizeRequest = `
{
	"resize": {
		"flavorRef": "2CPU-8GB"
	}
}
`

var CreateImageRequest = `
{
	"createImage": {
		"metadata": {
			"key": "create_image"
		},
		"name": "Test Create Image"
	}
}
`

var expectedServers = []servers.Server{expectedServer1, expectedServer2}

var expectedCreated, _ = time.Parse(eclcloud.RFC3339Milli, "2020-05-11T06:25:53Z")
var expectedUpdated, _ = time.Parse(eclcloud.RFC3339Milli, "2020-05-18T01:51:41Z")

var expectedServer1 = servers.Server{
	ID:         "707dbd55-b6bf-439d-804c-3002f49ac898",
	TenantID:   "1bc271e7a8af4d988ff91612f5b122f8",
	UserID:     "5e86848fbc63403daaeffc1b76b3a784",
	Name:       "Test Server2",
	Updated:    expectedUpdated,
	Created:    expectedCreated,
	HostID:     "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
	Status:     "ACTIVE",
	Progress:   0,
	AccessIPv4: "",
	Image: map[string]interface{}{
		"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
		"links": []map[string]interface{}{
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
				"rel":  "bookmark",
			},
		},
	},
	Flavor: map[string]interface{}{
		"id": "1CPU-4GB",
		"links": []map[string]interface{}{
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
				"rel":  "bookmark",
			},
		},
	},
	Addresses: map[string]interface{}{
		"IF-4831": []interface{}{
			map[string]interface{}{
				"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:f3:ed:05",
				"OS-EXT-IPS:type":         "fixed",
				"addr":                    "192.168.1.103",
				"version":                 float64(4),
			},
		},
	},
	Metadata: map[string]string{
		"vmha":       "false",
		"HA_Enabled": "false",
	},
	Links: []interface{}{
		map[string]interface{}{
			"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/707dbd55-b6bf-439d-804c-3002f49ac898",
			"rel":  "self",
		},
	},
	AdminPass:      "",
	SecurityGroups: nil,
	Fault:          servers.Fault{},
	ConfigDrive:    "True",
}

var expectedServer2 = servers.Server{
	ID:         "8e69a092-53f9-4225-bae6-57cfbb5d6857",
	TenantID:   "1bc271e7a8af4d988ff91612f5b122f8",
	UserID:     "5e86848fbc63403daaeffc1b76b3a784",
	Name:       "Test Server1",
	Updated:    expectedUpdated,
	Created:    expectedCreated,
	HostID:     "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
	Status:     "ACTIVE",
	Progress:   0,
	AccessIPv4: "",
	Image: map[string]interface{}{
		"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
		"links": []map[string]interface{}{
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
				"rel":  "bookmark",
			},
		},
	},
	Flavor: map[string]interface{}{
		"id": "1CPU-4GB",
		"links": []map[string]interface{}{
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
				"rel":  "bookmark",
			},
		},
	},
	Addresses: map[string]interface{}{
		"IF-4831": []interface{}{
			map[string]interface{}{
				"OS-EXT-IPS-MAC:mac_addr": "fa:16:3e:49:78:28",
				"OS-EXT-IPS:type":         "fixed",
				"addr":                    "192.168.1.101",
				"version":                 float64(4),
			},
		},
	},
	Metadata: map[string]string{
		"vmha":       "false",
		"HA_Enabled": "false",
	},
	Links: []interface{}{
		map[string]interface{}{
			"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
			"rel":  "self",
		},
	},
	//KeyName:        nil,
	AdminPass:      "",
	SecurityGroups: nil,
	Fault:          servers.Fault{},
	ConfigDrive:    "",
}

var serverNameUpdated = servers.Server{
	ID:         "8e69a092-53f9-4225-bae6-57cfbb5d6857",
	TenantID:   "1bc271e7a8af4d988ff91612f5b122f8",
	UserID:     "5e86848fbc63403daaeffc1b76b3a784",
	Name:       "Update Name",
	Updated:    expectedUpdated,
	Created:    expectedCreated,
	HostID:     "d7961f8a2cde3e49a3f5d3a0a95c6c1d9ce28a342285d4118a936247",
	Status:     "ACTIVE",
	Progress:   0,
	AccessIPv4: "",
	Image: map[string]interface{}{
		"id": "c11a6d55-70e9-4d04-a086-4451f07da0d7",
		"links": []map[string]interface{}{
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/images/c11a6d55-70e9-4d04-a086-4451f07da0d7",
				"rel":  "bookmark",
			},
		},
	},
	Flavor: map[string]interface{}{
		"id": "1CPU-4GB",
		"links": []map[string]interface{}{
			{
				"href": "https://nova-jp1-ecl.api.ntt.com/1bc271e7a8af4d988ff91612f5b122f8/flavors/1CPU-4GB",
				"rel":  "bookmark",
			},
		},
	},
	Addresses: map[string]interface{}{
		"IF-4831": []interface{}{
			map[string]interface{}{
				"addr":    "192.168.1.101",
				"version": float64(4),
			},
		},
	},
	Metadata: map[string]string{
		"vmha":       "false",
		"HA_Enabled": "false",
	},
	Links: []interface{}{
		map[string]interface{}{
			"href": "https://nova-jp1-ecl.api.ntt.com/v2/1bc271e7a8af4d988ff91612f5b122f8/servers/8e69a092-53f9-4225-bae6-57cfbb5d6857",
			"rel":  "self",
		},
	},
	AdminPass:      "",
	SecurityGroups: nil,
	Fault:          servers.Fault{},
	ConfigDrive:    "",
}

var expectMetadata = map[string]string{
	"vmha":       "false",
	"HA_Enabled": "false",
}

var expectMetadatum = map[string]string{
	"vmha": "false",
}

var expectCreateMetadatum = map[string]string{
	"key1": "val1",
}

var ecpectUpdateMetadata = map[string]string{
	"key1": "update_val",
}

var expectResetMetadata = map[string]string{
	"key1": "val1",
	"key2": "val2",
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
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
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
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, CreateResponse)
	})
}

// HandleDeleteServerSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server deletion.
func HandleDeleteServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s", expectedServer1.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleUpdateServerSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server update.
func HandleUpdateServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateResponse)
	})
}

// HandleGetMetadataSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that responds with a server metadata.
func HandleGetMetadataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/metadata", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, MetadataResult)
	})
}

// HandleGetMetadatumSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that responds with a server metadatum.
func HandleGetMetadatumSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/metadata/vmha", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, MetadatumResult)
	})
}

// HandleCreateMetadatumSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server metadata creation.
func HandleCreateMetadatumSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/metadata/key1", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateMetadatumRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateMetadatumResponse)
	})
}

// HandleDeleteMetadatumSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server metadata deletion.
func HandleDeleteMetadatumSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/metadata/vmha", expectedServer1.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleUpdateMetadataSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server metadata update.
func HandleUpdateMetadataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/metadata", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestJSONRequest(t, r, UpdateMetadataRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateMetadataResponse)
	})
}

// HandleResetMetadataSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server metadata reset.
func HandleResetMetadataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/metadata", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, ResetMetadataRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ResetMetadataResponse)
	})
}

// HandleResizeServerSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests server resize action.
func HandleResizeServerSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/action", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, ResizeRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})
}

// HandleCreateImageSuccessfully creates an HTTP handler at `/servers` on the
// test handler mux that tests create server image.
func HandleCreateImageSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/servers/%s/action", expectedServer2.ID), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateImageRequest)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
	})
}
