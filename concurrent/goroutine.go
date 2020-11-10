package concurrent

import (
	"fmt"
	"time"
)

func fn1(p string) {
	fmt.Printf("fn1 p = %s\n", p)
}

func GoroutineExperiment() {
	go fn1("first")
	fmt.Println("main")
	go fn1("second")
	time.Sleep(200 * time.Millisecond) // 如果不 sleep fn1 没有得到运行
}
