#### golang(排序

比如,给出一串数组[1,2,5,3,4],在利用golang自带库函数如何将它进行升序排序呢,实际上非常简单,代码如下
```
func arrSequenceAsc(nums []int) []int {
	sort.Ints(nums)
	return nums
}
```
看吧,非常简单就实现了一串数组的升序排序,那么问题来了,如何进行数组的降序排序呢(利用库函数)？
```
func arrSequenceDesc(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	return nums
}
```

> 为什么这样做?原理是什么?内部又使用了什么算法来排序？为啥sort.Ints就可以直接使用

首先我们来先看下`sort.Ints`的实现方式
```
// 升序排序
func Ints(a []int) { 
    Sort(IntSlice(a))
}

// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
type IntSlice []int

// Sort sorts data.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
// 这里使用快排来实现的
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}
```
实际上就是为了简便调用`golang`类库帮我们做了一些事情

而实际上我们调用的代码如下
```
type IntSlice []int
func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s IntSlice) Less(i, j int) bool { return s[i] > s[j] }
func arrSequenceTest(nums []int) []int {
	sort.Sort(IntSlice(nums)) //此处必须是slice类型
	return nums
}
```



##### 自定义排序
然而实际开发中会遇到各种各样的排序问题,比如`top k`问题
eg:给出一群学生,并按照学生的成绩排名,因为学生不可能只有分数属性,它还需要有性别等属性,所以属于根据一个结构体类型的Index进行排序
```
type Student struct {
	name  string
	score int
}
type Students []Student

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

```




