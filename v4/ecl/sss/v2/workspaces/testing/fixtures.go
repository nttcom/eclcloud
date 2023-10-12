package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/v4"
	"github.com/nttcom/eclcloud/v4/ecl/sss/v2/workspaces"
)

const contractID = "econ0000000001"

const workspaceID1 = "ws0000000001"
const workspaceID2 = "ws0000000002"

const nameWorkspace1 = "jp1_workspace01"
const nameWorkspace2 = "jp1_workspace02"

const descriptionWorkspace1 = "jp1 workspace01"
const descriptionWorkspace2 = "jp1 workspace02"

const startTime = "2020-01-01 00:00:00"

var workspaceStartTime, _ = time.Parse(eclcloud.RFC3339ZNoTNoZ, startTime)

var listResponse = fmt.Sprintf(`
{
    "contract_id": "%s",
    "workspaces": [
        {
            "workspace_id": "%s",
            "workspace_name": "%s",
            "description": "%s",
            "start_time": "%s"
        },
        {
            "workspace_id": "%s",
            "workspace_name": "%s",
            "description": "%s",
            "start_time": "%s"
        }
    ]
}
`,
	contractID,
	workspaceID1, nameWorkspace1, descriptionWorkspace1, startTime,
	workspaceID2, nameWorkspace2, descriptionWorkspace2, startTime,
)

var firstWorkspace = workspaces.Workspace{
	ContractID:    contractID,
	WorkspaceID:   workspaceID1,
	WorkspaceName: nameWorkspace1,
	Description:   descriptionWorkspace1,
	StartTime:     workspaceStartTime,
}

var secondWorkspace = workspaces.Workspace{
	ContractID:    contractID,
	WorkspaceID:   workspaceID2,
	WorkspaceName: nameWorkspace2,
	Description:   descriptionWorkspace2,
	StartTime:     workspaceStartTime,
}

var expectedWorkspacesSlice = []workspaces.Workspace{firstWorkspace, secondWorkspace}

var getResponse = fmt.Sprintf(`
{
	"contract_id": "%s",
	"workspace_id": "%s",
	"workspace_name": "%s",
	"description": "%s",
	"start_time": "%s",
	"regions": [
		{
			"region_name": "jp1",
			"tenant_id": "9a76dca6d8cd4391aac6f2ea052f10f4"
		},
		{
			"region_name": "jp2",
			"tenant_id": "27a58d42769141ff8e94920a99aeb44b"
		}
	],
	"users": [
		{
			"user_id": "ecid000000001",
			"contract_id": "econ0000000001",
			"contract_owner": true
		},
		{
			"user_id": "ecid000000002",
			"contract_id": "econ0000000002",
			"contract_owner": false
		}
	]
}`,
	contractID, workspaceID1, nameWorkspace1, descriptionWorkspace1, startTime)

var getResponseStruct = workspaces.Workspace{
	ContractID:    contractID,
	WorkspaceID:   workspaceID1,
	WorkspaceName: nameWorkspace1,
	Description:   descriptionWorkspace1,
	StartTime:     workspaceStartTime,
	Regions: []workspaces.Region{
		{
			RegionName: "jp1",
			TenantID:   "9a76dca6d8cd4391aac6f2ea052f10f4",
		},
		{
			RegionName: "jp2",
			TenantID:   "27a58d42769141ff8e94920a99aeb44b",
		},
	},
	Users: []workspaces.User{
		{
			UserID:        "ecid000000001",
			ContractID:    "econ0000000001",
			ContractOwner: true,
		},
		{
			UserID:        "ecid000000002",
			ContractID:    "econ0000000002",
			ContractOwner: false,
		},
	},
}

var createRequest = `
{
  "workspace_name": "sample_workspace",
  "description": "sample workspace",
  "contract_id": "econ0000000001"
}
`

var createResponse = fmt.Sprintf(`
{
  "workspace_id": "%s",
  "workspace_name": "sample_workspace",
  "description": "sample workspace",
  "contract_id": "%s"
}
`, workspaceID1, contractID)

var createdWorkspace = workspaces.Workspace{
	ContractID:    contractID,
	WorkspaceID:   workspaceID1,
	WorkspaceName: "sample_workspace",
	Description:   "sample workspace",
}

var updateRequest = `
{
 "description": "updated workspace"
}
`
