package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/vpc/v1/quotas"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
	"github.com/huaweicloud/huaweicloud-sdk-go/testhelper/client"
)

func TestList(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleListSuccessfully(t)

	actual, err := quotas.List(client.ServiceClient(), quotas.ListOpts{
		Type: "vpc",
	}).Extract()
	th.AssertNoErr(t, err)
	th.CheckDeepEquals(t, &ListResponse, actual)
}
