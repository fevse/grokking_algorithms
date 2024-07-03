package chapter4

func Maxn(arr []int) int {
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