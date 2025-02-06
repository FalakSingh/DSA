def bubble_sort(arr: list[int]):
    arr_len = len(arr)
    for i in range(arr_len - 1):
        print(arr)
        swapped = False
        for j in range(arr_len - 1 - i):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
                swapped = True
        if not swapped:
            break
    print(arr)


unsorted_arr = [7, 12, 9, 11, 3, 2, 11, 6, 14, 1]
bubble_sort(unsorted_arr)
