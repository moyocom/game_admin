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
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminUser", Appname: "会员管理"})
	app.R.HandleFunc("/adminUser/index", app.AppHandler(admin_User, 1))
	app.R.HandleFunc("/adminUser/login", app.AppHandler(admin_login))
	fmt.Println("load adminUser")
}
func admin_User(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminUser/index.html", true)
}

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

			return
		}
		fmt.Println(forms.Tomd5(pwd))
		if forms.Tomd5(pwd) != userData.UserPwd {
			ExitMsg(w, "密码错误")
			return
		}
		//49ba59abbe56e057
		sess := model.Gsession.SessionStart(w, r)
		sess.Set("User", strconv.Itoa(userData.ID))
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}
	fmt.Println("SQLDB", model.SQLDB)
	t, _ := template.ParseFiles("template/adminUser/login.html")
	t.Execute(w, nil)
}
