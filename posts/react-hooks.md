Title: React Hooks
Date: 03-Jan-2025

React Hooks are a powerful feature introduced in React 16.8 that allow functional components to access state and lifecycle methods, eliminating the need for class components in many cases.  They provide a more concise and readable way to manage component logic, making code easier to understand, test, and maintain.

**Why Hooks?**

Before hooks, functional components were stateless and couldn't participate in the component lifecycle.  This often led to developers using class components even for simple tasks, introducing unnecessary complexity. Hooks solve this by letting you "hook into" React state and lifecycle features directly from functional components.

**useState Hook:**

The `useState` hook is the most fundamental hook. It allows functional components to manage state. It takes an initial state value and returns an array containing the current state and a function to update it.

```javascript
import React, { useState } from 'react';

function Counter() {
  const [count, setCount] = useState(0);

  return (
    <div>
      <p>Count: {count}</p>
      <button onClick={() => setCount(count + 1)}>Increment</button>
    </div>
  );
}
```

In this example, `count` is the current state variable initialized to 0, and `setCount` is a function to update the `count` state.

**useEffect Hook:**

The `useEffect` hook is used for side effects in functional components.  Side effects are anything that affects something outside the scope of the component, like data fetching, DOM manipulation, or timers.  It replaces lifecycle methods like `componentDidMount`, `componentDidUpdate`, and `componentWillUnmount`.

```javascript
import React, { useState, useEffect } from 'react';

function DataFetcher() {
  const [data, setData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      const response = await fetch('https://api.example.com/data');
      const result = await response.json();
      setData(result);
    };

    fetchData();
  }, []); // Empty dependency array means this runs only once on mount

  return (
    <div>
      {data ? <pre>{JSON.stringify(data, null, 2)}</pre> : <p>Loading...</p>}
    </div>
  );
}
```

Here, `useEffect` fetches data when the component mounts. The empty dependency array ensures the effect runs only once, equivalent to `componentDidMount`.

**Custom Hooks:**

One of the greatest benefits of hooks is the ability to create custom hooks. This allows you to extract reusable logic into separate functions.

```javascript
import React, { useState, useEffect } from 'react';

function useWindowSize() {
  const [size, setSize] = useState({ width: window.innerWidth, height: window.innerHeight });

  useEffect(() => {
    const handleResize = () => {
      setSize({ width: window.innerWidth, height: window.innerHeight });
    };

    window.addEventListener('resize', handleResize);

    return () => window.removeEventListener('resize', handleResize); // Cleanup
  }, []);

  return size;
}

function MyComponent() {
  const windowSize = useWindowSize();

  return (
    <div>
      Window width: {windowSize.width}, height: {windowSize.height}
    </div>
  );
}

```

`useWindowSize` is a custom hook that encapsulates the logic for tracking window size. This promotes code reuse and makes components cleaner.


By understanding and utilizing hooks effectively, you can write cleaner, more maintainable React code.  They are a fundamental concept in modern React development.
