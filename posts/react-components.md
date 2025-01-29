Title: React Components
Date: 29-Jan-2025

React applications are built using reusable blocks of code called components.  These components encapsulate logic, styling, and markup, promoting modularity and maintainability.  Think of them like Lego bricks – individual pieces that combine to create complex structures.

There are two main types of React components: functional components and class components. While class components were prevalent in older React codebases, functional components, especially when combined with hooks, have become the preferred approach for modern React development.

**Functional Components:**

These are JavaScript functions that accept props (properties) as input and return JSX (JavaScript XML) describing the UI.

```javascript
function Welcome(props) {
  return <h1>Hello, {props.name}</h1>;
}

// Usage:
const element = <Welcome name="Sarah" />;
// Renders "Hello, Sarah"
```

This simple `Welcome` component takes a `name` prop and displays a personalized greeting.  It's concise and easy to understand.

**More Complex Example:**

```javascript
function Toggle(props) {
  const [isOn, setIsOn] = React.useState(false);

  const handleClick = () => {
    setIsOn(!isOn);
  };

  return (
    <div>
      <button onClick={handleClick}>
        {isOn ? 'ON' : 'OFF'}
      </button>
      {isOn && <p>Toggle is on!</p>}
    </div>
  );
}
```

This `Toggle` component uses the `useState` hook to manage its internal state, demonstrating how functional components handle dynamic behavior. Clicking the button toggles the `isOn` state, conditionally rendering the paragraph.


**Why Components?**

- **Reusability:** Write once, use anywhere. Components avoid code duplication and make maintenance easier.
- **Organization:** Break down complex UIs into manageable pieces. This improves code readability and collaboration.
- **State Management:**  Components can manage their own internal state, making it easier to handle user interactions and dynamic updates.
- **Testability:**  Smaller, isolated units of code are simpler to test thoroughly.


Components are the fundamental building blocks of any React application. Mastering their use is key to building robust and maintainable user interfaces.
