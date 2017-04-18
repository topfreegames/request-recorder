// request-recorder
// https://github.com/topfreegames/request-recorder
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package api

import (
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/topfreegames/request-recorder/models"
)

type App struct {
	Address string
	Logger  logrus.FieldLogger
	Router  *mux.Router
	Server  *http.Server
	Holder  models.Holder
}

func NewApp(
	host string,
	port int,
	logger logrus.FieldLogger,
) *App {
	app := &App{
		Address: fmt.Sprintf("%s:%d", host, port),
		Logger:  logger,
		Holder:  models.Holder{},
	}

	app.configureApp()

	return app
}

func (a *App) getRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/requests", &HolderHandler{
		App:    a,
		Method: "requests",
	}).Methods("GET")

	r.Handle("/", &HolderHandler{
		App:    a,
		Method: "record",
	}).Methods("POST")

	return r
}

func (a *App) configureApp() {
	a.configureLogger()
	a.configureServer()
}

func (a *App) configureLogger() {
	a.Logger = a.Logger.WithFields(logrus.Fields{
		"source":    "api/app.go",
		"operation": "initializeApp",
		"version":   "0.1.0",
	})
}

//ConfigureServer construct the routes
func (a *App) configureServer() {
	a.Router = a.getRouter()
	a.Server = &http.Server{Addr: a.Address, Handler: a.Router}
}

//ListenAndServe requests
func (a *App) ListenAndServe() (io.Closer, error) {
	listener, err := net.Listen("tcp", a.Address)
	if err != nil {
		return nil, err
	}

	err = a.Server.Serve(listener)
	if err != nil {
		listener.Close()
		return nil, err
	}

	return listener, nil
}