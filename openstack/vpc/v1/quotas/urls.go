package quotas

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
