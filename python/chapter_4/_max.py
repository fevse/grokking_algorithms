def _max(arr):
    if len(arr) == 0:
        return 0
    elif len(arr) == 1:
        return arr[0]
    else:
        if arr[0] > arr[1]:
            arr[1] = arr[0]
        return _max(arr[1:])
