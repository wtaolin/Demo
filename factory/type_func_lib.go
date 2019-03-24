package factory

import (
	"Demo/role"
	"fmt"
)

type props func() *person

func Demo(pro  props)  {
	fmt.Printf("pro :%#v \n ",pro)
}


type Tt struct {
	Buffer role.BufferQueue
}