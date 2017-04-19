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
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("http://%s:%d/requests", host, port)
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		bts, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		bodyJSON := make(map[string]interface{})
		err = json.Unmarshal(bts, &bodyJSON)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		b, err := json.MarshalIndent(bodyJSON, "", "  ")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println("Status:", res.StatusCode)
		fmt.Print(string(b))
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
