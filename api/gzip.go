// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package api

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func makeGzipHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}

		fz, err := gzip.NewReader(r.Body)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		defer fz.Close()

		s, err := ioutil.ReadAll(fz)
		if err != nil {
			writeResponse(w, http.StatusBadRequest, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		newReq, err := http.NewRequest("POST", r.URL.Path, bytes.NewReader(s))
		if err != nil {
			writeResponse(w, http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		fn(w, newReq)
	}
}
