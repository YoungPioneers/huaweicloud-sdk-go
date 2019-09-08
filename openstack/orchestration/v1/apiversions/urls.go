package apiversions

import "github.com/huaweicloud/huaweicloud-sdk-go"

func apiVersionsURL(c *gophercloud.ServiceClient) string {
	return c.Endpoint
}
