package adminGame

import (
	"app"
	"fmt"
	"net/http"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminGame", Appname: "游戏管理"})
	app.R.HandleFunc("/adminGame/index", app.AppHandler(admin_GameNotice, 1))
	app.R.HandleFunc("/adminGame/AddNotice", app.AppHandler(admin_AddNotice, 1))
	app.R.HandleFunc("/adminGame/SysMail", app.AppHandler(admin_SysMail, 1))
	fmt.Println("load adminGame")
}

//游戏公告
func admin_GameNotice(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminGame/Notice.html", true)
}

func admin_AddNotice(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		return
	}
	NoticeValue := r.FormValue("NoticeValue")
	getStr := "http://192.168.1.123:8888/sys/notice?action=add&" + NoticeValue
	http.Get(getStr)
}

func admin_SysMail(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminGame/SysMail.html", true)
}
