package main

import "fmt"

func SelectMax(arr []int) int {
	length := len(arr)
	if length == 1 {
		return arr[0]
	}

	max := arr[0]
	for i := 1; i < length; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func SelectSort(arr []int) []int {
	length := len(arr)
	if length == 1 {
		return arr
	}

	for i := 1; i < len(arr); i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if i != min {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	max := SelectMax(arr)
	selectsort := SelectSort(arr)
	fmt.Println(max)
	fmt.Println(selectsort)
}
