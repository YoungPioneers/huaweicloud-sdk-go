// +build acceptance networking provider

package extensions

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/clients"
	networking "github.com/huaweicloud/huaweicloud-sdk-go/acceptance/openstack/networking/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/tools"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/networking/v2/networks"
)

func TestNetworksProviderCRUD(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	if err != nil {
		t.Fatalf("Unable to create a network client: %v", err)
	}

	// Create a network
	network, err := networking.CreateNetwork(t, client)
	if err != nil {
		t.Fatalf("Unable to create network: %v", err)
	}
	defer networking.DeleteNetwork(t, client, network.ID)

	getResult := networks.Get(client, network.ID)
	newNetwork, err := getResult.Extract()
	if err != nil {
		t.Fatalf("Unable to extract network: %v", err)
	}

	tools.PrintResource(t, newNetwork)
}
