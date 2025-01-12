package main

import "fmt"

type Number interface {
	int64 | float64
}

func main() {
	// initialize int64 map
	ints := map[string]int64{
		"first":  34,
		"second": 26,
	}

	// initialize float64 map
	floats := map[string]float64{
		"first":  10.99,
		"second": 11.01,
	}

	// print out the sum of the maps using non-generic methods
	fmt.Printf("Non Generic Sums %v and %v\n", SumInts(ints), SumFloats(floats))

	// print out the sum of the maps using generic methods
	fmt.Printf("Generic Sums %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))

	// print out the sum of the maps using number type method
	fmt.Printf("Type Method Sums %v and %v\n", SumNumbers(ints), SumNumbers(floats))
}

// function which sums together the values in m (int64)
func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// function which sums together the values in m (float64)
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// function which sums together the generic values in m (int64 or float64)
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// function which sums together the number type values in m
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
