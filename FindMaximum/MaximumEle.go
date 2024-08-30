package FindMaximum

import "fmt"

func FindMaximumElement(arr []int) {

	max := getMaximum(arr)
	fmt.Printf("Max element from array : %d \n", max)

}

func getMaximum(arr []int) int {

	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		return arr[0]
	}

	max := arr[0]
	for i := 1; i < len(arr); i++ {

		if arr[i] > max {
			max = arr[i]
		}

	}
	return max

}
