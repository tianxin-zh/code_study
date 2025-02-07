package main

import (
	"testing"
	"time"
)

func TestSwitch(T *testing.T) {
	// 1、switch 如果case都没有走到才会走default
	// 如果走到一个case下面一整块都会输出，并且后续的case不会走
	switch time.Now().Add(24 * time.Hour).Weekday() {
	case time.Saturday, time.Sunday:
		T.Log("weekend")
		T.Log("会一起输出")
	default:
		T.Log("weekday")
	}
}
