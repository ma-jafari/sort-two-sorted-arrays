package main

import "fmt"

func main() {
	arrayA := []int{2, 6, 8, 10, 11}
	arrayB := []int{1, 4, 6, 12, 14, 16}
	resultArray := make([]int, len(arrayA)+len(arrayB))

	i, j, k := 0, 0, 0
	for i < len(arrayA) && j < len(arrayA) {
		if arrayA[i] < arrayB[j] {
			resultArray[k] = arrayA[i]
			i++
		} else {
			resultArray[k] = arrayB[j]
			j++
		}
		k++
	}

	for ; i < len(arrayA); i++ {
		resultArray[k] = arrayA[i]
		k++
	}

	for ; j < len(arrayB); j++ {
		resultArray[k] = arrayB[j]
		k++
	}
	fmt.Println(resultArray)
}

