package buildinfo

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
