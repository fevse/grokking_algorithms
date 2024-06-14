def _count(arr):
    if len(arr) == 1:
        return 1
    else:
        return 1 + _count(arr[1:])
