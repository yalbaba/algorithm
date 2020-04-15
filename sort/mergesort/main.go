package main

import "fmt"

func main() {
	arr := []int{2, 4, 10, 8, 1, 4, 3, 5, 7}
	fmt.Println(merge_sort(arr, 0, len(arr)-1))
}

func merge_sort(arr []int, start, end int) []int {
	mid := (start + end) / 2
	arr1 := merge_sort(arr, start, mid)
	arr2 := merge_sort(arr, mid+1, end)
	return merge(arr1, arr2)
}

func merge(arr1, arr2 []int) []int {
	var temp []int
	var (
		i = 0
		j = 0
	)
	for {
		if arr1[i] < arr2[j] {
			temp = append(temp, arr1[i])
			i++
		} else {
			temp = append(temp, arr1[j])
			j++
		}
		if i >= len(arr1) || j >= len(arr2) {
			break
		}
	}
	//放左边的剩余元素
	for {
		temp = append(temp, arr1[i])
		i++
		if i >= len(arr1) {
			break
		}
	}
	//放右边的剩余元素
	for {
		temp = append(temp, arr1[j])
		j++
		if j >= len(arr2) {
			break
		}
	}
	return temp
}
