package apiversions

import (
	"net/url"
	"strings"

	"github.com/huaweicloud/huaweicloud-sdk-go"
)

func getURL(c *gophercloud.ServiceClient, version string) string {
	return c.ServiceURL(strings.TrimRight(version, "/") + "/")
}

func listURL(c *gophercloud.ServiceClient) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"
	return u.String()
}
