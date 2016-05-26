package data_type

type TypList struct {
	count int
	Head  *TypListNode
}

func List(args ...interface{}) *TypList {
	RetList := &TypList{}
	var PreList *TypListNode
	for _, v := range args {
		curNode := &TypListNode{Value: v, Next: nil}
		if RetList.Head == nil {
			RetList.Head = curNode
			PreList = RetList.Head
			continue
		}

		PreList.Next = curNode
		PreList = curNode
		RetList.count++
	}
	return RetList
}

func (t *TypList) Count() int {
	return t.count
}

func (t *TypList) Empty() ICollection {
	return &TypList{}
}

func (t *TypList) Conj(val ...interface{}) {
	newVal := &TypListNode{Value: val[0], Next: t.Head}
	t.Head = newVal
}

func (t *TypList) GetIterator() IIterator {
	return &ListIterator{t.Head, true}
}

//====================================================================
type TypListNode struct {
	Value interface{}
	Next  *TypListNode
}

type ListIterator struct {
	CurElem *TypListNode
	isFirst bool
}

func (l *ListIterator) MoveNext() bool {
	if l.isFirst == true {
		l.isFirst = false
		if l.CurElem == nil {
			return false
		} else {
			return true
		}
	}
	if l.CurElem.Next != nil {
		l.CurElem = l.CurElem.Next
		return true
	}

	return false
}

func (l *ListIterator) Current() interface{} {
	return l.CurElem.Value
}
