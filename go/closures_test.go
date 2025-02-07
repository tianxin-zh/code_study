package main

import "testing"

// 闭包（Closure） 是引用外部变量的匿名函数，即使外部作用域的变量超出作用域，闭包仍然能访问这些变量。
func TestClosure(T *testing.T) {
	counter := Count()
	T.Log(counter())
	T.Log(counter())
}

func Count() func() int {
	count := 0
	return func() int { // 匿名函数，引用外部变量count
		count++
		return count
	}
}
