package main

import (
	max "PracticeProgram/FindMaximum"
	inter "PracticeProgram/Intersection"
	"PracticeProgram/MissingNumber"
	"PracticeProgram/RemoveDuplicates"
	second "PracticeProgram/SecondLargest"
)

func main() {

	//maximum element from array
	values := []int{95, 10, 95, 40, 60, 50, 80, 20, 90, 20, 15, 30, 90, 90, 80}
	max.FindMaximumElement(values)

	//Second Maximum
	second.SecondMaximumElement(values)

	//Missing number
	arr := []int{1, 2, 3, 6, 5, 8}
	MissingNumber.FindMissingNumber(arr, 8)

	//Intersectin of array
	arr1 := []int{6, 2, 3, 8, 4, 10}
	arr2 := []int{1, 2, 7, 8, 5, 9, 3, 4}
	inter.ArrayIntersection(arr1, arr2)

	//Remove dupliacte from array
	dup := []int{2, 4, 6, 2, 4, 8, 1, 3}
	RemoveDuplicates.RemoveDuplicatesFromArray(dup)
}
