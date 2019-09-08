package buildinfo

import "github.com/huaweicloud/huaweicloud-sdk-go"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("build_info")
}
