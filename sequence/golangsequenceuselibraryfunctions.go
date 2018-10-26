package main

import (
	"fmt"
	"sort"
)

func arrSequenceAsc(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func arrSequenceDesc(nums []int) []int {

	fmt.Println(sort.IntSlice(nums))
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums
}

type IntSlice []int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
func arrSequenceTest(nums []int) []int {
	sort.Sort(IntSlice(nums))
	return nums
}

type Student struct {
	name  string
	score int
}
type Students []Student

//    ========    //
func (s Students) Len() int {
	return len(s)
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Students) Less(i, j int) bool {
	if s[i].score < s[j].score {
		return true
	}
	return false
}

func main() {

	//res := arrSequenceAsc([]int{1, 23, 56, 2, 3, 45, 6, 2})
	//mt.Println(res)
	//res := arrSequenceDesc([]int{1, 23, 56, 2, 3, 45, 6, 2})
	//fmt.Println(res)
	//res := arrSequenceTest([]int{1, 23, 56, 2, 3, 45, 6, 2})
	//fmt.Println(res)
	//首先向学生结构数组里面增加学生
	s := make(Students, 0, 5)
	s = append(s, Student{name: "gwr", score: 100})
	s = append(s, Student{name: "ly", score: 98})
	s = append(s, Student{name: "zxy", score: 50})
	s = append(s, Student{name: "zzz", score: 3})
	s = append(s, Student{name: "lisa", score: 6})

	sort.Sort(s)

	for k, v := range s {
		fmt.Println(k, v)
	}

}
