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
var GameDB *sql.DB

var Gsession *session.Manager

var APIServer string
var APIServerId int = 2

func Init() {
	APIServer = "http://192.168.1.116:8888"
	var err error
	SQLDB, err = sql.Open("mysql", "game:ghgame@tcp(192.168.1.102:3306)/go_webadmin")
	CheckErr(err)
	err = SQLDB.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("数据库连接成功", SQLDB)
	GameDB, err = sql.Open("mysql", "game:ghgame@tcp(192.168.1.116:3306)/yulong_game")
	CheckErr(err)
	GameDB.Ping()

	Gsession, err = session.NewManager("memory", "gosessionid", 36000)
	CheckErr(err)
	go Gsession.GC()
}
