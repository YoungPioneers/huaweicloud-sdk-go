package extensions

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(
		c,
		ListExtensionURL(c),
		func(r pagination.PageResult) pagination.Page {
			return ExtensionPage{pagination.SinglePageBase(r)}
		},
	)
}
