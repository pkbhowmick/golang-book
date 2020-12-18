package main

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age int
}

type ByName []Person
type ByAge []Person

func (this ByName)Len() int {
	return len(this)
}

func (this ByName)Less(i,j int) bool {
	return this[i].Name < this[j].Name
}

func (this ByName)Swap(i,j int)  {
	this[i],this[j] = this[j],this[i]
}

func (this ByAge)Len() int {
	return len(this)
}

func (this ByAge)Less(i,j int) bool {
	return this[i].Age < this[j].Age
}

func (this ByAge)Swap(i,j int)  {
	this[i],this[j] = this[j],this[i]
}

func main()  {
	fmt.Println(
		// true
		strings.Contains("test", "es"),

		// 2
		strings.Count("test", "t"),

		// true
		strings.HasPrefix("test", "te"),

		// true
		strings.HasSuffix("test", "st"),

		// 1
		strings.Index("reset", "e"),

		// "a-b"
		strings.Join([]string{"a","b","c"}, "-"),

		// == "aaaaa"
		strings.Repeat("a", 5),

		// "bbaa"
		strings.Replace("aeaeaaaa", "a", "b", 3),

		// []string{"a","b","c","d","e"}
		strings.Split("a-b-c-d-e", "-"),

		// "test"
		strings.ToLower("TEST"),

		// "TEST"
		strings.ToUpper("test"),
	)
	// file io with os package
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Errorf("Error: %v\n",err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		fmt.Errorf("Error: %v\n",err)
		return
	}

	bs := make([]byte, stat.Size())
	_,err  = file.Read(bs)
	if err != nil {
		fmt.Errorf("Error: %v\n",err)
		return
	}

	str := string(bs)
	fmt.Println(str)

	// file io with io/ioutil package
	file2,err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Errorf("Error: %v\n",err)
		return
	}
	fmt.Println(string(file2))

	// read files from the directory
	dir, err := os.Open(".")
	if err != nil{
		fmt.Errorf("Error: %v\n",err)
		return
	}
	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil{
		fmt.Errorf("Error: %v\n",err)
		return
	}

	for _,fi :=range fileInfos{
		fmt.Println(fi.Name())
	}

	// container/list package
	var x list.List
	x.PushBack(2)
	x.PushFront(3)
	x.PushBack(1)

	for e:= x.Front(); e!=nil ; e=e.Next() {
		fmt.Println(e.Value)
	}

	// sort package
	xs := []int{2,1,3}
	sort.Ints(xs)
	fmt.Println(xs)


	kids := []Person{
		{"Bob",15},
		{"Alice", 20},
	}

	sort.Sort(ByName(kids))
	fmt.Println(kids)

	sort.Sort(ByAge(kids))
	fmt.Println(kids)


}