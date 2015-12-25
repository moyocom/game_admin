package adminHome

import (
	"app"
	"fmt"
	"io/ioutil"
	"net/http"
)

func init() {
	app.R.HandleFunc("/", app.AppHandler(admin_Home, 1))
	fmt.Println("Load adminHome")
}

func admin_Home(w http.ResponseWriter, r *http.Request) {
	req, err := http.Get("http://192.168.1.123:8888/user/online?type=count")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	HomeMap := map[string]interface{}{"OnlineNumber": string(body)}

	app.AdminTemplate(w, r, HomeMap, "template/adminHome/index.html", false)
}
