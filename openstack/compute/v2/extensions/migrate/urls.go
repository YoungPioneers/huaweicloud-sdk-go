package migrate

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
)

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
