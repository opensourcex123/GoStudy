package main

import "fmt"

func main() {
	arrList := []int{1, 2, 8, 11, 3, 6, 8, 4, 9, 343, 3}
	arrList = standardInsertSort(arrList)
	fmt.Println(arrList)
}

func standardInsertSort(list []int) []int {
	resultList := []int{}
	length := len(list)
	i := 1
	resultList = append(resultList, list[0])
	for i < length {
		for j := 0; j < len(resultList); j++ {
			if list[i] <= resultList[j] {
				resultList = insertList(resultList, j, list[i])
				break
			}
			if j == len(resultList)-1 && list[i] > resultList[j] {
				resultList = insertList(resultList, j+1, list[i])
				break
			}
		}
		i++
	}
	return resultList
}

func insertList(list []int, i int, x int) []int {
	resultList := []int{}
	n := 0
	if i == len(list) {
		resultList = append(list, x)
		return resultList
	}

	for n < len(list) {
		if n < i {
			resultList = append(resultList, list[n])
		} else if n == i {
			resultList = append(resultList, x)
			resultList = append(resultList, list[n])
		} else {
			resultList = append(resultList, list[n])
		}
		n++
	}
	return resultList
}
