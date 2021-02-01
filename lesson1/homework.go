package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func squareGenerator() func() int {
	var counter int = 0

	return func() int {
		counter += 1
		return counter * counter
	}
}

func Solution(number int) int {
	var bin string = strconv.FormatInt(int64(number), 2)

	var zeroRegexp = regexp.MustCompile(`0+`)
	var arrayOfZeros = zeroRegexp.FindAllString(bin, -1)

	var result float64 = 0

	for _, str := range arrayOfZeros {
		result = math.Max(float64(len(str)), float64(result))
	}

	return int(result)
}

func main() {
	var generator = squareGenerator()

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}

	fmt.Println(Solution(1024))
}
