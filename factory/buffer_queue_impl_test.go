package factory

import "testing"

func TestBufferQueue(t *testing.T){
	input:=[]int{1,3,4,2,6}
	BufferQueue:=NewBufferQueueImpl()

	for _,x:=range input {
		BufferQueue.EnQueue(x)
	}
	if BufferQueue.Len()!=len(input) {
		t.FailNow()
	}
	for _,x:=range input {
		if BufferQueue.DeQueue()!=x {
			t.FailNow()
		}
	}
	if !BufferQueue.IsEmpty() {
		t.FailNow()
	}
	for _,x:=range input {
		BufferQueue.EnQueue(x)
	}
	BufferQueue.Clear()
	if !BufferQueue.IsEmpty() {
		t.FailNow()
	}
}