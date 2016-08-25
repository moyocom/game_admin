package DataAnalysis

import (
	"fmt"
	. "model"
	. "moyoLogic"

	. "git.oschina.net/yangdao/extlib"

	"time"
)

/*
*DailyAnalysis*
Day   日新增用户
1       200
2       100
3       100
4       100
5       100
6       100
7       100

*DailyRemain*
RegisterDay  Day  Number Rate
1            1    100    50
1            2    100    50
1            3    50     25
2            1    100    100
2            2    50     50
*/
const (
	DaySecond = 86400
)

func DailyAnalysis() {
	AnalysisDB.Load()
	fmt.Println("DailyAnalysis??")
	serverLst := ServerData_Table()
	for i := 0; i < len(serverLst); i++ {
		RunDailyAnalysis(serverLst[i])
	}
}

func RunDailyAnalysis(serverData *ServerData) {
	fmt.Println("RunDailyAnalysis")
	AnalysisDB.Connect()
	curTimeDay := time.Now().UTC().Unix() / DaySecond
	//刷新增人数
	maxDaySql := "select max(Day) as maxNumber from Analysis_Daily where serverId = " + ToStr(serverData.Id)
	row := AnalysisDB.AnalysisDB.QueryRow(maxDaySql)
	var curMaxDay int64
	row.Scan(&curMaxDay)
	if curMaxDay == 0 {
		curMaxDay = serverData.StartTime / DaySecond
	}
	fmt.Println(curMaxDay)
	gameDB := NewGameDB(serverData.DBUser, serverData.DBPwd, serverData.IP)
	gameDB.Connect()
	subDay := int(curTimeDay - curMaxDay)
	for i := 1; i < subDay-1; i++ {
		minDay := int(curMaxDay+i-1) * DaySecond
		maxDay := (int(curMaxDay) + i) * DaySecond
		newPlayerNumber := gameDB.GetNewPlayerByDay(minDay, maxDay)
		payPlayerNumber := AnalysisDB.GetPayPlayerNumber(serverData.Id, minDay, maxDay)

		insertSql := "insert into Analysis_Daily(Day,DayNewPlayer,ServerId,PayPlayerNumber) values(?,?,?,?)"
		stmt, _ := AnalysisDB.AnalysisDB.Prepare(insertSql)
		_, err := stmt.Exec(int(curMaxDay)+i, newPlayerNumber, serverData.Id, payPlayerNumber)
		if err != nil {
			fmt.Println(err)
		}

		//生成每天的注册缓存数据.
		Fill_DailyRegisterPlayer(int(curMaxDay)+i, gameDB, serverData.Id)
	}

	//刷留存人数和留存率
	serverStartDay := serverData.StartTime / DaySecond
	openDay := int(curTimeDay - serverStartDay)
	for i := 0; i < openDay; i++ {
		//检测是否有当前天的留存数据
		checkSql := "select max(Day) from analysis_retention where RegisterDay =" + ToStr(int(serverStartDay)+i)
		row := AnalysisDB.AnalysisDB.QueryRow(checkSql)
		var dayNumer int64
		row.Scan(&dayNumer)
		if dayNumer == 0 {
			dayNumer = serverData.StartTime / DaySecond
		}
		//分析生成并插入留存数据
		var firstDayNumber int64 = 0
		firstDayNumber = ScanRegisterNumberByDay(int(serverStartDay)+i, serverData.Id)
		fmt.Println("firstDayNumber:", firstDayNumber, int(serverStartDay)+i)

		curDay := int(serverStartDay) + i
		for j := int(curTimeDay); j > int(curDay); j-- {
			Analysis_Retention(j, int(serverStartDay)+i, serverData, int(firstDayNumber))
		}
	}
}

func Analysis_Retention(day int, registerDay int, serverData *ServerData, firstDayNumber int) {

	//log type=203的var2是玩家Id
	querySql := "select count(distinct var2) from log_data where type  = 203 and serverid = " + ToStr(serverData.Id)
	querySql += " and time >" + ToStr(DayToSecond(day-1))
	querySql += " and time <" + ToStr(DayToSecond(day))
	row := AnalysisDB.AnalysisDB.QueryRow(querySql)
	fmt.Println("RetentionSql", querySql)
	var LoginPlayerNumber int64
	var RetentionPlayerNumber int64
	row.Scan(&LoginPlayerNumber)

	allPlayerSql := querySql + " and var2 in (select playerid from analysis_registerplayer where day = " + ToStr(registerDay) + ")"
	row = AnalysisDB.AnalysisDB.QueryRow(allPlayerSql)
	row.Scan(&RetentionPlayerNumber)
	fmt.Println(allPlayerSql)
	if firstDayNumber == 0 {
		firstDayNumber = -1
	}

	fmt.Println(RetentionPlayerNumber, firstDayNumber)
	insertSql := `insert into analysis_retention(RegisterDay,Day,RemainNumber,RemainRate) values(` +
		ToStr(registerDay) + `,` + ToStr(day) + `,` + ToStr(RetentionPlayerNumber) + `,` + ToStr(float32(RetentionPlayerNumber)/float32(firstDayNumber)) + `)`
	AnalysisDB.AnalysisDB.Exec(insertSql)
}

func ScanRegisterNumberByDay(day int, serverId int) int64 {
	//log type=203的var2是玩家Id
	querySql := "select count(distinct playerid) from analysis_registerplayer where day  = " + ToStr(day) + " and serverid = " + ToStr(serverId)
	row := AnalysisDB.AnalysisDB.QueryRow(querySql)
	var LoginPlayerNumber int64
	row.Scan(&LoginPlayerNumber)
	return LoginPlayerNumber
}

//填充每天的注册玩家中间表
func Fill_DailyRegisterPlayer(dayNumber int, gameDB *TypGameDB, serverId int) {

	querySql := "select id from player where reg_time > " + ToStr(dayNumber*DaySecond) + " and reg_time < " + ToStr((dayNumber+1)*DaySecond)
	rows, err := gameDB.GameDB.Query(querySql)
	if err != nil {
		fmt.Println(err)
	}
	insertStr := "insert into analysis_registerplayer(day,playerid,serverid) values(?,?,?)"
	tx, _ := AnalysisDB.AnalysisDB.Begin()
	stmt, err := tx.Prepare(insertStr)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var playerId int
		rows.Scan(&playerId)
		_, err := stmt.Exec(dayNumber, playerId, serverId)
		if err != nil {
			fmt.Println(err)
		}
	}
	stmt.Close()
	tx.Commit()
}

func DayToSecond(day int) int64 {
	sec := day * DaySecond
	curTime := time.Unix(int64(sec), 0)
	return time.Date(curTime.Year(), curTime.Month(), curTime.Day(), 0, 0, 0, 0, time.Local).Unix()
}
