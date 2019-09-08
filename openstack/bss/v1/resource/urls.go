package resource

import "github.com/YoungPioneers/huaweicloud-sdk-go"

func listURL(client *gophercloud.ServiceClient, domainId string) string {
	return client.ServiceURL(domainId, "common/order-mgr/resources/detail")
}
