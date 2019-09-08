package events

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
)

func createURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("events")
}
