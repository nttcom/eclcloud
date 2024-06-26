package testing

import (
	"fmt"

	"github.com/nttcom/eclcloud/v4/ecl/managed_load_balancer/v1/tls_policies"
)

const id = "497f6eca-6276-4993-bfeb-53cbbbba6f08"

var listResponse = fmt.Sprintf(`
{
    "tls_policies": [
        {
            "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
            "name": "TLSv1.2_202210_01",
            "description": "description",
            "default": true,
            "tls_protocols": [
                "TLSv1.2",
                "TLSv1.3"
            ],
            "tls_ciphers": [
                "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
                "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
                "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
                "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
                "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
                "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
                "TLS_AES_256_GCM_SHA384",
                "TLS_CHACHA20_POLY1305_SHA256",
                "TLS_AES_128_GCM_SHA256"
            ]
        }
    ]
}`)

func listResult() []tls_policies.TLSPolicy {
	var tLSPolicy1 tls_policies.TLSPolicy

	tLSPolicy1.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	tLSPolicy1.Name = "TLSv1.2_202210_01"
	tLSPolicy1.Description = "description"
	tLSPolicy1.Default = true
	tLSPolicy1.TLSProtocols = []string{"TLSv1.2", "TLSv1.3"}
	tLSPolicy1.TLSCiphers = []string{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256", "TLS_AES_256_GCM_SHA384", "TLS_CHACHA20_POLY1305_SHA256", "TLS_AES_128_GCM_SHA256"}

	return []tls_policies.TLSPolicy{tLSPolicy1}
}

var showResponse = fmt.Sprintf(`
{
    "tls_policy": {
        "id": "497f6eca-6276-4993-bfeb-53cbbbba6f08",
        "name": "TLSv1.2_202210_01",
        "description": "description",
        "default": true,
        "tls_protocols": [
            "TLSv1.2",
            "TLSv1.3"
        ],
        "tls_ciphers": [
            "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
            "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
            "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
            "TLS_AES_256_GCM_SHA384",
            "TLS_CHACHA20_POLY1305_SHA256",
            "TLS_AES_128_GCM_SHA256"
        ]
    }
}`)

func showResult() *tls_policies.TLSPolicy {
	var tLSPolicy tls_policies.TLSPolicy

	tLSPolicy.ID = "497f6eca-6276-4993-bfeb-53cbbbba6f08"
	tLSPolicy.Name = "TLSv1.2_202210_01"
	tLSPolicy.Description = "description"
	tLSPolicy.Default = true
	tLSPolicy.TLSProtocols = []string{"TLSv1.2", "TLSv1.3"}
	tLSPolicy.TLSCiphers = []string{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256", "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384", "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256", "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256", "TLS_AES_256_GCM_SHA384", "TLS_CHACHA20_POLY1305_SHA256", "TLS_AES_128_GCM_SHA256"}

	return &tLSPolicy
}
