package data_type

//ICollection ISequence IAssociative IIndexed IStack
type TypVector struct {
	buffers []interface{}
	count   int
}

func (t *TypVector) Count() int {
	return t.count
}

func (t *TypVector) Conj(val ...interface{}) {
	if t.buffers == nil || t.count == cap(t.buffers) {
		t.allocateMore()
	}
	t.buffers[t.count] = val[0]
	t.count++
}

func (t *TypVector) allocateMore() {
	if t.buffers == nil {
		t.buffers = make([]interface{}, 2)
	} else {
		bflen := len(t.buffers)
		if bflen >= 2 {
			newbuffer := make([]interface{}, bflen<<1)
			copy(newbuffer, t.buffers)
			t.buffers = newbuffer
		}
	}
}

func (t *TypVector) Empty() ICollection {
	return &TypVector{}
}

func (t *TypVector) Get(idx interface{}) interface{} {
	return t.buffers[idx.(int)]
}

func (t *TypVector) Assoc(k interface{}, v interface{}) {
	if k.(int) == t.Count() {
		t.Conj(v)
	} else {
		t.buffers[k.(int)] = v
	}
}

func (t *TypVector) Dissoc(k interface{}) {
	t.RemoveAt(k.(int))
}

func (t *TypVector) ContainsKey(k interface{}) bool {
	if k.(int) >= 0 && k.(int) < t.count {
		return true
	}
	return false
}

func (t *TypVector) Nth(k int) interface{} {
	return t.buffers[k]
}

func (t *TypVector) RemoveAt(index int) {
	if t.buffers != nil && index < len(t.buffers) {

		for i := index; i < t.count; i++ {
			t.buffers[i] = t.buffers[i+1]
			t.buffers[t.count] = nil
		}
		t.count = t.count - 1
	}
}

func (t *TypVector) Peek() interface{} {
	return t.buffers[t.count]
}
func (t *TypVector) Pop() interface{} {
	retVal := t.buffers[t.count]
	t.RemoveAt(t.count)
	return retVal
}

func Vector(args ...interface{}) *TypVector {
	RetVec := &TypVector{}
	for _, v := range args {
		RetVec.Conj(v)
	}
	return RetVec
}

//=============================================================
func (t *TypVector) GetIterator() IIterator {
	return &VectorIterator{0, t, true}
}

type VectorIterator struct {
	curIndex int
	iterVec  *TypVector
	isFirst  bool
}

func (v *VectorIterator) MoveNext() bool {
	if v.isFirst == true {
		v.isFirst = false
		if v.iterVec.Count() == 0 {
			return false
		}
		return true
	}
	v.curIndex++
	if v.curIndex >= v.iterVec.Count() {
		return false
	}

	return true
}

func (v *VectorIterator) Current() interface{} {
	return v.iterVec.buffers[v.curIndex]
}
