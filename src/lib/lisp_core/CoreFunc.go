package lisp_core

import (
	. "lib/lisp_core/data_type"
	"strconv"
)

type Action func(...interface{})
type Func func(...interface{}) interface{}

func Map(args ...interface{}) ISequence {
	fn := args[0].(func(...interface{}) interface{})

	return fn().(ISequence)
}

func Str(args ...interface{}) string {
	retStr := ""
	for _, arg := range args {
		switch val := arg.(type) {
		case int64:
			retStr += strconv.Itoa(int(val))
		case int:
			retStr += strconv.Itoa(val)
		case string:
			retStr += val
		case bool:
			if args[0].(bool) == true {
				return "true"
			} else {
				return "false"
			}
		}
	}
	return retStr
}

func Int(arg interface{}) int {
	switch val := arg.(type) {
	case int:
	case string:
		RetInt, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		return RetInt
	}
	return 0
}

func Filter(fn func(interface{}) bool, seqArg ISequence) ISequence {
	retlst := List()
	iter := seqArg.GetIterator()
	for iter.MoveNext() {
		if fn(iter.Current()) == true {
			retlst.Conj(iter.Current())
		}
	}
	return retlst
}
