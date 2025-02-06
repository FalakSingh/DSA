function bubbleSort(arr) {
  let arrLen = arr.length;

  for (let i = 0; i < arrLen - 1; i++) {
    let swapped = false;
    for (let j = 0; j < arrLen - 1 - i; j++) {
      if (arr[j] > arr[j + 1]) {
        [arr[j], arr[j + 1]] = [arr[j + 1], arr[j]];
        swapped = true;
      }
    }
    if (!swapped) {
      break;
    }
  }

  console.log(arr);
}

const unsortedArray = [7, 12, 9, 11, 3, 2, 11, 6, 14, 1];
bubbleSort(unsortedArray);
