// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/topfreegames/request-recorder/api"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Long:  `Starts the server.`,
	Run: func(cmd *cobra.Command, args []string) {
		ll := logrus.InfoLevel
		switch Verbose {
		case 0:
			ll = logrus.ErrorLevel
		case 1:
			ll = logrus.WarnLevel
		case 3:
			ll = logrus.DebugLevel
		}

		var log = logrus.New()
		log.Formatter = new(logrus.JSONFormatter)
		log.Level = ll

		cmdL := log.WithFields(logrus.Fields{
			"source":    "startCmd",
			"operation": "Run",
			"host":      host,
			"port":      port,
		})

		cmdL.Info("Creating application...")
		app := api.NewApp(
			host,
			port,
			log,
		)
		cmdL.Info("Application created successfully.")

		cmdL.Info("Starting application...")
		closer, err := app.ListenAndServe()
		if closer != nil {
			defer closer.Close()
		}
		if err != nil {
			cmdL.WithError(err).Fatal("Error running application.")
		}
	},
}

func init() {
	RootCmd.AddCommand(startCmd)
}
