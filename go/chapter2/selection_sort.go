package selection_sort

func findSmallest(arr []int) int {
	s := arr[0]
	si := 0
	for i, v := range arr {
		if v < s {
			s = v
			si = i
		}
	}
	return si
}

func SelectionSort(arr []int) ([]int) {
	res := make([]int, len(arr))
	if len(arr) == 0 {
		return res
	}
	l := len(arr)
	var s int
	for i := 0; i < l; i++ {
		s = findSmallest(arr)
		res[i] = arr[s]
		arr = append(arr[:s], arr[s+1:]...)
	}
	return res
}