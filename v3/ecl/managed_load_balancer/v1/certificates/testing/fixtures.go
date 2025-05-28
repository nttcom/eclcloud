package testing

import (
	"encoding/json"
	"fmt"

	"github.com/nttcom/eclcloud/v3/ecl/managed_load_balancer/v1/certificates"
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
                "status": "UPLOADED",
                "info": {
                    "issuer": {
                        "C": "JP",
                        "ST": "Tokyo",
                        "L": "Chiyoda-ku",
                        "O": "NTT Communications Corporation",
                        "CN": "example.com"
                    },
                    "subject": {
                        "C": "JP",
                        "ST": "Tokyo",
                        "L": "Chiyoda-ku",
                        "O": "NTT Communications Corporation",
                        "CN": "example.com"
                    },
                    "not_before": "2023-11-09 06:20:55",
                    "not_after": "2024-12-10 06:20:54",
                    "key_algorithm": "RSA-4096",
                    "serial": "e7:61:4a:49:85:aa:7c:f2",
                    "fingerprint": "db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd"
                }
            },
            "ssl_cert": {
                "status": "UPLOADED",
                "info": {
                    "issuer": {
                        "C": "JP",
                        "ST": "Tokyo",
                        "L": "Chiyoda-ku",
                        "O": "NTT Communications Corporation",
                        "CN": "example.com"
                    },
                    "subject": {
                        "C": "JP",
                        "ST": "Tokyo",
                        "L": "Chiyoda-ku",
                        "O": "NTT Communications Corporation",
                        "CN": "example.com"
                    },
                    "not_before": "2023-11-09 06:20:55",
                    "not_after": "2024-12-10 06:20:54",
                    "key_algorithm": "RSA-4096",
                    "serial": "d3:11:fe:4d:a3:71:4e:13",
                    "fingerprint": "46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2"
                }
            },
            "ssl_key": {
                "status": "UPLOADED",
                "info": {
                    "key_algorithm": "RSA-4096",
                    "passphrase": true
                }
            }
        }
    ]
}`)

func listResult() []certificates.Certificate {
	var certificate1 certificates.Certificate

	var sslKey1Info map[string]interface{}
	sslKey1InfoJson := `{"key_algorithm":"RSA-4096","passphrase":true}`
	err := json.Unmarshal([]byte(sslKey1InfoJson), &sslKey1Info)
	if err != nil {
		panic(err)
	}

	sslKey1 := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   sslKey1Info,
	}

	var sslCert1Info map[string]interface{}
	sslCert1InfoJson := `{"issuer":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"example` +
		`.com"},"subject":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"` +
		`example.com"},"not_before":"2023-11-09 06:20:55","not_after":"2024-12-10 06:20:54","key_algorithm":"` +
		`RSA-4096","serial":"d3:11:fe:4d:a3:71:4e:13","fingerprint":"46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:8` +
		`8:9f:3d:a7:c5:42:b2"}`
	err = json.Unmarshal([]byte(sslCert1InfoJson), &sslCert1Info)
	if err != nil {
		panic(err)
	}

	sslCert1 := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   sslCert1Info,
	}

	var caCert1Info map[string]interface{}
	caCert1InfoJson := `{"issuer":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"example` +
		`.com"},"subject":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"` +
		`example.com"},"not_before":"2023-11-09 06:20:55","not_after":"2024-12-10 06:20:54","key_algorithm":"` +
		`RSA-4096","serial":"e7:61:4a:49:85:aa:7c:f2","fingerprint":"db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:1` +
		`2:2c:37:4d:bb:7a:bd"}`
	err = json.Unmarshal([]byte(caCert1InfoJson), &caCert1Info)
	if err != nil {
		panic(err)
	}

	caCert1 := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   caCert1Info,
	}

	var tags1 map[string]interface{}
	tags1Json := `{"key":"value"}`
	err = json.Unmarshal([]byte(tags1Json), &tags1)
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
            "status": "NOT_UPLOADED",
            "info": {}
        },
        "ssl_cert": {
            "status": "NOT_UPLOADED",
            "info": {}
        },
        "ssl_key": {
            "status": "NOT_UPLOADED",
            "info": {}
        }
    }
}`)

func createResult() *certificates.Certificate {
	var certificate certificates.Certificate

	var sslKeyInfo map[string]interface{}
	sslKeyInfoJson := `{}`
	err := json.Unmarshal([]byte(sslKeyInfoJson), &sslKeyInfo)
	if err != nil {
		panic(err)
	}

	sslKey := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
		Info:   sslKeyInfo,
	}

	var sslCertInfo map[string]interface{}
	sslCertInfoJson := `{}`
	err = json.Unmarshal([]byte(sslCertInfoJson), &sslCertInfo)
	if err != nil {
		panic(err)
	}

	sslCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
		Info:   sslCertInfo,
	}

	var caCertInfo map[string]interface{}
	caCertInfoJson := `{}`
	err = json.Unmarshal([]byte(caCertInfoJson), &caCertInfo)
	if err != nil {
		panic(err)
	}

	caCert := certificates.FileInResponse{
		Status: "NOT_UPLOADED",
		Info:   caCertInfo,
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err = json.Unmarshal([]byte(tagsJson), &tags)
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
            "status": "UPLOADED",
            "info": {
                "issuer": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "subject": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "not_before": "2023-11-09 06:20:55",
                "not_after": "2024-12-10 06:20:54",
                "key_algorithm": "RSA-4096",
                "serial": "e7:61:4a:49:85:aa:7c:f2",
                "fingerprint": "db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd"
            }
        },
        "ssl_cert": {
            "status": "UPLOADED",
            "info": {
                "issuer": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "subject": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "not_before": "2023-11-09 06:20:55",
                "not_after": "2024-12-10 06:20:54",
                "key_algorithm": "RSA-4096",
                "serial": "d3:11:fe:4d:a3:71:4e:13",
                "fingerprint": "46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2"
            }
        },
        "ssl_key": {
            "status": "UPLOADED",
            "info": {
                "key_algorithm": "RSA-4096",
                "passphrase": true
            }
        }
    }
}`)

