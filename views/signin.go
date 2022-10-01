package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}

	session := uadmin.IsAuthenticated(r)
    user := uadmin.User{}

    if session != nil {
        http.Redirect(w, r, "/index", http.StatusSeeOther)
        return
    }

    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")

        uadmin.Get(&user, "username = ?", username)
        session := user.Login(password, "")

        if session != nil {
            http.SetCookie(w, &http.Cookie{
                Path:  "/",
                Name:  "session",
                Value: session.Key,
            })

            http.SetCookie(w, &http.Cookie{
                Path:  "/",
                Name:  "username",
                Value: username,
            })

            http.Redirect(w, r, "/index", http.StatusSeeOther)
        }
    }
	uadmin.RenderHTML(w, r, "templates/signin.html", data)
}
