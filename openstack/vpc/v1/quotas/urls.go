package quotas

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
)

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("quotas")
}
