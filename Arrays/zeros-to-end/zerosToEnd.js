function zerosToEnd(arr) {
  let writeIndex = 0;

  for (let readIndex = 0; readIndex < arr.length; readIndex++) {
    if (arr[readIndex] !== 0) {
      arr[writeIndex] = arr[readIndex];
      writeIndex++;
    }
  }

  while (writeIndex < arr.length) {
    arr[writeIndex] = 0;
    writeIndex++;
  }
}

const array = [0, 1, 0, 2, 3, 0, 4, 0, 5];
zerosToEnd(array);
console.log(array);
