package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func square(x int) int {
	return x * x
}

func main() {
	fmt.Print(square(sum(25, 9)))
}
