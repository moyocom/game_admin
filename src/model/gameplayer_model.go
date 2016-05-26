package model

import (
	"fmt"
	. "lib/lisp_core"
	. "lib/lisp_core/data_type"
)

type GamePlayer struct {
	Id       int
	AccName  string
	Name     string
	Vip      int
	Lvl      int
	State    int //2禁登录 3禁言
	GMId     int //对应GMCtrl的Id
	IsOnline bool
}

func GamePlayer_MaxNumber() int {
	var Number int
	row := GameDB.QueryRow("select COUNT(id) as Number from player")
	err := row.Scan(&Number)
	if err != nil {
		fmt.Println(err, "err")
	}
	return int(Number)
}

func GamePlayer_Table(start string, length string) *TypVector {
	newVec := Vector()
	sqlStr := "select id,name,vip,lvl from player limit " + start + "," + length
	rows, err := GameDB.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		playerData := &GamePlayer{}
		rows.Scan(&playerData.Id, &playerData.Name, &playerData.Vip, &playerData.Lvl)
		newVec.Conj(playerData)
	}
	//过滤掉封IP的
	ctrl_lst := GMCtrl_Table()
	ctrl23_lst := Filter(func(val interface{}) bool {
		ctrl_type := val.(IAssociative).Get("ctrl_type").(int32)
		if ctrl_type == 2 || ctrl_type == 3 {
			return true
		}
		return false
	}, ctrl_lst)

	//赋入GM状态
	iter := ctrl23_lst.GetIterator()
	for iter.MoveNext() {
		GM_Table := iter.Current().(IAssociative)
		id := GM_Table.Get("target").(int)
		iterPlayer := newVec.GetIterator()
		for iterPlayer.MoveNext() {
			player := iterPlayer.Current().(*GamePlayer)
			if player.Id == id {
				player.State = int(GM_Table.Get("ctrl_type").(int32))
				player.GMId = id
			}
		}
	}

	//赋入在线状态
	onlineSeq := GetOnLinePlayerId().Get(APIServerId).(ISequence)
	iterPlayer := newVec.GetIterator()
	for iterPlayer.MoveNext() {
		player := iterPlayer.Current().(*GamePlayer)
		onlineIter := onlineSeq.GetIterator()
		for onlineIter.MoveNext() {
			onlineId := onlineIter.Current().(int)
			if onlineId == player.Id {
				player.IsOnline = true
			}
		}
	}
	return newVec
}

func GamePlayerTable2Json(seq ISequence) string {
	retStr := ""
	iter := seq.GetIterator()
	for iter.MoveNext() {
		player := iter.Current().(*GamePlayer)
		retStr += "[" +
			Str(player.Id) + "," +
			"\"" + Str(player.Name) + "\"," +
			Str(player.Lvl) + "," +
			Str(player.Vip) + "," +
			Str(player.State) + "," +
			Str(player.IsOnline) + "," + "\"操作\"" +
			"],"
	}
	return retStr[:len(retStr)-1]
}

func GMCtrl_Table() ISequence {
	sqlStateStr := "select * from gm_ctrl"
	query, _ := GameDB.Query(sqlStateStr)
	lst := List()
	for query.Next() {
		var id, begin_time, end_time int
		var ctrl_type int32
		var target int
		query.Scan(&id, &ctrl_type, &begin_time, &end_time, &target)
		dic := HashMap("id", id, "ctrl_type", ctrl_type, "begin_time", begin_time, "end_time", end_time, "target", target)
		lst.Conj(dic)
	}
	return lst
}
