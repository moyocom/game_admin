package model

import (
	"fmt"
	"lib/Util"
	"net/http"
	"strconv"
	"time"
)

type AdminUser struct {
	ID            int
	UserName      string
	NickName      string
	UserPwd       string
	Email         string
	Contact       string
	UserType      int
	RegTime       int64
	LastLoginTime int64
}

func (user *AdminUser) LoginTimeStr() string {
	return Util.TimeInt2Str(user.LastLoginTime)
}

func (user *AdminUser) UserTypeStr() string {
	switch user.UserType {
	case 0:
		return "超级管理员"
	case 1:
		return "管理员"
	}
	return ""
}

func AdminUser_CurUserId(w http.ResponseWriter, r *http.Request) string {
	sess := Gsession.SessionStart(w, r)
	userid := sess.Get("User")
	if userid == nil {
		return ""
	}
	return userid.(string)
}

func AdminUser_CurUser(w http.ResponseWriter, r *http.Request) *AdminUser {
	idstr := AdminUser_CurUserId(w, r)
	if idstr != "" {
		return AdminUser_FindById(idstr)
	}
	return nil
}

func AdminUser_UpdateLoginTime(id int) {
	sqlstr := "update go_adminuser set LastLoginTime = " + strconv.Itoa(int(time.Now().Unix())) + " where ID=" + strconv.Itoa(id)
	SQLDB.Exec(sqlstr)
}

func AdminUser_FindById(id string) *AdminUser {
	sqlstr := "select * from go_adminuser where ID = " + id
	row := SQLDB.QueryRow(sqlstr)
	queryUser := &AdminUser{}
	err := row.Scan(&queryUser.ID,
		&queryUser.UserName,
		&queryUser.NickName,
		&queryUser.UserPwd,
		&queryUser.UserType,
		&queryUser.Email,
		&queryUser.Contact,
		&queryUser.RegTime,
		&queryUser.LastLoginTime)
	if err != nil {
		fmt.Println(err, sqlstr)
		return nil
	}
	return queryUser
}

func AdminUser_Find(userName string) *AdminUser {
	sqlstr := "select * from go_adminuser WHERE UserName = \"" + userName + "\""
	row := SQLDB.QueryRow(sqlstr)
	queryUser := &AdminUser{}
	err := row.Scan(&queryUser.ID,
		&queryUser.UserName,
		&queryUser.NickName,
		&queryUser.UserPwd,
		&queryUser.UserType,
		&queryUser.Email,
		&queryUser.Contact,
		&queryUser.RegTime,
		&queryUser.LastLoginTime)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return queryUser
}

func AdminUser_Table() []*AdminUser {
	rows, _ := SQLDB.Query("select * from go_adminuser")
	var RetUserlist []*AdminUser
	for rows.Next() {
		queryUser := &AdminUser{}
		rows.Scan(&queryUser.ID,
			&queryUser.UserName,
			&queryUser.NickName,
			&queryUser.UserPwd,
			&queryUser.UserType,
			&queryUser.Email,
			&queryUser.Contact,
			&queryUser.RegTime,
			&queryUser.LastLoginTime)
		RetUserlist = append(RetUserlist, queryUser)
	}
	return RetUserlist
}

func AdminUser_AddNewUser(newUser *AdminUser) {
	if newUser == nil {
		return
	}
	stmt, _ := SQLDB.Prepare("insert into go_adminuser(UserName,NickName,UserPwd,UserType,RegTime,LastLoginTime,Email,Contact)values(?,?,?,?,?,?,?,?)")
	_, err := stmt.Exec(newUser.UserName, newUser.NickName, newUser.UserPwd, newUser.UserType, newUser.RegTime, newUser.LastLoginTime, newUser.Email, newUser.Contact)
	if err != nil {
		fmt.Println(err)
	}
}

func AdminUser_Update(id string, newValDic map[string]interface{}) {
	SQLDB.Prepare("update go_adminuser set")
	setValStr := ""
	for k, v := range newValDic {
		setValStr += k + " = " + v.(string) + ","
	}
	setValStr = setValStr[:len(setValStr)-1]
	_, err := SQLDB.Exec("update go_adminuser set " + setValStr + " where ID=" + id)
	if err != nil {
		fmt.Println(err)
	}
}
