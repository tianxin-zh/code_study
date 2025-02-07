package main

import (
	"fmt"
	"testing"
)

type People struct {
	Name string
	Age  int
}

type Student struct {
	People // 结构的嵌入,如果带了名字就无法直接访问
	Score  int
}

func TestStruct(T *testing.T) {
	// 主要看一下结构的嵌入
	student := Student{
		People: People{
			Name: "Alice",
			Age:  20,
		},
		Score: 90,
	}

	fmt.Println(student.Name) // 结构体嵌入可以直接访问
	fmt.Println(student.People.Name)
}
