package main

import (
	"sync"
	"testing"
)

type ScoreMap struct {
	// 在 Go 语言中，sync.Mutex 是一种互斥锁，用于保护共享资源，防止多个 goroutine 同时访问导致数据竞争（race condition）。
	mu    sync.Mutex
	Score map[string]int
}

func (sm *ScoreMap) IncScore(name string, score int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.Score[name] += score
}

func TestMutexes(t *testing.T) {
	sm := ScoreMap{
		mu:    sync.Mutex{},
		Score: make(map[string]int),
	}

	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			sm.IncScore("A", 1)
			wg.Done()
		}()
	}

	wg.Wait()
	println(sm.Score["A"])
}
