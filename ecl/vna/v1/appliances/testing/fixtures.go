package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/ecl/vna/v1/appliances"
)

const applianceType = "ECL::VirtualNetworkAppliance::VSRX"
const idAppliance1 = "45db3e66-31af-45a6-8ad2-d01521726141"
const idAppliance2 = "45db3e66-31af-45a6-8ad2-d01521726142"

const idVirtualNetworkAppliancePlan = "6589b37a-cf82-4918-96fe-255683f78e76"

var listResponse = fmt.Sprintf(`
{
    "virtual_network_appliances": [
        {
            "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
            "availability_zone": "zone1-groupb",
            "default_gateway": "192.168.1.1",
            "description": "appliance_1_description",
            "id": "%s",
            "interfaces": {
                "interface_1": {
                    "allowed_address_pairs": [
                        {
                            "ip_address": "1.1.1.1",
                            "mac_address": "aa:bb:cc:dd:ee:f1",
                            "type": "vrrp",
                            "vrid": 123
                        }
                    ],
                    "description": "interface_1_description",
                    "fixed_ips": [
                        {
                            "ip_address": "192.168.1.51",
                            "subnet_id": "dummySubnetID"
                        }
                    ],
                    "name": "interface_1",
                    "network_id": "dummyNetworkID",
                    "tags": {},
                    "updatable": true
                },
                "interface_2": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_3": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_4": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_5": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_6": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_7": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_8": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                }
            },
            "name": "appliance_1",
            "operation_status": "COMPLETE",
            "os_login_status": "ACTIVE",
            "os_monitoring_status": "ACTIVE",
            "password": "Passw0rd",
            "tags": {
                "k1": "v1"
            },
            "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
            "username": "root",
            "virtual_network_appliance_plan_id": "%s",
            "vm_status": "ACTIVE"
        },
        {
            "appliance_type": "ECL::VirtualNetworkAppliance::VSRX",
            "availability_zone": "zone1-groupb",
            "default_gateway": "192.168.1.1",
            "description": "appliance_2_description",
            "id": "%s",
            "interfaces": {
                "interface_1": {
                    "allowed_address_pairs": [
                        {
                            "ip_address": "2.2.2.2",
                            "mac_address": "aa:bb:cc:dd:ee:f2",
                            "type": "",
                            "vrid": null
                        }
                    ],
                    "description": "interface_1_description",
                    "fixed_ips": [
                        {
                            "ip_address": "192.168.1.52",
                            "subnet_id": "dummySubnetID"
                        }
                    ],
                    "name": "interface_1",
                    "network_id": "dummyNetworkID",
                    "tags": {},
                    "updatable": true
                },
                "interface_2": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_3": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_4": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_5": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_6": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_7": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                },
                "interface_8": {
                    "allowed_address_pairs": [],
                    "description": "",
                    "fixed_ips": [],
                    "name": "",
                    "network_id": "",
                    "tags": {},
                    "updatable": true
                }
            },
            "name": "appliance_2",
            "operation_status": "COMPLETE",
            "os_login_status": "ACTIVE",
            "os_monitoring_status": "ACTIVE",
            "password": "Passw0rd",
            "tags": {
                "k1": "v1"
            },
            "tenant_id": "9ee80f2a926c49f88f166af47df4e9f5",
            "username": "root",
            "virtual_network_appliance_plan_id": "%s",
            "vm_status": "ACTIVE"
        }
    ]
}`,
	// for appliance1
	idAppliance1,
	idVirtualNetworkAppliancePlan,
	// for appliance2
	idAppliance2,
	idVirtualNetworkAppliancePlan,
)

var defaultInterface = appliances.InterfaceInResponse{
	Name:                "",
	Description:         "",
	NetworkID:           "",
	Updatable:           true,
	Tags:                map[string]string{},
	FixedIPs:            []appliances.FixedIPInResponse{},
	AllowedAddressPairs: []appliances.AllowedAddressPairInResponse{},
}

var appliance1 = appliances.Appliance{
	ID:                 idAppliance1,
	Name:               "appliance_1",
	ApplianceType:      applianceType,
	Description:        "appliance_1_description",
	DefaultGateway:     "192.168.1.1",
	AvailabilityZone:   "zone1-groupb",
	OSMonitoringStatus: "ACTIVE",
	OSLoginStatus:      "ACTIVE",
	VMStatus:           "ACTIVE",
	OperationStatus:    "COMPLETE",
	AppliancePlanID:    idVirtualNetworkAppliancePlan,
	TenantID:           "9ee80f2a926c49f88f166af47df4e9f5",
	Tags:               map[string]string{"k1": "v1"},
	Interfaces: appliances.InterfacesInResponse{
		Interface1: appliances.InterfaceInResponse{
			Name:        "interface_1",
			Description: "interface_1_description",
			NetworkID:   "dummyNetworkID",
			Tags:        map[string]string{},
			Updatable:   true,
			FixedIPs: []appliances.FixedIPInResponse{
				appliances.FixedIPInResponse{
					IPAddress: "192.168.1.51",
					SubnetID:  "dummySubnetID",
				},
			},
			AllowedAddressPairs: []appliances.AllowedAddressPairInResponse{
				appliances.AllowedAddressPairInResponse{
					IPAddress:  "1.1.1.1",
					MACAddress: "aa:bb:cc:dd:ee:f1",
					Type:       "vrrp",
					VRID:       float64(123),
				},
			},
		},
		Interface2: defaultInterface,
		Interface3: defaultInterface,
		Interface4: defaultInterface,
		Interface5: defaultInterface,
		Interface6: defaultInterface,
		Interface7: defaultInterface,
		Interface8: defaultInterface,
	},
}

var appliance2 = appliances.Appliance{
	ID:                 idAppliance2,
	Name:               "appliance_2",
	ApplianceType:      applianceType,
	Description:        "appliance_2_description",
	DefaultGateway:     "192.168.1.1",
	AvailabilityZone:   "zone1-groupb",
	OSMonitoringStatus: "ACTIVE",
	OSLoginStatus:      "ACTIVE",
	VMStatus:           "ACTIVE",
	OperationStatus:    "COMPLETE",
	AppliancePlanID:    idVirtualNetworkAppliancePlan,
	TenantID:           "9ee80f2a926c49f88f166af47df4e9f5",
	Tags:               map[string]string{"k1": "v1"},
	Interfaces: appliances.InterfacesInResponse{
		Interface1: appliances.InterfaceInResponse{
			Name:        "interface_1",
			Description: "interface_1_description",
			NetworkID:   "dummyNetworkID",
			Tags:        map[string]string{},
			Updatable:   true,
			FixedIPs: []appliances.FixedIPInResponse{
				appliances.FixedIPInResponse{
					IPAddress: "192.168.1.52",
					SubnetID:  "dummySubnetID",
				},
			},
			AllowedAddressPairs: []appliances.AllowedAddressPairInResponse{
				appliances.AllowedAddressPairInResponse{
					IPAddress:  "2.2.2.2",
					MACAddress: "aa:bb:cc:dd:ee:f2",
					Type:       "",
					VRID:       interface{}(nil),
				},
			},
		},
		Interface2: defaultInterface,
		Interface3: defaultInterface,
		Interface4: defaultInterface,
		Interface5: defaultInterface,
		Interface6: defaultInterface,
		Interface7: defaultInterface,
		Interface8: defaultInterface,
	},
}

var expectedAppliancesSlice = []appliances.Appliance{
	appliance1,
	appliance2,
}
