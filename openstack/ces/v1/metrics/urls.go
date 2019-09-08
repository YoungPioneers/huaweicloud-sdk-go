package metrics

import "github.com/huaweicloud/huaweicloud-sdk-go"

func getMetricsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("metrics")
}
