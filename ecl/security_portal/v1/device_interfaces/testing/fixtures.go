package testing

import (
	"github.com/nttcom/eclcloud/ecl/security_portal/v1/device_interfaces"
)

const deviceUUID = "cad6e00a-2af9-491c-b732-ca5688d147f5"

const listResponse = `
{
    "device_interfaces": [
        {
            "os_ip_address": "192.168.1.50",
            "msa_port_id": "port4",
            "os_port_name": "port4-CES11892",
            "os_port_id": "82ebe045-9c9a-4088-8b33-cb0d590079aa",
            "os_network_id": "5ef9c597-15fe-431c-84b9-88d00d567202",
            "os_port_status": "ACTIVE",
            "os_mac_address": "fa:16:3e:05:ff:66",
            "os_subnet_id": "48ea24c7-fe48-4a54-9ed0-528aa09cebc7"
        },
        {
            "os_ip_address": "192.168.2.50",
            "msa_port_id": "port7",
            "os_port_name": "port7-CES11892",
            "os_port_id": "82ebe045-9c9a-4088-8b33-cb0d590079aa",
            "os_network_id": "5ef9c597-15fe-431c-84b9-88d00d567203",
            "os_port_status": "ACTIVE",
            "os_mac_address": "fa:16:3e:05:ff:67",
            "os_subnet_id": "48ea24c7-fe48-4a54-9ed0-528aa09cebc8"
        }
    ]
}
`

var expectedDeviceInterfacesSlice = []device_interfaces.DeviceInterface{
	firstDeviceInterface, secondDeviceInterface,
}

var firstDeviceInterface = device_interfaces.DeviceInterface{
	OSIPAddress:  "192.168.1.50",
	MSAPortID:    "port4",
	OSPortName:   "port4-CES11892",
	OSPortID:     "82ebe045-9c9a-4088-8b33-cb0d590079aa",
	OSNetworkID:  "5ef9c597-15fe-431c-84b9-88d00d567202",
	OSPortStatus: "ACTIVE",
	OSMACAddress: "fa:16:3e:05:ff:66",
	OSSubnetID:   "48ea24c7-fe48-4a54-9ed0-528aa09cebc7",
}

var secondDeviceInterface = device_interfaces.DeviceInterface{
	OSIPAddress:  "192.168.2.50",
	MSAPortID:    "port7",
	OSPortName:   "port7-CES11892",
	OSPortID:     "82ebe045-9c9a-4088-8b33-cb0d590079aa",
	OSNetworkID:  "5ef9c597-15fe-431c-84b9-88d00d567203",
	OSPortStatus: "ACTIVE",
	OSMACAddress: "fa:16:3e:05:ff:67",
	OSSubnetID:   "48ea24c7-fe48-4a54-9ed0-528aa09cebc8",
}
