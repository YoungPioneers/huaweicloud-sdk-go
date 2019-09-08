package logs

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
)

func ListURL(c *gophercloud.ServiceClient, scalingGroupId string) string {
	return c.ServiceURL("scaling_activity_log", scalingGroupId)
}
