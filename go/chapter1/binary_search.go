package binarysearch

func BinarySearch(data []int, item int) (int, int) {
	var low, high, step int
	high = len(data)-1
	for {
		if low > high {
			return 0, 0
		}
		step += 1
		mid := (low + high) / 2
		if data[mid] == item {
			return mid, step
		} else if data[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}