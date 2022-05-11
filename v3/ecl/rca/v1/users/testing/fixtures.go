package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v3/ecl/rca/v1/users"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	"github.com/nttcom/eclcloud/v3/testhelper/client"
)

var (
	password        = "dummy_passw@rd"
	passwordUpdated = "dummy_passw@rd_updated"
)

// ListResult provides a single page of user results.
const ListResult = `
{
	"users": [
		{
			"name": "ef5778e553a24d789c15c689e30adf5d",
			"vpn_endpoints": [
				{
					"endpoint": "https://rca-sslvpn1-jp1.ecl.ntt.com",
					"type": "SSL-VPN"
				}
			]
		},
		{
			"name": "8bbe05d4bec747189e0dab81e486969f-1005",
			"vpn_endpoints": [
				{
					"endpoint": "https://rca-sslvpn1-jp1.ecl.ntt.com",
					"type": "SSL-VPN"
				}
			]
		}
	]
}
`

// GetResult provides a Get result.
const GetResult = `
{
	"user": {
		"name": "8bbe05d4bec747189e0dab81e486969f-1005",
		"vpn_endpoints": [
			{
				"endpoint": "https://rca-sslvpn1-jp1.ecl.ntt.com",
				"type": "SSL-VPN"
			}
		]
	}
}
`

// CreateRequest provides the input to a Create request.
const CreateRequest = `
{
	"user": {
		"password": "dummy_passw@rd"
	}
}
`

// CreateResponse provides the output from a Create request.
const CreateResponse = `
{
	"user": {
		"name": "8bbe05d4bec747189e0dab81e486969f-1005",
		"password": "dummy_passw@rd",
		"vpn_endpoints": [
			{
				"endpoint": "https://rca-sslvpn1-jp1.ecl.ntt.com",
				"type": "SSL-VPN"
			}
		]
	}
}
`

// UpdateRequest provides the input to as Update request.
const UpdateRequest = `
{
	"user": {
		"password": "dummy_passw@rd_updated"
	}
}
`

// UpdateResult provides an update result.
const UpdateResult = GetResult

// FirstUser is the first user in the List request.
var FirstUser = users.User{
	Name: "ef5778e553a24d789c15c689e30adf5d",
	VPNEndpoints: []users.VPNEndpoint{
		{
			Endpoint: "https://rca-sslvpn1-jp1.ecl.ntt.com",
			Type:     "SSL-VPN",
		},
	},
}

// SecondUser is the second user in the List request.
var SecondUser = users.User{
	Name: "8bbe05d4bec747189e0dab81e486969f-1005",
	VPNEndpoints: []users.VPNEndpoint{
		{
			Endpoint: "https://rca-sslvpn1-jp1.ecl.ntt.com",
			Type:     "SSL-VPN",
		},
	},
}

// SecondUserUpdated is how SecondUser should look after an Update.
var SecondUserUpdated = users.User{
	Name: "8bbe05d4bec747189e0dab81e486969f-1005",
	VPNEndpoints: []users.VPNEndpoint{
		{
			Endpoint: "https://rca-sslvpn1-jp1.ecl.ntt.com",
			Type:     "SSL-VPN",
		},
	},
}

// ExpectedUsersSlice is the slice of users expected to be returned from ListResult.
var ExpectedUsersSlice = []users.User{FirstUser, SecondUser}

// HandleListUsersSuccessfully creates an HTTP handler at `/users` on the
// test handler mux that responds with a list of two users.
func HandleListUsersSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, ListResult)
	})
}

// HandleGetUserSuccessfully creates an HTTP handler at `/users` on the
// test handler mux that responds with a single user.
func HandleGetUserSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/users/%s", SecondUser.Name), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, GetResult)
	})
}

// HandleCreateUserSuccessfully creates an HTTP handler at `/users` on the
// test handler mux that tests user creation.
func HandleCreateUserSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, CreateResponse)
	})
}

// HandleDeleteUserSuccessfully creates an HTTP handler at `/users` on the
// test handler mux that tests user deletion.
func HandleDeleteUserSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/users/%s", FirstUser.Name), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)

		w.WriteHeader(http.StatusOK)
	})
}

// HandleUpdateUserSuccessfully creates an HTTP handler at `/users` on the
// test handler mux that tests user update.
func HandleUpdateUserSuccessfully(t *testing.T) {
	th.Mux.HandleFunc(fmt.Sprintf("/users/%s", SecondUser.Name), func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, UpdateResult)
	})
}
