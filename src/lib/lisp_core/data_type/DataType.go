package data_type

type ICollection interface {
	Count() int          //count
	Empty() ICollection  //empty
	Conj(...interface{}) //conj
	//Equal(ICollection) bool // =
}

type IIterator interface {
	MoveNext() bool
	Current() interface{}
}

type ISequence interface {
	GetIterator() IIterator
}

type IAssociative interface {
	Assoc(interface{}, interface{}) //assosc
	Dissoc(interface{})
	Get(interface{}) interface{}
	ContainsKey(interface{}) bool
}

type IIndexed interface {
	Nth(int) interface{}
}

type IStack interface {
	Pop() interface{}
	Peek() interface{}
	Conj(interface{}) //conj
}

type ISorted interface {
	RSeq() ISorted
	SubSeq() ISorted
	RSubSeq() ISorted
}
