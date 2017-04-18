// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func writeResponse(w http.ResponseWriter, status int, res map[string]interface{}) {
	w.Header().Set("Content-Type", "application/json")
	bts, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
	}

	w.WriteHeader(status)
	w.Write(bts)
}
