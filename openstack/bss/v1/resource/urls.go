package resource

import "github.com/huaweicloud/huaweicloud-sdk-go"

func listURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/detail")
}
