package model

import (
	"database/sql"
	"fmt"
	. "lib/lisp_core"
	"strconv"
)

type Notice struct {
	Id       int
	Title    string
	Content  string
	ServerId int
}

func Notice_GetById(id int) *Notice {
	row := SQLDB.QueryRow("select * from go_notice where serverid =" + Str(id))
	retNotice := &Notice{}

	err := row.Scan(&retNotice.Id, &retNotice.Title, &retNotice.Content, &retNotice.ServerId)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return retNotice
}

func Notice_Add(notice *Notice) int {
	stmt, err := SQLDB.Prepare("insert into go_notice(title,content,serverid)values(?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	r, _ := stmt.Exec(notice.Title, notice.Content, notice.ServerId)
	id, _ := r.LastInsertId()
	return int(id)
}

func Notice_Update(notice *Notice, Id int) {
	stmt, err := SQLDB.Prepare("update go_notice set title= ? , content= ? where id=" + Str(Id))
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := stmt.Exec(notice.Title, notice.Content)
	if err2 != nil {
		fmt.Println(err)
	}
}

type RollNotice struct {
	Id       int
	PostTime int
	PlanTime int
	EndTime  int
	Period   int
	Status   int
	Title    string
	Content  string
	ServerID string
}

func RollNotice_Table() []*RollNotice {
	retRollNotice := make([]*RollNotice, 0)
	query, _ := GameDB.Query("select * from notice")
	for query.Next() {
		rollNotice := goQuery2RollNoticeStruct(query)
		retRollNotice = append(retRollNotice, rollNotice)
	}
	return retRollNotice
}

func RollNotice_Add(notice *RollNotice) int {
	stmt, err := GameDB.Prepare("insert into notice(post_time,plan_time,end_time,period,status,content)values(?,?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	r, _ := stmt.Exec(notice.PostTime, notice.PlanTime, notice.EndTime, notice.Period, notice.Status, notice.Content)
	id, _ := r.LastInsertId()
	return int(id)
}

func RollNotice_Del(id int) {
	GameDB.Exec("delete from notice where id = " + strconv.Itoa(id))
}

func RollNotice_Update(notice *RollNotice) {

}

func goQuery2RollNoticeStruct(rows *sql.Rows) *RollNotice {
	retRoll := &RollNotice{}
	rows.Scan(
		&retRoll.Id, &retRoll.PostTime,
		&retRoll.PlanTime, &retRoll.EndTime,
		&retRoll.Period, &retRoll.Status, &retRoll.Content)
	return retRoll
}
