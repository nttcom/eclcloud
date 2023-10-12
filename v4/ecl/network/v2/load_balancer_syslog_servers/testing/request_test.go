package testing

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/nttcom/eclcloud/v4/ecl/network/v2/common"
	"github.com/nttcom/eclcloud/v4/ecl/network/v2/load_balancer_syslog_servers"
	"github.com/nttcom/eclcloud/v4/pagination"
	th "github.com/nttcom/eclcloud/v4/testhelper"
)

func TestListLoadBalancerSyslogServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()
	count := 0

	load_balancer_syslog_servers.List(client, load_balancer_syslog_servers.ListOpts{}).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := load_balancer_syslog_servers.ExtractLoadBalancerSyslogServers(page)
		if err != nil {
			t.Errorf("Failed to extract Load Balancer Syslog Servers: %v", err)
			return false, nil
		}

		th.CheckDeepEquals(t, ExpectedLoadBalancerSlice, actual)

		return true, nil
	})

	if count != 1 {
		t.Errorf("Expected 1 page, got %d", count)
	}
}

func TestGetLoadBalancerSyslogServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers/6e9c7745-61f2-491f-9689-add8c5fc4b9a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, GetResponse)
	})

	s, err := load_balancer_syslog_servers.Get(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LoadBalancerSyslogServerDetail, s)
}

func TestCreateLoadBalancerSyslogServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, CreateRequest)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		fmt.Fprintf(w, CreateResponse)
	})

	priority := 20

	options := load_balancer_syslog_servers.CreateOpts{
		AclLogging:                  "DISABLED",
		AppflowLogging:              "DISABLED",
		DateFormat:                  "MMDDYYYY",
		Description:                 "test",
		IPAddress:                   "120.120.120.30",
		LoadBalancerID:              "9f872504-36ab-46af-83ce-a4991c669edd",
		LogFacility:                 "LOCAL3",
		LogLevel:                    "DEBUG",
		Name:                        "first_syslog_server",
		PortNumber:                  514,
		Priority:                    &priority,
		TcpLogging:                  "ALL",
		TenantID:                    "6a156ddf2ecd497ca786ff2da6df5aa8",
		TimeZone:                    "LOCAL_TIME",
		TransportType:               "UDP",
		UserConfigurableLogMessages: "NO",
	}
	s, err := load_balancer_syslog_servers.Create(fake.ServiceClient(), options).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &LoadBalancerSyslogServerDetail, s)
}

func TestRequiredCreateOptsLoadBalancerSyslogServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	res := load_balancer_syslog_servers.Create(fake.ServiceClient(), load_balancer_syslog_servers.CreateOpts{})
	if res.Err == nil {
		t.Fatalf("Expected error, got none")
	}
}

func TestUpdateLoadBalancerSyslogServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers/6e9c7745-61f2-491f-9689-add8c5fc4b9a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "PUT")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, UpdateRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, UpdateResponse)
	})

	aclLogging := "DISABLED"
	appflowLogging := "DISABLED"
	dateFormat := "MMDDYYYY"
	description := "test2"
	logFacility := "LOCAL3"
	logLevel := "DEBUG"
	priority := 20
	tcpLogging := "ALL"
	timeZone := "LOCAL_TIME"
	userConfigurableLogMessages := "NO"

	id := "6e9c7745-61f2-491f-9689-add8c5fc4b9a"

	ipAddress := "120.120.120.30"
	loadBalancerID := "9f872504-36ab-46af-83ce-a4991c669edd"
	name := "first_syslog_server"
	portNumber := 514
	status := "PENDING_UPDATE"
	tenantID := "6a156ddf2ecd497ca786ff2da6df5aa8"
	transportType := "UDP"

	options := load_balancer_syslog_servers.UpdateOpts{
		AclLogging:                  aclLogging,
		AppflowLogging:              appflowLogging,
		DateFormat:                  dateFormat,
		Description:                 &description,
		LogFacility:                 logFacility,
		LogLevel:                    logLevel,
		Priority:                    &priority,
		TcpLogging:                  tcpLogging,
		TimeZone:                    timeZone,
		UserConfigurableLogMessages: userConfigurableLogMessages,
	}

	s, err := load_balancer_syslog_servers.Update(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a", options).Extract()
	th.AssertNoErr(t, err)

	th.CheckEquals(t, aclLogging, s.AclLogging)
	th.CheckEquals(t, appflowLogging, s.AppflowLogging)
	th.CheckEquals(t, dateFormat, s.DateFormat)
	th.CheckEquals(t, description, s.Description)
	th.CheckEquals(t, id, s.ID)
	th.CheckEquals(t, logFacility, s.LogFacility)
	th.CheckEquals(t, logLevel, s.LogLevel)
	th.CheckEquals(t, priority, s.Priority)
	th.CheckEquals(t, tcpLogging, s.TcpLogging)
	th.CheckEquals(t, timeZone, s.TimeZone)
	th.CheckEquals(t, userConfigurableLogMessages, s.UserConfigurableLogMessages)
	th.CheckEquals(t, ipAddress, s.IPAddress)
	th.CheckEquals(t, loadBalancerID, s.LoadBalancerID)
	th.CheckEquals(t, name, s.Name)
	th.CheckEquals(t, portNumber, s.PortNumber)
	th.CheckEquals(t, status, s.Status)
	th.CheckEquals(t, tenantID, s.TenantID)
	th.CheckEquals(t, transportType, s.TransportType)

}

func TestDeleteLoadBalancerSyslogServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers/6e9c7745-61f2-491f-9689-add8c5fc4b9a", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := load_balancer_syslog_servers.Delete(fake.ServiceClient(), "6e9c7745-61f2-491f-9689-add8c5fc4b9a")
	th.AssertNoErr(t, res.Err)
}

func TestIDFromName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	expectedID := "6e9c7745-61f2-491f-9689-add8c5fc4b9a"
	actualID, err := load_balancer_syslog_servers.IDFromName(client, "first_syslog_server")

	th.AssertNoErr(t, err)
	th.AssertEquals(t, expectedID, actualID)
}

func TestIDFromNameNoResult(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponse)
	})

	client := fake.ServiceClient()

	_, err := load_balancer_syslog_servers.IDFromName(client, "syslog_server X")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}

}

func TestIDFromNameDuplicated(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/load_balancer_syslog_servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fake.TokenID)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, ListResponseDuplicatedNames)
	})

	client := fake.ServiceClient()

	_, err := load_balancer_syslog_servers.IDFromName(client, "first_syslog_server")

	if err == nil {
		t.Fatalf("Expected error, got none")
	}
}
