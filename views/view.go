package views

import (
	"net/http"

	"github.com/nabalmes/tesst/models"
	"github.com/uadmin/uadmin"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	Balmes := []models.Balmes{}
	uadmin.All(&Balmes)

	data["humanDetails"] = Balmes

	uadmin.RenderHTML(w, r, "templates/index.html", data)
}
