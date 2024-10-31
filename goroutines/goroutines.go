package main

import (
	"fmt"
	"time"
)

func routine1() {
	for i := 0; i < 5; i++ {
		fmt.Println("Routine 1 - Counting", i)
		time.Sleep(1 * time.Second)
	}
}

func routine2() {
	for i := 0; i < 5; i++ {
		fmt.Println("Routine 2 - Counting", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go routine1()
	go routine2()

	time.Sleep(6 * time.Second)
	fmt.Println("All routines finished.")
}