package main

import "fmt"

func main() {
	fmt.Println(paosort([]int{1, 3, 7, 4, 2, 9, 8}))
}

func paosort(arr []int) []int {
	maxi := len(arr) - 1
	for {
		if maxi == 1 {
			break
		}
		for i := 0; i < maxi; i++ {
			temp := arr[i]
			if arr[i] > arr[i+1] {
				arr[i] = arr[i+1]
				arr[i+1] = temp
			}
		}
		maxi--
		fmt.Println(maxi)

	}

	return arr
}
