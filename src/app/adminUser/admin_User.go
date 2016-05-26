package adminUser

import (
	"app"

	"fmt"
	"html/template"
	. "lib/Util"
	"lib/forms"
	"model"
	"net/http"
	"strconv"
	"time"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminUser", Appname: "会员管理"})
	app.R.HandleFunc("/adminUser/addNewUser", app.AppHandler(admin_AddNewUser, 1))
	app.R.HandleFunc("/adminUser/index", app.AppHandler(admin_User, 1))
	app.R.HandleFunc("/adminUser/login", app.AppHandler(admin_login))
	app.R.HandleFunc("/adminUser/userinfo", app.AppHandler(admin_UserInfo))
	app.R.HandleFunc("/adminUser/deleteUser", app.AppHandler(admin_DeleteUser, 1))
	app.R.HandleFunc("/adminUser/updateUser", app.AppHandler(admin_UpdateUser, 1))
	fmt.Println("load adminUser")
}

/*============================/adminUser/index==========================================*/
func admin_User(w http.ResponseWriter, r *http.Request) {

	app.AdminTemplate(w, r, map[string]interface{}{"userlist": model.AdminUser_Table()}, "template/adminUser/index.html", true)
}

/*============================/adminUser/login==========================================*/
func admin_login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		userName := r.FormValue("userName")
		pwd := r.FormValue("pwd")
		if !forms.Reg_user(userName) {
			ExitMsg(w, "账号错误")
			return
		}
		userData := model.AdminUser_Find(userName)
		if userData == nil {
			ExitMsg(w, "账号错误")
			return
		}
		if forms.Tomd5(pwd) != userData.UserPwd {
			ExitMsg(w, "密码错误")
			return
		}
		//49ba59abbe56e057
		model.AdminUser_UpdateLoginTime(userData.ID)

		sess := model.Gsession.SessionStart(w, r)
		sess.Set("User", strconv.Itoa(userData.ID))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("template/adminUser/login.html")
	t.Execute(w, nil)
}

/*============================/adminUser/userinfo==========================================*/
func admin_UserInfo(w http.ResponseWriter, r *http.Request) {
	var curUser *model.AdminUser
	if r.FormValue("id") == "" {
		curUser = model.AdminUser_CurUser(w, r)
	} else {
		curUser = model.AdminUser_FindById(r.FormValue("id"))
	}

	if r.Method == "POST" && curUser != nil {
		NickNameSql := ""
		if forms.Reg_user(r.FormValue("NickName")) {
			NickNameSql = "NickName='" + r.FormValue("NickName") + "'"
		}
		pwdSql, _ := getChangePwdSql(curUser.UserPwd, r.FormValue("OldPwd"), r.FormValue("CheckNewPwd"), r.FormValue("NewPwd"))

		model.SQLDB.Exec("update go_adminuser set " +
			NickNameSql + pwdSql +
			",Email='" + r.FormValue("Email") + "'" +
			",Contact='" + r.FormValue("Contact") + "'" + " where ID = " + strconv.Itoa(curUser.ID))
		ExitMsg(w, "修改成功", "/adminUser/userinfo")
		return
	}

	app.AdminTemplate(w, r, map[string]interface{}{
		"User": curUser,
	}, "template/adminUser/userinfo.html", true)
}

func getChangePwdSql(trueOldPwd string, oldPwd string, checkNewPwd string, newPwd string) (string, string) {
	if oldPwd == "" {
		return "", ""
	}
	if forms.Tomd5(oldPwd) != trueOldPwd {
		return "", "旧密码错误"
	}
	if newPwd != oldPwd {
		return "", "两次密码输入不一致"
	}

	return " ,UserPwd= '" + forms.Tomd5(newPwd) + "'", ""
}

/*============================/adminUser/addNewUser==========================================*/
func admin_AddNewUser(w http.ResponseWriter, r *http.Request) {

	if !forms.Reg_user(r.FormValue("UserName")) || model.AdminUser_Find(r.FormValue("UserName")) != nil {
		fmt.Fprintf(w, "非法用户名")
		return
	}

	if !forms.Reg_user(r.FormValue("NickName")) {
		fmt.Fprintf(w, "非法昵称")
		return
	}

	if (r.FormValue("Email") != "" && !forms.Reg_email(r.FormValue("Email"))) ||
		(r.FormValue("Contact") != "" && !forms.Reg_user(r.FormValue("Contact"))) {
		fmt.Fprintf(w, "非法的邮箱或联系方式")
		return
	}

	model.AdminUser_AddNewUser(&model.AdminUser{
		UserName:      r.FormValue("UserName"),
		NickName:      r.FormValue("NickName"),
		UserPwd:       r.FormValue("UserPwd"),
		UserType:      forms.Toint(r.FormValue("UserType")),
		Email:         r.FormValue("Email"),
		Contact:       r.FormValue("Contact"),
		RegTime:       time.Now().Unix(),
		LastLoginTime: time.Now().Unix(),
	})
	fmt.Fprintf(w, "ok")
}

/*============================/adminUser/deleteUser==========================================*/
func admin_DeleteUser(w http.ResponseWriter, r *http.Request) {
	delId, err := strconv.Atoi(r.FormValue("Id"))
	if err != nil {
		ExitMsg(w, "删除失败")
		return
	}
	delSql := "delete from go_adminuser where ID=" + strconv.Itoa(delId)
	model.SQLDB.Exec(delSql)
	http.Redirect(w, r, "/adminUser/index", http.StatusFound)
}

/*============================/adminUser/updateUser==========================================*/
func admin_UpdateUser(w http.ResponseWriter, r *http.Request) {
	if !forms.Reg_user(r.FormValue("UpdateNickName")) ||
		(r.FormValue("UpdateEmail") != "" && !forms.Reg_email(r.FormValue("UpdateEmail"))) ||
		(r.FormValue("UpdateContact") != "" && !forms.Reg_user(r.FormValue("UpdateContact"))) {
		ExitMsg(w, "格式错误")
		return
	}

	updateMap := map[string]interface{}{
		"NickName": "\"" + r.FormValue("UpdateNickName") + "\"",
		"Email":    "\"" + r.FormValue("UpdateEmail") + "\"",
		"Contact":  "\"" + r.FormValue("UpdateContact") + "\"",
		"UserType": "'" + r.FormValue("UpdateUserType") + "'",
	}
	fmt.Println(updateMap)
	model.AdminUser_Update(r.FormValue("id"), updateMap)
	http.Redirect(w, r, "/adminUser/index", http.StatusFound)
}
