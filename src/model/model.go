package model

import (
	"database/sql"
	"fmt"
	. "lib/Util"
	_ "lib/mysql"
	"lib/session"
	_ "lib/session/memory"
)

var SQLDB *sql.DB
var Gsession *session.Manager

func Init() {
	var err error
	SQLDB, err = sql.Open("mysql", "game:ghgame@tcp(192.168.1.102:3306)/go_webadmin")
	CheckErr(err)
	err = SQLDB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	SQLDB.SetMaxOpenConns(2000)
	SQLDB.SetMaxIdleConns(1000)
	fmt.Println("数据库连接成功", SQLDB)

	Gsession, err = session.NewManager("memory", "gosessionid", 3600)
	CheckErr(err)
	go Gsession.GC()
}
