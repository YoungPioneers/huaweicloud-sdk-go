package limits

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
