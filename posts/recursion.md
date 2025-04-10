Title: Recursion
Date: 10-Apr-2025

Recursion is a powerful programming technique where a function calls itself within its own definition. This allows for elegant solutions to problems that can be broken down into smaller, self-similar subproblems.  A recursive function typically consists of two parts:

1. **Base Case:**  The condition that stops the recursion. Without a base case, the function would call itself indefinitely, leading to a stack overflow error.

2. **Recursive Step:** The part where the function calls itself with a modified input, moving closer to the base case with each call.

Let's illustrate with a classic example: calculating the factorial of a number.  The factorial of a non-negative integer `n`, denoted by `n!`, is the product of all positive integers less than or equal to `n`.

```javascript
function factorial(n) {
  // Base case: factorial of 0 is 1
  if (n === 0) {
    return 1;
  }

  // Recursive step: n! = n * (n-1)!
  return n * factorial(n - 1);
}

console.log(factorial(5)); // Output: 120
```

In this example, `factorial(5)` calculates `5 * factorial(4)`. `factorial(4)` then calculates `4 * factorial(3)`, and so on.  This continues until `factorial(0)` is reached, which returns 1 (the base case). The values then unwind, multiplying back up the call stack to produce the final result.

Another example is traversing a tree-like data structure:

```javascript
const tree = {
  value: 1,
  children: [
    { value: 2, children: [] },
    { value: 3, children: [{ value: 4, children: [] }] },
  ],
};

function printTree(node) {
  console.log(node.value); // Print the current node's value

  // Recursive step: iterate through children and call printTree on each
  for (const child of node.children) {
    printTree(child);
  }
}

printTree(tree); // Output: 1 2 3 4
```

Here, `printTree` visits each node and recursively calls itself on its children. The base case is implicitly when a node has no children – the loop simply doesn't execute.

While recursion can be powerful, it's crucial to design the base case carefully to avoid infinite loops.  Excessive recursion depth can also lead to stack overflow errors.  However, for problems that exhibit self-similarity, recursion offers an elegant and concise solution.
