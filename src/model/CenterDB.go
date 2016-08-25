package model

import (
	"database/sql"

	cfg "lib/config"
	_ "lib/mysql"

	"git.oschina.net/yangdao/extlib/sqldb_helper"
)

var CenterDB typCenterDB

type typCenterDB struct {
	CenterDB *sql.DB
}

func (this *typCenterDB) Connect() {
	if this.CenterDB != nil && this.CenterDB.Stats().OpenConnections > 0 {
		return
	}
	strCenterAddr := cfg.Get()["centerDB"]
	this.CenterDB = sqldb_helper.ConnectDB(strCenterAddr)
}

func (this *typCenterDB) GetServerTable() []*ServerData {
	return ServerData_Table()
}
