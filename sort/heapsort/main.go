package main

import "fmt"

func main() {
	arr := []int{4, 6, 8, 5, 9}
	var temp int
	//1、将无序切片构建成第一个大顶堆
	for i := len(arr)/2 - 1; i >= 0; i-- {
		buildHeap(arr, i, len(arr))
	}
	// fmt.Println(arr)
	//2、将堆顶元素（最大元素）沉到切片最末端,最末端元素放最前面，从根节点开始循环构建堆
	for j := len(arr) - 1; j > 0; j-- {
		// j表示最末端
		temp = arr[j]
		arr[j] = arr[0]
		arr[0] = temp
		fmt.Printf("arr::::%v,j:::::::%d\n", arr, j)
		buildHeap(arr, 0, j)
	}

	// fmt.Println(arr)
}

func buildHeap(arr []int, i, length int) {
	temp := arr[i] //先保存当前需要局部调整结构的子堆的根节点

	// for循环比较一个字堆中的三个节点的大小
	for k := 2*i + 1; k < length; k = 2*k + 1 {
		if arr[k] < arr[k+1] && k+1 < length {
			// 表示左子节点小于右子节点，此时要将k指向右子节点
			k++
		}
		if arr[k] > temp {
			// 表示该子堆的根节点小于右子节点，右子节点与根节点交换位置
			arr[i] = arr[k]
			i = k
		} else {
			break
		}
	}
	// 把根节点移动到最终于根节点交换位置的子节点上
	arr[i] = temp
}
