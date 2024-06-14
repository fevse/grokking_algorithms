def find_smallest(arr):
    smallest = arr[0]
    smallest_index = 0
    for i in range(1, len(arr)):
        if arr[i] < smallest:
            smallest = arr[i]
            smallest_index = i
    return smallest_index

def selection_sort(arr):
    res = []
    for i in range(len(arr)):
        smallest = find_smallest(arr)
        res.append(arr.pop(smallest))
    return res

data = input('Введите элементы массива через пробел: ')
data = list(map(int, data.split()))
print(selection_sort(data))

