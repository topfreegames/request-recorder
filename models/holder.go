// request-recorder
// https://github.com/topfreegames/mystack-controller
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package models

type Holder map[string][]map[string]interface{}

func NewHolder() map[string][]map[string]interface{} {
	return make(map[string][]map[string]interface{})
}

func (holder Holder) Add(path string, body map[string]interface{}) {
	if ar, ok := holder[path]; ok {
		holder[path] = append(ar, body)
		return
	}

	holder[path] = []map[string]interface{}{body}
}
