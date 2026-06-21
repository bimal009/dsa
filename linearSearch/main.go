package main

import "fmt"

func linearSearch(arr []int, value int) bool {
	for _, n := range arr {
		if n == value {
			return true
		}
	}
	return false
}

func main() {
	arr := make([]int, 8)
	arr[1] = 10
	fmt.Println(linearSearch(arr, 10))
}
