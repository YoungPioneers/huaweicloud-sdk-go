package hypervisors

import "github.com/huaweicloud/huaweicloud-sdk-go"

func hypervisorsListDetailURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-hypervisors", "detail")
}

func hypervisorsStatisticsURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("os-hypervisors", "statistics")
}

func hypervisorsGetURL(c *gophercloud.ServiceClient, hypervisorID string) string {
	return c.ServiceURL("os-hypervisors", hypervisorID)
}

func hypervisorsUptimeURL(c *gophercloud.ServiceClient, hypervisorID string) string {
	return c.ServiceURL("os-hypervisors", hypervisorID, "uptime")
}
