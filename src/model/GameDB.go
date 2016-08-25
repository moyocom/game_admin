package model

import (
	"database/sql"
	"fmt"
	_ "lib/mysql"
	"strconv"
)

type TypGameDB struct {
	ConnStr string
	GameDB  *sql.DB
}

func NewGameDB(dbUser string, dbPwd string, host string) *TypGameDB {
	newGameDB := &TypGameDB{}
	newGameDB.ConnStr = dbUser + ":" + dbPwd + "@tcp(" + host + ":3306)/yulong_game"
	return newGameDB
}

func (this *TypGameDB) Connect() {

	GameDB, err := sql.Open("mysql", this.ConnStr)
	fmt.Println(this.ConnStr)
	if err != nil {
		fmt.Println("游戏数据库连接失败")
		fmt.Println(err)
	} else {
		fmt.Println("游戏数据库连接成功")
	}
	this.GameDB = GameDB
}

func (this *TypGameDB) Close() {
	this.GameDB.Close()
}

func (this *TypGameDB) GetNewPlayerByDay(minTime int, maxTime int) int {

	querySql := "select count(id) from player where reg_time > " + strconv.Itoa(minTime) + " and reg_time <" + strconv.Itoa(int(maxTime))
	var NewPlayerNumber int
	row := this.GameDB.QueryRow(querySql)
	row.Scan(&NewPlayerNumber)
	//fmt.Println(querySql)
	return NewPlayerNumber

}
