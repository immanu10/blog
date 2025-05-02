Title: TypeScript Mapped Types
Date: 02-May-2025

Mapped types in TypeScript allow you to create new types based on existing ones by iterating over their properties and applying a transformation.  They provide a concise way to manipulate and transform type definitions, reducing boilerplate and improving code maintainability.

Let's explore how mapped types work and how they can be useful.

### Basic Usage

Consider an interface representing a user:

```typescript
interface User {
  firstName: string;
  lastName: string;
  age: number;
  isActive: boolean;
}
```

Suppose you need a new type that represents the same user data but with all properties made optional. You could manually create a new interface, but mapped types offer a more elegant solution:

```typescript
type PartialUser = {
  [P in keyof User]?: User[P];
};
```

Let's break down this code:

* `[P in keyof User]` iterates over each property key (`firstName`, `lastName`, etc.) in the `User` interface, assigning each key to the type variable `P`.
* `User[P]` accesses the type of each property using the key `P`.
* `?` makes each property optional in the new `PartialUser` type.

### Built-in Mapped Types

TypeScript provides several handy built-in mapped types:

* `Partial<T>`: Makes all properties of `T` optional.
* `Required<T>`: Makes all properties of `T` required.
* `Readonly<T>`: Makes all properties of `T` read-only.
* `Pick<T, K>`: Creates a new type by picking specific properties `K` from `T`.
* `Omit<T, K>`: Creates a new type by omitting specific properties `K` from `T`.
* `Record<K, T>`: Defines a new type with keys of type `K` and values of type `T`.

Using the built-in `Partial` would simplify our previous example:

```typescript
type PartialUser = Partial<User>;
```

### Transforming Property Types

Mapped types can also transform property types. Let's say you want a type where all string properties of `User` are converted to uppercase:

```typescript
type UppercaseUser = {
  [P in keyof User]: User[P] extends string ? Uppercase<User[P]> : User[P];
};
```

Here, we use a conditional type to check if the property type `User[P]` extends `string`. If it does, we apply `Uppercase<User[P]>` to convert it to uppercase. Otherwise, the original property type is retained.


### Practical Application

Mapped types are particularly useful when working with forms, data transformations, and creating utility types. For instance, you can use them to define types for API responses, create types for form inputs with validation, or easily transform data between different formats.

By mastering mapped types, you can write more concise, type-safe, and maintainable TypeScript code. They provide powerful tools for manipulating and transforming types, reducing boilerplate, and improving the overall structure of your projects.
