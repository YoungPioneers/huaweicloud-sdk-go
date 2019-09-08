package evacuate

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
)

func actionURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("servers", id, "action")
}
