package apiversions

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
	"github.com/YoungPioneers/huaweicloud-sdk-go/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users
func ListVersions(c *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(c, apiVersionsURL(c), func(r pagination.PageResult) pagination.Page {
		return APIVersionPage{pagination.SinglePageBase(r)}
	})
}
