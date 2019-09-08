package events

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}
