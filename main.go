package main

import (
	"net/http"

	"github.com/nabalmes/tesst/models"
	"github.com/nabalmes/tesst/views"
	"github.com/uadmin/uadmin"
)

func main() {
	dataHandler()
	RegisterModels()
	ServerPort()

}

func ServerPort() {
	uadmin.Port = 6969
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "uAdmin TM"
	uadmin.BindIP = "localhost"
	uadmin.StartServer()
}

func RegisterModels() {
	uadmin.Register(
		models.Balmes{},
	)
}

func dataHandler() {
	http.HandleFunc("/index/", uadmin.Handler(views.IndexHandler))
	http.HandleFunc("/", uadmin.Handler(views.SigninHandler))
}
