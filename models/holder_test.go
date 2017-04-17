// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package models_test

import (
	. "github.com/topfreegames/request-recorder/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Holder", func() {
	It("should save an json request", func() {
		holder := Holder{}
		body := map[string]interface{}{
			"field1": "value1",
			"filed2": 2,
			"field3": map[string]int{"a": 1, "b": 2},
		}
		path := "/some/random/path"
		holder.Add(path, body)

		Expect(holder).To(HaveLen(1))
		Expect(holder[path]).To(HaveLen(1))
		Expect(holder[path][0]).To(BeEquivalentTo(body))
	})
})
