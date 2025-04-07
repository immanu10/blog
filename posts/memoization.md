Title: Memoization
Date: 07-Apr-2025

Memoization is an optimization technique that speeds up applications by storing the results of expensive function calls and returning the cached result when the same inputs occur again.  It's a trade-off of space (to store the cached results) for time (saved by not recalculating).

Think of it like a lookup table. The first time you call a function with specific inputs, the result is calculated and stored in the table.  Subsequent calls with the same inputs skip the calculation and retrieve the result directly from the table.

Here's a JavaScript example demonstrating memoization:

```javascript
function expensiveCalculation(n) {
  console.log(`Calculating ${n}!`); // Simulate an expensive operation
  let result = 1;
  for (let i = 2; i <= n; i++) {
    result *= i;
  }
  return result;
}

function memoize(func) {
  const cache = new Map(); // Use a Map to store cached results

  return function(...args) { // Use rest parameters to handle any number of arguments
    const key = JSON.stringify(args); // Create a key from the arguments
    if (cache.has(key)) {
      console.log(`Fetching ${args} from cache`);
      return cache.get(key);
    } else {
      const result = func(...args);
      cache.set(key, result);
      return result;
    }
  };
}


const memoizedCalculation = memoize(expensiveCalculation);

console.log(memoizedCalculation(5)); // Calculates and caches 5!
console.log(memoizedCalculation(5)); // Retrieves 5! from cache
console.log(memoizedCalculation(10)); // Calculates and caches 10!
console.log(memoizedCalculation(5)); // Retrieves 5! from cache again
console.log(memoizedCalculation(10)); // Retrieves 10! from cache
```

In this example, `memoize` is a higher-order function that takes a function (`func`) and returns a memoized version of it.  The `cache` object stores the results of previous calculations.  Before performing the expensive calculation, the memoized function checks if the result for the given arguments is already in the cache. If so, it returns the cached result. Otherwise, it calculates the result, stores it in the cache, and then returns it.

The output demonstrates how the expensive calculation is only performed once for each unique set of inputs:

```
Calculating 5!
120
Fetching 5 from cache
120
Calculating 10!
3628800
Fetching 5 from cache
120
Fetching 10 from cache
3628800
```

Memoization is particularly effective for functions with:

* **Expensive computations:** Functions that involve complex calculations, network requests, or database queries can benefit greatly.
* **Repeated calls with the same inputs:** If a function is called multiple times with the same arguments, memoization avoids redundant computations.  This is often the case in recursive functions or functions used in loops.

However, keep in mind that memoization consumes memory to store the cached results.  It's most beneficial when the cost of recalculating is significantly higher than the cost of storing the cached values.  For functions with a large number of possible input combinations, the cache could grow very large, potentially negating the performance benefits.
