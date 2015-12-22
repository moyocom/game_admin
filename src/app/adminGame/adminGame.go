package adminGame

import (
	"app"
	"fmt"
	"net/http"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminGame", Appname: "游戏管理"})
	app.R.HandleFunc("/adminGame/index", app.AppHandler(admin_GameNotice, 1))
	fmt.Println("load adminGame")
}

//游戏公告
func admin_GameNotice(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminGame/Notice.html", true)
}
