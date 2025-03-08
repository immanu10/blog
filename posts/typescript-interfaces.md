Title: TypeScript Interfaces
Date: 08-Mar-2025

Interfaces in TypeScript are a powerful tool for defining the shape of objects. They act as contracts, ensuring that objects adhere to a specific structure.  This is crucial for maintainability and catching errors during development, especially in larger projects.  Interfaces primarily enhance code clarity and type safety by explicitly specifying the expected properties and methods of objects.

**Basic Usage:**

```typescript
interface Person {
  name: string;
  age: number;
  greet(message: string): string; 
}

let user: Person = {
  name: "John Doe",
  age: 30,
  greet: function(message: string): string {
    return `${message}, ${this.name}!`;
  }
};

console.log(user.greet("Hello")); // Output: Hello, John Doe!


// Error: Property 'greet' is missing in type '{ name: string; age: number; }' but required in type 'Person'.
let badUser: Person = { 
    name: "Jane Doe", 
    age: 25
}; 
```

In this example, `Person` is an interface defining the structure for a person object.  It mandates that any object assigned to a variable of type `Person` *must* have a string property `name`, a number property `age`, and a method `greet` with a specific signature.  Trying to assign an object missing `greet` (like `badUser`) results in a compile-time error.

**Optional Properties:**

Sometimes, not all properties are required. You can mark properties as optional using the `?` symbol:

```typescript
interface Product {
  name: string;
  description?: string; // Optional property
  price: number;
}

let product1: Product = {
  name: "Awesome Gadget",
  price: 99.99
};

let product2: Product = {
  name: "Another Gadget",
  description: "This gadget is even more awesome!",
  price: 129.99
};
```

**Readonly Properties:**

To prevent modification after initialization, use `readonly`:

```typescript
interface Point {
  readonly x: number;
  readonly y: number;
}

let point: Point = { x: 10, y: 20 };

// Error: Cannot assign to 'x' because it is a read-only property.
point.x = 30; 
```

**Extending Interfaces:**

Interfaces can inherit from each other using the `extends` keyword. This promotes code reuse and creates a clear hierarchy.

```typescript
interface Shape {
    color: string;
}

interface Square extends Shape {
    sideLength: number;
}

let square: Square = {
    color: "red",
    sideLength: 10
};
```

**Index Signatures:**

For objects with an arbitrary number of properties, use index signatures:

```typescript
interface StringDictionary {
  [index: string]: string;
}

let myDictionary: StringDictionary = {
  "key1": "value1",
  "key2": "value2"
};

console.log(myDictionary["key1"]); // Output: value1
```

**Function Types:**

Interfaces can also describe function types:

```typescript
interface MathFunction {
  (x: number, y: number): number;
}

let add: MathFunction = (x: number, y: number) => x + y;
let subtract: MathFunction = (x: number, y: number) => x - y;

console.log(add(5, 3)); // Output: 8
```

By leveraging these features, TypeScript interfaces significantly improve code organization, readability, and maintainability, making them essential for building robust and scalable applications.
