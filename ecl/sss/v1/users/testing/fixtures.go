package testing

import (
	"fmt"
	"time"

	"github.com/nttcom/eclcloud/v2"
	"github.com/nttcom/eclcloud/v2/ecl/sss/v1/users"
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
		"contract_id": "%s",
		"start_time": "%s"
	}, {
		"user_id": "%s",
		"login_id": "login_id_2",
		"mail_address": "user2@example.com",
		"contract_id": "%s",
		"start_time": "%s"
	}]
}`, contractID,
	idUser1, contractID, startTime,
	idUser2, contractID, startTime)

var expectedUsersSlice = []users.User{firstUser, secondUser}

var userStartTime, _ = time.Parse(eclcloud.RFC3339ZNoTNoZ, startTime)

var firstUser = users.User{
	UserID:      idUser1,
	LoginID:     "login_id_1",
	MailAddress: "user1@example.com",
	ContractID:  contractID,
	StartTime:   userStartTime,
}

var secondUser = users.User{
	UserID:      idUser2,
	LoginID:     "login_id_2",
	MailAddress: "user2@example.com",
	ContractID:  contractID,
	StartTime:   userStartTime,
}

var getResponse = fmt.Sprintf(`
{
	"user_id": "%s",
	"login_id": "login_id_1",
	"mail_address": "user1@example.com",
	"contract_owner": false,
	"super_user": false,
	"sss_endpoint": "http://sss.com",
	"keystone_endpoint": "http://keystone.com",
	"keystone_name": "keystonename1",
	"keystone_password": "keystonepassword1",
	"start_time": "%s",
	"contract_id": "%s",
	"login_integration": "",
	"external_reference_id": "econ0000009999",
	"brand_id": "ecl2",
	"otp_activation": false
}`, idUser1,
	startTime,
	contractID,
)

var getResponseStruct = users.User{
	UserID:              idUser1,
	LoginID:             "login_id_1",
	MailAddress:         "user1@example.com",
	ContractOwner:       false,
	SSSEndpoint:         "http://sss.com",
	KeystoneEndpoint:    "http://keystone.com",
	KeystoneName:        "keystonename1",
	KeystonePassword:    "keystonepassword1",
	StartTime:           userStartTime,
	ContractID:          contractID,
	LoginIntegration:    "",
	ExternalReferenceID: "econ0000009999",
	BrandID:             "ecl2",
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
	"keystone_name": "keystonename1",
	"keystone_password": "keystonepassword1",
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
	KeystoneName:     "keystonename1",
	KeystonePassword: "keystonepassword1",
	KeystoneEndpoint: "http://keystone.com",
	SSSEndpoint:      "http://sss.com",
	Password:         "Passw0rd",
}

var updateRequest = `{
	"login_id": "login_id_1_update",
	"mail_address": "user1_update@example.com",
	"new_password": "NewPassw0rd"
}`
