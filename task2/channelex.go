package task2

import (
	"fmt"
	"sync"
)

/*
*
1.题目 ：编写一个程序，使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/
func Channel1() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
			fmt.Printf("发送了%d\n", i)
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for data := range ch {
			fmt.Printf("接收到%d\n", data)
		}
	}()

	fmt.Println("模拟channel1 start....")
	wg.Wait()
	fmt.Println("模拟channel1 end....")
}

/**
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
*/

func Channel2() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)

	go producer(ch, &wg)
	go consumer(ch, &wg)
	fmt.Println("模拟channel2 start....")
	wg.Wait()
	fmt.Println("模拟channel2 end....")
}
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Printf("发送了%d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("接收到%d\n", data)
	}
}
