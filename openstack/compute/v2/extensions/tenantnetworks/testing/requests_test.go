package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/compute/v2/extensions/tenantnetworks"
	"github.com/huaweicloud/huaweicloud-sdk-go/pagination"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	count := 0
	err := tenantnetworks.List(client.ServiceClient()).EachPage(func(page pagination.Page) (bool, error) {
		count++
		actual, err := tenantnetworks.ExtractNetworks(page)
		th.AssertNoErr(t, err)
		th.CheckDeepEquals(t, ExpectedNetworkSlice, actual)

		return true, nil
	})
	th.AssertNoErr(t, err)
	th.CheckEquals(t, 1, count)
}

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	actual, err := tenantnetworks.Get(client.ServiceClient(), "20c8acc0-f747-4d71-a389-46d078ebf000").Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &SecondNetwork, actual)
}
