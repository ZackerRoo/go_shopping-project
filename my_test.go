package main

import (
	"fmt"
	"sort"
	"sync"
	"testing"
)

// type Counter struct {
// 	count int
// }

// // 值接收者方法
// func (c Counter) increment() {
// 	c.count++
// }

// // 指针接收者方法
// func (c *Counter) decrement() {
// 	c.count--
// }

// func main() {
// 	// 值接收者方法不会改变原始接收者的值
// 	// c1 := Counter{count: 0}
// 	// c1.increment()
// 	// fmt.Println(c1.count) // 输出 0

// 	// // 指针接收者方法会改变原始接收者的值
// 	// c2 := Counter{count: 0}
// 	// c2.decrement()
// 	// fmt.Println(c2.count) // 输出 -1
// 	TestSliceConcurrencySafe(nil)
// }

/**
* 切片非并发安全
* 多次执行，每次得到的结果都不一样
* 可以考虑使用 channel 本身的特性 (阻塞) 来实现安全的并发读写
 */
func TestSliceConcurrencySafe(t *testing.T) {
	a := make([]int, 0)
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(i int) {
			a = append(a, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	t.Log(len(a))
	fmt.Printf("len(a): %v\n", len(a))
	// not equal 10000
}

func TestMapRange(t *testing.T) {
	m := map[int]string{1: "a", 2: "b", 3: "c"}
	t.Log("first range:")
	for i, v := range m {
		t.Logf("m[%v]=%v ", i, v)
	}
	t.Log("second range:")
	for i, v := range m {
		t.Logf("m[%v]=%v ", i, v)
	}

	// 实现有序遍历
	var sl []int
	// 把 key 单独取出放到切片
	for k := range m {
		sl = append(sl, k)
	}
	// 排序切片
	sort.Ints(sl)
	// 以切片中的 key 顺序遍历 map 就是有序的了
	for _, k := range sl {
		t.Log(k, m[k])
	}
}

func TestConcurrentMap(t *testing.T) {
	s := make(map[int]int)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s[i] = i
		}(i)
	}
	wg.Wait()

	for i := 0; i < 100; i++ {
		if s[i] != i {
			t.Errorf("Expected map value at index %d to be %d, but got %d", i, i, s[i])
		}
	}
}
