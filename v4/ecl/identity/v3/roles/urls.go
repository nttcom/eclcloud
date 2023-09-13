package roles

import "github.com/nttcom/eclcloud/v4"

const (
	rolePath = "roles"
)

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func getURL(client *eclcloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func updateURL(client *eclcloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func deleteURL(client *eclcloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func listAssignmentsURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

func listAssignmentsOnResourceURL(client *eclcloud.ServiceClient, targetType, targetID, actorType, actorID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath)
}

func assignURL(client *eclcloud.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath, roleID)
}
