// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package cmd

import (
	"github.com/Sirupsen/logrus"
)

func createLog() *logrus.Logger {
	ll := logrus.InfoLevel
	switch Verbose {
	case 0:
		ll = logrus.ErrorLevel
		break
	case 1:
		ll = logrus.WarnLevel
		break
	case 3:
		ll = logrus.DebugLevel
		break
	default:
		ll = logrus.InfoLevel
	}

	log := logrus.New()
	log.Level = ll

	return log
}
