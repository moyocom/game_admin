package moyoLogic

import (
	"bufio"
	//"fmt"
	//"io"
	. "git.oschina.net/yangdao/extlib/data_type"
	"os"
	"path/filepath"
)

var local_log_loader typLocalTestLoader

type typLocalTestLoader struct {
	linefn func(int, string) *TypMap
}

func (t *typLocalTestLoader) SetLineParse(fn func(int, string) *TypMap) {
	t.linefn = fn
}

func (this *typLocalTestLoader) Load() {
	filepath.Walk(`D:\Gopro\game_admin\src\main\log`, func(path string, f os.FileInfo, err error) error {
		if f == nil || f.IsDir() {
			return nil
		}
		fi, err := os.Open(`D:\Gopro\game_admin\src\main\log\` + f.Name())
		if err != nil {
			panic(err)
		}
		rd := bufio.NewReader(fi)
		for {
			bydata, _, err := rd.ReadLine()
			if err != nil {
				break
			}
			if this.linefn != nil {
				this.linefn(-1, string(bydata))
			}
		}
		fi.Close()
		return nil
	})
}
