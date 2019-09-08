package migrate

import (
	"github.com/YoungPioneers/huaweicloud-sdk-go"
)

// MigrateResult is the response from a Migrate operation. Call its ExtractErr
// method to determine if the request suceeded or failed.
type MigrateResult struct {
	gophercloud.ErrResult
}
