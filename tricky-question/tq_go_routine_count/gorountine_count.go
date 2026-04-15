package main

import (
	"fmt"
	"runtime"
	"time"
)

func square(c chan int){
	for i:=0;i<4;i++ {
		num := <-c
		fmt.Println(num*num)
	}
}


func main(){
	c:= make(chan int, 3)
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
	go square(c)
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
	c<-1
	c<-2
	c<-3
	c<-4
	fmt.Println("runtime.NumGoroutine()", runtime.NumGoroutine())
	time.Sleep(30*time.Millisecond)
}