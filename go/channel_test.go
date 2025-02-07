package main

import (
	"testing"
	"time"
)

// 单向通道函数
func ping(msg chan<- string) {
	time.Sleep(5 * time.Second)
	msg <- "ping"

}

func pong(msg <-chan string) {
	println(<-msg)
}

func TestChannel(t *testing.T) {

	// channel 是 Go 语言的并发通信机制，用于 在 goroutine 之间传递数据，可以避免使用共享内存时的竞争问题。
	recMsg := make(chan string)
	go ping(recMsg)
	pong(recMsg)
}
