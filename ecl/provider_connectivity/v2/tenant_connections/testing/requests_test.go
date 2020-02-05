package testing

import (
	"github.com/nttcom/eclcloud/ecl/provider_connectivity/v2/tenant_connections"
	"github.com/nttcom/eclcloud/pagination"
	th "github.com/nttcom/eclcloud/testhelper"
	"github.com/nttcom/eclcloud/testhelper/client"
	"testing"
)

func TestListTenantConnections(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantConnectionsSuccessfully(t)

	count := 0
	err := tenant_connections.List(client.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++

		actual, err := tenant_connections.ExtractTenantConnections(page)
		th.AssertNoErr(t, err)

		th.AssertDeepEquals(t, ExpectedTenantConnectionsSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.AssertEquals(t, count, 1)
}

func TestListTenantConnectionsAllPages(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListTenantConnectionsSuccessfully(t)

	allPages, err := tenant_connections.List(client.ServiceClient(), nil).AllPages()
	th.AssertNoErr(t, err)
	actual, err := tenant_connections.ExtractTenantConnections(allPages)
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, ExpectedTenantConnectionsSlice, actual)
}

func TestGetTenantConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetTenantConnectionSuccessfully(t)

	actual, err := tenant_connections.Get(client.ServiceClient(), SecondTenantConnection.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnection, *actual)
}

func TestCreateTenantConnectionAttachServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateTenantConnectionAttachServerSuccessfully(t)

	createOpts := tenant_connections.CreateOpts{
		Name:                      "test_name_1",
		Description:               "test_desc_1",
		Tags:                      map[string]string{"test_tags1": "test1"},
		TenantConnectionRequestID: "21b344d8-be11-11e7-bf3c-0050569c850d",
		DeviceType:                "ECL::Compute::Server",
		DeviceID:                  "8c235a3b-8dee-41a1-b81a-64e06edc0986",
		DeviceInterfaceID:         "",
		AttachmentOpts: tenant_connections.Server{
			AllowedAddressPairs: []tenant_connections.AddressPair{
				{
					IPAddress:  "192.168.1.2",
					MACAddress: "11:22:33:aa:bb:cc",
				},
			},
			FixedIPs: []tenant_connections.ServerFixedIPs{
				{
					SubnetID:  "1f424165-2202-4022-ad70-0fa6f9ec99e1",
					IPAddress: "192.168.1.1",
				},
			},
		},
	}

	actual, err := tenant_connections.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &FirstTenantConnection, actual)
}

func TestCreateTenantConnectionAttachVna(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleCreateTenantConnectionAttachVnaSuccessfully(t)

	createOpts := tenant_connections.CreateOpts{
		Name:                      "attach_vna_name",
		Description:               "attach_vna_desc",
		Tags:                      map[string]string{"test_tags1": "test1"},
		TenantConnectionRequestID: "67d76b00-3804-11ea-8088-525400060300",
		DeviceType:                "ECL::VirtualNetworkAppliance::VSRX",
		DeviceID:                  "c291f4c4-a680-4db0-8b88-7e579f0aaa37",
		DeviceInterfaceID:         "interface_2",
		AttachmentOpts: tenant_connections.Vna{
			FixedIPs: []tenant_connections.VnaFixedIPs{
				{
					IPAddress: "192.168.1.3",
				},
			},
		},
	}

	actual, err := tenant_connections.Create(client.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, &CreateTenantConnectionAttachVna, actual)
}

func TestDeleteTenantConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleDeleteTenantConnectionSuccessfully(t)

	res := tenant_connections.Delete(client.ServiceClient(), FirstTenantConnection.ID)
	th.AssertNoErr(t, res.Err)
}

func TestUpdateTenantConnection(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleUpdateTenantConnectionSuccessfully(t)

	updateOpts := tenant_connections.UpdateOpts{
		Name:             "update_name",
		Description:      "update_desc",
		Tags:             map[string]string{"update_tags": "update"},
		NameOther:        "update_name_other",
		DescriptionOther: "update_desc_other",
		TagsOther:        map[string]string{"test_tags_other": "update"},
	}

	actual, err := tenant_connections.Update(client.ServiceClient(), SecondTenantConnection.ID, updateOpts).Extract()
	th.AssertNoErr(t, err)
	th.AssertDeepEquals(t, SecondTenantConnectionUpdated, *actual)
}
