package health

import (
	"net/http"

	"github.com/seanarwa/ds/util"
)

func Handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		util.WriteHTTPResponse(w, map[string]interface{}{
			"health": "HEALTHY",
		}, http.StatusAccepted)
	}
}
