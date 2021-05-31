package main

import (
	"fmt"
	"time"
)

func main() {
	for i:=0; i<10; i++ {
	        go func(n int) {
	                for {
				time.Sleep(time.Second)
	                        fmt.Printf("Hi, I'm Goroutine-%d\n", n)
	                }
	        }(i)
	}
	time.Sleep(time.Second*10)
}

