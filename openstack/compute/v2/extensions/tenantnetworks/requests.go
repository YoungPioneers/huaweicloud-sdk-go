package tenantnetworks

import (
	"github.com/huaweicloud/huaweicloud-sdk-go"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
)

// List returns a Pager that allows you to iterate over a collection of Networks.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	return pagination.NewPager(client, listURL(client), func(r pagination.PageResult) pagination.Page {
		return NetworkPage{pagination.SinglePageBase(r)}
	})
}

// Get returns data about a previously created Network.
func Get(client *gophercloud.ServiceClient, id string) (r GetResult) {
	_, r.Err = client.Get(getURL(client, id), &r.Body, nil)
	return
}
