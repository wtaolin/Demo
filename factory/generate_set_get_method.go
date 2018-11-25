package factory

import (
	"Demo/role"
	"fmt"
	"io/ioutil"
	log2 "log"
	"os"
	"path/filepath"
	"reflect"
)

type person struct {
	Name string
	Age int
}

func NewPerson()*person  {
	return &person{}
}
type uer struct {
	Aa int
	Bb string
	Person person
}

func GenerateSetAndGetMethod() {
	file, i := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR, 0)
	if i != nil {
		fmt.Println("open file failed")
	}
	logger := log2.New(file, "\r\n", log2.Ltime|log2.Ldate)
	path, e := filepath.Abs(".")
	fmt.Println(path)
	if e != nil {
		logger.Printf("err is %+v", e)
	}
	//filename, err := os.Create(path + "Generate.go")
	//if err != nil {
	//	logger.Printf("create Generate failed err is \n",err)
	//}
	uer:=uer{}
	uerType:=reflect.TypeOf(uer)
	fmt.Printf("uer type is %+v",uerType)
	//uerValue:=reflect.ValueOf(uer)
	fmt.Println(uerType.NumField())
	s:=""
	for i:=0;i<uerType.NumField() ;i++  {
		sub:="func (m *uer)set"+uerType.Field(i).Name+"("+uerType.Field(i).Name+" "+uerType.Field(i).Type.String()+")"+"{ \n" +
			"m."+uerType.Field(i).Name+"="+uerType.Field(i).Name+"\n"+"} \n\n"
		s+=sub
	}
	ioutil.WriteFile(path + "/Generate.go",[]byte(s),0777)
}

func TestReflecType( queue role.BufferQueue){
	roleType:=reflect.TypeOf(queue)
	fmt.Printf("%#v \n",roleType)
	fmt.Printf("%#v \n",roleType.String())
	//fmt.Printf("%#v \n",roleType.Field(0).Name)
}
