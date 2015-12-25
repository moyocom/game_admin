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

/*
http://127.0.0.1:8888/sys/notice?action=add&id=15&post_time=1305698204&plan_time=1305700000&end_time=1305910000&period=30&status=1&content=nihao
*/
