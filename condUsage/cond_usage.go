package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync"
	"sync/atomic"
	"time"
)

var status int64

func main() {
	//c := sync.NewCond(&sync.Mutex{})
	//for i := 0; i < 10; i++ {
	//	go listen(c)
	//}
	//time.Sleep(1 * time.Second)
	//go broadcast(c)
	//
	//ch := make(chan os.Signal, 1)
	//signal.Notify(ch, os.Interrupt)
	//<-ch

	//errGroup()

	//chansend()

	a := "information"
	b := "信息     "
	fmt.Println(len(a), len(b))
}

func errGroup() {
	var g errgroup.Group
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	errgroup.WithContext(ctx)

	for i := 0; i < 60; i++ {
		time.Sleep(1 * time.Second)
		g.Go(get)
	}
	if err := g.Wait(); err != nil {
		fmt.Println(err)
	}
}

func get() error {
	n := time.Now().Second()
	if n == 11 {
		return errors.New("not 11")
	}
	fmt.Println(n)
	return nil
}

func listen(c *sync.Cond) {
	c.L.Lock()
	for atomic.LoadInt64(&status) != 1 { // 未达到某种条件
		c.Wait() // 自动地对与该条件变量关联的互斥锁进行解锁，并阻塞该Goroutine，等待发送通知解除阻塞。
	}
	fmt.Println("listen")
	c.L.Unlock()
}

func chansend() {
	c := make(chan int, 3)
	for i := 0; i < 3; i++ {
		c <- i
	}
	close(c)
	for data := range c {
		fmt.Println("----")
		go printC(data)
	}
	fmt.Println("abc")
}

func printC(c int) {
	fmt.Println(c)
}

func broadcast(c *sync.Cond) {
	c.L.Lock()
	atomic.StoreInt64(&status, 1)
	c.Broadcast()
	c.L.Unlock()
}
