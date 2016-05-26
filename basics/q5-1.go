package main

import (
	"fmt"
)

func main() {
	float_list := []float64{2.5, 3.5, 4.5, 5.5}

	sum := 0.0
	avg := 0.0

	switch len(float_list) > 0 {
	case true:
		for _, v := range float_list {
			sum += v
		}
		avg = sum / float64(len(float_list))
	}
	fmt.Println(avg)
}
