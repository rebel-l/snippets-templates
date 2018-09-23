package main

import "fmt"

func main() {
	var num uint32
	for num = 0; num < 10; num++ {
		fmt.Printf("Left Shift: %d\n", 1<<num)
	}

	for num = 0; num < 10; num++ {
		fmt.Printf("Right Shift: %d\n", 512>>num)
	}
}
