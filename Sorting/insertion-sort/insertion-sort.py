def insertion_sort(arr):
    arr_len = len(arr)
    for i in range(1, arr_len):
        current_el = arr[i]
        j = i - 1
        while j >= 0 and arr[j] > current_el:
            arr[j + 1] = arr[j]
            j = j - 1
        arr[j + 1] = current_el


unsorted_array = [64, 34, 25, 12, 22, 11, 90, 5]
insertion_sort(unsorted_array)
