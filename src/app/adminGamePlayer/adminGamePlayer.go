package adminGamePlayer

import (
	"app"
	"fmt"
	"io/ioutil"
	. "lib/Util"
	"model"
	. "model"
	"net/http"
	"strconv"
	"strings"
	//"time"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminGamePlayer", Appname: "玩家管理"})
	app.R.HandleFunc("/adminGamePlayer/index", app.AppHandler(admin_GamePlayerList, 1))
	app.R.HandleFunc("/adminGamePlayer/playerList", app.AppHandler(admin_GamePlayerList, 1))
	app.R.HandleFunc("/adminGamePlayer/Query", app.AppHandler(admin_GamePlayerQuery, 1))
	app.R.HandleFunc("/adminGamePlayer/playerEditorOP", app.AppHandler(admin_GamePlayerEditorHandler, 1))
	app.R.HandleFunc("/adminGamePlayer/QueryPlayerList", app.AppHandler(admin_QueryplayerList))
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

//查询玩家列表
func admin_QueryplayerList(w http.ResponseWriter, r *http.Request) {

	start := r.FormValue("start")
	length := r.FormValue("length")
	draw := r.FormValue("draw")

	retStr := "{"
	retStr += `"draw":` + draw + ","
	maxNumber := strconv.Itoa(model.GamePlayer_MaxNumber())
	retStr += `"recordsTotal":` + maxNumber + `,"recordsFiltered":"` + maxNumber + `",`
	retStr += `"data":[`

	players := GamePlayer_Table(start, length)
	jsonStr := GamePlayerTable2Json(players)
	retStr += jsonStr
	retStr += "]}"
	fmt.Fprint(w, retStr)
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
	for i := 0; i < len(strArrs); i++ {
		//禁言.
		if strArrs[i] == "vJinYan" {
			getStr := GetUrlString(APIServer + "/api/user?msg=1032&id=" + playerID + "&time=" + OPTime)
			if getStr == "0" {
				fmt.Fprint(w, "ok")
			} else {
				fmt.Fprint(w, getStr)
			}
			return
		}

		//封号
		if strArrs[i] == "vFengHao" {
			getStr := GetUrlString(APIServer + "/api/user?msg=1034&id=" + playerID + "&time=" + OPTime)
			if getStr == "0" {
				fmt.Fprint(w, "ok")
			} else {
				fmt.Fprint(w, getStr)
			}
			return
		}
		//封IP
		if strArrs[i] == "vFengIP" {

		}
		//踢人
		if strArrs[i] == "vTiRen" {
			resp, _ := http.Get(APIServer + "/api/user?msg=1031&id=" + r.FormValue("playerID"))
			bydata, _ := ioutil.ReadAll(resp.Body)
			if string(bydata) == "0" {
				model.UpdateOnlinePlayers()
				fmt.Fprint(w, "ok")
			} else {
				fmt.Fprint(w, string(bydata))
			}
			return
		}

		if strArrs[i] == "vReJinyan" {
			getStr := GetUrlString(APIServer + "/api/user?msg=1033&id=" + playerID)
			if getStr == "0" {
				fmt.Fprint(w, "ok")
			} else {
				fmt.Fprint(w, getStr)
			}
			return
		}

		if strArrs[i] == "vReFengHao" {
			getStr := GetUrlString(APIServer + "/api/user?msg=1035&id=" + playerID)
			if getStr == "0" {
				fmt.Fprint(w, "ok")
			} else {
				fmt.Fprint(w, getStr)
			}
			return
		}

		if strArrs[i] == "vReFengIP" {

		}
	}

	fmt.Fprint(w, "ok")
}
