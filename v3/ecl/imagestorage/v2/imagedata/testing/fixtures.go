package testing

import (
	"io/ioutil"
	"net/http"
	"testing"

	th "github.com/nttcom/eclcloud/v3/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v3/testhelper/client"
)

// HandlePutImageDataSuccessfully setup
func HandlePutImageDataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/0cb9328d-dd8c-41bb-b378-404b854b93b9/file", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Unable to read request body: %v", err)
		}

		th.AssertByteArrayEquals(t, []byte{5, 3, 7, 24}, b)

		w.WriteHeader(http.StatusNoContent)
	})
}

// HandleGetImageDataSuccessfully setup
func HandleGetImageDataSuccessfully(t *testing.T) {
	th.Mux.HandleFunc("/images/100f4d2d-dcb5-472e-b93f-b4e13d888604/file", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusOK)

		_, err := w.Write([]byte{34, 87, 0, 23, 23, 23, 56, 255, 254, 0})
		th.AssertNoErr(t, err)
	})
}
