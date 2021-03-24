package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/nttcom/eclcloud/v2/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v2/testhelper/client"
)

// HandleCreateImageMemberSuccessfully setup
func HandleCreateImageMemberSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/54d63e39-4ee1-4a62-8704-0ae5025a0deb/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `{"member": "f6a818c3d4aa458798ed86892e7150c0"}`)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		    "created_at": "2013-09-20T19:22:19Z",
		    "image_id": "54d63e39-4ee1-4a62-8704-0ae5025a0deb",
		    "member_id": "f6a818c3d4aa458798ed86892e7150c0",
		    "schema": "/v2/schemas/member",
		    "status": "pending",
		    "updated_at": "2013-09-20T19:25:31Z"
			}`)

	})
}

// HandleImageMemberList happy path setup
func HandleImageMemberList(t *testing.T) {
	th.Mux.HandleFunc("/images/54d63e39-4ee1-4a62-8704-0ae5025a0deb/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
		    "members": [
		        {
		            "created_at": "2013-10-07T17:58:03Z",
		            "image_id": "54d63e39-4ee1-4a62-8704-0ae5025a0deb",
		            "member_id": "f6a818c3d4aa458798ed86892e7150c0",
		            "schema": "/v2/schemas/member",
		            "status": "pending",
		            "updated_at": "2013-10-07T17:58:03Z"
		        },
		        {
		            "created_at": "2013-10-07T17:58:55Z",
		            "image_id": "54d63e39-4ee1-4a62-8704-0ae5025a0deb",
		            "member_id": "1efb79fe4437490aab966b57da5b9f05",
		            "schema": "/v2/schemas/member",
		            "status": "accepted",
		            "updated_at": "2013-10-08T12:08:55Z"
		        }
		    ],
		    "schema": "/v2/schemas/members"
		}`)
	})
}

// HandleImageMemberEmptyList happy path setup
func HandleImageMemberEmptyList(t *testing.T) {
	th.Mux.HandleFunc("/images/54d63e39-4ee1-4a62-8704-0ae5025a0deb/members", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, `{
		    "members": [],
		    "schema": "/v2/schemas/members"
		}`)
	})
}

// HandleImageMemberDetails setup
func HandleImageMemberDetails(t *testing.T) {
	th.Mux.HandleFunc("/images/54d63e39-4ee1-4a62-8704-0ae5025a0deb/members/f6a818c3d4aa458798ed86892e7150c0", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{
		    "status": "pending",
		    "created_at": "2013-11-26T07:21:21Z",
		    "updated_at": "2013-11-26T07:21:21Z",
		    "image_id": "54d63e39-4ee1-4a62-8704-0ae5025a0deb",
		    "member_id": "f6a818c3d4aa458798ed86892e7150c0",
		    "schema": "/v2/schemas/member"
		}`)
	})
}

// HandleImageMemberDeleteSuccessfully setup
func HandleImageMemberDeleteSuccessfully(t *testing.T) *CallsCounter {
	var counter CallsCounter
	th.Mux.HandleFunc("/images/54d63e39-4ee1-4a62-8704-0ae5025a0deb/members/f6a818c3d4aa458798ed86892e7150c0", func(w http.ResponseWriter, r *http.Request) {
		counter.Counter = counter.Counter + 1

		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})
	return &counter
}

// HandleImageMemberUpdate setup
func HandleImageMemberUpdate(t *testing.T) *CallsCounter {
	var counter CallsCounter
	th.Mux.HandleFunc("/images/54d63e39-4ee1-4a62-8704-0ae5025a0deb/members/f6a818c3d4aa458798ed86892e7150c0", func(w http.ResponseWriter, r *http.Request) {
		counter.Counter = counter.Counter + 1

		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		th.TestJSONRequest(t, r, `{"status": "accepted"}`)

		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, `{
		    "status": "accepted",
		    "created_at": "2013-11-26T07:21:21Z",
		    "updated_at": "2013-11-26T07:21:21Z",
		    "image_id": "54d63e39-4ee1-4a62-8704-0ae5025a0deb",
		    "member_id": "f6a818c3d4aa458798ed86892e7150c0",
		    "schema": "/v2/schemas/member"
		}`)
	})
	return &counter
}

// CallsCounter for checking if request handler was called at all
type CallsCounter struct {
	Counter int
}
