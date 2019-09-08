package extensions

import "github.com/YoungPioneers/huaweicloud-sdk-go"

// ListExtensionURL generates the URL for the extensions resource collection.
func ListExtensionURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("extensions")
}
