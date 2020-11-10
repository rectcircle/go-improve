package concurrent

import (
	"fmt"
	"sync"
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
}
