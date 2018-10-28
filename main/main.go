package main

import (
	"fmt"
	"math/rand"
	"sort"
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
a:=[]int{2,1,4,9,3}
sort.Ints(a)

fmt.Println(a)

	}
