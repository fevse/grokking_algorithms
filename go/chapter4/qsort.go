package chapter4

func Qsort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		mid := len(arr)/2
		pivot := arr[mid]
		less := make([]int, 0)
		greater := make([]int, 0)
		arrwp := make([]int, 0)
		arrwp = append(arrwp, arr[:mid]...)
		arrwp = append(arrwp, arr[mid+1:]...)
		for _, v := range arrwp {
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