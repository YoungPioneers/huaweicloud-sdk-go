// +build acceptance compute usage

package v2

import (
	"strings"
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/clients"
	"github.com/huaweicloud/huaweicloud-sdk-go/acceptance/tools"
	"github.com/huaweicloud/huaweicloud-sdk-go/openstack/compute/v2/extensions/usage"
)

func TestUsageSingleTenant(t *testing.T) {
	client, err := clients.NewComputeV2Client()
	if err != nil {
		t.Fatalf("Unable to create a compute client: %v", err)
	}

	endpointParts := strings.Split(client.Endpoint, "/")
	tenantID := endpointParts[4]

	page, err := usage.SingleTenant(client, tenantID, nil).AllPages()
	if err != nil {
		t.Fatal(err)
	}

	tenantUsage, err := usage.ExtractSingleTenant(page)
	if err != nil {
		t.Fatal(err)
	}

	tools.PrintResource(t, tenantUsage)
}
