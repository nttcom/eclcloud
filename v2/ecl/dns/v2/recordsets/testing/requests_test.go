package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nttcom/eclcloud/ecl/dns/v2/recordsets"
	"github.com/nttcom/eclcloud/pagination"

	th "github.com/nttcom/eclcloud/testhelper"
	fakeclient "github.com/nttcom/eclcloud/testhelper/client"
)

func TestListDNSRecordSet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	prepareMuxForListResponse(t)

	count := 0
	err := recordsets.ListByZone(fakeclient.ServiceClient(), idZone, nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := recordsets.ExtractRecordSets(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedRecordSetSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListDNSRecordSetLimited(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	prepareMuxForListResponse(t)

	count := 0
	listOpts := recordsets.ListOpts{
		Limit:  1,
		Marker: idRecordSet1,
	}
	err := recordsets.ListByZone(fakeclient.ServiceClient(), idZone, listOpts).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := recordsets.ExtractRecordSets(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedRecordSetSliceLimited, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestListDNSRecordSetAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	prepareMuxForListResponse(t)

	allPages, err := recordsets.ListByZone(fakeclient.ServiceClient(), idZone, nil).AllPages()
	th.AssertNoErr(t, err)
	allRecordSets, err := recordsets.ExtractRecordSets(allPages)
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 2, len(allRecordSets))
}

func prepareMuxForListResponse(t *testing.T) {
	url := fmt.Sprintf("/zones/%s/recordsets", idZone)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		r.ParseForm()

		marker := r.Form.Get("marker")
		switch marker {
		case idRecordSet1:
			fmt.Fprintf(w, ListResponseLimited)
		case "":
			fmt.Fprintf(w, ListResponse)
		}
	})
}

func TestGetDNSRecordSet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s/recordsets/%s", idZone, idRecordSet1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, GetResponse)
	})

	actual, err := recordsets.Get(fakeclient.ServiceClient(), idZone, idRecordSet1).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &FirstRecordSet, actual)
}

func TestNextPageURL(t *testing.T) {
	var page recordsets.RecordSetPage
	var body map[string]interface{}
	err := json.Unmarshal([]byte(NextPageRequest), &body)
	if err != nil {
		t.Fatalf("Error unmarshaling data into page body: %v", err)
	}
	page.Body = body
	expected := nextURL
	actual, err := page.NextPageURL()
	th.AssertNoErr(t, err)
	th.CheckEquals(t, expected, actual)
}

func TestCreateDNSRecordSet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s/recordsets", idZone)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, CreateRequest)

		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, CreateResponse)
	})

	createOpts := recordsets.CreateOpts{
		Name:        nameRecordSet1,
		Type:        "A",
		TTL:         TTLRecordSet1,
		Description: descriptionRecordSet1,
		Records:     []string{ipRecordSet1},
	}

	// Clone FirstRecord into CreatedRecordSet(Created result struct)
	CreatedRecordSet := FirstRecordSet
	CreatedRecordSet.CreatedAt = time.Time{}
	CreatedRecordSet.UpdatedAt = time.Time{}
	CreatedRecordSet.Status = ""

	actual, err := recordsets.Create(fakeclient.ServiceClient(), idZone, createOpts).ExtractCreatedRecordSet()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &CreatedRecordSet, actual)
}

func TestUpdateDNSRecordSet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s/recordsets/%s", idZone, idRecordSet1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestJSONRequest(t, r, UpdateRequest)

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, UpdateResponse)
	})

	name := nameRecordSet1Update
	ttl := TTLRecordSet1Update
	description := descriptionRecordSet1Update
	records := []string{ipRecordSet1Update}

	updateOpts := recordsets.UpdateOpts{
		Name:        &name,
		TTL:         &ttl,
		Description: &description,
		Records:     &records,
	}

	actual, err := recordsets.Update(
		fakeclient.ServiceClient(), idZone, idRecordSet1, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &UpdatedRecordSet, actual)
}

func TestDelete(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/zones/%s/recordsets/%s", idZone, idRecordSet1)
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.WriteHeader(http.StatusNoContent)
	})

	rs := recordsets.Delete(fakeclient.ServiceClient(), idZone, idRecordSet1)
	th.AssertNoErr(t, rs.Err)
}
