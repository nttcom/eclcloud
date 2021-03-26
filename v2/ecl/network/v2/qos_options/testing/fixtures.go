package testing

import "github.com/nttcom/eclcloud/v2/ecl/network/v2/qos_options"

const ListResponse = `
{
	"qos_options": [
		{
			"aws_service_id"	  : null,
			"azure_service_id"	  : "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
			"bandwidth"			  : "20",
			"description"		  : "20M-guarantee-menu-for-azure",
			"fic_service_id"	  : null,
			"gcp_service_id"	  : null,
			"id"				  : "a6b91294-8870-4f2c-b9e9-a899acada723",
			"interdc_service_id"  : null,
			"internet_service_id" : null,
			"name"				  : "20M-GA-AZURE",
			"qos_type"			  : "guarantee",
			"service_type"		  : "azure",
			"status"			  : "ACTIVE",
			"vpn_service_id"	  : null
		},
		{
			"aws_service_id"	  : null,
			"azure_service_id"	  : "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
			"bandwidth"			  : "500",
			"description"		  : "500M-guarantee-menu-for-azure",
			"fic_service_id"	  : null,
			"gcp_service_id"	  : null,
			"id"				  : "aa776ce4-08a8-4cc1-9a2c-bb95e547916b",
			"interdc_service_id"  : null,
			"internet_service_id" : null,
			"name"				  : "500M-GA-AZURE",
			"qos_type"			  : "guarantee",
			"service_type"		  : "azure",
			"status"			  : "ACTIVE",
			"vpn_service_id"	  : null
		}
	]
}
`

const GetResponse = `
{
	"qos_option": {
		"aws_service_id"	  : null,
		"azure_service_id"	  : "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
		"bandwidth"			  : "20",
		"description"		  : "20M-guarantee-menu-for-azure",
		"fic_service_id"	  : null,
		"gcp_service_id"	  : null,
		"id"				  : "a6b91294-8870-4f2c-b9e9-a899acada723",
		"interdc_service_id"  : null,
		"internet_service_id" : null,
		"name"				  : "20M-GA-AZURE",
		"qos_type"			  : "guarantee",
		"service_type"		  : "azure",
		"status"			  : "ACTIVE",
		"vpn_service_id"	  : null
	}
}
`

var Qos1 = qos_options.QoSOption{
	AWSServiceID:      "",
	AzureServiceID:    "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
	Bandwidth:         "20",
	Description:       "20M-guarantee-menu-for-azure",
	FICServiceID:      "",
	GCPServiceID:      "",
	ID:                "a6b91294-8870-4f2c-b9e9-a899acada723",
	InterDCServiceID:  "",
	InternetServiceID: "",
	Name:              "20M-GA-AZURE",
	QoSType:           "guarantee",
	ServiceType:       "azure",
	Status:            "ACTIVE",
	VPNServiceID:      "",
}

var Qos2 = qos_options.QoSOption{
	AWSServiceID:      "",
	AzureServiceID:    "d4006e79-9f60-4b72-9f86-5f6ef8b4e9e9",
	Bandwidth:         "500",
	Description:       "500M-guarantee-menu-for-azure",
	FICServiceID:      "",
	GCPServiceID:      "",
	ID:                "aa776ce4-08a8-4cc1-9a2c-bb95e547916b",
	InterDCServiceID:  "",
	InternetServiceID: "",
	Name:              "500M-GA-AZURE",
	QoSType:           "guarantee",
	ServiceType:       "azure",
	Status:            "ACTIVE",
	VPNServiceID:      "",
}

var ExpectedQosSlice = []qos_options.QoSOption{Qos1, Qos2}
