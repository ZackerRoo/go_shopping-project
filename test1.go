// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // 全局变量，用于保存正在处理的任务
// var (
// 	currentTask int
// 	taskMutex   sync.Mutex
// )

// func producer(tasks chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	// 生产 10 个任务
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("producer producing task %d\n", i)
// 		tasks <- i
// 		time.Sleep(time.Second)
// 	}

// 	// 关闭任务通道
// 	close(tasks)
// }

// func consumer(id int, tasks <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for task := range tasks {
// 		fmt.Printf("consumer %d processing task %d\n", id, task)
// 		// 模拟处理任务的耗时
// 		time.Sleep(time.Second)

// 		// 交出任务，使用 hand off 机制
// 		taskMutex.Lock()
// 		currentTask = task
// 		taskMutex.Unlock()
// 		// done <- true // 这个是这句话为什么去掉的原因似乎每个消费者在完成一个任务后尝试向done通道发送信号，但如果没有其他协程在接收这个信号，那么这些消费者协程就会在发送信号后阻塞（因为在无缓冲的通道上发送操作是阻塞的，直到另一端有协程进行接收）。如果所有的消费者都处于等待状态，并且没有新的协程来接收done通道中的信号，那么程序就会死锁，因为所有的协程都在等待一些操作，而这些操作永远不会发生。
// 	}

// 	fmt.Printf("consumer %d has processed all tasks\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// 任务通道
// 	tasks := make(chan int)

// 	// done 通道，用于实现 hand off 机制
// 	// done := make(chan bool)

// 	// 启动 3 个 consumer goroutine
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go consumer(i, tasks, &wg)
// 		// wg.Done()
// 	}

// 	// 启动 producer goroutine
// 	wg.Add(1)
// 	go producer(tasks, &wg)

// 	// 等待所有 goroutine 执行完毕
// 	wg.Wait()

//		// 所有任务处理完毕后，输出最后一个交出任务的 consumer ID 和任务 ID
//		fmt.Printf("last consumer to hand off task: %d, task ID: %d\n", currentTask%3+1, currentTask)
//	}
package main

// func main() {
// 	go func() {
// 		for {
// 			fmt.Println("Goroutine 1 is running")
// 			runtime.Gosched()
// 		}
// 	}()

// 	go func() {
// 		for {
// 			fmt.Println("Goroutine 2 is running")
// 			runtime.Gosched()
// 		}
// 	}()

// 	for {
// 		fmt.Println("Main Goroutine is running")
// 		time.Sleep(time.Second)
// 	}
// }
