package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
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
	HHH name = iota
	eee
	hehe
	hehh
)

type element struct {
	a []int
	b []string
}

func add(n []element) {
	num := 10
	for _, v := range n {
		for i := 0; i < num; i++ {
			v.a = append(v.a, i)
		}
	}
}



func main() {
}

func TestSlice() {
	n := []element{
		{
			a: make([]int, 2, 100),
		},
		{
			a: make([]int, 2, 100),
		},
	}
	add(n)
	for _, v := range n {
		for _, i := range v.a {
			fmt.Printf("%d ", i)
		}
	}
}

func httpGetTest() {
	url := "http://www.w3school.com.cn/example/html_examples.asp"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln("get rsp err")
	}
	if resp.StatusCode != 200 {
		log.Fatalln("statusCode is wrong")
	}
	str, _ := filepath.Abs(".")
	file, err := os.Create(str + "/text.txt")
	if err != nil {
		log.Fatalln("create file failed")
	}
	defer file.Close()
	dst := io.MultiWriter(os.Stdout, file)
	written, err := io.Copy(dst, resp.Body)
	fmt.Println(written)
	if err := resp.Body.Close(); err != nil {
		log.Fatalln("resp.body close failed")
	}
	//ioutil.WriteFile(str+"/text.txt",[]byte(resp),0777)
}
