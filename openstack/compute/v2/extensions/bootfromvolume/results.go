package bootfromvolume

import (
	os "github.com/huaweicloud/huaweicloud-sdk-go/openstack/compute/v2/servers"
)

// CreateResult temporarily contains the response from a Create call.
// It embeds the standard servers.CreateResults type and so can be used the
// same way as a standard server request result.
type CreateResult struct {
	os.CreateResult
}
