package testing

import (
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_interfaces"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_syslog_servers"
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancers"
)

const ListResponse = `
{
  "load_balancers": [
    {
      "admin_username": "user-admin",
      "availability_zone": "zone1-groupa",
      "default_gateway": "100.127.253.1",
      "description": "Load Balancer 1 Description",
      "id": "5f3cae7c-58a5-4124-b622-9ca3cfbf2525",
      "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
      "name": "Load Balancer 1",
      "status": "ACTIVE",
      "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
      "user_username": "user-read"
    },
    {
      "admin_username": "user-admin",
      "availability_zone": "zone1_groupa",
      "default_gateway": null,
      "description": "abcdefghijklmnopqrstuvwxyz",
      "id": "601665cf-c161-4e80-87f0-a3c0925d07a0",
      "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
      "name": "Load Balancer 2",
      "status": "PENDING_CREATE",
      "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
      "user_username": "user-read"
    }
  ]
}
`
const GetResponse = `
{
  "load_balancer": {
    "admin_username": "user-admin",
    "availability_zone": "zone1-groupa",
    "default_gateway": "100.127.253.1",
    "description": "Load Balancer 1 Description",
    "id": "5f3cae7c-58a5-4124-b622-9ca3cfbf2525",
    "interfaces": [
      {
        "id": "ee335c69-b50f-4a32-9d0f-f44cef84a456",
        "ip_address": "100.127.253.173",
        "name": "Interface 1/1",
        "network_id": "c7f88fab-573e-47aa-b0b4-257db28dae23",
        "slot_number": 1,
        "status": "ACTIVE",
        "type": "user",
        "virtual_ip_address": "100.127.253.174",
        "virtual_ip_properties": {
			"protocol": "vrrp",
			"vrid": 10
		}
      },
      {
        "id": "b39b61e4-00b1-4698-aed0-1928beb90abe",
        "ip_address": "192.168.110.1",
        "name": "Interface 1/2",
        "network_id": "1839d290-721c-49ba-99f1-3d7aa37811b0",
        "slot_number": 2,
        "status": "ACTIVE",
        "type": "user",
        "virtual_ip_address": null,
        "virtual_ip_properties": null
      }
    ],
    "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
    "name": "Load Balancer 1",
    "status": "ACTIVE",
    "syslog_servers": [
      {
        "id": "11001101-2edf-1844-1ff7-12ba5b7e566a",
        "ip_address": "177.77.07.215",
        "log_facility": "LOCAL0",
        "log_level": "ALERT|INFO|ERROR",
        "name": "syslog_server_main",
        "port_number": 514,
        "status": "ACTIVE"
      },
      {
        "id": "22002202-2edf-1844-1ff7-12ba5b7e566a",
        "ip_address": "177.77.07.211",
        "log_facility": "LOCAL1",
        "log_level": "ERROR",
        "name": "syslog_server_backup_fst",
        "port_number": 514,
        "status": "ACTIVE"
      }
    ],
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "user_username": "user-read"
  }
}
  `
