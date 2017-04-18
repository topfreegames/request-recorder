package api_test

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	"github.com/topfreegames/request-recorder/api"
	"github.com/topfreegames/request-recorder/models"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Api Suite")
}

var app *api.App
var handler *api.HolderHandler

var _ = BeforeSuite(func() {
	l := logrus.New()
	l.Level = logrus.FatalLevel

	app = &api.App{
		Address: fmt.Sprintf("%s:%d", "0.0.0.0", 8889),
		Logger:  l,
		Holder:  models.Holder{},
	}
})

var _ = BeforeEach(func() {
	app.Holder = models.Holder{}
	handler = &api.HolderHandler{
		App: app,
	}
})
