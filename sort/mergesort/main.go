package main

import "fmt"

var assit []int

func main() {
	arr := []int{2, 4, 10, 8, 1, 4, 3, 5, 7}
	assit = make([]int, len(arr))
	merge_sort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func merge_sort(arr []int, left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	merge_sort(arr, left, mid)
	merge_sort(arr, mid+1, right)
	merge(arr, left, mid, right)
}

func merge(arr []int, left, mid, right int) {
	i := left
	p1 := left
	p2 := mid + 1

	for p1 <= mid && p2 <= right {
		// 两个归并的数组都没有越界，则开始合并
		if arr[p1] < arr[p2] {
			assit[i] = arr[p1]
			p1++
			i++
		} else {
			assit[i] = arr[p2]
			p2++
			i++
		}
	}
	// 如果是左子组的元素先放完
	for p2 <= right {
		assit[i] = arr[p2]
		i++
		p2++
	}
	// 如果是右子组先放完
	for p1 <= mid {
		assit[i] = arr[p1]
		i++
		p1++
	}
	// 将assit复制到原数组中
	for index := left; index <= right; index++ {
		arr[index] = assit[index]
	}
}
