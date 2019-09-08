package flavor

import "github.com/huaweicloud/huaweicloud-sdk-go"

// GET list url
func getListUrl(sc *gophercloud.ServiceClient) string {
	return sc.ServiceURL("cloudservers","flavors")
}

// Resize url
func resizeURL(sc *gophercloud.ServiceClient,serverId string) string {
	return sc.ServiceURL("cloudservers",serverId,"resize")
}