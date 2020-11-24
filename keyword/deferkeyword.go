package keyword

import (
	"fmt"
	"time"
)

func deferInFor() {
	// defer 执行顺序为栈的顺序
	// 不建议在 for 中使用，因为性能较弱
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// 输出 43210
}

func deferSeq() {
	{
		defer fmt.Println("defer runs")
		fmt.Println("block ends")
	}

	fmt.Println("main ends")
	// 输出
	// block ends
	// main ends
	// defer runs
}

func deferParamPrecompute() {
	startedAt := time.Now()
	defer fmt.Println(time.Since(startedAt))
	// 输出远远小于1秒，说明defer 函数的参数已经被拷贝了
	time.Sleep(time.Second)
}

func deferParamPrecompute2() {
	startedAt := time.Now()
	defer func() { fmt.Println(time.Since(startedAt)) }()
	// 输出1秒左右
	time.Sleep(time.Second)
}


func DeferExperiment() {
	deferInFor()
	deferSeq()
	deferParamPrecompute()
	deferParamPrecompute2()
}
