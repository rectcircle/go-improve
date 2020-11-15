package concurrent

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var i int32 = 0
var iOnce sync.Once

func InitI() {
	i = 1
	fmt.Println("init I = 1")
}

func SyncOnce() {
	iOnce.Do(InitI)
	fmt.Println(i)
	iOnce.Do(InitI)
	iOnce.Do(InitI)
	iOnce.Do(InitI)
}

func SyncWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(10 * time.Millisecond)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("wait finish")
}

func SyncPool() {
	var bytePool = sync.Pool{
		New: func() interface{} {
			fmt.Println("New")
			return make([]byte, 1024)
		},
	}

	bytePool.Get()
	ba := bytePool.Get()
	bytePool.Put(ba)
	bytePool.Get()
	// 输出
	// New
	// New
}

func Atomic() {
	var a int32 = 1
	atomic.AddInt32(&a, 1)
	fmt.Println(a)
}


func ContextContext() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println(ctx.Deadline())
	cancel()
	fmt.Println(ctx.Deadline())

	ctx1, cancel1 := context.WithTimeout(context.Background(), 50 * time.Millisecond)
	fmt.Println(ctx1.Deadline())
	// cancel1()
	fmt.Println(cancel1)
	fmt.Println(ctx1.Deadline())
	time.Sleep(100 * time.Millisecond)
	fmt.Println(ctx1.Deadline())

	ctx2, cancel2 := context.WithCancel(ctx1);
	fmt.Println(ctx2.Deadline())
	fmt.Println(cancel2)
	fmt.Println(ctx2.Err())
	fmt.Println(ctx.Err())
}