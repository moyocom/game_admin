package moyoLogic

import (
	"fmt"

	. "git.oschina.net/yangdao/extlib"
	. "git.oschina.net/yangdao/extlib/data_type"

	"database/sql"
	_ "lib/mysql"
)

var LogDataBase *TypLogDataBase

func GetMoYoLog() *TypLogDataBase {
	if LogDataBase == nil {
		LogDataBase = &TypLogDataBase{
			Loader:  &local_log_loader,
			LogData: HashMap(),
		}
		LogDataBase.Load()
	}
	return LogDataBase
}

type ILogLoader interface {
	SetLineParse(func(int, string) *TypMap)
	Load()
}

type TypLogDataBase struct {
	Loader      ILogLoader
	LogData     *TypMap
	LogTypeDesc *TypMap
}

func (this *TypLogDataBase) Load() {
	this.Loader.SetLineParse(this.parseLine)
	this.Loader.Load()
}

func (this *TypLogDataBase) ExportToSql(sqladdr string) {
	var err error
	LogDB, err := sql.Open("mysql", sqladdr)
	if err != nil {
		fmt.Println(err)
	}
	err = LogDB.Ping()
	if err != nil {
		fmt.Println(err)
	}

	ForEach(this.LogData, func(v interface{}) {
		//mapk := v.(IIndexed).Nth(0)
		mapv := v.(IIndexed).Nth(1)
		//createTableSql := `CREATE TABLE log_` + Str(mapk) + `(`
		firstMap := mapv.(*TypList).Head.Value.(*TypMap)
		ForEach(firstMap, func(v interface{}) {
			curlogK := v.(IIndexed).Nth(0)
			fmt.Println(curlogK)
		})

	})
}
