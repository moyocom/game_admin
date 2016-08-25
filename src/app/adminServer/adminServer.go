package adminServer

import (
	"app"
	"database/sql"
	"encoding/json"
	"fmt"
	. "git.oschina.net/yangdao/extlib"
	"io/ioutil"
	. "lib/Util"
	"model"
	"net/http"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminServer", Appname: "服务器管理"})
	app.R.HandleFunc("/adminServer/index", app.AppHandler(admin_ServerList, 1))
	app.R.HandleFunc("/adminServer/ChangeServer", app.AppHandler(admin_ChangeServer, 1))
	app.R.HandleFunc("/adminServer/DeleteServer", app.AppHandler(admin_deleteServer, 1))
	app.R.HandleFunc("/adminServer/AddNewServer", app.AppHandler(admin_AddNewServer, 1))
	fmt.Println("load adminServer")
}

//服务器状态
func admin_GameServerState(w http.ResponseWriter, r *http.Request) {
	//http:127.0.0.1:8888/sys/stats
	retMap := make(map[string]interface{})
	resp, err := http.Get(model.APIServer + "/sys/stats")
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		CheckErr(err)
		retData := &model.ServerState{}
		err = json.Unmarshal(body, retData)
		retMap["SeverState"] = retData
		CheckErr(err)
	}
	app.AdminTemplate(w, r, retMap, "template/adminServer/ServerState.html", true)
}

func admin_ServerList(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{"server_list": model.ServerData_Table()}, "template/adminServer/ServerList.html", true)
}

func admin_ChangeServer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.FormValue("id"))
	serverData := model.ServerData_ById(Int(r.FormValue("id")))
	model.GameDB.Close()
	var err error
	model.GameDB, err = sql.Open("mysql", serverData.DBUser+":"+serverData.DBPwd+"@tcp("+serverData.IP+":3306)/yulong_game")
	CheckErr(err)
	model.GameDB.Ping()
	model.APIServer = serverData.IP + ":8888"
	model.APIServerId = serverData.Id
	fmt.Println(serverData)
	fmt.Println(model.APIServer)
	ExitMsg(w, "切换成功", "/adminServer/index")
}

func admin_deleteServer(w http.ResponseWriter, r *http.Request) {

	model.CenterDB.CenterDB.Exec("delete from go_server_list where id = " + r.FormValue("id"))
	ExitMsg(w, "删除成功", "/adminServer/index")
}

func admin_AddNewServer(w http.ResponseWriter, r *http.Request) {
	model.ServerData_AddServer(&model.ServerData{
		Id:     Int(r.FormValue("ServerId")),
		Name:   r.FormValue("ServerName"),
		Desc:   r.FormValue("ServerDesc"),
		IP:     r.FormValue("ServerIP"),
		Port:   r.FormValue("ServerPort"),
		DBUser: r.FormValue("DBUser"),
		DBPwd:  r.FormValue("DBPwd"),
	})
	fmt.Fprint(w, "ok")
}
