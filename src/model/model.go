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
var APIServerId int = 1

func Init() {
	APIServer = "http://60.205.95.15:8888"
	var err error
	CenterDB.Connect()
	SQLDB = CenterDB.CenterDB
	fmt.Println("CetnerDB 连接成功")

	Gsession, err = session.NewManager("memory", "gosessionid", 36000)
	CheckErr(err)
	go Gsession.GC()
}
