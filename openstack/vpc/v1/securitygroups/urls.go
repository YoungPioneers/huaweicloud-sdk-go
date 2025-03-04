package securitygroups

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
)

func CreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-groups")
}

func DeleteURL(c *gophercloud.ServiceClient, securityGroupId string) string {
	return c.ServiceURL("security-groups", securityGroupId)
}

func GetURL(c *gophercloud.ServiceClient, securityGroupId string) string {
	return c.ServiceURL("security-groups", securityGroupId)
}

func ListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("security-groups")
}
