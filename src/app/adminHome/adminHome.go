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

	//req, err := http.Get(model.APIServer + "/user/online?type=count")
	OnlineNumber := "读取服务器数据失败"

	HomeMap := map[string]interface{}{"OnlineNumber": OnlineNumber}
	app.AdminTemplate(w, r, HomeMap, "template/adminHome/index.html", false)
}
