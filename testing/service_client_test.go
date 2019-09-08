package testing

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go"
	th "github.com/huaweicloud/huaweicloud-sdk-go/testhelper"
)

func TestServiceURL(t *testing.T) {
	c := &gophercloud.ServiceClient{Endpoint: "http://123.45.67.8/"}
	expected := "http://123.45.67.8/more/parts/here"
	actual := c.ServiceURL("more", "parts", "here")
	th.CheckEquals(t, expected, actual)
}
