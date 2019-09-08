package limits

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
)

const resourcePath = "limits"

func getURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}
