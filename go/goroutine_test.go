package main

import (
	"sync"
	"testing"
)

// 协程（Goroutine） 是 Go 语言中用于并发执行的机制。
// 协程是轻量级的线程，由 Go 运行时管理。
// 协程之间通过通道（Channel）进行通信。
// 协程的创建和销毁开销较小，适合处理大量的并发任务。
// 协程可以在一个进程中并发执行，而线程需要在不同的进程中执行。

func Print(s string) {
	for i := 0; i < 5; i++ {
		println(s)
	}
}

func TestGoroutine(t *testing.T) {
	println("start")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Print("hello") // 启动一个协程
	}()
	wg.Wait()
	// go Print("hello") // 启动一个协程
	Print("world") // 主协程,不会等待子协程执行完毕
	println("end")
}
