package main

import (
	"fmt"
	"runtime"

	"github.com/carlescere/scheduler"
)

func hello() {
	fmt.Println("hello, world")
}

func runSuccess() {
	// 1秒に1回 success!! と出力させる
	scheduler.Every(1).Seconds().Run(printSuccess)
	runtime.Goexit()
}

func printSuccess() {
	fmt.Printf("success!! \n")
}
