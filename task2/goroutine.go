package task2

import (
	"fmt"
	"sync"
	"time"
)

/**
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
*/

func PrintOddEven() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start")
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("end")
}

/**
2.题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
*/

func TaskScheduler() {
	//定义任务
	tasks := []func(){
		func() {
			time.Sleep(1000 * time.Millisecond)
			fmt.Println("任务1执行")
		},
		func() {
			time.Sleep(2000 * time.Millisecond)
			fmt.Println("任务2执行")
		},
		func() {
			time.Sleep(1500 * time.Millisecond)
			fmt.Println("任务3执行")
		},
	}

	var wg sync.WaitGroup
	wg.Add(len(tasks))
	for _, task := range tasks {
		go func(task func()) {
			defer wg.Done()
			start := time.Now()
			task()
			duration := time.Since(start)
			fmt.Printf("任务执行完成，耗时：%v\n", duration)
		}(task)
	}
	wg.Wait()
	fmt.Println("所有任务执行完毕")
}
