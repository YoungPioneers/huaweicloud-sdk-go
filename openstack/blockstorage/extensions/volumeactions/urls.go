package volumeactions

import "github.com/huaweicloud/huaweicloud-sdk-go"

func actionURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("volumes", id, "action")
}
