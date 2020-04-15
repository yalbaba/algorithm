package main

import "fmt"

func main() {
	arr := []int{3, 5, 1, 7, 2, 4, 9}
	for i := 0; i < len(arr)-1; i++ {
		minpos := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minpos] {
				minpos = j
			}
		}
		temp := arr[i]
		arr[i] = arr[minpos]
		arr[minpos] = temp
	}

	fmt.Println(arr)
}
