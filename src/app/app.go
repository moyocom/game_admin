package app

import (
	"fmt"
	"html/template"
	. "lib/Util"
	"lib/mux"
	"model"
	"net/http"
	"strings"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()
}

var Gapps []Apps

type Apps struct {
	Pkgname string
	Appname string
}

func AppHandler(fn http.HandlerFunc, role ...int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(role) > 0 {
			if role[0] == 1 {
				id := model.AdminUser_CurUserId(w, r)
				if id == "" {
					ExitMsg(w, "请先登录", "/adminUser/login")
					return
				}
			}
		}
		/*
			defer func() {
				if err, ok := recover().(error); ok {
					fmt.Fprint(w, err)
				}
			}()
		*/
		fn(w, r)
	}
}

type View struct {
	HideLeftMenu bool
}

func AdminTemplate(w http.ResponseWriter, r *http.Request, data map[string]interface{}, file string, ShowLeftMenu bool) {
	t, err := template.ParseFiles("template/adminbase.html", file)
	if err != nil {
		fmt.Println(err)
	}
	data["apps"] = Gapps
	data["view"] = &View{}

	if r.RequestURI == "/" {
		data["curapp"] = "index"
	} else {
		arr := strings.Split(r.RequestURI, "/")
		data["curapp"] = arr[1]
	}

	if ShowLeftMenu == true {
		data["ShowLeftMenu"] = true
	}

	data["TimeInt2Str"] = TimeInt2Str
	data["CurServer"] = model.APIServer

	data["user"] = model.AdminUser_CurUser(w, r)
	t.Execute(w, data)
}

func (this *View) UtcToString(curt interface{}) string {
	return "啊♂"
}
