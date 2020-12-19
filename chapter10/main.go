package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println("routine:", n, "print", i)
		time.Sleep(time.Millisecond * 250)
	}
}

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	//for i:=0; i<10; i++ {
	//	go f(i)
	//}
	//var c chan string = make(chan string)
	//go pinger(c)
	//go ponger(c)
	//go printer(c)
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				//default:
				//	fmt.Println("nothing is ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
