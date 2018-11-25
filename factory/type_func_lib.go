package factory

import (
	"fmt"
)

type props func() *person

func Demo(pro  props)  {
	fmt.Printf("pro :%#v \n ",pro)
}

