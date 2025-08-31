const removeDuplicates = (arr) => {
  if (arr.length === 0) return;
  let writeIndex = 1;

  for (readIndex = 1; readIndex < arr.length; readIndex++) {
    if (arr[readIndex] !== arr[readIndex - 1]) {
      arr[writeIndex] = arr[readIndex];
      writeIndex++;
    }
  }

  arr.length = writeIndex;
  return writeIndex;
};

const sortedArr = [1, 1, 2, 2, 3, 4, 4, 5]
removeDuplicates(sortedArr);
console.log(sortedArr)