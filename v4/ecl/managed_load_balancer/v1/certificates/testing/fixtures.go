package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/certificates"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "certificates": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "certificate",
            "description": "description",
            "tags": {
                "key": "value"
            },
            "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
            "ca_cert": {
                "status": "NOT_UPLOADED"
            },
            "ssl_cert": {
                "status": "NOT_UPLOADED"
            },
            "ssl_key": {
                "status": "NOT_UPLOADED"
            }
        }
    ]
}`)

func listResult() []certificates.Certificate {
	var certificate1 certificates.Certificate

	sslKey1 := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	sslCert1 := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	caCert1 := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err := json.Unmarshal([]byte(tags1Json), &tags1)
	if err != nil {
		panic(err)
	}

	certificate1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	certificate1.Name = "certificate"
	certificate1.Description = "description"
	certificate1.Tags = tags1
	certificate1.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	certificate1.CACert = caCert1
	certificate1.SSLCert = sslCert1
	certificate1.SSLKey = sslKey1

	return []certificates.Certificate{certificate1}
}

var createRequest = fmt.Sprintf(`
{
    "certificate": {
        "name": "certificate",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var createResponse = fmt.Sprintf(`
{
    "certificate": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "certificate",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ca_cert": {
            "status": "NOT_UPLOADED"
        },
        "ssl_cert": {
            "status": "NOT_UPLOADED"
        },
        "ssl_key": {
            "status": "NOT_UPLOADED"
        }
    }
}`)

func createResult() *certificates.Certificate {
	var certificate certificates.Certificate

	sslKey := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	sslCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	caCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	certificate.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	certificate.Name = "certificate"
	certificate.Description = "description"
	certificate.Tags = tags
	certificate.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	certificate.CACert = caCert
	certificate.SSLCert = sslCert
	certificate.SSLKey = sslKey

	return &certificate
}

var showResponse = fmt.Sprintf(`
{
    "certificate": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "certificate",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ca_cert": {
            "status": "NOT_UPLOADED"
        },
        "ssl_cert": {
            "status": "NOT_UPLOADED"
        },
        "ssl_key": {
            "status": "NOT_UPLOADED"
        }
    }
}`)

func showResult() *certificates.Certificate {
	var certificate certificates.Certificate

	sslKey := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	sslCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	caCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	certificate.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	certificate.Name = "certificate"
	certificate.Description = "description"
	certificate.Tags = tags
	certificate.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	certificate.CACert = caCert
	certificate.SSLCert = sslCert
	certificate.SSLKey = sslKey

	return &certificate
}

var updateRequest = fmt.Sprintf(`
{
    "certificate": {
        "name": "certificate",
        "description": "description",
        "tags": {
            "key": "value"
        }
    }
}`)

var updateResponse = fmt.Sprintf(`
{
    "certificate": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "certificate",
        "description": "description",
        "tags": {
            "key": "value"
        },
        "tenant_id": "34f5c98ef430457ba81292637d0c6fd0",
        "ca_cert": {
            "status": "NOT_UPLOADED"
        },
        "ssl_cert": {
            "status": "NOT_UPLOADED"
        },
        "ssl_key": {
            "status": "NOT_UPLOADED"
        }
    }
}`)

func updateResult() *certificates.Certificate {
	var certificate certificates.Certificate

	sslKey := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	sslCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}
	caCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err := json.Unmarshal([]byte(tagsJson), &tags)
	if err != nil {
		panic(err)
	}

	certificate.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	certificate.Name = "certificate"
	certificate.Description = "description"
	certificate.Tags = tags
	certificate.TenantID = "34f5c98ef430457ba81292637d0c6fd0"
	certificate.CACert = caCert
	certificate.SSLCert = sslCert
	certificate.SSLKey = sslKey

	return &certificate
}

var uploadFileRequest = fmt.Sprintf(`
{
    "type": "ca-cert",
    "content": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCjAxMjM0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEKMjM0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMwo0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1CjY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1NjcKODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OQpBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCCkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0QKRUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRgpHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdICklKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUoKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTApNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OCk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1AKUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUgpTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUClVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVYKV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWApZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaCmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaYWIKY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaYWJjZAplZmdoaWprbG1ub3BxcnN0dXZ3eHl6VjAxMjM0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlCmZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg=="
}`)
