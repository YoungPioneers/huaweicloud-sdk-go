package serviceassets

import "github.com/huaweicloud/huaweicloud-sdk-go"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
