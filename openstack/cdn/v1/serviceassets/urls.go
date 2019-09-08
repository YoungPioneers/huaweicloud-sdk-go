package serviceassets

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("services", id, "assets")
}
