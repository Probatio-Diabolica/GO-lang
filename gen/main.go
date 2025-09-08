package main

import "fmt"

func main(){
	ints := map[string]int64{
		"f" : 10,
		"s" : 200,
	}	

	floats := map[string]float64{
		"f" : 10.01,
		"s" : 90.678,
	}
	fmt.Printf("sum of floats is %v while som of ints is %v", SumFloats(floats),SumInts(ints))
}

// SumInts returns sum of all keys over m
func SumInts(m map[string]int64) int64{
	var sum int64
	for _,v := range m{
		sum+=v
	}
	return sum
}

func SumFloats(m map[string]float64) float64{
	var sum float64
	for _,v := range m{
		sum+=v
	}
	return sum
}

