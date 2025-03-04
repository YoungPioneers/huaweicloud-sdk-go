package startstop

import "github.com/huaweicloud/huaweicloud-sdk-go"

// StartResult is the response from a Start operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type StartResult struct {
	gophercloud.ErrResult
}

// StopResult is the response from Stop operation. Call its ExtractErr
// method to determine if the request succeeded or failed.
type StopResult struct {
	gophercloud.ErrResult
}
