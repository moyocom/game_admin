package model

import (
	"fmt"
	"io/ioutil"
	. "lib/lisp_core/data_type"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var OnLinePlayerId *TypMap = HashMap()
var LastOnlinePullTime int64

func GetOnLinePlayerId() IAssociative {
	if LastOnlinePullTime == 0 || (time.Now().Unix()-LastOnlinePullTime) > int64(10) {
		fmt.Println("GetOnlinePlayer")
		UpdateOnlinePlayers()
	}
	return OnLinePlayerId
}

func fillOnLinePlayerId() *TypMap {
	servers := ServerData_Table()
	for _, server := range servers {
		if server.State == 1 {
			OnLinePlayerId.Assoc(server.Id, GetOnlineIdSeq("http://"+server.IP+":8888"))
		}
	}
	return OnLinePlayerId
}

func GetOnlineIdSeq(apiaddr string) ISequence {
	resp, err := http.Get(apiaddr + "/user/online?type=list")
	if err != nil {
		//fmt.Println("在线玩家数据获取失败", apiaddr+"/user/online?type=list")
		return List()
	}
	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	retLst := List()
	bodyStr = bodyStr[1 : len(bodyStr)-1]
	arr := strings.Split(bodyStr, ",")

	for _, v := range arr {
		intId, _ := strconv.Atoi(v)
		retLst.Conj(intId)
	}

	return retLst
}

func UpdateOnlinePlayers() {
	LastOnlinePullTime = time.Now().Unix()
	OnLinePlayerId = fillOnLinePlayerId()
}
