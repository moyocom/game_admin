package adminStats

import (
	"app"
	"fmt"
	. "git.oschina.net/yangdao/extlib"
	. "git.oschina.net/yangdao/extlib/data_type"
	. "moyoLogic"
	"net/http"
)

func init() {
	app.R.HandleFunc("/adminStats/roleInfo", app.AppHandler(adminStats_roleInfo, 1))
	app.R.HandleFunc("/adminStats/roleInfoQuery", app.AppHandler(adminStats_roleInfoQuery, 1))
}

func adminStats_roleInfo(w http.ResponseWriter, r *http.Request) {
	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminStats/RoleInfoStats.html", true)
}

func adminStats_roleInfoQuery(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	length := r.FormValue("length")
	draw := r.FormValue("draw")
	strType := r.FormValue("type")
	var StrTypeId string
	switch strType {
	case "login":
		StrTypeId = " type = 203"
	case "level":
		StrTypeId = " type = 301"
	case "die":
		StrTypeId = " type = 302"
	case "changemap":
		StrTypeId = " type = 305"
	case "changePK":
		StrTypeId = " type = 304"
	}
	querySql := "select * from log_data where " + StrTypeId + " limit " + start + `,` + length
	maxNumber := AnalysisDB.GetTableCount("log_data", StrTypeId)
	putData := AnalysisDB.QuerySql(querySql)
	fmt.Fprint(w, `{"draw":`+draw+`, "recordsTotal":`+Str(maxNumber)+`,"recordsFiltered":`+Str(maxNumber)+`,"data":`+roleInfoLog2Json(putData)+` }`)
}

func roleInfoLog2Json(seqData ISequence) string {
	retStr := "["
	ForEach(seqData, func(data interface{}) {
		curData := data.(IAssociative)
		var0 := curData.Get(":var0").(string)
		var var1 string = "null"
		if curData.Get(":var1") != nil {
			var1 = curData.Get(":var1").(string)
		}
		retStr += "[\"" + Str(curData.Get(":time").(int)) + "\"," + var0 + "," + var1 + "],"
	})
	if retStr != "[" {
		return retStr[:len(retStr)-1] + "]"
	} else {
		return "[]"
	}
}
