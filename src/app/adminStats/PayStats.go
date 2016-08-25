package adminStats

import (
	"app"
	//"html/template"
	"fmt"
	. "moyoLogic"
	"net/http"

	//"strconv"

	//. "git.oschina.net/yangdao/SGLisp"
	. "git.oschina.net/yangdao/extlib"
	. "git.oschina.net/yangdao/extlib/data_type"
)

func init() {
	app.Gapps = append(app.Gapps, app.Apps{Pkgname: "adminStats", Appname: "运营分析"})
	app.R.HandleFunc("/adminStats/index", app.AppHandler(adminStats_Pay, 1))
	app.R.HandleFunc("/adminStats/QueryPayStats", app.AppHandler(adminStats_QueryPay))
	fmt.Println("load adminStats")
}

func adminStats_Pay(w http.ResponseWriter, r *http.Request) {

	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminStats/PayStats.html", true)
}

func adminStats_QueryPay(w http.ResponseWriter, r *http.Request) {
	AnalysisDB.Load()

	start := r.FormValue("start")
	length := r.FormValue("length")
	draw := r.FormValue("draw")
	searchValue := r.FormValue("search[value]")
	querySql := `select * from log_data where type = 331 ` + searchValue + ` limit ` + start + `,` + length
	fmt.Println(querySql)
	ListData331 := AnalysisDB.QuerySql(querySql)
	maxNumber := AnalysisDB.GetTableCount("log_data", "type=331 "+searchValue)

	fmt.Fprint(w, `{"draw":`+draw+`,"recordsTotal":`+Str(maxNumber)+`,
		"recordsFiltered":`+Str(maxNumber)+`,"data":`+PayLog2JsonStr(ListData331)+`}`)
}

func PayLog2JsonStr(seqData *TypList) string {
	retJsonStr := "["
	i := 0
	ForEach(seqData, func(v interface{}) {
		table := v.(IAssociative)
		retJsonStr += `["` + Str(table.Get(":time").(int)) + `",` + `"` + table.Get(":var0").(string) + `",` + `"` +
			table.Get(":var1").(string) + `",` + `"` + Str(table.Get(":var2")) + `",` + `"` + Str(table.Get(":var3"))
		i++
		if i < seqData.Count() {
			retJsonStr += `"],`
		} else {
			retJsonStr += `"]`
		}
	})
	retJsonStr += "]"
	return retJsonStr
}
