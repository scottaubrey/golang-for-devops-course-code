package main

import (
	"fmt"
)

func main() {
	printString("one")

	c := make(chan bool)

	go func() { printString("two"); c <- true }()
	printString("three")
	<-c

	f := func() { printString("four"); c <- true }
	go f()
	printString("five")
	<-c
}

func printString(str string) {
	fmt.Println(str)
}
