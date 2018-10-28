package main

import (
	list2 "container/list"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
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
list:=list2.New()
	for i:=0;i<7 ;i++  {
		list.PushBack(i)
	}
	fmt.Printf("list len :%v \n",list.Len())
	for e:=list.Front();e!=nil ;e=e.Next()  {
		fmt.Printf("%v \n",e.Value)
	}
	a:=element{}
	b:=element{2,1}
	if reflect.DeepEqual(a,b) {
		fmt.Println("a==b")
	}
	bb:=reflect.TypeOf(b)
	fmt.Println(bb.String())
	s:="heafafafjajflaj"
	if 	strings.Contains(s,"he") {
		fmt.Printf("contain this substr \n")
	}
	fmt.Println(strings.Count(s,"afa"))
	}
