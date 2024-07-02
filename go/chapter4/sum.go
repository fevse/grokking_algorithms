package chapter4

func Sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	} else {
		return Sum(arr[1:]) + arr[0]
	}
}