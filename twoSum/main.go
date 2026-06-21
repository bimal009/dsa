package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	result := twoSum(nums, 9)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)

	for i, v := range nums {
		if j, ok := table[target-v]; ok {
			return []int{j, i}
		}

		table[v] = i
	}

	return nil
}
