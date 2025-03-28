package main

import (
	"fmt"
	"sync"
	"time"
)

// 使用RWMutex保护的数据结构
type SafeCounter struct {
	mu     sync.RWMutex
	values map[string]int
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{
		values: make(map[string]int),
	}
}

// 写操作 - 使用写锁
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()         // 获取写锁
	defer c.mu.Unlock() // 确保释放锁

	// 修改共享数据
	c.values[key]++
}

// 读操作 - 使用读锁
func (c *SafeCounter) Value(key string) int {
	c.mu.RLock()         // 获取读锁
	defer c.mu.RUnlock() // 确保释放锁

	// 读取共享数据
	return c.values[key]
}

func main() {
	counter := NewSafeCounter()

	// 启动多个goroutine进行读写操作
	var wg sync.WaitGroup

	// 启动5个写goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				key := fmt.Sprintf("key%d", id)
				counter.Inc(key)
				time.Sleep(time.Millisecond * 10)
			}
		}(i)
	}

	// 启动3个读goroutine
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 15; j++ {
				key := fmt.Sprintf("key%d", id%5)
				value := counter.Value(key)
				fmt.Printf("Reader %d: %s = %d\n", id, key, value)
				time.Sleep(time.Millisecond * 20)
			}
		}(i)
	}

	wg.Wait()

	// 最终结果
	fmt.Println("\nFinal Counter Values:")
	for i := 0; i < 5; i++ {
		key := fmt.Sprintf("key%d", i)
		fmt.Printf("%s: %d\n", key, counter.Value(key))
	}
}
