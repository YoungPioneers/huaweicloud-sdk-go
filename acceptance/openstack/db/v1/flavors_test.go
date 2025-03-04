// +build acceptance db

package v1

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/clients"
	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/tools"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/db/v1/flavors"
)

func TestFlavorsList(t *testing.T) {
	client, err := clients.NewDBV1Client()
	if err != nil {
		t.Fatalf("Unable to create a DB client: %v", err)
	}

	allPages, err := flavors.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve flavors: %v", err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		t.Fatalf("Unable to extract flavors: %v", err)
	}

	for _, flavor := range allFlavors {
		tools.PrintResource(t, &flavor)
	}
}

func TestFlavorsGet(t *testing.T) {
	client, err := clients.NewDBV1Client()
	if err != nil {
		t.Fatalf("Unable to create a DB client: %v", err)
	}

	allPages, err := flavors.List(client).AllPages()
	if err != nil {
		t.Fatalf("Unable to retrieve flavors: %v", err)
	}

	allFlavors, err := flavors.ExtractFlavors(allPages)
	if err != nil {
		t.Fatalf("Unable to extract flavors: %v", err)
	}

	if len(allFlavors) > 0 {
		flavor, err := flavors.Get(client, allFlavors[0].StrID).Extract()
		if err != nil {
			t.Fatalf("Unable to get flavor: %v", err)
		}

		tools.PrintResource(t, flavor)
	}
}
