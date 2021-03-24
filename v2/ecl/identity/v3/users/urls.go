package users

import "github.com/nttcom/eclcloud/v2"

func listURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func getURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func createURL(client *eclcloud.ServiceClient) string {
	return client.ServiceURL("users")
}

func updateURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func changePasswordURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "password")
}

func deleteURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID)
}

func listGroupsURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "groups")
}

func addToGroupURL(client *eclcloud.ServiceClient, groupID, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func isMemberOfGroupURL(client *eclcloud.ServiceClient, groupID, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func removeFromGroupURL(client *eclcloud.ServiceClient, groupID, userID string) string {
	return client.ServiceURL("groups", groupID, "users", userID)
}

func listProjectsURL(client *eclcloud.ServiceClient, userID string) string {
	return client.ServiceURL("users", userID, "projects")
}

func listInGroupURL(client *eclcloud.ServiceClient, groupID string) string {
	return client.ServiceURL("groups", groupID, "users")
}
