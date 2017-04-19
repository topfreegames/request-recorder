// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package cmd

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var file, msg, route string

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post",
	Short: "post helper",
	Long:  `post helper because it is a long post request`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(file) == 0 && len(msg) == 0 {
			fmt.Printf("Inform file or message")
			os.Exit(1)
		}

		var err error
		var bts []byte

		if len(file) == 0 {
			bts = []byte(msg)
		} else {
			bts, err = ioutil.ReadFile(file)
		}

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		var b bytes.Buffer
		w := gzip.NewWriter(&b)
		w.Write(bts)
		w.Close()

		url := fmt.Sprintf("http://%s:%d%s", host, port, route)
		req, err := http.NewRequest("POST", url, &b)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		req.Header.Add("-X", "POST")
		req.Header.Add("-H", "Accept-Encoding: gzip")
		req.Header.Add("-H", "Content-Type: application/gzip")

		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		bts, err = ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Status:", res.StatusCode)
		fmt.Println("Body:  ", string(bts))
	},
}

func init() {
	RootCmd.AddCommand(postCmd)
	postCmd.Flags().StringVarP(&file, "file", "f", "", "Path to file to be sent on POST body (set either this or msg)")
	postCmd.Flags().StringVarP(&msg, "msg", "m", "", "Message to be sent on POST body (set either this or file)")
	postCmd.Flags().StringVarP(&route, "route", "r", "", "Request route to execute request")
}
