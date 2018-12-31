package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	fmt.Println(cap(ch))
	ch <- 10
	ch <- 20
	fmt.Println(len(ch))

	for v := range ch {
		fmt.Println(v)
	}
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
