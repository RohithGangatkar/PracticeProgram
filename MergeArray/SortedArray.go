package MergeArray

import "sort"

func MergeSortedArray(arr1 []int, m int, arr2 []int, n int) {

	var mergeArr []int

	mergeArr = append(mergeArr, arr1...)

	mergeArr = append(mergeArr, arr2...)

	sort.Ints(mergeArr)

}
