package testing

import (
	"testing"

	"github.com/YoungPioneers/huaweicloud-sdk-go/openstack/compute/v2/extensions/suspendresume"
	th "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
	"github.com/YoungPioneers/huaweicloud-sdk-go/testhelper/client"
)

const serverID = "{serverId}"

func TestSuspend(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockSuspendServerResponse(t, serverID)

	err := suspendresume.Suspend(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}

func TestResume(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	mockResumeServerResponse(t, serverID)

	err := suspendresume.Resume(client.ServiceClient(), serverID).ExtractErr()
	th.AssertNoErr(t, err)
}