func showResult() *certificates.Certificate {
	var certificate certificates.Certificate

	var sslKeyInfo map[string]interface{}
	sslKeyInfoJson := `{"key_algorithm":"RSA-4096","passphrase":true}`
	err := json.Unmarshal([]byte(sslKeyInfoJson), &sslKeyInfo)
	if err != nil {
		panic(err)
	}

	sslKey := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   sslKeyInfo,
	}

	var sslCertInfo map[string]interface{}
	sslCertInfoJson := `{"issuer":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"example` +
		`.com"},"subject":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"` +
		`example.com"},"not_before":"2023-11-09 06:20:55","not_after":"2024-12-10 06:20:54","key_algorithm":"` +
		`RSA-4096","serial":"d3:11:fe:4d:a3:71:4e:13","fingerprint":"46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:8` +
		`8:9f:3d:a7:c5:42:b2"}`
	err = json.Unmarshal([]byte(sslCertInfoJson), &sslCertInfo)
	if err != nil {
		panic(err)
	}

	sslCert := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   sslCertInfo,
	}

	var caCertInfo map[string]interface{}
	caCertInfoJson := `{"issuer":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"example` +
		`.com"},"subject":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"` +
		`example.com"},"not_before":"2023-11-09 06:20:55","not_after":"2024-12-10 06:20:54","key_algorithm":"` +
		`RSA-4096","serial":"e7:61:4a:49:85:aa:7c:f2","fingerprint":"db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:1` +
		`2:2c:37:4d:bb:7a:bd"}`
	err = json.Unmarshal([]byte(caCertInfoJson), &caCertInfo)
	if err != nil {
		panic(err)
	}

	caCert := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   caCertInfo,
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err = json.Unmarshal([]byte(tagsJson), &tags)
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
            "status": "UPLOADED",
            "info": {
                "issuer": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "subject": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "not_before": "2023-11-09 06:20:55",
                "not_after": "2024-12-10 06:20:54",
                "key_algorithm": "RSA-4096",
                "serial": "e7:61:4a:49:85:aa:7c:f2",
                "fingerprint": "db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:12:2c:37:4d:bb:7a:bd"
            }
        },
        "ssl_cert": {
            "status": "UPLOADED",
            "info": {
                "issuer": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "subject": {
                    "C": "JP",
                    "ST": "Tokyo",
                    "L": "Chiyoda-ku",
                    "O": "NTT Communications Corporation",
                    "CN": "example.com"
                },
                "not_before": "2023-11-09 06:20:55",
                "not_after": "2024-12-10 06:20:54",
                "key_algorithm": "RSA-4096",
                "serial": "d3:11:fe:4d:a3:71:4e:13",
                "fingerprint": "46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:88:9f:3d:a7:c5:42:b2"
            }
        },
        "ssl_key": {
            "status": "UPLOADED",
            "info": {
                "key_algorithm": "RSA-4096",
                "passphrase": true
            }
        }
    }
}`)

func updateResult() *certificates.Certificate {
	var certificate certificates.Certificate

	var sslKeyInfo map[string]interface{}
	sslKeyInfoJson := `{"key_algorithm":"RSA-4096","passphrase":true}`
	err := json.Unmarshal([]byte(sslKeyInfoJson), &sslKeyInfo)
	if err != nil {
		panic(err)
	}

	sslKey := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   sslKeyInfo,
	}

	var sslCertInfo map[string]interface{}
	sslCertInfoJson := `{"issuer":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"example` +
		`.com"},"subject":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"` +
		`example.com"},"not_before":"2023-11-09 06:20:55","not_after":"2024-12-10 06:20:54","key_algorithm":"` +
		`RSA-4096","serial":"d3:11:fe:4d:a3:71:4e:13","fingerprint":"46:06:c5:ed:f0:e6:9f:c5:e3:bd:06:63:54:8` +
		`8:9f:3d:a7:c5:42:b2"}`
	err = json.Unmarshal([]byte(sslCertInfoJson), &sslCertInfo)
	if err != nil {
		panic(err)
	}

	sslCert := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   sslCertInfo,
	}

	var caCertInfo map[string]interface{}
	caCertInfoJson := `{"issuer":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"example` +
		`.com"},"subject":{"C":"JP","ST":"Tokyo","L":"Chiyoda-ku","O":"NTT Communications Corporation","CN":"` +
		`example.com"},"not_before":"2023-11-09 06:20:55","not_after":"2024-12-10 06:20:54","key_algorithm":"` +
		`RSA-4096","serial":"e7:61:4a:49:85:aa:7c:f2","fingerprint":"db:b1:49:84:f6:2e:ec:c9:41:fc:a1:30:26:1` +
		`2:2c:37:4d:bb:7a:bd"}`
	err = json.Unmarshal([]byte(caCertInfoJson), &caCertInfo)
	if err != nil {
		panic(err)
	}

	caCert := certificates.FileInResponse{
		Status: "UPLOADED",
		Info:   caCertInfo,
	}

	var tags map[string]interface{}
	tagsJson := `{"key":"value"}`
	err = json.Unmarshal([]byte(tagsJson), &tags)
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
    "type": "ssl-key",
    "content": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCjAxMjM0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEKMjM0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMwo0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1CjY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1NjcKODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OQpBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCCkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0QKRUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRgpHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdICklKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUoKS0xNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTApNTk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OCk9QUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1AKUVJTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUgpTVFVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUClVWV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVYKV1hZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWApZWmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaCmFiY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaYWIKY2RlZmdoaWprbG1ub3BxcnN0dXZ3eHl6MDEyMzQ1Njc4OUFCQ0RFRkdISUpLTE1OT1BRUlNUVVZXWFlaYWJjZAplZmdoaWprbG1ub3BxcnN0dXZ3eHl6VjAxMjM0NTY3ODlBQkNERUZHSElKS0xNTk9QUVJTVFVWV1hZWmFiY2RlCmZnaGlqa2xtbm9wcXJzdHV2d3h5ejAxMgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==",
    "passphrase": "passphrase"
}`)
