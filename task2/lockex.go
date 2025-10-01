package task2

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/**
1.题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

type Counter struct {
	mu    sync.Mutex
	count int
}

func Lock1() {
	var c Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				c.mu.Lock()
				c.count++
				c.mu.Unlock()
			}
		}()
	}
	fmt.Println("模拟Lock1 start....")
	wg.Wait()
	fmt.Println(c.count)
	fmt.Println("模拟Lock1 end....")
}

/**
2.题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/

type Counter2 struct {
	count int64
}

func Lock2() {
	var c Counter2
	var wg sync.WaitGroup
	wg.Add(10)
	fmt.Println("模拟Lock2 start....")
	l2(&c, &wg)
	wg.Wait()
	fmt.Println(getCount(&c))
	fmt.Println("模拟Lock2 end....")

}

func l2(c *Counter2, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&c.count, 1)
			}
		}()
	}
}

func getCount(c *Counter2) int64 {
	return atomic.LoadInt64(&c.count)
}
