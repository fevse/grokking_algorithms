def quicksort(arr):
    if len(arr) < 2:
        return arr
    else:
        pivot_index = len(arr) // 2
        pivot = arr[pivot_index]
        less = [i for i in arr[:pivot_index] + arr[pivot_index+1:] if i <= pivot]
        greater = [i for i in arr[:pivot_index] + arr[pivot_index+1:] if i > pivot]
        return quicksort(less) + [pivot] + quicksort(greater)


data = input('Введите элементы массива через пробел: ')
data = list(map(int, data.split()))
print(quicksort(data))

