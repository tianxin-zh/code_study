package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	// 1、基础字段排序，int64、string
	strArray := []string{"ab", "123", "hdja"}
	slices.Sort(strArray)
	fmt.Println(strArray)

	slices.SortFunc(strArray, func(a, b string) int {
		return cmp.Compare(b, a)
	})
	fmt.Println(strArray)

	int64Array := []int64{163781368913618931, 2, 3}
	slices.Sort(int64Array)
	fmt.Println(int64Array)

	fmt.Println(slices.IsSorted(int64Array))

	fmt.Println(slices.Min(strArray))
	fmt.Println(slices.Max(strArray))
	// 2、结构体排序
	type Student struct {
		Name string
		Age  int64
	}

	studentArray := []Student{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 22},
	}
	slices.SortFunc(studentArray, func(a, b Student) int {
		return int(a.Age - b.Age)
	})
	fmt.Println(studentArray)

	// 3、旧版本sort排序
	sort.Slice(studentArray, func(i, j int) bool {
		return studentArray[i].Age > studentArray[j].Age
	})
	fmt.Println(studentArray)
}
