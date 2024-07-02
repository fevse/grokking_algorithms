package chapter4

func Count(arr []int) int {
	if len(arr) == 1 {
		return 1
	} else {
		return Count(arr[1:]) + 1
	}
}