package adminGamePlayer

import (
	"app"
	"fmt"
	"model"
	. "model"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminGamePlayer", Appname: "游戏玩家管理"})
	app.R.HandleFunc("/adminGamePlayer/index", app.AppHandler(admin_GamePlayerEditor, 1))
	app.R.HandleFunc("/adminGamePlayer/playerList", app.AppHandler(admin_GamePlayerList, 1))
	app.R.HandleFunc("/adminGamePlayer/Query", app.AppHandler(admin_GamePlayerQuery, 1))
	app.R.HandleFunc("/adminGamePlayer/playerEditorOP", app.AppHandler(admin_GamePlayerEditorHandler, 1))
	fmt.Println("load adminGamePlayer")
}

//玩家操作界面
func admin_GamePlayerEditor(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminGamePlayer/PlayerEditor.html", true)
}

//玩家列表
func admin_GamePlayerList(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminGamePlayer/PlayerList.html", true)
}

//玩家查询
func admin_GamePlayerQuery(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("playerID")
	userName := r.FormValue("playerName")
	var sqlstr string
	if userID == "" {
		sqlstr = "select id from player where name =\"" + userName + "\""
		var playerid int
		model.GameDB.QueryRow(sqlstr).Scan(&playerid)
		fmt.Fprintf(w, strconv.Itoa(playerid))
	} else {
		sqlstr = "select name from player where id=" + userID
		var playerName string
		model.GameDB.QueryRow(sqlstr).Scan(&playerName)
		fmt.Fprintf(w, playerName)
	}
}

//玩家操作处理
func admin_GamePlayerEditorHandler(w http.ResponseWriter, r *http.Request) {
	OPStr := r.FormValue("OPStr")
	strArrs := strings.Split(OPStr, ",")
	playerID := r.FormValue("playerID")
	OPTime := r.FormValue("OPTime")
	fmt.Println(playerID)
	if OPTime == "" {
		OPTime = "1"
	}
	OPTimeInt, _ := strconv.Atoi(OPTime)
	for i := 0; i < len(strArrs); i++ {
		//禁言.
		if strArrs[i] == "vJinYan" {
			endTime := time.Now().Add(time.Hour * time.Duration(OPTimeInt))
			execStr := APIServer + "/api/user?msg=1032&end_time=" + strconv.FormatInt(endTime.Unix(), 10) + "&id=" + playerID
			http.Get(execStr)
		}
		//封号
		if strArrs[i] == "vFengHao" {

		}
		//封IP
		if strArrs[i] == "vFengIP" {

		}
		//踢人
		if strArrs[i] == "vTiRen" {

		}

		if strArrs[i] == "vReJinyan" {

		}

		if strArrs[i] == "vReFengHao" {

		}

		if strArrs[i] == "vReFengIP" {

		}
	}

	fmt.Fprint(w, "ok")
}