const CreateResponse = `
{
  "load_balancer": {
    "admin_username": "user-admin",
    "availability_zone": "zone1-groupa",
    "default_gateway": "100.127.253.1",
    "description": "Load Balancer 1 Description",
    "id": "5f3cae7c-58a5-4124-b622-9ca3cfbf2525",
    "interfaces": [
      {
        "id": "ee335c69-b50f-4a32-9d0f-f44cef84a456",
        "ip_address": "100.127.253.173",
        "name": "Interface 1/1",
        "network_id": "c7f88fab-573e-47aa-b0b4-257db28dae23",
        "slot_number": 1,
        "status": "ACTIVE",
        "type": "user",
        "virtual_ip_address": "100.127.253.174",
        "virtual_ip_properties": {
			"protocol": "vrrp",
			"vrid": 10
		}
      },
      {
        "id": "b39b61e4-00b1-4698-aed0-1928beb90abe",
        "ip_address": "192.168.110.1",
        "name": "Interface 1/2",
        "network_id": "1839d290-721c-49ba-99f1-3d7aa37811b0",
        "slot_number": 2,
        "status": "ACTIVE",
        "type": "user",
        "virtual_ip_address": null,
        "virtual_ip_properties": null
      }
    ],
    "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
    "name": "Load Balancer 1",
    "status": "ACTIVE",
    "syslog_servers": [
      {
        "id": "11001101-2edf-1844-1ff7-12ba5b7e566a",
        "ip_address": "177.77.07.215",
        "log_facility": "LOCAL0",
        "log_level": "ALERT|INFO|ERROR",
        "name": "syslog_server_main",
        "port_number": 514,
        "status": "ACTIVE"
      },
      {
        "id": "22002202-2edf-1844-1ff7-12ba5b7e566a",
        "ip_address": "177.77.07.211",
        "log_facility": "LOCAL1",
        "log_level": "ERROR",
        "name": "syslog_server_backup_fst",
        "port_number": 514,
        "status": "ACTIVE"
      }
    ],
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "user_username": "user-read"
  }
}
  `
const CreateRequest = `
{
  "load_balancer": {
    "availability_zone": "zone1-groupa",
    "description": "abcdefghijklmnopqrstuvwxyz",
    "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
    "name": "abcdefghijklmnopqrstuvwxyz",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8"
  }
}
`
const UpdateResponse = `
{
  "load_balancer": {
    "admin_username": "user-admin",
    "availability_zone": "zone1-groupa",
    "default_gateway": "100.127.253.1",
    "description": "UPDATED",
    "id": "5f3cae7c-58a5-4124-b622-9ca3cfbf2525",
    "interfaces": [
      {
        "id": "ee335c69-b50f-4a32-9d0f-f44cef84a456",
        "ip_address": "100.127.253.173",
        "name": "Interface 1/1",
        "network_id": "c7f88fab-573e-47aa-b0b4-257db28dae23",
        "slot_number": 1,
        "status": "ACTIVE",
        "virtual_ip_address": null,
        "virtual_ip_properties": null
      },
      {
        "id": "b39b61e4-00b1-4698-aed0-1928beb90abe",
        "ip_address": "192.168.110.1",
        "name": "Interface 1/2",
        "network_id": "1839d290-721c-49ba-99f1-3d7aa37811b0",
        "slot_number": 2,
        "status": "ACTIVE",
        "virtual_ip_address": null,
        "virtual_ip_properties": null
      }
    ],
    "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
    "name": "abcdefghijklmnopqrstuvwxyz",
    "status": "PENDING_UPDATE",
    "syslog_servers": [
      {
        "id": "11001101-2edf-1844-1ff7-12ba5b7e566a",
        "ip_address": "177.77.07.215",
        "log_facility": "LOCAL0",
        "log_level": "ALERT|INFO|ERROR",
        "name": "syslog_server_main",
        "port_number": 514,
        "status": "ACTIVE"
      },
      {
        "id": "22002202-2edf-1844-1ff7-12ba5b7e566a",
        "ip_address": "177.77.07.211",
        "log_facility": "LOCAL1",
        "log_level": "ERROR",
        "name": "syslog_server_backup_fst",
        "port_number": 514,
        "status": "ACTIVE"
      }
    ],
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "user_username": "user-read"
  }
}
`
const UpdateRequest = `
{
  "load_balancer": {
    "default_gateway": "100.127.253.1",
    "description": "UPDATED",
    "load_balancer_plan_id": "bd12784a-c66e-4f13-9f72-5143d64762b6",
    "name": "abcdefghijklmnopqrstuvwxyz"
  }
}
`

