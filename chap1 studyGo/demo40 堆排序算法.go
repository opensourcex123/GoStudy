package main

import "fmt"

func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 63, 234, 12}
	fmt.Println(HeapSort(arr))
}

func HeapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1
	for i := depth; i >= 0; i-- {
		topmax := i // 假设最大的位置在i
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] {
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}

func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}
