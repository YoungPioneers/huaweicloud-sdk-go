package aggregates

import "github.com/huaweicloud/huaweicloud-sdk-go"

func aggregatesListURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-aggregates")
}

func aggregatesCreateURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-aggregates")
}

func aggregatesDeleteURL(c *gophercloud.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesGetURL(c *gophercloud.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID)
}

func aggregatesUpdateURL(c *gophercloud.ServiceClient, aggregateID string) string {
	return c.ServiceURL("os-aggregates", aggregateID)
}
