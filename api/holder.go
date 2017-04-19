// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type HolderHandler struct {
	App    *App
	Method string
}

func (h HolderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch h.Method {
	case "requests":
		h.requests(w, r)
	case "record":
		h.record(w, r)
	}
}

func (h HolderHandler) requests(w http.ResponseWriter, r *http.Request) {
	writeResponse(w, http.StatusOK, map[string]interface{}{
		"routes": h.App.Holder,
	})
}

func (h *HolderHandler) record(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	path := r.URL.Path

	bodyBts, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	bodyJSON := make(map[string]interface{})
	err = json.Unmarshal(bodyBts, &bodyJSON)
	if err != nil {
		writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
		return
	}

	h.App.Holder.Add(path, bodyJSON)
	writeResponse(w, http.StatusOK, map[string]interface{}{
		"success": "ok",
	})
}
