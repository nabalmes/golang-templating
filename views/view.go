package views

import (
	"net/http"
	"strings"

	"github.com/nabalmes/tesst/models"
	"github.com/uadmin/uadmin"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	Balmes := []models.Balmes{}
	uadmin.All(&Balmes)

	context := map[string]interface{}{}
	context["humanDetails"] = Balmes

	// uadmin.RenderHTML(w, r, "templates/index.html", data)

	username_cookie, _ := r.Cookie("username")
	username := username_cookie.Value
	user := uadmin.User{}

	uadmin.Get(&user, "username = ?", username)
	context["User"] = user

	session_cookie, _ := r.Cookie("session")
	if session_cookie != nil {
		http.SetCookie(w, &http.Cookie{
			Name:  "session",
			Value: session_cookie.Value,
			Path:  "/",
		})
	}

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/index")
	uadmin.RenderHTML(w, r, "templates/index.html", context)
}