var LoadBalancer1 = load_balancers.LoadBalancer{
	ID:                 "5f3cae7c-58a5-4124-b622-9ca3cfbf2525",
	AdminUsername:      "user-admin",
	AvailabilityZone:   "zone1-groupa",
	DefaultGateway:     &DetailDefaultGateway,
	Description:        "Load Balancer 1 Description",
	LoadBalancerPlanID: "bd12784a-c66e-4f13-9f72-5143d64762b6",
	Name:               "Load Balancer 1",
	Status:             "ACTIVE",
	TenantID:           "6a156ddf2ecd497ca786ff2da6df5aa8",
	UserUsername:       "user-read",
}

var LoadBalancer2 = load_balancers.LoadBalancer{
	ID:                 "601665cf-c161-4e80-87f0-a3c0925d07a0",
	AdminUsername:      "user-admin",
	AvailabilityZone:   "zone1_groupa",
	Description:        "abcdefghijklmnopqrstuvwxyz",
	LoadBalancerPlanID: "bd12784a-c66e-4f13-9f72-5143d64762b6",
	Name:               "Load Balancer 2",
	Status:             "PENDING_CREATE",
	TenantID:           "6a156ddf2ecd497ca786ff2da6df5aa8",
	UserUsername:       "user-read",
}

var DetailDefaultGateway = "100.127.253.1"
var DetailIPAddress1 = "100.127.253.173"
var DetailNetworkID1 = "c7f88fab-573e-47aa-b0b4-257db28dae23"
var DetailVirtualIPAddress1 = "100.127.253.174"

var DetailIPAddress2 = "192.168.110.1"
var DetailNetworkID2 = "1839d290-721c-49ba-99f1-3d7aa37811b0"

var LoadBalancerDetail = load_balancers.LoadBalancer{
	ID:               "5f3cae7c-58a5-4124-b622-9ca3cfbf2525",
	AdminUsername:    "user-admin",
	AvailabilityZone: "zone1-groupa",
	DefaultGateway:   &DetailDefaultGateway,
	Description:      "Load Balancer 1 Description",
	Interfaces: []load_balancer_interfaces.LoadBalancerInterface{
		{
			ID:               "ee335c69-b50f-4a32-9d0f-f44cef84a456",
			IPAddress:        &DetailIPAddress1,
			Name:             "Interface 1/1",
			NetworkID:        &DetailNetworkID1,
			SlotNumber:       1,
			Status:           "ACTIVE",
			Type:             "user",
			VirtualIPAddress: &DetailVirtualIPAddress1,
			VirtualIPProperties: &load_balancer_interfaces.VirtualIPProperties{
				Protocol: "vrrp",
				Vrid:     10,
			},
		},
		{
			ID:         "b39b61e4-00b1-4698-aed0-1928beb90abe",
			IPAddress:  &DetailIPAddress2,
			Name:       "Interface 1/2",
			NetworkID:  &DetailNetworkID2,
			SlotNumber: 2,
			Status:     "ACTIVE",
			Type:       "user",
		},
	},
	LoadBalancerPlanID: "bd12784a-c66e-4f13-9f72-5143d64762b6",
	Name:               "Load Balancer 1",
	Status:             "ACTIVE",
	SyslogServers: []load_balancer_syslog_servers.LoadBalancerSyslogServer{
		{
			ID:          "11001101-2edf-1844-1ff7-12ba5b7e566a",
			IPAddress:   "177.77.07.215",
			LogFacility: "LOCAL0",
			LogLevel:    "ALERT|INFO|ERROR",
			Name:        "syslog_server_main",
			PortNumber:  514,
			Status:      "ACTIVE",
		},
		{
			ID:          "22002202-2edf-1844-1ff7-12ba5b7e566a",
			IPAddress:   "177.77.07.211",
			LogFacility: "LOCAL1",
			LogLevel:    "ERROR",
			Name:        "syslog_server_backup_fst",
			PortNumber:  514,
			Status:      "ACTIVE",
		},
	},
	TenantID:     "6a156ddf2ecd497ca786ff2da6df5aa8",
	UserUsername: "user-read",
}

var ExpectedLoadBalancerSlice = []load_balancers.LoadBalancer{LoadBalancer1, LoadBalancer2}
