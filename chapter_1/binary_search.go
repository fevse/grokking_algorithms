package main

import "fmt"

func main() {
	var n, l, i int
	fmt.Println("Введите длинну массива: ")
	fmt.Scan(&l)
	data := make([]int, 0, l)
	fmt.Print("Введите элементы массива: ")
	for i := 0; i < l; i++ {
		fmt.Scan(&n)
		data = append(data, n)
	}
	fmt.Println("Введите искомый элемент: ")
	fmt.Scan(&i)
	fmt.Println("Вычисляю результат ...")
	index, step := binarySearch(data, i)
	fmt.Printf("Индекс = %v, число шагов = %v\n", index, step)
}

func binarySearch(data []int, item int) (int, int) {
	var low, high, step int
	high = len(data)-1
	for {
		step += 1
		mid := (low + high) / 2
		if data[mid] == item {
			return mid, step
		} else if data[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
		if low > high {
			return 0, 0
		}
	}
}