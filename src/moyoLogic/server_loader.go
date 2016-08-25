package moyoLogic

import (
	. "git.oschina.net/yangdao/extlib/data_type"
)

type typServerLoader struct {
	lineFn func(int, string) *TypMap
}

func (this *typServerLoader) SetLineParse(fn func(int, string) *TypMap) {
	this.lineFn = fn
}

func (this *typServerLoader) Load() {

}

/*
 109
   16072219.log
   16072220.log
   16072221.log
   16072222.log
 110
   16072219.log
   16072220.log
   16072221.log
   16072222.log
 111
   16072219.log
   16072220.log
   16072221.log
   16072222.log
*/
