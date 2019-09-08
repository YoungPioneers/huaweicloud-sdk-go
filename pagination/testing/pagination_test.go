package testing

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
	"github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
)

func createClient() *gophercloud.ServiceClient {
	return &gophercloud.ServiceClient{
		ProviderClient: &gophercloud.ProviderClient{TokenID: "abc123"},
		Endpoint:       testhelper.Endpoint(),
	}
}
