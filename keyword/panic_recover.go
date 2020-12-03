package keyword

import (
	"fmt"
	"time"
)

func GoroutineWithPanic() {
	defer println("main goroutine")
	go func() {
		defer println("in sub goroutine")
		panic("")
	}()
	time.Sleep(1 * time.Second)

	// 只会输出 in sub goroutine
	// main 协程直接退出
}

func RecoverNotInDefer() {
	defer fmt.Println("main goroutine")
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	panic("unknown err")
	// 只会输出 main goroutine
}

func PanicNested() {
	defer fmt.Println("in main")
	defer func() {
		defer func() {
			panic("panic again and again")
		}()
		panic("panic again")
	}()

	panic("panic once")
	// in main
	// --- FAIL: TestPanicNested (0.00s)
	// panic: panic once
	// 	panic: panic again
	// 	panic: panic again and again [recovered]
	// 	panic: panic again and again
}

func PanicRecoverWithNameReturn() (a int32) {
	a = 1
	defer func() {
		recover()
	}()
	panic("")
}

func PanicRecoverWithReturn() (int32, string, *string) {
	defer func() {
		recover()
	}()
	s := "123"
	if true {
		panic("")
	}
	return 1, "123", &s
}

func RecoverExperiment() {
	a := PanicRecoverWithNameReturn()
	// 返回 1
	fmt.Println("return PanicRecoverWithNameReturn() = ", a)
	b1, b2, b3 := PanicRecoverWithReturn()
	// 返回 0, "", nil
	fmt.Println("return PanicRecoverWithReturn() = ", b1, b2, b3)
}
