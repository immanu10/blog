Title: React Props
Date: 19-Mar-2025

React props are a fundamental concept for building dynamic and reusable components.  They allow you to pass data from a parent component to its children, enabling communication and customization within your application. Think of props as arguments you pass to a function; they determine how the component renders and behaves.

Here's a breakdown of how props work in React:

**Passing Props:**

Props are passed as attributes within the opening tag of a child component.  The value of the prop can be a variety of JavaScript data types, including strings, numbers, arrays, objects, and even functions.

```javascript
// Parent component
function ParentComponent() {
  const name = "John Doe";
  const age = 30;

  return (
    <div>
      <ChildComponent name={name} age={age} />
    </div>
  );
}

// Child component
function ChildComponent(props) {
  return (
    <div>
      <p>Name: {props.name}</p>
      <p>Age: {props.age}</p>
    </div>
  );
}
```

In this example, `name` and `age` are props passed from `ParentComponent` to `ChildComponent`. Inside `ChildComponent`, these props are accessed as properties of the `props` object.

**Destructuring Props:**

You can destructure props within the child component's parameter list for cleaner syntax:

```javascript
function ChildComponent({ name, age }) {
  return (
    <div>
      <p>Name: {name}</p>
      <p>Age: {age}</p>
    </div>
  );
}

```

**Default Props:**

You can define default values for props to handle cases where a parent component doesn't provide a value:

```javascript
ChildComponent.defaultProps = {
  name: "Guest User",
  age: 0
};
```

**Prop Types:**

While not mandatory, using `prop-types` (a separate library) is highly recommended. It allows you to define the expected data type for each prop, providing helpful warnings during development if incorrect types are passed.

```javascript
import PropTypes from 'prop-types';

ChildComponent.propTypes = {
  name: PropTypes.string.isRequired,
  age: PropTypes.number
};
```

**Key Takeaway:**

Props are a crucial mechanism for building reusable and flexible React components. They facilitate data flow from parent to child, enabling dynamic rendering based on the provided input.  Understanding how to use props effectively is essential for structuring well-organized and maintainable React applications.
