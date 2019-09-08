package publicips

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func CreateURL(c *gophercloud.ServiceClient)string{
	return c.ServiceURL("publicips")
}