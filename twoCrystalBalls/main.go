package main

import (
	"fmt"
	"math"
)

func main() {
	floors := []bool{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		true,
		true,
		true,
		true,
	}

	fmt.Println(findBreakingPoint(floors))
}

func findBreakingPoint(floors []bool) int {
	jumpAmount := int(math.Sqrt(float64(len(floors))))

	i := jumpAmount

	for ; i < len(floors); i += jumpAmount {
		if floors[i] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < jumpAmount && i+j < len(floors); j++ {
		if floors[i+j] {
			return i + j
		}
	}

	return -1
}
