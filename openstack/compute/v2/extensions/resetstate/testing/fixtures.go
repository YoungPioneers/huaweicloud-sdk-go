package testing

import (
	"fmt"
	"net/http"
	"testing"

	th "github.com/YoungPioneers/huaweicloud-sdk-go/testhelper"
	"github.com/YoungPioneers/huaweicloud-sdk-go/testhelper/client"
)

func mockResetStateResponse(t *testing.T, id string, state string) {
	th.Mux.HandleFunc("/servers/"+id+"/action", func(w http.ResponseWriter, r *http.Request) {
		th.TestMethod(t, r, "POST")
		th.TestHeader(t, r, "X-Auth-Token", client.TokenID)
		th.TestJSONRequest(t, r, fmt.Sprintf(`{"os-resetState": {"state": "%s"}}`, state))
		w.WriteHeader(http.StatusAccepted)
	})
}
