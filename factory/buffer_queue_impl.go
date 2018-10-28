package factory

import (
	"Demo/role"
	"container/list"
)

type BufferQueueImp struct {
	list *list.List
}

func NewBufferQueueImpl() role.BufferQueue {
	return &BufferQueueImp{list.New()}
}

func (q *BufferQueueImp) EnQueue(v interface{}) {
	q.list.PushBack(v)
}

func (q *BufferQueueImp) DeQueue() interface{} {
	v := q.list.Front().Value
	q.list.Remove(q.list.Front())
	return v
}

func (q *BufferQueueImp) IsEmpty() bool {
	return q.list.Len()==0
}

func (q *BufferQueueImp) Clear() {
	for e:= q.list.Front(); e != nil;  {
		tmp:=e.Next()
		q.list.Remove(e)
		e=tmp
	}
}

func (q *BufferQueueImp) Len() int {
	return q.list.Len()
}
