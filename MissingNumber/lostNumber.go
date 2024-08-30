package MissingNumber

import "fmt"

func FindMissingNumber(arr []int, n int) {

	//Method 1
	value := getMissingNumber(arr, n)
	fmt.Printf("Missing value from method 1 : %v \n", value)

	//Method 2 (applicable for only one missing number)
	// missValue := getNumberMethod2(arr, n)
	// fmt.Printf("Missing value from method 2 : %v \n", missValue)

}

func getMissingNumber(arr []int, n int) []int {

	var newArr []int
	for i := 1; i <= n; i++ {
		newArr = append(newArr, i)
	}

	var missArr []int
	var flag bool
	for _, oriArr := range newArr {
		flag = false
		for _, eachVal := range arr {
			if eachVal == oriArr {
				flag = true
			}
		}
		if !flag {
			missArr = append(missArr, oriArr)
		}
	}
	return missArr

}

func getNumberMethod2(arr []int, n int) int {

	totalValue := n * (n + 1) / 2

	sum := 0

	for _, eachVal := range arr {
		sum += eachVal
	}

	return totalValue - sum

}
