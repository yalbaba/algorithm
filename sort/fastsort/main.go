package main

import "fmt"

func main() {
	arr := []int{1, 3, 2, 8, 5, 9, 2, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, start, end int) {
	if start > end {
		return
	}
	// 定义基准数
	base := arr[start]
	i, j := start, end
	// 开始排序
	for i != j {
		for arr[j] >= base && i < j {
			j--
		}
		for arr[i] <= base && i < j {
			i++
		}
		// 当两边都停止排序，交互两边位置
		arr[i], arr[j] = arr[j], arr[i]
	}
	// 当两个指针相遇则本次排序结束，交换基准数和相遇位置的值
	arr[start] = arr[i]
	arr[i] = base
	// 然后交换相遇点的左边剩余数组
	quickSort(arr, start, i-1)
	// 交换相遇点右边剩余数组
	quickSort(arr, j+1, end)
}
