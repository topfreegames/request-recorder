// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string
var host string
var port int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "request-recorder",
	Short: "Is a request mock",
	Long:  `Save requests is memory and retrive them.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

// Verbose determines how verbose mystack will run under
var Verbose int

func init() {
	RootCmd.PersistentFlags().IntVarP(
		&Verbose, "verbose", "v", 2,
		"Verbosity level => v0: Error, v1=Warning, v2=Info, v3=Debug",
	)
	RootCmd.PersistentFlags().StringVarP(&host, "bind", "b", "0.0.0.0", "Host to bind mystack to")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8080, "Port to bind mystack to")
}
