package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	user := uadmin.User{}

	user.FirstName = r.FormValue("name")
	username := r.FormValue("username")
	password := r.FormValue("password")
	uadmin.Trail(uadmin.DEBUG, "FNAME: %v", r.FormValue("name"))
	uadmin.Trail(uadmin.DEBUG, "USERNAME: %v", username)
	uadmin.Trail(uadmin.DEBUG, "PASSWD: %v", password)

	user.Username = username
	user.Password = password
	user.Active = true
	user.Save()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
