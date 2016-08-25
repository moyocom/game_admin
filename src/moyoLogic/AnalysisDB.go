package moyoLogic

import (
	"database/sql"
	cfg "lib/config"
	_ "lib/mysql"

	"fmt"

	. "git.oschina.net/yangdao/extlib"
	. "git.oschina.net/yangdao/extlib/data_type"
	"git.oschina.net/yangdao/extlib/sqldb_helper"
)

var AnalysisDB typAnalysisDB

type typAnalysisDB struct {
	AnalysisDB *sql.DB
}

func (this *typAnalysisDB) Load() {
	GameLogLoader.Load()
}

func (this *typAnalysisDB) Connect() {
	if this.AnalysisDB != nil && this.AnalysisDB.Stats().OpenConnections > 0 {
		return
	}
	AnalysisDBAddr := cfg.Get()["AnalysisDB"]
	this.AnalysisDB = sqldb_helper.ConnectDB(AnalysisDBAddr)
}

func (this *typAnalysisDB) Close() {
	this.AnalysisDB.Close()
}

func (this *typAnalysisDB) GenLogTable() {
	this.AnalysisDB.Exec(`Create Table Log_Index
		(
           id        int primary key auto_increment,
           serverid  int,
           logkey    int
		)`)

	this.AnalysisDB.Exec(`Create Table Log_Data
		(
             id int primary key auto_increment,
             logkey    int,
             serverid  int,
             time      int,
             type      int,
             var2      varchar(100),
             var3      varchar(100),
             var4      varchar(100),
             var5      varchar(100),
             var6      varchar(100),
             var7      varchar(100),
             var8      varchar(100),
             var9      varchar(100),
            var10      varchar(100)
	    )`)
}

func (this *typAnalysisDB) ParseSql2Data(rows *sql.Rows) *TypList {
	retLst := List()

	for rows.Next() {
		var NotRead, Time, ServerId, typeId int
		var StrVars [9]string
		rows.Scan(&NotRead, &NotRead, &ServerId, &Time, &typeId, &StrVars[0], &StrVars[1], &StrVars[2],
			&StrVars[3], &StrVars[4], &StrVars[5], &StrVars[6], &StrVars[7], &StrVars[8])
		newMap := HashMap(":serverid", ServerId, ":time", Time, ":type", typeId)
		for i := 0; i < 9; i++ {
			if StrVars[i] != "" {
				newMap.Assoc(":var"+Str(i), StrVars[i])
			}
		}
		retLst.Conj(newMap)
	}
	return retLst
}

func (this *typAnalysisDB) QuerySql(sqlStr string) *TypList {
	rows, err := AnalysisDB.AnalysisDB.Query(sqlStr)
	if err != nil {
		panic(err)
	}

	return this.ParseSql2Data(rows)
}

func (this *typAnalysisDB) GetTableCount(tableName string, whereStr string) int {
	row := this.AnalysisDB.QueryRow(`select count(*) from ` + tableName + " where " + whereStr)
	var countNumber int
	row.Scan(&countNumber)
	return countNumber
}

func (this *typAnalysisDB) GetPayPlayerNumber(serverId int, minDay int, maxDay int) int {

	querySql := `select count(id) from log_data	where serverid=` + ToStr(serverId) +
		`and type = 331 and  var5 = 1 and time >` + ToStr(minDay) + " and time < " + ToStr(maxDay)
	row, err := this.AnalysisDB.Query(querySql)
	if err != nil {
		fmt.Println(err)
	}
	var PlayerNumber int
	row.Scan(&PlayerNumber)
	return PlayerNumber
}
