package chapter4

func Qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		less := make([]int, 0)
		greater := make([]int, 0)
		for _, v := range arr[1:] {
			if v <= pivot {
				less = append(less, v)
			} else {
				greater = append(greater, v)
			}
		}
		res := make([]int, 0)
		
		if len(less) > 0 {
			res = append(res, Qsort(less)...)
		}
		
		res = append(res, pivot)

		if len(greater) > 0 {
			res = append(res, Qsort(greater)...)
		}

		return res
	}
}