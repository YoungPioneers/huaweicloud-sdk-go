package roles

import "github.com/huaweicloud/huaweicloud-sdk-go"

const (
	rolePath = "roles"
)

func listURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func getURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func createURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL(rolePath)
}

func updateURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func deleteURL(client *gophercloud.ServiceClient, roleID string) string {
	return client.ServiceURL(rolePath, roleID)
}

func listAssignmentsURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("role_assignments")
}

func assignURL(client *gophercloud.ServiceClient, targetType, targetID, actorType, actorID, roleID string) string {
	return client.ServiceURL(targetType, targetID, actorType, actorID, rolePath, roleID)
}
