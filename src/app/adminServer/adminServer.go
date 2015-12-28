package adminServer

import (
	"app"
	"encoding/json"
	"fmt"
	"io/ioutil"
	. "lib/Util"
	"model"
	"net/http"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminServer", Appname: "服务器管理"})
	app.R.HandleFunc("/adminServer/index", app.AppHandler(admin_GameServerState, 1))
	fmt.Println("load adminServer")
}

//服务器状态
func admin_GameServerState(w http.ResponseWriter, r *http.Request) {
	//http:127.0.0.1:8888/sys/stats
	resp, err := http.Get(model.APIServer + "/sys/stats")
	CheckErr(err)
	body, err := ioutil.ReadAll(resp.Body)
	CheckErr(err)
	retData := &model.ServerState{}
	err = json.Unmarshal(body, retData)
	CheckErr(err)
	app.AdminTemplate(w, r, map[string]interface{}{"SeverState": retData}, "template/adminServer/ServerState.html", true)
}
