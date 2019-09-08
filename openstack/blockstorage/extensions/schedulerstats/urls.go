package schedulerstats

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func storagePoolsListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("scheduler-stats", "get_pools")
}
