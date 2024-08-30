package Intersection

import "fmt"

func ArrayIntersection(arr []int, otherArr []int) {

	var commonEle []int
	for _, eachArr := range arr {
		for _, eachOtherArr := range otherArr {
			if eachArr == eachOtherArr {
				commonEle = append(commonEle, eachArr)
				break
			}
		}

	}

	fmt.Printf("Intersection array:  %v \n", commonEle)
}
