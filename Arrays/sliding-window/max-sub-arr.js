function maxSubArray(arr, k) {
  let maxSum = 0;

  for (let i = 0; i < k; i++) {
    maxSum += arr[i];
  }
  let windowSum = maxSum;
  for (let i = k; i < arr.length; i++) {
    windowSum = windowSum - arr[i - k] + arr[i];
    if (windowSum > maxSum) {
      maxSum = windowSum;
    }
  }
  return maxSum;
}

maxSubArray([1, 2, 3, 4, 5, 1, 2, 3], 3);
