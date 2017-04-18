// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package testing

import (
	"bytes"
	"encoding/json"
	"io"
)

func JSONFor(res map[string]interface{}) io.Reader {
	bts, _ := json.Marshal(res)
	return bytes.NewBuffer(bts)
}
