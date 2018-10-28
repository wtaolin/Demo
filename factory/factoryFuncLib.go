package factory

import (
	"container/heap"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

type (
	ss struct {
		a int
		b int
	}
	yy struct {
		n int
		x int
	}
	sss struct {
		//	test.Q
		c int
	}
)
type dd map[string]int
type (
	hh interface {
		insert(a string, b int)
	}
	he struct {
		User dd
	}
)

func (c *he) insert(a string, b int) {
	c.User[a] = b
}
func test1(aa hh) {
	aa.insert("hhh", 2)
	fmt.Println(aa)
}

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type wendu uint

func practise() {
	var1 := new(int32)
	*var1 = 2
	fmt.Printf("%T   %v", var1, var1)
	var2 := (*int32)(var1)
	fmt.Printf("%T   %v", var2, var2)
}
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // success
		}
		log.Printf("server not responding (%s);retrying…", err)
		time.Sleep(time.Second << uint(tries)) // exponential back-off
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}

func errTest(a int) error {
	if a == 0 {
		return errors.New("some wrongs happend")
	}
	return nil
}

// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func Squares(a int) func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

//topo排序的例子
// prereqs记录了每个课程的前置课程
var Prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func TopoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

//函数跟踪，记录函数的起始时刻和持续时间
func BigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
func IsPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for ; start < end; start, end = start+1, end-1 {
		if s[start] != s[end] {
			return false
		}
	}
	return true
}
func FmtTest(s string) {
	fmt.Printf("%v", s)
}

//fuzzy query
func fuzzyQuery(str string, strs []string) []string {

	lens := len(str)
	queryResult := make([]string, 0)
	for _, subStr := range strs {
		for i := 0; i+lens < len(subStr)+1; i++ {
			if subStr[i:i+lens] == str {
				queryResult = append(queryResult, subStr)
				break
			}
		}
	}
	sort.Strings(queryResult)
	return queryResult
}

func PractiseWriteRead() {
	file, err := os.OpenFile("text.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalln("read file failed", err)
	}
	defer file.Close()
	//fileReader:=bufio.NewReader(file)
	all, err := ioutil.ReadAll(file)
	fmt.Println(string(all))
	s := "are you ok"
	file.WriteString(s)

	logFile, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE, 0)
	if err != nil {
		log.Fatalln("read log.txt failed", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "\r\n", log.Ldate|log.Ltime)
	logger.Printf("prepare to write log")

}

type maskEvent uint64
func maskCnt(mask maskEvent) (uint32, error) {
	if mask <= 0 {
		return 0, errors.New("mask is invalid")
	}
	var cnt uint32
	for {
		if mask&1 != 1 {
			cnt++
			mask >>= 1
		} else {
			break
		}
	}
	return cnt, nil
}
func RemoveElement(nums []int, val int) int {
	cnt:=0
	for i:=0;i<len(nums) ;i++  {
		if nums[i]==val {
			nums=append(nums[:i],nums[i:]...)
			i--
			cnt++
		}
	}
	return len(nums)-cnt
}

func longestPalindrome(s string) string {
	maxLenOfStr:=0
	longestPliStr:=""
	for i:=0;i<len(s) ;i++  {
		for j:=i+1;j<len(s) ;j++  {
			if IsPalindrome(s[i:j]) {
				if maxLenOfStr<=j-i {
					maxLenOfStr=j-i
					longestPliStr=s[i:j]
				}
			}
		}
	}
	return longestPliStr
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// copied from golang doc
// mininum setup of integer min heap
