package role


type  BufferQueue interface {
	EnQueue(v interface{})
	DeQueue()interface{}
	IsEmpty()bool
	Clear()
	Len()int
}