package moyoLogic

import (
	"bufio"
	//"database/sql"
	"io/ioutil"
	cfg "lib/config"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	. "git.oschina.net/yangdao/extlib"
	"git.oschina.net/yangdao/extlib/sqldb_helper"
)

var GameLogLoader typGameLogLoader

type typGameLogLoader struct {
}

func (this *typGameLogLoader) Load() {
	AnalysisDB.Connect()
	if sqldb_helper.IsExitTable(AnalysisDB.AnalysisDB, "LogIndex") == false {
		AnalysisDB.GenLogTable()
	}

	strLogPath := cfg.Get()["logPath"]
	folders, _ := ioutil.ReadDir(strLogPath)
	for _, folderinfo := range folders {
		if !folderinfo.IsDir() {
			continue
		}
		curServerId := Int(folderinfo.Name())
		filepath.Walk(strLogPath+"/"+Str(folderinfo.Name()), func(path string, f os.FileInfo, err error) error {
			if f.IsDir() {
				return nil
			}
			//判断是否已经加载过
			logKey := strings.Split(f.Name(), ".")[0]
			reqStr := "select count(id) from log_index where serverid = " + folderinfo.Name() + " and logkey = " + logKey
			row := AnalysisDB.AnalysisDB.QueryRow(reqStr)
			var cur_index_number int
			row.Scan(&cur_index_number)
			if cur_index_number > 0 {
				return nil
			}

			if _, err := strconv.Atoi(logKey); err == nil {
				this.LoadLogFile2Sql(path, curServerId, logKey)
			}
			return nil
		})
	}
}

func (this *typGameLogLoader) LoadLogFile2Sql(path string, serverid int, logkey string) {
	insertIndexSql := "insert into log_index(serverid,logkey) values (" + Str(serverid) + "," + Str(logkey) + ")"
	AnalysisDB.AnalysisDB.Exec(insertIndexSql)

	tx, _ := AnalysisDB.AnalysisDB.Begin()
	stmt, _ := tx.Prepare(`insert into 
		log_data(logkey,serverid,time,type,var2,var3,var4,var5,var6,var7,var8,var9,var10) 
		  values(?,?,?,?,?,?,?,?,?,?,?,?,?)`)
	fi, _ := os.Open(path)
	rd := bufio.NewReader(fi)
	for {
		bydata, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		lineArr := strings.Split(string(bydata), "|")
		intTime := Int(lineArr[0])
		intType := Int(lineArr[1])
		getStrArrVal := func(index int) string {
			if index < len(lineArr) {
				return lineArr[index]
			}
			return ""
		}
		stmt.Exec(Int(logkey), serverid, intTime, intType,
			getStrArrVal(2), getStrArrVal(3), getStrArrVal(4), getStrArrVal(5),
			getStrArrVal(6), getStrArrVal(7),
			getStrArrVal(8), getStrArrVal(9), getStrArrVal(10))
	}
	fi.Close()
	stmt.Close()
	tx.Commit()

}
