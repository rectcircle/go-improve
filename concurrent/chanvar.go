package concurrent

import (
	"fmt"
	"strconv"
	"time"
)

type Message struct {
	id    uint64
	value string
}

func consumer(name string, done <-chan bool, messageChannel <-chan Message, commitChannel chan<- uint64) {
	fmt.Printf("consumer %s start\n", name)
SELECT:
	for {
		select {
		case <-done:
			break SELECT
		case message := <-messageChannel:
			fmt.Printf("consumer %s handle %s\n", name, message.value)
			commitChannel <- message.id
		}
	}
	fmt.Printf("consumer %s exit\n", name)
}

func selectMultiChannel(a <-chan string, b <-chan string) {
	for i := 0; i < 10; i++ {
		select {
		case ad := <-a:
			fmt.Println(i, ad)
		case bd := <-b:
			fmt.Println(i, bd)
		}
	}
}

func channelTimeout(a <-chan string) {
	select {
	case ad := <-a:
		fmt.Println(ad)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

func selectWriteChannel(a chan<- string) {
	select {
	case a <- "从 select中写入 a":
		fmt.Println("写入 a 成功")
	default:
		fmt.Println("Default 写入 a 失败")
	}
}

func Chanvar() {
	done := make(chan bool)
	messageChannel := make(chan Message)
	commitChannel := make(chan uint64)
	// 创建消费者
	for i := 0; i < 10; i++ {
		go consumer(strconv.FormatInt(int64(i), 10), done, messageChannel, commitChannel)
	}
	// 发送消息给消费者
	for i := 0; i < 20; i++ {
		id := uint64(i)
		messageChannel <- Message{id, strconv.FormatUint(uint64(i), 10)}
		time.Sleep(10 * time.Millisecond)
		committedID := <-commitChannel
		fmt.Printf("consumer committed message(id=%d)\n", committedID)
	}
	// 销毁消费者
	for i := 0; i < 10; i++ {
		done <- true
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("测试 close")
	closedCh := make(chan int32)
	go func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("close(closedCh)")
		close(closedCh)
	}()
	data, ok := <-closedCh
	fmt.Println(data, ok)
	// closedCh <- 1 // panic: send on closed channel [recovered]
	data, ok = <-closedCh
	fmt.Println("receive closedCh")
	fmt.Println(data, ok)

	fmt.Println("测试 select 多channel")
	a := make(chan string, 1)
	b := make(chan string, 1)
	b <- "b"
	a <- "a"
	go selectMultiChannel(a, b)
	// 输出
	// 0 a
	// 1 b
	time.Sleep(1000 * time.Millisecond)

	fmt.Println("select 超时")
	a2 := make(chan string, 1)
	go channelTimeout(a2)
	time.Sleep(600 * time.Millisecond)
	// 输出 Timeout
	a2 <- "message a"
	time.Sleep(100 * time.Millisecond)

	fmt.Println("select 写入操作")
	a3 := make(chan string, 2)
	go selectWriteChannel(a3)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(cap(a3), len(a3))
	go selectWriteChannel(a3)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(cap(a3), len(a3))
	go selectWriteChannel(a3)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(cap(a3), len(a3))
	// 输出
	// 写入 a 成功
	// 2 1
	// 写入 a 成功
	// 2 2
	// Default 写入 a 失败
	// 2 2
}
