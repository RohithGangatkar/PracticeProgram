package SecondLargest

import (
	"fmt"
	"math"
)

func SecondMaximumElement(arr []int) {

	max := getSecondMaximum(arr)
	fmt.Printf("Second max Max element from array : %d \n", max)

}

func getSecondMaximum(arr []int) int {

	if len(arr) < 2 {
		return -1
	}

	//95, 10, 40, 60, 50, 80, 20, 15, 30, 90
	largest, secondLargest := math.MinInt, math.MinInt
	for _, eachNum := range arr {
		if eachNum > largest {
			largest = eachNum
		} else if eachNum > secondLargest && largest != eachNum {
			secondLargest = eachNum
		}

	}

	if secondLargest == math.MinInt {
		fmt.Println("No second largest number found")
		return math.MinInt
	}
	return secondLargest

}
