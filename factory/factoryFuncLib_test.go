package factory

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	input := "heeh"
	if !IsPalindrome(input) {
		t.Error("IsPalindrome(hello)==false")
	}
	//t.Error("IsPalindrome(hello)==true")
	///t.FailNow()
}

func TestFuzzyQuery(t *testing.T) {
	strs := []string{"test1", "Test", "ttestu", "atest"}
	str := "test"
	result := []string{"atest","test1", "ttestu"}
	queryResult:=fuzzyQuery(str,strs)
	fmt.Println(queryResult)
	for i,subStr:= range queryResult   {
		if result[i]!=subStr {
			t.Errorf(" \n want \n %#v \n got \n %#v \n",result[i],subStr)
		}
	}
}

func TestMaskCnt(t *testing.T)  {
	input:=[]uint64{1,2,4,8,16,4096}
	result:=[]uint32{0,1,2,3,4,12}
	for i,x:=range input {
		y,_:=maskCnt(maskEvent(x))
		if y!=result[i] {
			t.Errorf("\n want %v \n got %v \n",result[i],y)
		}
	}
}
func TestRmoveElement(t *testing.T)  {
	nums:=[]int{3,2,2,3}
	result:=[]int{2,2}
	val:=3
	if RemoveElement(nums,val)!=2 {
		t.FailNow()
	}
	if !reflect.DeepEqual(nums,result) {
		t.FailNow()
	}
}