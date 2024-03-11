package main

import (
	"fmt"
)

// func main() {
// 	// ch是长度为4的带缓冲的channel
// 	// 初始hchan结构体重的buf为空，sendx和recvx均为0
// 	ch := make(chan string, 4)
// 	fmt.Println(ch, unsafe.Sizeof(ch))
// 	go sendTask(ch)
// 	go receiveTask(ch)
// 	time.Sleep(1 * time.Second)
// }

// G1是发送者
// 当G1向ch里发送数据时，首先会对buf加锁，然后将task存储的数据copy到buf中，然后sendx++，然后释放对buf的锁
func sendTask(ch chan string) {
	taskList := []string{"this", "is", "a", "demo"}
	for _, task := range taskList {
		ch <- task //发送任务到channel
	}
}

// G2是接收者
// 当G2消费ch的时候，会首先对buf加锁，然后将buf中的数据copy到task变量对应的内存里，然后recvx++,并释放锁
func receiveTask(ch chan string) {
	for {
		task := <-ch                  //接收任务
		fmt.Println("received", task) //处理任务
	}
}
