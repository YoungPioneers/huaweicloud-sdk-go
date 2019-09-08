package testing

import (
	"testing"

	"github.com/YoungPioneers/huaweicloud-sdk-go/openstack/compute/v2/extensions/lockunlock"
	th "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
	"github.com/YoungPioneers/huaweicloud-sdk-go/testhelper/client"
)

const serverID = "{serverId}"

func TestLock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStartServerResponse(t, serverID)

	err := lockunlock.Lock(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestUnlock(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockStopServerResponse(t, serverID)

	err := lockunlock.Unlock(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
