// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get recorded requests",
	Long:  `get returns the bodies requested in each path of this mock`,
	Run: func(cmd *cobra.Command, args []string) {
		log := createLog()

		url := fmt.Sprintf("http://%s:%d/requests", host, port)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		bts, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		bodyJSON := make(map[string]interface{})
		err = json.Unmarshal(bts, &bodyJSON)
		if err != nil {
			log.Fatal(err)
		}

		b, err := json.MarshalIndent(bodyJSON, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		log.WithFields(logrus.Fields(map[string]interface{}{
			"Status": res.StatusCode,
		})).Info(string(b))
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
