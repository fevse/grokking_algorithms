package chapter3

func countdown(i int) int {
	if i == 1 {
		return i
	} else {
		return countdown(i-1) + i
	}
}