package Util

import (
	"html/template"
	"net/http"
)

func ExitMsg(w http.ResponseWriter, args ...string) {
	t, _ := template.ParseFiles("template/msg.html")
	if len(args) == 1 {
		t.Execute(w, map[string]interface{}{"Msg": args[0], "Url": nil})
	}
	if len(args) == 2 {
		t.Execute(w, map[string]interface{}{"Msg": args[0], "Url": args[1]})
	}
}

//检测错误
func CheckErr(err interface{}) {
	if err != nil {
		panic(err)
	}
}
