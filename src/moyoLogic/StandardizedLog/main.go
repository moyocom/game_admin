package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"time"

	. "git.oschina.net/yangdao/extlib"
	. "git.oschina.net/yangdao/extlib/tools"
)

var OutFile map[int]*os.File

func main() {
	fmt.Println("Run ReLog", os.Args[0])
	LogPath := GetCurrPath()
	OutFile = make(map[int]*os.File)
	if len(os.Args) > 1 {
		LogPath = os.Args[1]
	}
	fmt.Println(LogPath)

	folders, _ := ioutil.ReadDir(LogPath)
	for _, folderinfo := range folders {
		//curServerId := Int(folderinfo.Name())
		ScanFolder(LogPath + "/" + folderinfo.Name())
	}
}

func ScanFolder(path string) {
	filepath.Walk(path, func(filepath string, f os.FileInfo, err error) error {
		if !strings.HasSuffix(f.Name(), ".log") {
			return nil
		}
		ScanLogFile(filepath, path)
		return nil
	})
}

func ScanLogFile(path string, folderPath string) {
	limitTime := time.Date(2016, 8, 1, 0, 0, 0, 0, time.Local).Unix()
	fi, _ := os.Open(path)
	rd := bufio.NewReader(fi)
	for {
		bydata, _, err := rd.ReadLine()
		if err != nil {
			break
		}
		strLine := string(bydata)
		lineArr := strings.Split(strLine, "|")
		intTime := Int(lineArr[0])
		intType := Int(lineArr[1])

		if intType == 331 && lineArr[5] != "12" && lineArr[5] != "4" {
			fmt.Println(strLine)
		}
		if int64(intTime) < limitTime {
			continue
		}

		timeHour := intTime / (60 * 60)
		curFile, ok := OutFile[timeHour]
		if !ok {

			f, _ := os.Create(folderPath + "/" + ToStr(timeHour))
			OutFile[timeHour] = f

		}
		curFile.WriteString(strLine + "\r\n")
	}

	fi.Close()
}
