package model

import (
	"fmt"
	"net/http"
)

type AdminUser struct {
	ID       int
	UserName string
	UserPwd  string
	UserType int
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

func AdminUser_FindById(id string) *AdminUser {
	sqlstr := "select * from go_adminuser where ID = " + id
	row := SQLDB.QueryRow(sqlstr)
	queryUser := &AdminUser{}
	err := row.Scan(&queryUser.ID, &queryUser.UserName, &queryUser.UserPwd, &queryUser.UserType)
	if err != nil {
		fmt.Println(err, sqlstr)
		return nil
	}
	return queryUser
}

func AdminUser_Find(userName string) *AdminUser {
	sqlstr := "SELECT * FROM go_adminuser WHERE UserName = \"" + userName + "\""

	row := SQLDB.QueryRow(sqlstr)

	queryUser := &AdminUser{}
	err := row.Scan(&queryUser.ID, &queryUser.UserName, &queryUser.UserPwd, &queryUser.UserType)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return queryUser
}
