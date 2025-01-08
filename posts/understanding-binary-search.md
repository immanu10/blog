Title: Understanding Binary Search
Date: 08-Jan-2025

Binary search is an efficient algorithm for finding a target value within a *sorted* array. It follows a divide-and-conquer approach, repeatedly halving the search space until the target is found or the search space is exhausted.

**How it works:**

1. **Start with the middle element:**  Compare the middle element of the array with the target value.

2. **Three possibilities:**
    - If the middle element *equals* the target, the search is successful.
    - If the middle element is *greater* than the target, discard the right half of the array and repeat the search on the left half.
    - If the middle element is *less* than the target, discard the left half of the array and repeat the search on the right half.

3. **Repeat:** Continue this process until the target is found or the remaining search space is empty.

**Example (JavaScript):**

```javascript
function binarySearch(arr, target) {
  let left = 0;
  let right = arr.length - 1;

  while (left <= right) {
    const mid = Math.floor((left + right) / 2);

    if (arr[mid] === target) {
      return mid; // Target found at index mid
    } else if (arr[mid] < target) {
      left = mid + 1; // Search in the right half
    } else {
      right = mid - 1; // Search in the left half
    }
  }

  return -1; // Target not found
}

const sortedArray = [2, 5, 8, 12, 16, 23, 38, 56, 72, 91];
const targetValue = 23;

const result = binarySearch(sortedArray, targetValue);

if (result !== -1) {
  console.log(`Target ${targetValue} found at index ${result}`);
} else {
  console.log(`Target ${targetValue} not found in the array`);
}


const targetValue2 = 24; //testing element not present
const result2 = binarySearch(sortedArray, targetValue2);
if (result2 !== -1) {
  console.log(`Target ${targetValue2} found at index ${result2}`);
} else {
  console.log(`Target ${targetValue2} not found in the array`);
}
```

**Time Complexity:** O(log n) - Since the search space is halved with each comparison, the time complexity is logarithmic. This makes binary search significantly faster than linear search (O(n)) for large sorted arrays.

**Space Complexity:** O(1) - Binary search uses a constant amount of extra space, making it very space-efficient.

**Key Requirement:** The input array *must* be sorted for binary search to work correctly.  If the array is not sorted, you'll need to sort it first or use a different search algorithm.
