def binary_search(data, item):
    low = 0
    high = len(data) - 1
    step = 0
    while low <= high:
        mid = (low + high) // 2
        guess = data[mid]
        step += 1
        if guess == item:
            return f'Индекс = {mid}, число шагов = {step}'
        if guess > item:
            high = mid - 1
        else:
            low = mid + 1
    return 'Элемент не найден'


assert binary_search(
    [1, 2, 3, 4, 5, 6, 7, 8, 9], 5) == 'Индекс = 4, число шагов = 1'
assert binary_search(
    [1, 2, 3, 4, 5, 6, 7, 8, 9], 6) == 'Индекс = 5, число шагов = 3'
assert binary_search(
    [1, 2, 3, 4, 5, 6, 7, 8, 9], 1) == 'Индекс = 0, число шагов = 3'
assert binary_search(
    [1, 2, 3, 4, 5, 6, 7, 8, 9], 55) == 'Элемент не найден'


data = input('Введите элементы массива через пробел: ')
data = list(map(int, data.split()))
item = int(input('Введите искомый элемент: '))
print(binary_search(data, item))
