package services

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-services")
}
