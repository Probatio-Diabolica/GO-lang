package main

import (
	"fmt"
)

// returning two values
func devide(a int, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func exp1() {
	result, err := devide(800, 8)
	fmt.Println(result, err)

	_, ero := devide(800, 0)
	fmt.Println(ero)
}

// --------------------------------------------------------------------------------------------------------------------------
func minMax(nums []int) (int, int) {
	min, max := nums[0], nums[0]

	for _, val := range nums {
		if min > val {
			min = val
		}
		if max < val {
			max = val
		}
	}
	return min, max
}

// sugar syntax: it just means if, two params have same type, no need to specify them again and again

func sugar(x, y int) int {
	return x + y
}

func sugar2(x, y int64, a, b string) (int64, string) {
	return x + y, a + " " + b
}

func exp2() {
	nums := []int{3, 7, 8, 9, 5}
	low, high := minMax(nums)
	fmt.Println("LOW => ", low, "HIGH =>", high)
}

func exp3() {
	fmt.Println(sugar(90, 90))
	fmt.Println(sugar2(89, 89, "hi", "hello"))
}

// --------------------------------------------------------------------------------------------------------------------------
// valid but doesn't works
// func f(ab func(int, int) int, b int) int {
// 	return ab(90, 90) + b
// }

// ---------------------------------------------------------------------------------------------------------------------------

// pass by value nature
func increment(x int) {
	x++
}

// pass by pointer
func inc(a *int) {
	*(a)++
}

func exp4() {
	x := 89
	increment(x) //no increment because its just a copy
	fmt.Println(x)
	inc(&x) //actually increases it
	fmt.Println(x)

}

// ---------------------------------------------------------------------------------------------------------------------------
func main() {
	exp1()
	exp2()
	exp3()
	// f(sugar(89, 56), 999)
	exp4()
}
