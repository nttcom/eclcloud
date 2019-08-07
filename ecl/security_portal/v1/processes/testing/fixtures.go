package testing

import (
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/processes"
)

const processID = "85385"

const getResponse = `
{
  "processInstance": {
    "processId": {
      "id": 85385,
      "lastExecNumber": 1,
      "name": "ntt/FortiVA_Port_Management/Process_Manage_UTM_Interfaces/Process_Manage_UTM_Interfaces",
      "submissionType": "RUN"
    },
    "serviceId": {
      "id": 19382,
      "name": "FortiVA_Port_Management",
      "serviceReference": "PORT_MNGT_CES11892",
      "state": null
    },
    "status": {
      "comment": "Ping Monitoring started for the device 11892.",
      "duration": 0,
      "endingDate": "2019-07-26 04:34:56.0",
      "execNumber": 1,
      "processInstanceId": 85385,
      "processName": "ntt/FortiVA_Port_Management/Process_Manage_UTM_Interfaces/Process_Manage_UTM_Interfaces",
      "startingDate": "2019-07-26 04:24:45.0",
      "status": "RUNNING",
      "taskStatusList": [
        {
          "comment": "IP Address inputs verified successfully.",
          "endingDate": "2019-07-26 04:24:48.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:24:45.0",
          "status": "ENDED",
          "taskId": 1,
          "taskName": "Verify IP Address, MTU Inputs"
        },
        {
          "comment": "Ping Monitoring stopped for the device 11892.",
          "endingDate": "2019-07-26 04:26:49.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:24:48.0",
          "status": "ENDED",
          "taskId": 2,
          "taskName": "Stop Ping Monitoring"
        },
        {
          "comment": "Openstack Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3 stopped successfully.\nServer Status : SHUTOFF\nTask State : -\nPower State : Shutdown\n",
          "endingDate": "2019-07-26 04:27:03.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:26:49.0",
          "status": "ENDED",
          "taskId": 3,
          "taskName": "Stop the UTM"
        },
        {
          "comment": "IP Address 100.76.96.230 is now unreachable from MSA.\nPING Status : Destination Host Unreachable\n",
          "endingDate": "2019-07-26 04:27:13.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:27:03.0",
          "status": "ENDED",
          "taskId": 4,
          "taskName": "Wait for UTM Ping unreachability from MSA"
        },
        {
          "comment": "Ports deleted successfully.",
          "endingDate": "2019-07-26 04:28:29.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:27:13.0",
          "status": "ENDED",
          "taskId": 5,
          "taskName": "Delete Ports"
        },
        {
          "comment": "Ports created successfully.\nPort Id : 34c7389d-1428-4f98-a37c-9c2e32aab255\nPort Id : 3d09053b-fad8-45c4-bf71-501c0fc2b58a\nPort Id : 0262d90c-6056-4308-8b76-8e851f0132f5\nPort Id : 5fcabdf2-8a20-4337-bd10-02f5c5000ca1\nPort Id : 53211b09-f82b-40d5-bf5b-7289a298cbdf\nPort Id : 9ce2d3b7-7ae0-400d-8e41-16dc9b94f95e\nPort Id : a36493fe-43d2-4dc1-a39e-c96898e9c0be\n",
          "endingDate": "2019-07-26 04:29:50.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:28:29.0",
          "status": "ENDED",
          "taskId": 6,
          "taskName": "Create Ports"
        },
        {
          "comment": "Ports attached successfully to the Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3.",
          "endingDate": "2019-07-26 04:31:33.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:29:50.0",
          "status": "ENDED",
          "taskId": 7,
          "taskName": "Attach Ports"
        },
        {
          "comment": "Openstack Server 158eb01a-8d45-45c8-a9ff-1fba8f1ab7e3 started successfully.\nServer Status : ACTIVE\nTask State : -\nPower State : Running\n",
          "endingDate": "2019-07-26 04:31:47.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:31:33.0",
          "status": "ENDED",
          "taskId": 8,
          "taskName": "Start the UTM"
        },
        {
          "comment": "IP Address 100.76.96.230 is now reachable from MSA.\nPING Status : OK\n",
          "endingDate": "2019-07-26 04:32:30.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:31:47.0",
          "status": "ENDED",
          "taskId": 9,
          "taskName": "Wait for UTM Ping reachability from MSA"
        },
        {
          "comment": "OK LICENSE IS VALID",
          "endingDate": "2019-07-26 04:32:56.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:32:30.0",
          "status": "ENDED",
          "taskId": 10,
          "taskName": "Verify License Validity"
        },
        {
          "comment": "Ports updated successfully on Fortigate Device 11892.\n",
          "endingDate": "2019-07-26 04:33:17.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:32:56.0",
          "status": "ENDED",
          "taskId": 11,
          "taskName": "Update UTM"
        },
        {
          "comment": "Device 11892 Backup completed successfully.\nBackup Status : ENDED\nBackup Message : BACKUP  processed\n\nBackup Revision Id : 209408\n",
          "endingDate": "2019-07-26 04:33:28.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:33:17.0",
          "status": "ENDED",
          "taskId": 12,
          "taskName": "Device Backup"
        },
        {
          "comment": "Ping Monitoring started for the device 11892.",
          "endingDate": "2019-07-26 04:34:56.0",
          "execNumber": 1,
          "newParameters": {},
          "processInstanceId": 85385,
          "startingDate": "2019-07-26 04:33:28.0",
          "status": "ENDED",
          "taskId": 13,
          "taskName": "Start Ping Monitoring"
        }
      ]
    }
  }
}`

var expectedProcess = processes.ProcessInstance{
	Status: processes.ProcessStatus{
		Status: "RUNNING",
	},
}
