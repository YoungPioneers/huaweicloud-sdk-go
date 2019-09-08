package metrics

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
