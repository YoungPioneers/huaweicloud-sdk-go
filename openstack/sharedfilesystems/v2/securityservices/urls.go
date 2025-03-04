package securityservices

import "github.com/huaweicloud/huaweicloud-sdk-go"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-services")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("security-services", id)
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-services", "detail")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}
