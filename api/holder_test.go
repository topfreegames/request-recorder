package api_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"net/http"
	"net/http/httptest"

	"github.com/topfreegames/request-recorder/testing"
)

var _ = Describe("Holder", func() {
	var recorder *httptest.ResponseRecorder
	var request *http.Request
	var path string

	Describe("record", func() {
		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			handler.Method = "record"
		})

		It("should return status 200 when recording", func() {
			path = "/some/path"
			body := map[string]interface{}{
				"field1": "value1",
				"field2": 2,
				"field3": map[string]string{"a": "b", "c": "d"},
			}
			request, _ = http.NewRequest("POST", path, testing.JSONFor(body))
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			Expect(recorder.Header().Get("Content-Type")).To(Equal("application/json"))
			Expect(recorder.Body.String()).To(Equal(`{"success":"ok"}`))
			Expect(app.Holder).To(HaveKey(path))
			Expect(app.Holder[path]).To(HaveLen(1))
			Expect(app.Holder[path][0]).To(HaveKeyWithValue("field1", "value1"))
			Expect(app.Holder[path][0]["field2"]).To(BeEquivalentTo(2))

			resBody := app.Holder[path][0]["field3"].(map[string]interface{})
			Expect(resBody).To(HaveKeyWithValue("a", "b"))
			Expect(resBody).To(HaveKeyWithValue("c", "d"))
		})
	})

	Describe("request", func() {
		BeforeEach(func() {
			recorder = httptest.NewRecorder()
		})

		It("should return list of requests and theis bodies", func() {
			path1 := "/some/path/1"
			handler.Method = "record"
			body1 := map[string]interface{}{
				"field1": "value1",
			}
			request, _ = http.NewRequest("POST", path1, testing.JSONFor(body1))
			handler.ServeHTTP(recorder, request)

			path2 := "/some/path/2"
			body2 := map[string]interface{}{
				"field2": "value2",
			}
			request, _ = http.NewRequest("POST", path2, testing.JSONFor(body2))
			recorder = httptest.NewRecorder()
			handler.ServeHTTP(recorder, request)

			body3 := map[string]interface{}{
				"field3": "value3",
			}
			request, _ = http.NewRequest("POST", path1, testing.JSONFor(body3))
			recorder = httptest.NewRecorder()
			handler.ServeHTTP(recorder, request)

			handler.Method = "requests"
			recorder = httptest.NewRecorder()
			request, _ = http.NewRequest("GET", "/requests", nil)
			handler.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(http.StatusOK))
			bodyJSON := make(map[string]map[string][]map[string]interface{})
			json.Unmarshal(recorder.Body.Bytes(), &bodyJSON)

			Expect(bodyJSON["routes"][path1]).To(HaveLen(2))
			Expect(bodyJSON["routes"][path1][0]["field1"]).To(Equal("value1"))
			Expect(bodyJSON["routes"][path1][1]["field3"]).To(Equal("value3"))

			Expect(bodyJSON["routes"][path2]).To(HaveLen(1))
			Expect(bodyJSON["routes"][path2][0]["field2"]).To(Equal("value2"))
		})
	})
})
