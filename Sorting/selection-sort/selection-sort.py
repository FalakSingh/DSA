def selection_sort(arr: list[int]) -> None:
    arr_len = len(arr)

    for i in range(arr_len):
        min_index = i
        for j in range(i, arr_len):
            if arr[j] < arr[min_index]:
                min_index = j
        arr[i], arr[min_index] = arr[min_index], arr[i]

    print(arr)


unsorted_arr = [7, 12, 9, 11, 3, 2, 11, 6, 14, 1]

selection_sort(unsorted_arr)
