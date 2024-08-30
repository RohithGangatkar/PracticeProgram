package RemoveDuplicates

import "fmt"

func RemoveDuplicatesFromArray(arr []int) {

	arrMap := make(map[int]bool)
	var newArr []int
	for _, eachArr := range arr {
		if _, ok := arrMap[eachArr]; !ok {
			arrMap[eachArr] = true
			newArr = append(newArr, eachArr)
		}
	}

	fmt.Printf("Removed duplicates from array :  %v \n", newArr)

}
