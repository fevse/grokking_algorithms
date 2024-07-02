package chapter4

import "fmt"

func Maxn(arr []int) int {
	fmt.Println(arr)
	if len(arr) == 1 {
		return arr[0]
	} else {
		if arr[0] > arr[len(arr)-1] {
			return Maxn(arr[:len(arr)-1])
		} else {
			return Maxn(arr[1:])
		}
	}
}