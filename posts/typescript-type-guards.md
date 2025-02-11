Title: TypeScript Type Guards
Date: 11-Feb-2025

Type guards are a powerful TypeScript feature that helps you gain more control over types within your conditional logic.  They allow the compiler to understand which type is currently in use within a specific block of code, enabling better type checking and autocompletion.

Let's say you have a variable that can be either a string or a number:

```typescript
let value: string | number = "hello";
```

If you try to access string-specific properties without checking the type, TypeScript will complain:

```typescript
console.log(value.toUpperCase()); // Error: Property 'toUpperCase' does not exist on type 'string | number'.
```

This is where type guards come in. They help you narrow down the type within an `if` statement or other conditional logic.

**1. `typeof` operator:** The simplest type guard is the `typeof` operator.

```typescript
if (typeof value === "string") {
  console.log(value.toUpperCase()); // Works!
} else {
  console.log(value.toFixed(2)); // Works assuming value is a number in this branch
}
```

Within the `if` block, TypeScript now knows that `value` is a string, and in the `else` block, it infers `value` as a number.

**2. `instanceof` operator:** Use `instanceof` to check if a value is an instance of a particular class.

```typescript
class Dog {
  bark() {
    console.log("Woof!");
  }
}

class Cat {
  meow() {
    console.log("Meow!");
  }
}

let pet: Dog | Cat = new Dog();

if (pet instanceof Dog) {
  pet.bark(); // Works!
} else {
  pet.meow(); // Works!
}

```


**3. User-defined type guards:**  For more complex scenarios, you can create custom type guard functions. These functions must return a type predicate, which is a special return type that tells TypeScript what type a variable is if the function returns `true`.

```typescript
interface Car {
  wheels: number;
  model: string;
}

interface Bicycle {
  wheels: number;
  type: string;
}

function isCar(vehicle: Car | Bicycle): vehicle is Car {
  return (vehicle as Car).model !== undefined;
}

let vehicle: Car | Bicycle = { wheels: 4, model: "Sedan" };

if (isCar(vehicle)) {
  console.log(vehicle.model); // Works!
} else {
  console.log(vehicle.type); // Works!
}

```

The `vehicle is Car` syntax is the type predicate. It asserts that if `isCar` returns `true`, the `vehicle` variable is of type `Car`.

Using type guards allows you to write more robust and type-safe code, making your TypeScript projects easier to maintain and debug. They empower you to leverage TypeScript's full type checking capabilities even when dealing with complex conditional logic.
