package main

import (
	"fmt"
	"math/rand"
)

func initArray(N int) []int {
	var result []int = make([]int, N, N)

	for i := range result {
		result[i] = rand.Intn(N)
	}

	return result
}

func Move(array []int, N int) []int {
	var temp []int = make([]int, len(array))
	var moveIndex = N
	if N < 0 {
		moveIndex = len(array) + moveIndex
	}

	for idx, value := range array {
		temp[(idx+moveIndex)%len(array)] = value
	}

	return temp
}

func GetUniqueCounts(array []int) map[int]int {
	var result map[int]int = make(map[int]int, len(array))

	for _, value := range array {
		result[value] += 1
	}

	return result
}

func main() {
	var array []int = initArray(10)

	fmt.Println("Initial array:", array)

	var moveTo int = 5
	array = Move(array, moveTo)

	fmt.Println("Move array to", moveTo, ":", array)

	fmt.Println(GetUniqueCounts([]int{1, 2, 3, 4, 4}))
}
