package adminHome

import (
	"app"
	"fmt"
	"net/http"
)

func init() {
	app.R.HandleFunc("/", app.AppHandler(admin_Home, 1))
	fmt.Println("Load adminHome")
}

func admin_Home(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminHome/index.html", false)
}
