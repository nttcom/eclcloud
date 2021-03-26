package testing

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/nttcom/eclcloud/v2/ecl/baremetal/v2/servers"
	"github.com/nttcom/eclcloud/v2/pagination"
	th "github.com/nttcom/eclcloud/v2/testhelper"
	fakeclient "github.com/nttcom/eclcloud/v2/testhelper/client"
)

func TestListServers(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers/detail", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, listResponse)
	})

	count := 0
	err := servers.List(fakeclient.ServiceClient(), nil).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := servers.ExtractServers(page)
		th.AssertNoErr(t, err)
		fmt.Printf("person[%%#v] -> %#v\n", actual)
		th.CheckDeepEquals(t, expectedServers, actual)
		return true, nil
	})

	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestGetServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/servers/%s", "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79")
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "GET")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)

		w.Header().Add("Content-Type", "application/json")
		fmt.Fprintf(w, getResponse)
	})

	actual, err := servers.Get(fakeclient.ServiceClient(), "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &expectedServer1, actual)
}

func TestCreateServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		th.TestHeader(t, r, "Content-Type", "application/json")
		th.TestHeader(t, r, "Accept", "application/json")
		th.TestJSONRequest(t, r, createRequest)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, createResponse)
	})

	createOpts := servers.CreateOpts{
		Name: "server-test-1",
		Networks: []servers.CreateOptsNetwork{
			{
				UUID:    "d32019d3-bc6e-4319-9c1d-6722fc136a22",
				FixedIP: "10.0.0.100",
			},
		},
		AdminPass:        "aabbccddeeff",
		ImageRef:         "b5660a6e-4b46-4be3-9707-6b47221b454f",
		FlavorRef:        "05184ba3-00ba-4fbc-b7a2-03b62b884931",
		AvailabilityZone: "zone1-groupa",
		UserData:         []byte("user_data"),
		RaidArrays: []servers.CreateOptsRaidArray{
			{
				PrimaryStorage: true,
				Partitions: []servers.CreateOptsPartition{
					{
						LVM:            true,
						PartitionLabel: "primary-part1",
					},
					{
						Size:           "100G",
						PartitionLabel: "var",
					},
				},
			},
			{
				RaidCardHardwareID: "raid_card_uuid",
				DiskHardwareIDs: []string{
					"disk1_uuid",
					"disk2_uuid",
					"disk3_uuid",
					"disk4_uuid",
				},
				Partitions: []servers.CreateOptsPartition{
					{
						LVM:            true,
						PartitionLabel: "secondary-part1",
					},
				},
				RaidLevel: 10,
			},
		},
		LVMVolumeGroups: []servers.CreateOptsLVMVolumeGroup{
			{
				VGLabel: "VG_root",
				PhysicalVolumePartitionLabels: []string{
					"primary-part1",
					"secondary-part1",
				},
				LogicalVolumes: []servers.CreateOptsLogicalVolume{
					{
						Size:    "300G",
						LVLabel: "LV_root",
					},
					{
						Size:    "2G",
						LVLabel: "LV_swap",
					},
				},
			},
		},
		Filesystems: []servers.CreateOptsFilesystem{
			{
				Label:      "LV_root",
				FSType:     "xfs",
				MountPoint: "/",
			},
			{
				Label:      "var",
				FSType:     "xfs",
				MountPoint: "/var",
			},
			{
				Label:  "LV_swap",
				FSType: "swap",
			},
		},
		Metadata: map[string]string{
			"foo": "bar",
		},
	}
	server, err := servers.Create(fakeclient.ServiceClient(), createOpts).Extract()
	th.AssertNoErr(t, err)

	th.AssertEquals(t, server.AdminPass, "aabbccddeeff")
}

func TestDeleteServer(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	url := fmt.Sprintf("/servers/%s", "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79")
	th.Mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "DELETE")
		th.TestHeader(t, r, "X-Auth-Token", fakeclient.TokenID)
		w.WriteHeader(http.StatusNoContent)
	})

	res := servers.Delete(fakeclient.ServiceClient(), "cebf8bb5-74cf-4a53-bca5-b90d4bbe8d79")
	th.AssertNoErr(t, res.Err)
}
