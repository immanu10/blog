Title: TypeScript keyof Operator
Date: 30-Apr-2025

The `keyof` operator in TypeScript is a powerful tool for working with object types.  It extracts the keys of an object type as a union of string literals. This allows you to create type-safe code that interacts with object properties dynamically.

Let's say you have an interface defining a user:

```typescript
interface User {
  id: number;
  name: string;
  email?: string; // Optional property
}
```

Using `keyof`, you can create a type that represents the valid property names of `User`:

```typescript
type UserKey = keyof User;
```

`UserKey` will be equivalent to the type `'id' | 'name' | 'email'`.  This means any variable of type `UserKey` can only hold one of those three strings.  This becomes incredibly useful when you need to write functions that access object properties dynamically.

Here's an example:

```typescript
function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}

const user: User = {
  id: 123,
  name: 'John Doe',
  email: 'john.doe@example.com'
};

const userId = getProperty(user, 'id'); // userId is of type number
const userName = getProperty(user, 'name'); // userName is of type string

// The following line will cause a compile-time error because 'age' is not a key of User
// const userAge = getProperty(user, 'age'); 
```

In this example, `getProperty` is a generic function that takes an object `obj` of type `T` and a key `key` of type `K`. The `K extends keyof T` constraint ensures that `key` is a valid property name of `obj`.  This provides type safety and prevents accessing non-existent properties at compile time.

Another useful application of `keyof` is in creating lookup types:

```typescript
interface Product {
  name: string;
  price: number;
  description: string;
}

type ProductAttributes = {
  [K in keyof Product]: string;
};
```

`ProductAttributes` is a type that has the same keys as `Product`, but all the values are strings. This can be useful when you need to transform the values of an object while maintaining its keys. For instance, if you need to represent all product attributes as strings for a search index.

The `keyof` operator helps write more robust and type-safe code by ensuring that you are only working with valid property names of an object, ultimately making your TypeScript code easier to maintain and debug.
