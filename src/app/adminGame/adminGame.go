package adminGame

import (
	"app"
	"fmt"
	. "lib/Util"
	. "lib/lisp_core"
	. "lib/lisp_core/data_type"
	"model"
	"net/http"
	"strconv"
	"time"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminGame", Appname: "游戏管理"})
	app.R.HandleFunc("/adminGame/index", app.AppHandler(admin_GameNotice, 1))
	app.R.HandleFunc("/adminGame/AddNotice", app.AppHandler(admin_AddNotice, 1))
	app.R.HandleFunc("/adminGame/SysMail", app.AppHandler(admin_SysMail, 1))
	app.R.HandleFunc("/adminGame/RollNotice", app.AppHandler(admin_RollNotice, 1))
	app.R.HandleFunc("/adminGame/RollNoticeOpt", app.AppHandler(admin_NoticeOpt, 1))
	app.R.HandleFunc("/adminGame/AddMail", app.AppHandler(admin_AddMail, 1))
	fmt.Println("load adminGame")
}

//游戏公告
func admin_GameNotice(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.FormValue("Title")
		r.FormValue("Content")
		r.FormValue("Id")

		newNotice := &model.Notice{
			Title:    r.FormValue("Title"),
			Content:  r.FormValue("Content"),
			ServerId: model.APIServerId,
		}
		if r.FormValue("Id") == "0" {
			model.Notice_Add(newNotice)
		} else {
			intID := Int(r.FormValue("Id"))
			model.Notice_Update(newNotice, intID)
		}
		ExitMsg(w, "修改成功", "/adminGame/index")
		return
	}
	notice := model.Notice_GetById(model.APIServerId)
	if notice == nil {
		notice = &model.Notice{}
	}
	app.AdminTemplate(w, r, map[string]interface{}{"notice": notice}, "template/adminGame/Notice.html", true)
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

func admin_RollNotice(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{"notice_list": model.RollNotice_Table()}, "template/adminGame/RollNotice.html", true)
}

func admin_NoticeOpt(w http.ResponseWriter, r *http.Request) {
	strType := r.FormValue("Type")

	switch strType {
	case "Add":
		AddNoticeData := reqParam2Struct(r)
		fmt.Println(AddNoticeData)
		insertId := model.RollNotice_Add(AddNoticeData)
		retStr := GetUrlString(model.APIServer + "/sys/notice?action=add&id=" + Str(insertId) + "&" +
			"post_time=" + Str(AddNoticeData.PostTime) + "&plan_time=" + Str(AddNoticeData.PlanTime) + "&end_time=" +
			Str(AddNoticeData.EndTime) + "&period=" + Str(AddNoticeData.Period) + "&status=" + Str(AddNoticeData.Status) + "&content=" +
			Str(AddNoticeData.Content))
		fmt.Println(retStr)
		ExitMsg(w, "添加成功", "/adminGame/RollNotice")
	case "Del":
		id, _ := strconv.Atoi(r.FormValue("Id"))
		model.RollNotice_Del(id)
		retStr := GetUrlString(model.APIServer + "/sys/notice?action=delete&id=" + r.FormValue("Id"))
		fmt.Println(retStr)
		ExitMsg(w, "删除成功", "/adminGame/RollNotice")
	case "Update":
	}
}

func admin_AddMail(w http.ResponseWriter, r *http.Request) {
	strRecv := ""
	MailMap := HashMap()
	recvRoleType := r.FormValue("RecvRoleType")
	recvCareer := r.FormValue("RecvCareer")
	if recvRoleType == "0" {
		if recvCareer == "-1" {
			strRecv = `["all",0,0]`
		} else {
			strRecv = `["career",` + recvCareer + `,0,0]`
		}
	} else {
		strRecv = `[` + r.FormValue("RecvIdList") + `]`
	}
	strTime := ""
	if r.FormValue("SendTimeType") == "0" {
		strTime = "0"
	} else {
		strTime = r.FormValue("SenTime")
	}

	strGoods := "0"
	if r.FormValue("goodsId") != "" {
		strGoods = "[" + r.FormValue("goodsId") + "," + r.FormValue("goodsNumber") + "]"
		MailMap.Assoc("goods_type_id", r.FormValue("goodsId"))
		MailMap.Assoc("goods_num", r.FormValue("goodsNumber"))
	} else {
		MailMap.Assoc("goods_type_id", 0)
		MailMap.Assoc("goods_num", 0)
	}
	strMoney := "0"
	if r.FormValue("MoneyNumber") != "" {
		strMoney = "[" + r.FormValue("MoneyType") + "," + r.FormValue("MoneyNumber") + "]"
	}
	ExecAPIString := model.APIServer + "/sys/mail?action=add&recvs=" + strRecv + "&post_time=" + strTime + "&content=\"" + r.FormValue("content") +
		"\"&goods=" + strGoods + "&money=" + strMoney
	MailMap.Assoc("recv_ids", strRecv)
	MailMap.Assoc("post_time", Int(strTime))
	MailMap.Assoc("content", r.FormValue("content"))

	MailMap.Assoc("money_type", r.FormValue("MoneyType"))
	MailMap.Assoc("money_count", r.FormValue("MoneyNumber"))
	MailMap.Assoc("status", 0)

	id := model.Mail_Add(MailMap)
	fmt.Println(GetUrlString(ExecAPIString + "&id=" + Str(id)))
	ExitMsg(w, "添加成功", "/adminGame/SysMail")
}

func reqParam2Struct(r *http.Request) *model.RollNotice {
	rollNotice := &model.RollNotice{}
	if r.FormValue("Id") != "" {
		rollNotice.Id, _ = strconv.Atoi(r.FormValue("Id"))
	}
	if r.FormValue("Title") != "" {
		rollNotice.Title = r.FormValue("Title")
	}
	if r.FormValue("Content") != "" {
		rollNotice.Content = r.FormValue("Content")
	}

	rollNotice.PostTime = int(time.Now().Unix())

	if r.FormValue("PlanTime") != "" {
		intTime, _ := strconv.Atoi(r.FormValue("PlanTime"))
		rollNotice.PlanTime = int(intTime)
	}
	if r.FormValue("Period") != "" {
		intTime, _ := strconv.Atoi(r.FormValue("Period"))
		rollNotice.Period = intTime
	}
	if r.FormValue("EndTime") != "" {
		intTime, _ := strconv.Atoi(r.FormValue("EndTime"))
		rollNotice.EndTime = int(intTime)
	}
	rollNotice.Status = 1
	return rollNotice
}
