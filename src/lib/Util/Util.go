package Util

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"time"
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

func TimeInt2Str(intTime int64) string {
	str_time := time.Unix(int64(intTime), 0).Format("2006-01-02 15:04:05")
	return str_time
}

func GetUrlString(url string) string {
	resp, _ := http.Get(url)
	bytedata, _ := ioutil.ReadAll(resp.Body)
	return string(bytedata)
}
