package main

import "fmt"

func main() {
	// 初始化value为int64的map
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// 初始化value为float64的map
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		sumInts(ints),
		sumFloats(floats))
	fmt.Printf("Generic Sums: %v and %v\n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))
}
