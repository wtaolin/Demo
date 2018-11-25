package main

import (
	"Demo/factory"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

type bb struct {
	a int
}
type name int

const (
	HHH name=iota
	eee
	hehe
	hehh

)
type element struct {
	a int
	b int
}
func main() {
//	factory.Json2Struct()
	//factory.GenerateSetAndGetMethod()

	queue:=&factory.BufferQueueImp{}
	factory.TestReflecType(queue)
	}
