package apiversions

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
