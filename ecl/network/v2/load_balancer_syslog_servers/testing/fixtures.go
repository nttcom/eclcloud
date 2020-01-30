package testing

import (
	"github.com/nttcom/eclcloud/ecl/network/v2/load_balancer_syslog_servers"
)

const ListResponse = `
{
  "load_balancer_syslog_servers": [
    {
      "description": "test",
      "id": "6e9c7745-61f2-491f-9689-add8c5fc4b9a",
      "ip_address": "120.120.120.30",
      "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
      "log_facility": "LOCAL3",
      "log_level": "DEBUG",
      "name": "first_syslog_server",
      "port_number": 514,
      "status": "ACTIVE",
      "transport_type": "UDP"
    },
    {
      "description": "My second backup server",
      "id": "c7de2dee-73a0-4a9b-acdf-8a348c242a30",
      "ip_address": "120.120.122.30",
      "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
      "log_facility": "LOCAL2",
      "log_level": "ERROR",
      "name": "second_syslog_server",
      "port_number": 514,
      "status": "ACTIVE",
      "transport_type": "UDP"
    }
  ]
}
`
const GetResponse = `
{
  "load_balancer_syslog_server": {
    "acl_logging": "DISABLED",
    "appflow_logging": "DISABLED",
    "date_format": "MMDDYYYY",
    "description": "test",
    "id": "6e9c7745-61f2-491f-9689-add8c5fc4b9a",
    "ip_address": "120.120.120.30",
    "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
    "log_facility": "LOCAL3",
    "log_level": "DEBUG",
    "name": "first_syslog_server",
    "port_number": 514,
    "priority": 20,
    "status": "ACTIVE",
    "tcp_logging": "ALL",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "time_zone": "LOCAL_TIME",
    "transport_type": "UDP",
    "user_configurable_log_messages": "NO"
  }
}
`
const CreateResponse = `
{
  "load_balancer_syslog_server": {
    "acl_logging": "DISABLED",
    "appflow_logging": "DISABLED",
    "date_format": "MMDDYYYY",
    "description": "test",
    "id": "6e9c7745-61f2-491f-9689-add8c5fc4b9a",
    "ip_address": "120.120.120.30",
    "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
    "log_facility": "LOCAL3",
    "log_level": "DEBUG",
    "name": "first_syslog_server",
    "port_number": 514,
    "priority": 20,
    "status": "ACTIVE",
    "tcp_logging": "ALL",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "time_zone": "LOCAL_TIME",
    "transport_type": "UDP",
    "user_configurable_log_messages": "NO"
  }
}
`
const CreateRequest = `
{
  "load_balancer_syslog_server": {
    "acl_logging": "DISABLED",
    "appflow_logging": "DISABLED",
    "date_format": "MMDDYYYY",
    "description": "test",
    "ip_address": "120.120.120.30",
    "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
    "log_facility": "LOCAL3",
    "log_level": "DEBUG",
    "name": "first_syslog_server",
    "port_number": 514,
    "priority": 20,
    "tcp_logging": "ALL",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "time_zone": "LOCAL_TIME",
    "transport_type": "UDP",
    "user_configurable_log_messages": "NO"
  }
}
`
const UpdateResponse = `
{
  "load_balancer_syslog_server": {
    "acl_logging": "DISABLED",
    "appflow_logging": "DISABLED",
    "date_format": "MMDDYYYY",
    "description": "test2",
    "id": "6e9c7745-61f2-491f-9689-add8c5fc4b9a",
    "ip_address": "120.120.120.30",
    "load_balancer_id": "9f872504-36ab-46af-83ce-a4991c669edd",
    "log_facility": "LOCAL3",
    "log_level": "DEBUG",
    "name": "first_syslog_server",
    "port_number": 514,
    "priority": 20,
    "status": "PENDING_UPDATE",
    "tcp_logging": "ALL",
    "tenant_id": "6a156ddf2ecd497ca786ff2da6df5aa8",
    "time_zone": "LOCAL_TIME",
    "transport_type": "UDP",
    "user_configurable_log_messages": "NO"
  }
}
`
const UpdateRequest = `
{
  "load_balancer_syslog_server": {
    "acl_logging": "DISABLED",
    "appflow_logging": "DISABLED",
    "date_format": "MMDDYYYY",
    "description": "test2",
    "log_facility": "LOCAL3",
    "log_level": "DEBUG",
    "priority": 20,
    "tcp_logging": "ALL",
    "time_zone": "LOCAL_TIME",
    "user_configurable_log_messages": "NO"
  }
}
`

var LoadBalancerSyslogServer1 = load_balancer_syslog_servers.LoadBalancerSyslogServer{
	Description:    "test",
	ID:             "6e9c7745-61f2-491f-9689-add8c5fc4b9a",
	IPAddress:      "120.120.120.30",
	LoadBalancerID: "9f872504-36ab-46af-83ce-a4991c669edd",
	LogFacility:    "LOCAL3",
	LogLevel:       "DEBUG",
	Name:           "first_syslog_server",
	PortNumber:     514,
	Status:         "ACTIVE",
	TransportType:  "UDP",
}

var LoadBalancerSyslogServer2 = load_balancer_syslog_servers.LoadBalancerSyslogServer{
	Description:    "My second backup server",
	ID:             "c7de2dee-73a0-4a9b-acdf-8a348c242a30",
	IPAddress:      "120.120.122.30",
	LoadBalancerID: "9f872504-36ab-46af-83ce-a4991c669edd",
	LogFacility:    "LOCAL2",
	LogLevel:       "ERROR",
	Name:           "second_syslog_server",
	PortNumber:     514,
	Status:         "ACTIVE",
	TransportType:  "UDP",
}

var LoadBalancerSyslogServerDetail = load_balancer_syslog_servers.LoadBalancerSyslogServer{
	AclLogging:                  "DISABLED",
	AppflowLogging:              "DISABLED",
	DateFormat:                  "MMDDYYYY",
	Description:                 "test",
	ID:                          "6e9c7745-61f2-491f-9689-add8c5fc4b9a",
	IPAddress:                   "120.120.120.30",
	LoadBalancerID:              "9f872504-36ab-46af-83ce-a4991c669edd",
	LogFacility:                 "LOCAL3",
	LogLevel:                    "DEBUG",
	Name:                        "first_syslog_server",
	PortNumber:                  514,
	Priority:                    20,
	Status:                      "ACTIVE",
	TcpLogging:                  "ALL",
	TenantID:                    "6a156ddf2ecd497ca786ff2da6df5aa8",
	TimeZone:                    "LOCAL_TIME",
	TransportType:               "UDP",
	UserConfigurableLogMessages: "NO",
}

var ExpectedLoadBalancerSlice = []load_balancer_syslog_servers.LoadBalancerSyslogServer{LoadBalancerSyslogServer1, LoadBalancerSyslogServer2}
