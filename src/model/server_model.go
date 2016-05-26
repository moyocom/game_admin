package model

import (
	"database/sql"
	"fmt"
	. "lib/lisp_core"
)

//{"memory_total":143883448,"memory_processes":96757568,"memory_processes_used":96751880,"memory_system":47125880,
//"memory_atom":695185,"memory_atom_used":668905,"memory_binary":4148080,"memory_code":16526191,"memory_ets":12952992,
//"process_count":435,"run_queue":1,
//"context_switch":301544229,"io_input_total":2296049775,"io_output_total":1573391919,"max_fds":102400,"wordsize":8}

type ServerState struct {
	Memory_total          int64 `json:"memory_total"`
	Memory_processes      int64 `json:"memory_processes"`
	Memory_processes_used int64 `json:"memory_processes_used"`
	Memory_system         int64 `json:"memory_system"`
	Memory_atom           int64 `json:"memory_atom"`
	Memory_atom_used      int64 `json:"memory_atom_used"`
	Memory_binary         int64 `json:"memory_binary"`
	Memory_code           int64 `json:"memory_code"`
	Memory_ets            int64 `json:"memory_ets"`
	Process_count         int64 `json:"process_count"`
	Run_queue             int64 `json:"run_queue"`
	Context_switch        int64 `json:"context_switch"`
	Io_input_total        int64 `json:"io_input_total"`
	Io_output_total       int64 `json:"io_output_total"`
	Max_fds               int64 `json:"max_fds"`
	Wordsize              int64 `json:"wordsize"`
}

type ServerData struct {
	Id      int
	Name    string
	Desc    string
	IP      string
	Port    string
	DBUser  string
	DBPwd   string
	State   int
	AddTime int64
}

func ServerData_Table() []*ServerData {
	query, _ := SQLDB.Query("select * from go_server_list")
	retData := make([]*ServerData, 0)
	for query.Next() {
		serverData := goQueryServerData2Struct(query)
		retData = append(retData, serverData)
	}
	return retData
}

func ServerData_ById(id int) *ServerData {
	sqlStr := "select * from go_server_list where id = " + Str(id)
	fmt.Println(sqlStr)
	query := SQLDB.QueryRow(sqlStr)
	retServerData := &ServerData{}
	err := query.Scan(&retServerData.Id, &retServerData.Name, &retServerData.Desc, &retServerData.IP, &retServerData.Port, &retServerData.DBUser,
		&retServerData.DBPwd, &retServerData.State, &retServerData.AddTime)
	if err != nil {
		fmt.Println(err, "query err")
	}
	return retServerData
}

func goQueryServerData2Struct(row *sql.Rows) *ServerData {
	retServerData := &ServerData{}
	row.Scan(&retServerData.Id, &retServerData.Name, &retServerData.Desc, &retServerData.IP, &retServerData.Port, &retServerData.DBUser,
		&retServerData.DBPwd, &retServerData.State, &retServerData.AddTime)
	return retServerData
}
