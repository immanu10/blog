Title: React Context
Date: 17-Mar-2025

React Context provides a way to pass data through the component tree without having to pass props down manually at every level.  This is particularly useful for sharing data that is globally relevant to your application, like user authentication, theme preferences, or locale.  Think of it as a global state for your React components.

Before Context, the typical way to pass data down was through "prop drilling": passing props from parent to child, then child to grandchild, and so on. This can become cumbersome and difficult to maintain in complex applications.

Here's how React Context works and how to use it:

**Creating a Context:**

```javascript
const MyContext = React.createContext(defaultValue);
```

`createContext` creates a Context object. The `defaultValue` is optional and is used when a component consumes the context but doesn't have a matching provider above it.

**Providing Context:**

```javascript
import MyContext from './MyContext'; // Import the context

function MyProvider({ children }) {
  const [theme, setTheme] = useState('light');

  const contextValue = {
    theme,
    toggleTheme: () => setTheme(theme === 'light' ? 'dark' : 'light'),
  };

  return (
    <MyContext.Provider value={contextValue}>
      {children}
    </MyContext.Provider>
  );
}
```

The `<MyContext.Provider>` component makes the `contextValue` available to all consuming components within its subtree.  Here, we are providing both the current `theme` and a function `toggleTheme` to change it.

**Consuming Context:**

```javascript
import MyContext from './MyContext'; // Import the context

function MyComponent() {
  const { theme, toggleTheme } = useContext(MyContext);

  return (
    <div style={{ backgroundColor: theme === 'light' ? 'white' : 'gray' }}>
      <button onClick={toggleTheme}>Toggle Theme</button>
      <p>Current theme: {theme}</p>
    </div>
  );
}
```

The `useContext` hook allows components to consume the context.  It takes the Context object (`MyContext`) as an argument and returns the current context value.  Now, `MyComponent` can access both `theme` and `toggleTheme` without receiving them as props.

**Example Structure:**

```javascript
<MyProvider>
  <App>
    <ComponentA />
    <ComponentB>
      <ComponentC />
    </ComponentB>
  </App>
</MyProvider>
```

If `ComponentC` needs the context value, it doesn't need to receive it as props from `ComponentB` and then `App`.  It can directly consume it using `useContext(MyContext)`.

**Key Advantages of Using Context:**

* **Avoids prop drilling:**  Simplifies passing data to deeply nested components.
* **Centralized state management:**  Provides a single source of truth for shared data.
* **Improved code readability and maintainability:** Makes it easier to understand and modify the data flow.


React Context offers a convenient and powerful way to manage and share state within your application.  It promotes cleaner code, reduces complexity, and enhances maintainability.
