package model

import (
	"fmt"
	. "lib/lisp_core/data_type"
)

func Mail_Add(addMap IAssociative) int {
	stmt, err2 := GameDB.Prepare("insert into sys_mail(recv_ids,post_time,content,goods_type_id,goods_num,money_type,money_count,status) values(?,?,?,?,?,?,?,?)")
	if err2 != nil {
		fmt.Println(err2)
	}
	re, err := stmt.Exec(addMap.Get("recv_ids"), addMap.Get("post_time"), addMap.Get("content"),
		addMap.Get("goods_type_id"), addMap.Get("goods_num"), addMap.Get("money_type"), addMap.Get("money_count"), addMap.Get("status"))
	if err != nil {
		fmt.Println(err)
	}
	id, err3 := re.LastInsertId()
	if err3 != nil {
		fmt.Println(err3)
	}
	return int(id)
}
