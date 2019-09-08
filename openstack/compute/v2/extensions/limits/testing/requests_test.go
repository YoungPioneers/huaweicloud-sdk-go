package testing

import (
	"testing"

	"github.com/YoungPioneers/huaweicloud-sdk-go/openstack/compute/v2/extensions/limits"
	th "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
	"github.com/YoungPioneers/huaweicloud-sdk-go/testhelper/client"
)

func TestGet(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t)

	getOpts := limits.GetOpts{
		TenantID: TenantID,
	}

	actual, err := limits.Get(client.ServiceClient(), getOpts).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &LimitsResult, actual)
}
