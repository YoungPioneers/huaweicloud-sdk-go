package volumes

import "github.com/huaweicloud/huaweicloud-sdk-go"

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes")
}

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("volumes", "detail")
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return deleteURL(c, id)
}
