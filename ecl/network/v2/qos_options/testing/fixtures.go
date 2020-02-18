package testing

import "github.com/nttcom/eclcloud/ecl/network/v2/qos_options"

const ListResponse = `
{
	"qos_options": [
	  	{
			  "aws_service_id"		: "3472d029-8378-424c-aaaf-150edc08947b",
			  "azure_service_id"	: "",
			  "bandwidth"			: "400",
			  "description"			: "this is test1",
			  "fic_service_id"		: "23456",
			  "gcp_service_id"		: "7891011",
			  "id"					: "2c649b8e-f007-4d90-b208-9b8710937a94",
			  "interdc_service_id"	: "34567",
			  "internet_service_id" : "89101112",
			  "name"				: "100Gbps-Guaranteed",
			  "qos_type"			: "guarantee",
			  "service_type"		: "vpn",
			  "status"				: "ACTIVE",
			  "vpn_service_id"		: ""
	  	},
	  	{
			  "aws_service_id"		: "",
			  "azure_service_id"	: "99b91166-2166-4187-81e0-23d957773257",
			  "bandwidth"			: "100",
			  "description"			: "this is test2",
			  "fic_service_id"		: "34567",
			  "gcp_service_id"		: "89101112",
			  "id"					: "e39cd764-a30b-4b65-8dd7-b908b1665085",
			  "interdc_service_id"	: "45678",
			  "internet_service_id" : "910111213",
			  "name"				: "100Gbps-BestEffort",
			  "qos_type"			: "besteffort",
			  "service_type"		: "gcp",
			  "status"				: "ACTIVE",
			  "vpn_service_id"		: "613b6f71-ec65-4d51-93f8-ed43eaf4fd15"
		}
	]
}`
const GetResponse = `{
	"qos_option": {
		  "aws_service_id"		: "3472d029-8378-424c-aaaf-150edc08947b",
		  "azure_service_id"	: "",
		  "bandwidth"			: "400",
		  "description"			: "this is test1",
		  "fic_service_id"		: "23456",
		  "gcp_service_id"		: "7891011",
		  "id"					: "2c649b8e-f007-4d90-b208-9b8710937a94",
		  "interdc_service_id"	: "34567",
		  "internet_service_id" : "89101112",
		  "name"				: "100Gbps-Guaranteed",
		  "qos_type"			: "guarantee",
		  "service_type"		: "vpn",
		  "status"				: "ACTIVE",
		  "vpn_service_id"		: ""
	}
}`

var Qos1 = qos_options.QoSOpt{
	AWSServiceID:      "3472d029-8378-424c-aaaf-150edc08947b",
	AzureServiceID:    "",
	Bandwidth:         "400",
	Description:       "this is test1",
	FICServiceID:      "23456",
	GCPServiceID:      "7891011",
	ID:                "2c649b8e-f007-4d90-b208-9b8710937a94",
	InterDCServiceID:  "34567",
	InternetServiceID: "89101112",
	Name:              "100Gbps-Guaranteed",
	QoSType:           "guarantee",
	ServiceType:       "vpn",
	Status:            "ACTIVE",
	VPNServiceID:      "",
}

var Qos2 = qos_options.QoSOpt{
	AWSServiceID:      "",
	AzureServiceID:    "99b91166-2166-4187-81e0-23d957773257",
	Bandwidth:         "100",
	Description:       "this is test2",
	FICServiceID:      "34567",
	GCPServiceID:      "89101112",
	ID:                "e39cd764-a30b-4b65-8dd7-b908b1665085",
	InterDCServiceID:  "45678",
	InternetServiceID: "910111213",
	Name:              "100Gbps-BestEffort",
	QoSType:           "besteffort",
	ServiceType:       "gcp",
	Status:            "ACTIVE",
	VPNServiceID:      "613b6f71-ec65-4d51-93f8-ed43eaf4fd15",
}

var ExpectedQosSlice = []qos_options.QoSOpt{Qos1, Qos2}
