package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/v3"
	"github.com/nttcom/eclcloud/v3/ecl/sss/v2/users"
)

const contractID = "econ8000008888"

const idUser1 = "ecid1000000001"
const idUser2 = "ecid1000000002"

const startTime = "2018-07-26 08:40:01"

var listResponse = fmt.Sprintf(`
{
	"contract_id": "%s",
	"users": [{
		"user_id": "%s",
		"login_id": "login_id_1",
		"mail_address": "user1@example.com",
		"start_time": "%s",
		"api_availability": true,
		"contract_owner": true,
		"super_user": true
	}, {
		"user_id": "%s",
		"login_id": "login_id_2",
		"mail_address": "user2@example.com",
		"start_time": "%s",
		"api_availability": true,
		"contract_owner": true,
		"super_user": true
	}]
}`, contractID,
	idUser1, startTime,
	idUser2, startTime)

var expectedUsersSlice = []users.User{firstUser, secondUser}

var userStartTime, _ = time.Parse(eclcloud.RFC3339ZNoTNoZ, startTime)

var firstUser = users.User{
	UserID:          idUser1,
	LoginID:         "login_id_1",
	MailAddress:     "user1@example.com",
	ContractID:      contractID,
	StartTime:       userStartTime,
	ApiAvailability: true,
	ContractOwner:   true,
	Superuser:       true,
}

var secondUser = users.User{
	UserID:          idUser2,
	LoginID:         "login_id_2",
	MailAddress:     "user2@example.com",
	ContractID:      contractID,
	StartTime:       userStartTime,
	ApiAvailability: true,
	ContractOwner:   true,
	Superuser:       true,
}

var getResponse = fmt.Sprintf(`
{
	"user_id": "%s",
	"login_id": "login_id_1",
	"mail_address": "user1@example.com",
	"contract_owner": false,
	"super_user": false,
	"api_availability": true,
	"sss_endpoint": "http://sss.com",
	"keystone_endpoint": "http://keystone.com",
	"keystone_name": "keystonename1",
	"keystone_password": "keystonepassword1",
	"start_time": "%s",
	"contract_id": "%s",
	"login_integration": "",
	"external_reference_id": "econ0000009999",
	"brand_id": "ecl2",
	"auto_role_assignment_flag": false,
	"external_user_type": "iop",
	"otp_activation": false
}`, idUser1,
	startTime,
	contractID,
)

var getResponseStruct = users.User{
	UserID:                 idUser1,
	LoginID:                "login_id_1",
	MailAddress:            "user1@example.com",
	ContractOwner:          false,
	Superuser:              false,
	ApiAvailability:        false,
	SSSEndpoint:            "http://sss.com",
	KeystoneEndpoint:       "http://keystone.com",
	KeystoneName:           "keystonename1",
	KeystonePassword:       "keystonepassword1",
	StartTime:              userStartTime,
	ContractID:             contractID,
	LoginIntegration:       "",
	ExternalReferenceID:    "econ0000009999",
	BrandID:                "ecl2",
	AutoRoleAssignmentFlag: false,
	ExternalUserType:       "iop",
	OtpActivation:          false,
}

var createRequest = `{
	"login_id": "login_id_1",
	"mail_address": "user1@example.com",
	"notify_password": "false",
	"password": "Passw0rd"
}`

var createResponse = fmt.Sprintf(`{
	"login_id": "login_id_1",
	"mail_address": "user1@example.com",
	"user_id": "%s",
	"contract_id": "%s",
	"keystone_endpoint": "http://keystone.com",
	"sss_endpoint": "http://sss.com",
	"password": "Passw0rd"
}
`, idUser1,
	contractID,
)

var createdUser = users.User{
	LoginID:          "login_id_1",
	UserID:           idUser1,
	ContractID:       contractID,
	MailAddress:      "user1@example.com",
	KeystoneEndpoint: "http://keystone.com",
	SSSEndpoint:      "http://sss.com",
	Password:         "Passw0rd",
}

var updateRequest = `{
	"login_id": "login_id_1_update",
	"mail_address": "user1_update@example.com",
	"new_password": "NewPassw0rd"
}`
