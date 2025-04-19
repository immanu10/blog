Title: Monotonic Stack
Date: 19-Apr-2025

A monotonic stack is a special type of stack where the elements are either strictly increasing or strictly decreasing. This property allows for efficient solutions to certain problems involving finding the next greater or smaller element in an array.

Let's explore how a monotonic stack works with an example of finding the next greater element for each element in an array.

Consider the array `[2, 1, 5, 6, 2, 3]`. We want to find the next element to the right that is greater than each element.

1. **Initialization:** We start with an empty stack.

2. **Iteration:** We iterate through the array from left to right.

3. **Stack Operations:**
    - If the stack is empty or the current element is less than or equal to the top element of the stack, we push the current element and its index onto the stack. This maintains the monotonically decreasing property.
    - If the current element is greater than the top element of the stack, we pop elements from the stack until we find an element greater than the current element or the stack becomes empty. For each popped element, the current element is its next greater element.  We then push the current element onto the stack.

4. **Result:** After iterating through the array, the elements remaining in the stack don't have a next greater element to their right.

Here's a JavaScript implementation:

```javascript
function nextGreaterElement(arr) {
  const stack = [];
  const result = new Array(arr.length).fill(-1); // Initialize result with -1

  for (let i = 0; i < arr.length; i++) {
    while (stack.length > 0 && arr[i] > arr[stack[stack.length - 1]]) {
      const topIndex = stack.pop();
      result[topIndex] = arr[i];
    }
    stack.push(i);
  }

  return result;
}

const arr = [2, 1, 5, 6, 2, 3];
const result = nextGreaterElement(arr);
console.log(result); // Output: [5, 5, 6, -1, 3, -1]
```

In this example, for the element `2` at index `0`, the next greater element is `5`. For `1` at index `1`, it's also `5`.  `6` doesn't have a next greater element, so the result is `-1`.  And so on.


Monotonic stacks can also be used to find the next smaller element, previous greater element, and previous smaller element.  They are particularly useful in problems involving histograms, largest rectangle in a histogram, and other similar scenarios. They provide an efficient way to maintain a sorted sub-sequence while processing elements, achieving a time complexity of O(n) in many cases.
