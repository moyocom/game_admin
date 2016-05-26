package adminHome

import (
	"app"
	"fmt"
	"io/ioutil"
	"model"
	"net/http"
)

func init() {
	app.R.HandleFunc("/", app.AppHandler(admin_Home, 1))
	fmt.Println("Load adminHome")
}

func admin_Home(w http.ResponseWriter, r *http.Request) {

	req, err := http.Get(model.APIServer + "/user/online?type=count")
	OnlineNumber := "读取服务器数据失败"
	if err != nil {
		fmt.Println(err)
	} else {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			OnlineNumber = string(body)
		}
	}
	HomeMap := map[string]interface{}{"OnlineNumber": OnlineNumber}
	app.AdminTemplate(w, r, HomeMap, "template/adminHome/index.html", false)
}
