package adminGamePlayer

import (
	"app"
	"fmt"
	"net/http"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminGamePlayer", Appname: "游戏玩家管理"})
	app.R.HandleFunc("/adminGamePlayer/index", app.AppHandler(admin_GamePlayerEditor, 1))
	fmt.Println("load adminGamePlayer")
}

func admin_GamePlayerEditor(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminGamePlayer/PlayerEditor.html", true)
}
