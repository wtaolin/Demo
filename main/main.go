package main

import (
	"Demo/factory"
	"fmt"
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
	q:=3
	p:=2
	fmt.Println(factory.MirrorReflection(q,p))
	}
