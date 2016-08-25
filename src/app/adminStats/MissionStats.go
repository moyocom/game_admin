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
	app.R.HandleFunc("/adminStats/mission", app.AppHandler(adminStats_Mission, 1))
	app.R.HandleFunc("/adminStats/missionQuery", app.AppHandler(adminStats_MissionQuery, 1))
}

func adminStats_Mission(w http.ResponseWriter, r *http.Request) {

	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminStats/MissionStats.html", true)
}

func adminStats_MissionQuery(w http.ResponseWriter, r *http.Request) {

	start := r.FormValue("start")
	length := r.FormValue("length")
	draw := r.FormValue("draw")
	searchValue := r.FormValue("search[value]")
	strType := r.FormValue("type")
	var typeIdSqlStr string
	switch strType {
	case "GetTask":
		typeIdSqlStr = " type = 311 "
	case "TiaoJian":
		typeIdSqlStr = " type = 312 "
	case "WanCheng":
		typeIdSqlStr = " type = 313 "
	case "QuXiao":
		typeIdSqlStr = " type = 314 "
	case "ShiBai":
		typeIdSqlStr = " type = 315 "
	}
	querySql := `select * from log_data where ` + typeIdSqlStr + ` ` + searchValue + ` limit ` + start + `,` + length
	putData := AnalysisDB.QuerySql(querySql)
	maxNumber := AnalysisDB.GetTableCount("log_data", typeIdSqlStr+searchValue)
	//fmt.Println(strType, maxNumber)
	fmt.Fprint(w, `{ "draw":`+draw+`, "recordsTotal":`+Str(maxNumber)+` ,"recordsFiltered":`+Str(maxNumber)+` ,
	 "data":`+MissionLog2Json(putData)+` }`)
}

func MissionLog2Json(seqData ISequence) string {
	retStr := "["
	ForEach(seqData, func(data interface{}) {
		curData := data.(IAssociative)
		retStr += "[\"" + Str(curData.Get(":time").(int)) + "\"," + curData.Get(":var0").(string) + "," + curData.Get(":var1").(string) + "],"
	})
	if retStr != "[" {
		return retStr[:len(retStr)-1] + "]"
	} else {
		return "[]"
	}
}
