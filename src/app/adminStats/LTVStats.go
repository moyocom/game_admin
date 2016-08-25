package adminStats

import (
	"app"
	. "moyoLogic"
	"moyoLogic/DataAnalysis"
	"net/http"
)

func init() {
	app.R.HandleFunc("/adminStats/LTVStats", app.AppHandler(adminStats_LTV))
}

func adminStats_LTV(w http.ResponseWriter, r *http.Request) {
	AnalysisDB.Load()
	DataAnalysis.DailyAnalysis()

	app.AdminTemplate(w, r, map[string]interface{}{}, "template/adminStats/LTVStats.html", true)
}
