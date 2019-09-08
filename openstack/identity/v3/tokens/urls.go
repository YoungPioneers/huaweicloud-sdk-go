package tokens

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func tokenURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("auth", "tokens")
}
