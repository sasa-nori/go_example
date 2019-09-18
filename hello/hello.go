package hello

import (
    "fmt"
    "runtime"

    "github.com/carlescere/scheduler"
)

func sayHello() {
    fmt.Println("hello, world")
}

func runSuccess() {
    // 5秒に1回 success!! と出力させる
    scheduler.Every(5).Seconds().Run(printSuccess)
    runtime.Goexit()
}

func printSuccess() {
    fmt.Printf("success!! \n")
}
