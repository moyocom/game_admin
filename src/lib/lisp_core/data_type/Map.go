package data_type

type TypMap struct {
	Vals map[interface{}]interface{}
}

func (t *TypMap) Conj(args ...interface{}) {

}

/////ICollection

func HashMap(args ...interface{}) *TypMap {
	retMap := &TypMap{}
	retMap.Vals = make(map[interface{}]interface{})
	for i := 0; i < len(args); i += 2 {
		retMap.Vals[args[i]] = args[i+1]
	}
	return retMap
}

func (t *TypMap) Count() int {
	return len(t.Vals)
}

func (t *TypMap) Empty() ICollection {
	retMap := &TypMap{Vals: make(map[interface{}]interface{})}
	return retMap
}

/////IAssociative
func (t *TypMap) Assoc(k interface{}, v interface{}) {
	t.Vals[k] = v
}

func (t *TypMap) Dissoc(k interface{}) {
	delete(t.Vals, k)
}

func (t *TypMap) Get(k interface{}) interface{} {
	return t.Vals[k]
}

func (t *TypMap) ContainsKey(k interface{}) bool {
	_, ok := t.Vals[k]
	return ok
}

/////ISequence
func (t *TypMap) GetIterator() IIterator {
	return genMapIterator(t)
}

type MapIterator struct {
	curMap  *TypMap
	keys    []interface{}
	keyi    int
	isFirst bool
}

func genMapIterator(setMap *TypMap) *MapIterator {
	retMapIterator := &MapIterator{}
	retMapIterator.curMap = setMap
	retMapIterator.keyi = 0
	retMapIterator.isFirst = true
	retMapIterator.keys = make([]interface{}, 0, len(setMap.Vals))
	for k, _ := range setMap.Vals {
		retMapIterator.keys = append(retMapIterator.keys, k)
	}
	return retMapIterator
}

func (m *MapIterator) MoveNext() bool {
	if m.isFirst == true {
		m.isFirst = false
		if len(m.keys) == 0 {
			return false
		} else {
			return true
		}
	}
	m.keyi++
	if m.keyi >= len(m.keys) {
		return false
	}
	return true
}

func (m *MapIterator) Current() interface{} {
	return m.curMap.Vals[m.keys[m.keyi]]
}
