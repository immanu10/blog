Title: TypeScript Decorators
Date: 21-Feb-2025

Decorators provide a way to add metadata and modify classes, methods, properties, and parameters in TypeScript. They offer a declarative approach to enhancing code functionality without directly altering the original code structure.

Think of decorators as wrappers around your code elements. They allow you to inject additional logic before, after, or around the execution of the decorated element.  This can be useful for various tasks, including:

* **Logging:** Track method calls and arguments.
* **Access Control:** Restrict access to certain methods or properties.
* **Validation:** Enforce data integrity before processing.
* **Instrumentation:** Measure performance or gather usage statistics.

Here's a breakdown of how decorators work in TypeScript:

**1. Declaration:**

Decorators are prefixed with the `@` symbol followed by the decorator function.  They can be applied to classes, methods, properties, and parameters.

```typescript
function myDecorator(target: any, propertyKey: string | symbol) {
  console.log(`Decorating property: ${String(propertyKey)}`);
}

class MyClass {
  @myDecorator
  myProperty: string = "Hello";
}
```

**2. Decorator Factories:**

For more complex scenarios, you can use decorator factories. These are functions that return a decorator function, allowing you to pass arguments to the decorator.

```typescript
function logDecorator(message: string) {
  return function (target: any, propertyKey: string | symbol) {
    const originalMethod = target[propertyKey];

    target[propertyKey] = function (...args: any[]) {
      console.log(message);
      const result = originalMethod.apply(this, args);
      return result;
    };
  };
}

class MyClass {
  @logDecorator("Calling myMethod")
  myMethod(arg1: string, arg2: number): string {
    return `${arg1} ${arg2}`;
  }
}

const instance = new MyClass();
instance.myMethod("Test", 123); // Output: "Calling myMethod"
```

**3. Method Decorators:**

Method decorators receive three arguments:

* `target`: The prototype of the class.
* `propertyKey`: The name of the method.
* `descriptor`: A `PropertyDescriptor` containing information about the method.

**4. Class Decorators:**

Class decorators receive one argument:

* `constructor`: The constructor function of the class.

**5. Property Decorators:**

Property decorators receive two arguments:

* `target`: The prototype of the class or the constructor function if applied to a static property.
* `propertyKey`: The name of the property.

**Example: Validation Decorator**

```typescript
function validate(validator: (value: any) => boolean) {
  return function (target: any, propertyKey: string | symbol) {
    let value: any;

    Object.defineProperty(target, propertyKey, {
      get() {
        return value;
      },
      set(newValue) {
        if (!validator(newValue)) {
          throw new Error("Invalid value");
        }
        value = newValue;
      },
    });
  };
}

class User {
  @validate((value: string) => value.length > 0)
  name: string = "";
}

const user = new User();
user.name = "John"; // Valid
// user.name = ""; // Throws "Invalid value"
```


By leveraging decorators effectively, you can significantly improve code readability, maintainability, and extensibility in your TypeScript projects. They provide a powerful mechanism for separating concerns and applying cross-cutting logic in a clean and concise manner.
