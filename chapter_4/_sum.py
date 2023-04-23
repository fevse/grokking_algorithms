def _sum(arr):
    if len(arr) == 1:
        return arr[0]
    else:
        return arr[0] + _sum(arr[1:])
