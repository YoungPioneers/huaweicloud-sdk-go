package apiversions

import (
	"net/url"
	"strings"

	"github.com/huaweicloud/huaweicloud-sdk-go"
)

// getURL generates URL fot getting detail of specified version
func getURL(c *gophercloud.ServiceClient, version string) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/" + strings.TrimRight(version, "/") + "/"
	return u.String()
}

// listURL generates URL for list versions
func listURL(c *gophercloud.ServiceClient) string {
	u, _ := url.Parse(c.ServiceURL(""))
	u.Path = "/"
	return u.String()
}
