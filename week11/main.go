package main

import (
	"fmt"
	"week/greeting"
)

func main() {
	greeting.Hi("Inha")
	greeting.Hello("Harvard")
}

func Hi() {
	fmt.Printf("Hi!\n")
}

func Hello() {
	fmt.Printf("H!\n")
}
