package services

import "github.com/huaweicloud/huaweicloud-sdk-go"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
