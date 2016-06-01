package main

import "fmt"

func main() {
	fmt.Println(fibonacci(10))
}

func fibonacci(n int) []int {
	series := make([]int, n)
	series[0], series[1] = 0, 1

	for i := 2; i < n; i++ {
		series[i] = series[i-1] + series[i-2]
	}
	return series
}
