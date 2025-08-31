function isPalindrom(str) {
  let [left, right] = [0, str.length - 1];
  while (left < right) {
    if (str[left] !== str[right]) {
      return false;
    }
    left++;
    right--;
  }
  return true;
}
