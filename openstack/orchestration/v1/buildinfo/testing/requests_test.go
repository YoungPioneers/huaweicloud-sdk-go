package testing

import (
	"testing"

	"github.com/YoungPioneers/huaweicloud-sdk-go/openstack/orchestration/v1/buildinfo"
	th "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
	fake "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper/client"
)

func TestGetTemplate(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()
	HandleGetSuccessfully(t, GetOutput)

	actual, err := buildinfo.Get(fake.ServiceClient()).Extract()
	th.AssertNoErr(t, err)

	expected := GetExpected
	th.AssertDeepEquals(t, expected, actual)
}
